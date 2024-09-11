package routes

import (
	"os"
	"path/filepath"

	bash "github.com/tkdeng/gobash"
)

func init() {
	addCompiler("go", func(root, path string, isDir bool) {
		dir, file, out := getPaths(root, path, isDir, "go")
		if out == "" {
			return
		}

		os.MkdirAll(filepath.Dir(out), 0755)
		if isDir {
			// bash.Run([]string{`go`, `build`, `-o`, strings.TrimSuffix(out, ".go")}, dir, nil, true)
			bash.Run([]string{`go`, `build`, `-o`, out + ".bin"}, dir, nil, true)
		} else {
			// bash.Run([]string{`go`, `build`, `-o`, strings.TrimSuffix(out, ".go"), file}, dir, nil, true)
			bash.Run([]string{`go`, `build`, `-o`, out + ".bin", file}, dir, nil, true)
		}
	})
}
