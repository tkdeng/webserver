package routes

import "fmt"

func init() {
	addCompiler("js", func(root, path string, isDir bool) {
		//todo: create js route in dist
		// note: for node.js, keep a directory in dist
		// may handle files completely differently

		dir, file, out := getPaths(root, path, isDir, "js")
		if out == "" {
			return
		}

		fmt.Println(dir, file, out)
	})
}
