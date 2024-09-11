package routes

import (
	bash "github.com/tkdeng/gobash"
)

func init() {
	addRoute("go", func(root, path string, isDir bool) {
		dir, file, out := getPaths(root, path, isDir, "go")

		if out != "" {
			if isDir {
				// bash.Run([]string{`go`, `build`, `-o`, strings.TrimSuffix(out, ".go")}, dir, nil, true)
				bash.Run([]string{`go`, `build`, `-o`, out + ".bin"}, dir, nil, true)
			} else {
				// bash.Run([]string{`go`, `build`, `-o`, strings.TrimSuffix(out, ".go"), file}, dir, nil, true)
				bash.Run([]string{`go`, `build`, `-o`, out + ".bin", file}, dir, nil, true)
			}
		}

	})
}
