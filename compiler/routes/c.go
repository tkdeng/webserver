package routes

import (
	bash "github.com/tkdeng/gobash"
)

func init() {
	addRoute("c", func(root, path string, isDir bool) {
		dir, file, out := getPaths(root, path, isDir, "c")

		if out != "" {
			// bash.Run([]string{`gcc`, `-o`, strings.TrimSuffix(out, ".c"), file}, dir, nil, true)
			bash.Run([]string{`gcc`, `-o`, out + ".bin", file}, dir, nil, true)
		}
	})
}
