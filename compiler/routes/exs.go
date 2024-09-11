package routes

import "fmt"

func init() {
	addRoute("exs", func(root, path string, isDir bool) {
		//todo: create exs route in dist

		dir, file, out := getPaths(root, path, isDir, "exs")
		fmt.Println(dir, file, out)
	})
}
