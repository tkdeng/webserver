package routes

import (
	"os"

	bash "github.com/tkdeng/gobash"
)

func init() {
	addCompiler("scala", func(root, path string, isDir bool) {
		_, _, out := getPaths(root, path, isDir, "scala")
		if out == "" {
			return
		}

		// may treat scala file like a directory (object defines path)
		if !isDir {
			os.MkdirAll(out, 0755)
			bash.Run([]string{`scalac`, path}, out, nil, true)
		}

		//todo: add directory handler for scala
	})
}
