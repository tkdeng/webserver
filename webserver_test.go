package webserver

import (
	"testing"
)

func Test(t *testing.T) {
	app, err := New("./test")
	if err != nil {
		t.Error(err)
	}

	err = app.Listen()
	if err != nil {
		t.Error(err)
	}
}
