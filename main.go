package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/tkdeng/goutil"
)

//!htmlc: --src="src/templates" --dist="dist"

func main() {
	//todo: handle server
	// also compile with separate htmlc module
	// routes will be compiled by server (not by htmlc)

	// compile()

	//todo: create method to run dist files in 'routes' directory
	// also include dist route handlers in routes.go

	out, err := runRoute("index")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(out))
}

var dynamicRouteReg = regexp.MustCompile(`\[([\w_-]+)(\?|)\]`)

func runRoute(url string) ([]byte, error) {
	path, err := goutil.JoinPath("./dist/routes")
	if err != nil {
		return []byte{}, err
	}

	args := []string{}

	url = strings.Trim(url, "/")
	for _, page := range strings.Split(url, "/") {
		p, err := goutil.JoinPath(path, page)
		if err != nil {
			return []byte{}, err
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
						return []byte{}, err
					}

					path = p
					hasPath = true
					break
				}
			}
		}

		if !hasPath {
			return []byte{}, os.ErrNotExist
		}
	}

	stat, err := os.Stat(path)
	if err != nil || stat.IsDir() {
		return []byte{}, os.ErrNotExist
	}

	cmd := exec.Command(path, args...)
	out, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

/* func runRoute(name string) ([]byte, error) {
	path, err := goutil.JoinPath("./dist/routes", name)
	if err != nil {
		return []byte{}, err
	}

	var cmd *exec.Cmd
	if stat, err := os.Stat(path + ".bin"); err == nil && !stat.IsDir() {
		cmd = exec.Command(path + ".bin")
	} else if stat, err := os.Stat(path + ".class"); err == nil && !stat.IsDir() {
		cmd = exec.Command(`scala`, filepath.Base(path))
		cmd.Dir = filepath.Dir(path)
	}

	if cmd == nil {
		return []byte{}, errors.New("page not found")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
} */
