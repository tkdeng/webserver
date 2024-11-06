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
	buf := runRouteTemplate

	argStr := [][]byte{}
	for _, arg := range args {
		argStr = append(argStr, regex.JoinBytes('`', goutil.HTML.EscapeArgs([]byte(arg), '`'), '`'))
	}

	buf = bytes.ReplaceAll(buf, []byte("{DIR}"), []byte(goutil.HTML.EscapeArgs([]byte(dir), '`')))
	buf = bytes.ReplaceAll(buf, []byte("{CMD}"), []byte(goutil.HTML.EscapeArgs([]byte(cmd), '`')))
	buf = bytes.ReplaceAll(buf, []byte("`{ARGS}`"), bytes.Join(argStr, []byte{','}))

	if lastArgFirst {
		buf = bytes.ReplaceAll(buf, []byte("/*{LASTARGFIRST}*/ false"), []byte("true"))
	} else {
		buf = bytes.ReplaceAll(buf, []byte("/*{LASTARGFIRST}*/ false"), []byte("false"))
	}

	file, err := os.CreateTemp("/tmp", "*.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf)
	file.Sync()
	bash.Run([]string{`go`, `build`, `-o`, out, file.Name()}, "", nil, true)
}
