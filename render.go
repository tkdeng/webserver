package webserver

import (
	"bytes"
	"encoding/base64"
	"os"
	"os/exec"
	"regexp"
	"strings"

	regex "github.com/tkdeng/goregex"
	"github.com/tkdeng/goutil"
)

type PageRoute struct {
	Page   string
	Layout string
	Args   map[string]any

	// for debugging
	buf []byte
}

var dynamicRouteReg = regexp.MustCompile(`\[([\w_-]+)(\?|)\]`)

func getRoute(url string) (PageRoute, error) {
	resPage := PageRoute{Page: "404", Layout: "layout", Args: map[string]any{}}

	if url == "/" || url == "" {
		url = "index"
	}

	// try json
	if path, err := goutil.JoinPath(Config.Root+"/routes.bin", url); err == nil {
		if stat, err := os.Stat(path + ".json"); err == nil && !stat.IsDir() {
			if err := goutil.ReadJson(path+".json", &resPage); err == nil {
				return resPage, nil
			}
		} else if stat, err := os.Stat(path + "/index.json"); err == nil && !stat.IsDir() {
			if err := goutil.ReadJson(path+"/index.json", &resPage); err == nil {
				return resPage, nil
			}
		}
	}

	path, err := goutil.JoinPath(Config.Root + "/routes.bin")
	if err != nil {
		return resPage, err
	}

	args := []string{}

	url = strings.Trim(url, "/")
	for _, page := range strings.Split(url, "/") {
		p, err := goutil.JoinPath(path, page)
		if err != nil {
			return resPage, err
		}

		hasPath := false
		if _, err := os.Stat(p); err == nil {
			path = p
			hasPath = true
		} else if files, err := os.ReadDir(path); err == nil {
			for _, file := range files {
				if dynamicRouteReg.MatchString(file.Name()) {
					args = append(args, page)

					p, err = goutil.JoinPath(path, file.Name())
					if err != nil {
						return resPage, err
					}

					path = p
					hasPath = true
					break
				}
			}
		}

		if !hasPath {
			return resPage, os.ErrNotExist
		}
	}

	stat, err := os.Stat(path)
	if err != nil {
		return resPage, os.ErrNotExist
	}

	if stat.IsDir() {
		stat, err = os.Stat(path + "/index")
		if err != nil || stat.IsDir() {
			if files, err := os.ReadDir(path); err == nil {
				for _, file := range files {
					if dynamicRouteReg.MatchString(file.Name()) {
						args = append(args, "index")

						p, err := goutil.JoinPath(path, file.Name())
						if err != nil {
							return resPage, err
						}

						path = p
						break
					}
				}
			}
		}
	}

	stat, err = os.Stat(path)
	if err != nil || stat.IsDir() {
		return resPage, os.ErrNotExist
	}

	cmd := exec.Command(path, args...)
	out, err := cmd.Output()
	if err != nil {
		return resPage, err
	}

	regex.Comp(`(?m)^@(PAGE|LAYOUT|ARGS):(.*);\r?\n?$`).RepFunc(out, func(data func(int) []byte) []byte {
		buf := data(2)
		if bytes.Equal(data(1), []byte("PAGE")) {
			resPage.Page = string(buf)
		} else if bytes.Equal(data(1), []byte("LAYOUT")) {
			resPage.Layout = string(buf)
		} else if bytes.Equal(data(1), []byte("ARGS")) && len(buf) != 0 {
			if !(buf[0] == '{' && buf[len(buf)-1] == '}') {
				if b, err := base64.StdEncoding.DecodeString(string(buf)); err == nil {
					buf = b
				}
			}

			if json, err := goutil.JSON.Parse(buf); err == nil {
				goutil.JoinMap(&resPage.Args, &json)
			}
		}
		return []byte{}
	})

	// for debugging
	resPage.buf = out

	return resPage, nil
}
