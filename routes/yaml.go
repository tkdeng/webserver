package routes

import (
	"os"
	"strings"

	regex "github.com/tkdeng/goregex"
	"github.com/tkdeng/goutil"
)

type PageRoute struct {
	Page   string
	Layout string
	Args   map[string]any
}

func init() {
	addCompiler("yml", func(src, dist, path string, isDir bool) {
		compYml("yml", src, dist, path, isDir)
	})

	addCompiler("yaml", func(src, dist, path string, isDir bool) {
		compYml("yaml", src, dist, path, isDir)
	})

	addCompiler("json", func(src, dist, path string, isDir bool) {
		page := PageRoute{Page: "index", Layout: "layout", Args: map[string]any{}}
		if err := goutil.ReadJson(src+"/"+path, &page); err == nil {
			if args, err := goutil.JSON.Stringify(page.Args, 2); err == nil {
				os.WriteFile(dist+"/"+path, regex.JoinBytes(
					'{', '\n',
					`  "page": "`, page.Page, `",`, '\n',
					`  "layout": "`, page.Layout, `",`, '\n',
					`  "args": `, args, '\n',
					'}', '\n',
				), 0755)
			}
		}
	})
}

func compYml(ext string, src, dist, path string, isDir bool) {
	page := PageRoute{Page: "index", Layout: "layout", Args: map[string]any{}}
	if err := goutil.ReadYaml(src+"/"+path, &page); err == nil {
		if args, err := goutil.JSON.Stringify(page.Args, 2); err == nil {
			os.WriteFile(dist+"/"+strings.TrimSuffix(path, "."+ext)+".json", regex.JoinBytes(
				'{', '\n',
				`  "page": "`, page.Page, `",`, '\n',
				`  "layout": "`, page.Layout, `",`, '\n',
				`  "args": `, args, '\n',
				'}', '\n',
			), 0755)
		}
	}
}
