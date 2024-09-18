# Web Server

[<img src="./assets/icon.png" alt="icon" height="100"/>](./assets/icon.png)

A Web Server Written In... Any Programming Language?

Develop a website with your favorite programming language.
Currently Supported Languages: C, Go, Scala.

This module runs a server with a golang base for handling http.
Templates are rendered in [htmlc](https://github.com/tkdeng/htmlc) which compiles to elixir.

> Notice: This Project Is Still In Beta.

## Installation

```shell
go get github.com/tkdeng/webserver
```

## Usage

```go
func main(){
  // create new server
  app, err := webserver.New("./app")
  
  // do normal gofiber stuff (optional)
  app.Get("/", func(c fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  app.Get("/", func(c fiber.Ctx) error {
    // note: c.Render is not used
    // you can use a secondary template with gofiber
    return webserver.Render(c, "index", htmlc.Map{"args": 1}, "layout")
  })

  // listen with openssl (default port: [http: 8080, ssl: 8443])
  err = app.Listen()
}
```

## Inside App Directory

### config.yml

```yaml
title: "Web Server"
app_title: "WebServer"
desc: "A Web Server."

public_uri: "/public/"

port_http: 8080
port_ssl: 8443

origins: [
  "localhost",
  "example.com",
]

proxies: [
  "127.0.0.1",
  "192.168.0.1",
]

DebugMode: no
```

### templates

Templates are compiled by the [htmlc](https://github.com/tkdeng/htmlc) module.

### routes

Route files laid out in a nextjs like format.
These files can be written in a variety of different programming languages.

Print `@PAGE:<name>;`, `@LAYOUT:<name>;`, and `@ARGS:<json|base64>;` to specify what page to render.

Here is an example in C.

```c
#include <stdio.h>

int main(){
  printf("@PAGE:%s;\n", "index");
  printf("@ARGS:%s;\n", "{\"args\": 1}");
  printf("@LAYOUT:%s;\n", "layout");

  return 0;
}
```

YAML and JSON can also be used for simpler templates.

```yaml
page: index
layout: layout
args: {
  title: "Hello, Yaml!",
}
```

In Scala, the file name represents a base route, and an object represents a sub path.
For example, if the file is named `test.scala`.

```scala
object index { // `/test`
  def main(args: Array[String]) = {
    println("Hello, Scala!")
  }
}

object about { // `/test/about`
  def main(args: Array[String]) = {
    println("About, Scala!")
  }
}
```

We can also name a directory with a languages extension, to have the folder compiled as a single file.

For example, in go we can name out directory `about.go/`.
Then we can write our go module inside this directory.

In this example, `main.go` will be compiled with it's `go.mod` file, and any other files or subdirectories.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, Golang!")
}
```
