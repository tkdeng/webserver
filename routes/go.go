package routes

import (
	"strings"

	bash "github.com/tkdeng/gobash"
)

func init() {
	addCompiler("go", func(src, dist, path string, isDir bool) {
		if isDir {
			bash.Run([]string{`go`, `build`, `-o`, strings.TrimSuffix(dist+"/"+path, ".go")}, src+"/"+path, nil, true)
		} else {
			bash.Run([]string{`go`, `build`, `-o`, strings.TrimSuffix(dist+"/"+path, ".go"), path}, src, nil, true)
		}
	})
}
