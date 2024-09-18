package webserver

import (
	"fmt"
	"testing"
)

func TestRoutes(t *testing.T) {
	loadConfig("./test")
	compRoutes(Config.Root+"/src/routes", Config.Root+"/dist/routes", "")

	page, err := getRoute("/")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(page.buf))
}

func TestServer(t *testing.T) {
	return

	app, err := New("./test")
	if err != nil {
		t.Error(err)
	}

	app.Listen()
}
