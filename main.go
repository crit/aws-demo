package main

import (
	"html"
	"net/http"
	"os"

	"github.com/crit/critical-go/cacher"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/polds/MyIP"
)

func main() {
	m := martini.Classic()

	// Template rendering options
	m.Use(render.Renderer(render.Options{
		Directory:  os.Getenv("TEMPLATES"),
		Extensions: []string{".html"},
		Layout:     "layout",
	}))

	// Database initialization and table migration
	if err := storer.InitDb(); err == nil {
		MigratePerson()
	}

	// register storer watcher
	// to catch if the database goes off line
	m.Use(storageCheck())

	// Cacher initialization.
	cacher.InitCache(cacher.Options{
		Hosts:  os.Getenv("MEMCACHE"),
		Engine: cacher.MEMCACHE,
	})

	// HTML UI to client
	m.Get("", func(out render.Render) {
		ip, _ := myip.GetMyIP()
		out.HTML(200, "index", ip)
	})

	// JSON list of registered people to client
	m.Get("/api/people", func(out render.Render) {
		out.JSON(200, map[string][]Person{"people": PersonList()})
	})

	// Save new person to registered people
	m.Put("/api/people", func(out render.Render, req *http.Request) {
		name := html.EscapeString(req.FormValue("name"))
		email := html.EscapeString(req.FormValue("email"))

		if err := PersonCreate(name, email); err != nil {
			out.JSON(500, err.Error())
			return
		}

		out.Status(201)
	})

	// handle client unknown requests from the client
	m.NotFound(func(out render.Render) {
		out.HTML(404, "404", nil)
	})

	// start the server and listen for client requests
	m.Run()
}

func storageCheck() martini.Handler {
	return func(out render.Render) {
		if storer.MissingDb() {
			out.HTML(500, "error", "Database is unavailable.")
		}
	}
}
