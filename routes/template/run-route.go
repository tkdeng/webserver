package main

import (
	"os"
	"os/exec"
)

const lastArgFirst bool = /*{LASTARGFIRST}*/ false

func main() {
	args := []string{`{ARGS}`}

	if lastArgFirst && len(os.Args) > 1 {
		args = append([]string{os.Args[1]}, append(args, os.Args[2:]...)...)
	} else {
		args = append(args, os.Args[1:]...)
	}

	cmd := exec.Command(`{CMD}`, args...)
	cmd.Dir = `{DIR}`

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
