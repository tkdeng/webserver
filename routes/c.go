package routes

import (
	"os"
	"path/filepath"

	bash "github.com/tkdeng/gobash"
)

func init() {
	addCompiler("c", func(root, path string, isDir bool) {
		dir, file, out := getPaths(root, path, isDir, "c")
		if out == "" {
			return
		}

		os.MkdirAll(filepath.Dir(out), 0755)
		// bash.Run([]string{`gcc`, `-o`, strings.TrimSuffix(out, ".c"), file}, dir, nil, true)
		bash.Run([]string{`gcc`, `-o`, out + ".bin", file}, dir, nil, true)
	})
}
