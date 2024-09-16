package routes

import (
	"bytes"
	_ "embed"
	"os"

	bash "github.com/tkdeng/gobash"
	regex "github.com/tkdeng/goregex"
	"github.com/tkdeng/goutil"
)

//todo: add routes for: python, haskell, c#, assembly (maybe), and other functional programming languages

//go:embed template/run-route.go
var runRouteTemplate []byte

var RouteCompiler = map[string]func(src, dist, path string, isDir bool){}

func addCompiler(name string, cb func(src, dist, path string, isDir bool)) {
	RouteCompiler[name] = cb
}

func genBinary(out string, dir string, lastArgFirst bool, cmd string, args ...string) {
	tmp := runRouteTemplate

	argStr := [][]byte{}
	for _, arg := range args {
		argStr = append(argStr, regex.JoinBytes('`', goutil.HTML.EscapeArgs([]byte(arg), '`'), '`'))
	}

	tmp = bytes.ReplaceAll(tmp, []byte("{DIR}"), []byte(goutil.HTML.EscapeArgs([]byte(dir), '`')))
	tmp = bytes.ReplaceAll(tmp, []byte("{CMD}"), []byte(goutil.HTML.EscapeArgs([]byte(cmd), '`')))
	tmp = bytes.ReplaceAll(tmp, []byte("`{ARGS}`"), bytes.Join(argStr, []byte{','}))

	if lastArgFirst {
		tmp = bytes.ReplaceAll(tmp, []byte("`{LASTARGFIRST}`"), []byte("true"))
	}else{
		tmp = bytes.ReplaceAll(tmp, []byte("`{LASTARGFIRST}`"), []byte("false"))
	}

	file, err := os.CreateTemp("/tmp", "*.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(tmp)
	file.Sync()
	bash.Run([]string{`go`, `build`, `-o`, out, file.Name()}, "", nil, true)
}
