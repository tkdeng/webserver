package webserver

import (
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/tkdeng/goutil"
	"github.com/tkdeng/htmlc"
)

type ConfigData struct {
	Title    string
	AppTitle string
	Desc     string

	PublicURI string

	Origins []string
	Proxies []string

	OriginErrHandler func(c fiber.Ctx, err error) error

	PortHTTP uint16
	PortSSL  uint16

	DebugMode bool

	Root string
}

var Config = ConfigData{
	Title:    "Web Server",
	AppTitle: "WebServer",
	Desc:     "A Web Server.",

	PortHTTP: 8080,
	PortSSL:  8443,
}

var Engine *htmlc.ExsEngine

type App struct {
	*fiber.App
}

// New loads a new server
func New(root string) (App, error) {
	// load config file
	loadConfig(root)

	// compile src
	compile()

	var err error
	Engine, err = htmlc.Engine(Config.Root + "/templates.exs")
	if err != nil {
		return App{}, err
	}

	app := fiber.New(fiber.Config{
		PassLocalsToViews:       true,
		AppName:                 Config.AppTitle,
		ServerHeader:            Config.Title,
		TrustedProxies:          Config.Proxies,
		EnableTrustedProxyCheck: true,
		EnableIPValidation:      true,
	})

	compressAssets := !Config.DebugMode
	app.Get("/theme/*", static.New(Config.Root+"/theme", static.Config{Compress: compressAssets}))
	// app.Get("/assets/wasm/*", static.New(Config.Root+"/wasm", static.Config{Compress: compressAssets}))
	app.Get("/assets/*", static.New(Config.Root+"/assets", static.Config{Compress: compressAssets}))
	if Config.PublicURI != "" {
		app.Get(Config.PublicURI, static.New(Config.Root+"/public", static.Config{Compress: compressAssets, Browse: true}))
	}

	if Config.OriginErrHandler == nil {
		Config.OriginErrHandler = func(c fiber.Ctx, err error) error {
			c.SendStatus(403)
			return c.SendString(err.Error())
		}
	}

	// enforce specific domain and ip origins
	app.Use(VerifyOrigin(Config.Origins, Config.Proxies, Config.OriginErrHandler))

	// auto redirect http to https
	if Config.PortSSL != 0 {
		app.Use(RedirectSSL(Config.PortHTTP, Config.PortSSL))
	}

	return App{app}, nil
}

// Listen to both http and https ports and
// auto generate a self signed ssl certificate
// (will also auto renew every year)
//
// by using self signed certs, you can use a proxy like cloudflare and
// not have to worry about verifying a certificate athority like lets encrypt
func (app *App) Listen() error {
	app.Use(func(c fiber.Ctx) error {
		url := goutil.Clean(c.Path())
		// method := goutil.Clean(c.Method())

		/* if method == "POST" {
			//todo: render apis
		} */

		return RenderPage(c, url)
	})

	return ListenAutoTLS(app.App, Config.PortHTTP, Config.PortSSL, Config.Root+"/db/ssl/auto_ssl")
}

func loadConfig(root string) {
	// load config file
	if path, err := filepath.Abs(root); err == nil {
		root = path
	}
	root = strings.TrimSuffix(root, "/")

	goutil.ReadConfig(root+"/config.yml", &Config)
	Config.Root = root
}

func Render(c fiber.Ctx, name string, args htmlc.Map, layout ...string) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	buf, err := Engine.Render(name, args, layout...)
	if err != nil {
		c.SendStatus(404)

		if buf, err = Engine.Render("404", args, layout...); err == nil {
			return c.Send(goutil.Clean(buf))
		}

		return c.Send([]byte("<h1>Error 404</h1><h2>Page Not Found!</h2>"))
	}

	c.SendStatus(200)
	return c.Send(goutil.Clean(buf))
}

func RenderPage(c fiber.Ctx, url string) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	page, err := getRoute(url)
	if err != nil {
		c.SendStatus(404)
		if page, err = getRoute("/404"); err == nil {
			if buf, err := Engine.Render(page.Page, page.Args, page.Layout); err == nil {
				return c.Send(goutil.Clean(buf))
			}
		}

		if buf, err := Engine.Render("404", page.Args, page.Layout); err == nil {
			return c.Send(goutil.Clean(buf))
		}

		return c.Send([]byte("<h1>Error 404</h1><h2>Page Not Found!</h2>"))
	}

	buf, err := Engine.Render(page.Page, page.Args, page.Layout)
	if err != nil {
		c.SendStatus(404)

		if buf, err = Engine.Render("404", page.Args, page.Layout); err == nil {
			return c.Send(goutil.Clean(buf))
		}

		return c.Send([]byte("<h1>Error 404</h1><h2>Page Not Found!</h2>"))
	}

	c.SendStatus(200)
	return c.Send(goutil.Clean(buf))
}
