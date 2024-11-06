package routes

import (
	"os"
	"strings"

	bash "github.com/tkdeng/gobash"
)

func init() {
	addCompiler("scala", func(src, dist, path string, isDir bool) {
		if isDir {
			outDir := strings.TrimSuffix(dist+"/"+path, ".scala")
			os.MkdirAll(outDir, 0755)
			bash.Run([]string{`scalac`, src + "/" + path + "/index.scala"}, outDir, nil, true)
			genBinary(outDir+"/[index]", outDir, true, `scala`)
		} else {
			outDir := strings.TrimSuffix(dist+"/"+path, ".scala")
			os.MkdirAll(outDir, 0755)
			bash.Run([]string{`scalac`, src + "/" + path}, outDir, nil, true)
			genBinary(outDir+"/[index]", outDir, true, `scala`)
		}
	})
}
