package routes

import (
	"os"
	"path/filepath"
	"strings"

	bash "github.com/tkdeng/gobash"
)

func init() {
	addCompiler("c", func(src, dist, path string, isDir bool) {
		if isDir {
			os.MkdirAll(filepath.Dir(dist+"/"+path), 0755)
			bash.Run([]string{`gcc`, `-o`, strings.TrimSuffix(path, ".c"), src + "/" + path + "/index.c"}, dist, nil, true)
		} else {
			os.MkdirAll(filepath.Dir(dist+"/"+path), 0755)
			bash.Run([]string{`gcc`, `-o`, strings.TrimSuffix(path, ".c"), src + "/" + path}, dist, nil, true)
		}
	})
}
