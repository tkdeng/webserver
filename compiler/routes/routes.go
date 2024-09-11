package routes

import (
	"path/filepath"
	"strings"

	"github.com/tkdeng/goutil"
)

var RouteCompiler = map[string]func(root string, path string, isDir bool){}

func addRoute(name string, cb func(root string, path string, isDir bool)) {
	RouteCompiler[name] = cb
}

func getPaths(root, path string, isDir bool, ext string) (dir string, file string, out string) {
	if isDir {
		dir = path
		file = "index." + ext

		if p, err := goutil.JoinPath("./dist/routes", strings.Replace(path, root, "", 1)); err == nil {
			out = p
			if !strings.HasSuffix(out, "."+ext) {
				out += "." + ext
			}
		}
	} else {
		dir = filepath.Dir(path)
		file = filepath.Base(path)
		if !strings.HasSuffix(file, "."+ext) {
			file += "." + ext
		}

		if p, err := goutil.JoinPath("./dist/routes", strings.Replace(path, root, "", 1)); err == nil {
			out = p
			if !strings.HasSuffix(out, "."+ext) {
				out += "." + ext
			}
		}
	}

	return dir, file, out
}
