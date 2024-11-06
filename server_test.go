package webserver

import (
	"fmt"
	"testing"
)

func TestRoutes(t *testing.T) {
	// return

	loadConfig("./test")
	compRoutes(Config.Root+"/src/routes", Config.Root+"/routes.bin", "")

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
