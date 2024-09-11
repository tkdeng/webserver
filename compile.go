package main

import (
	"fmt"
	"os"
	"server/compiler/routes"
	"strings"
	"time"

	regex "github.com/tkdeng/goregex"
	"github.com/tkdeng/goutil"
	"github.com/tkdeng/htmlc"
)

func compile() {
	os.MkdirAll("./dist", 0755)

	os.RemoveAll("./dist/routes")
	if err := os.Mkdir("./dist/routes", 0755); err != nil {
		panic(err)
	}

	compRoutes("./src/routes", "")

	// compTemplates()
}

func compRoutes(root string, dir string) {
	if path, err := goutil.JoinPath(root); err == nil {
		root = path
	}

	if dir == "" {
		dir = root
	}

	if files, err := os.ReadDir(dir); err == nil {
		for _, file := range files {
			if file.IsDir() {
				if path, err := goutil.JoinPath(dir, file.Name()); err == nil {
					if regex.Comp(`\.([\w_-]+)$`).Match([]byte(file.Name())) {
						lang := ""
						regex.Comp(`\.([\w_-]+)$`).RepFunc([]byte(file.Name()), func(data func(int) []byte) []byte {
							lang = string(data(1))
							return nil
						})

						if lang != "" {
							if cb, ok := routes.RouteCompiler[lang]; ok {
								cb(root, path, true)
								continue
							}
						}
					}

					compRoutes(root, strings.Replace(path, root, "", 1))
				}
			} else if path, err := goutil.JoinPath(dir, file.Name()); err == nil {
				lang := ""
				regex.Comp(`\.([\w_-]+)$`).RepFunc([]byte(file.Name()), func(data func(int) []byte) []byte {
					lang = string(data(1))
					return nil
				})

				if lang != "" {
					if cb, ok := routes.RouteCompiler[lang]; ok {
						cb(root, path, false)
						continue
					}
				}
			}
		}
	}
}

func compTemplates() {
	if err := htmlc.Compile("./src/templates", "./dist/templates.exs"); err != nil {
		fmt.Println(err)
	}

	lastUpdate := time.Now().UnixMilli()

	fw := goutil.FileWatcher()
	fw.OnAny = func(path, op string) {
		if now := time.Now().UnixMilli(); now-lastUpdate > 1000 {
			lastUpdate = now
			if err := htmlc.Compile("./src/templates", "./dist/templates.exs"); err != nil {
				fmt.Println(err)
			}
		}
	}
	fw.WatchDir("./src/templates")
}
