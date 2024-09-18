package webserver

import (
	"fmt"
	"os"
	"time"
	"webserver/routes"

	regex "github.com/tkdeng/goregex"
	"github.com/tkdeng/goutil"
	"github.com/tkdeng/htmlc"
)

func compile() {
	PrintMsg("warn", "Compilling Server Routes...", 50, false)

	os.MkdirAll(Config.Root+"/dist", 0755)

	os.RemoveAll(Config.Root + "/dist/routes")
	if err := os.Mkdir(Config.Root+"/dist/routes", 0755); err != nil {
		panic(err)
	}

	compRoutes(Config.Root+"/src/routes", Config.Root+"/dist/routes", "")

	//todo: listen for file changes in routes
	// for better performance, simply recompile over existing dist files as needed
	// will need to detect when a route is removed separately, and remove it from dist

	PrintMsg("warn", "Compilling Server Templates...", 50, false)

	// temp: turnned off compile templates
	// remember to turn back on
	compTemplates()

	PrintMsg("confirm", "Compiled Server!", 50, true)
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
	if err := htmlc.Compile(Config.Root+"/src/templates", Config.Root+"/dist/templates.exs"); err != nil {
		fmt.Println(err)
	}

	lastUpdate := time.Now().UnixMilli()

	fw := goutil.FileWatcher()
	fw.OnAny = func(path, op string) {
		if now := time.Now().UnixMilli(); now-lastUpdate > 1000 {
			lastUpdate = now
			if err := htmlc.Compile(Config.Root+"/src/templates", Config.Root+"/dist/templates.exs"); err != nil {
				fmt.Println(err)
			}
		}
	}
	fw.WatchDir(Config.Root + "/src/templates")
}
