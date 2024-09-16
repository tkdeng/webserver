package main

import (
	"fmt"
	"os"
	"server/routes"
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

	compRoutes("./src/routes", "./dist/routes", "")

	//todo: listen for file changes in routes
	// for better performance, simply recompile over existing dist files as needed
	// will need to detect when a route is removed separately, and remove it from dist

	// temp: turnned off compile templates
	// remember to turn back on
	// compTemplates()
}

func compRoutes(src, dist, dir string) {
	if path, err := goutil.JoinPath(src); err == nil {
		src = path
	}

	if path, err := goutil.JoinPath(dist); err == nil {
		dist = path
	}

	fullDir := src
	if dir != "" {
		if path, err := goutil.JoinPath(src, dir); err == nil {
			fullDir = path
		}
	}

	if files, err := os.ReadDir(fullDir); err == nil {
		for _, file := range files {
			if file.IsDir() {
				if regex.Comp(`\.([\w_-]+)$`).Match([]byte(file.Name())) {
					lang := ""
					regex.Comp(`\.([\w_-]+)$`).RepFunc([]byte(file.Name()), func(data func(int) []byte) []byte {
						lang = string(data(1))
						return nil
					})

					if lang != "" {
						if cb, ok := routes.RouteCompiler[lang]; ok {
							if dir != "" {
								dir += "/"
							}
							cb(src, dist, dir+file.Name(), true)
							continue
						}
					}
				}

				if dir != "" {
					dir += "/"
				}
				compRoutes(src, dist, dir+file.Name())
			} else if regex.Comp(`\.([\w_-]+)$`).Match([]byte(file.Name())) {
				lang := ""
				regex.Comp(`\.([\w_-]+)$`).RepFunc([]byte(file.Name()), func(data func(int) []byte) []byte {
					lang = string(data(1))
					return nil
				})

				if lang != "" {
					if cb, ok := routes.RouteCompiler[lang]; ok {
						if dir != "" {
							dir += "/"
						}
						cb(src, dist, dir+file.Name(), false)
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
