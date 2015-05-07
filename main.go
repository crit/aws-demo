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

	m.Use(render.Renderer(render.Options{
		Directory:  os.Getenv("TEMPLATES"),
		Extensions: []string{".html"},
		Layout:     "layout",
	}))

	storer.InitDb()
	MigratePerson()

	m.Use(storageCheck())

	cacher.InitCache(cacher.Options{
		Hosts:  os.Getenv("MEMCACHE"),
		Engine: cacher.MEMCACHE,
	})

	m.Get("", func(out render.Render) {
		ip, _ := myip.GetMyIP()
		out.HTML(200, "index", ip)
	})

	m.Get("/api/people", func(out render.Render) {
		out.JSON(200, map[string][]Person{"people": PersonList()})
	})

	m.Put("/api/people", func(out render.Render, req *http.Request) {
		name := html.EscapeString(req.FormValue("name"))
		email := html.EscapeString(req.FormValue("email"))

		if err := PersonCreate(name, email); err != nil {
			out.JSON(500, err.Error())
		} else {
			out.Status(201)
		}
	})

	m.NotFound(func(out render.Render) {
		out.HTML(404, "404", nil)
	})

	m.Run()
}

func storageCheck() martini.Handler {
	return func(out render.Render) {
		if storer.MissingDb() {
			out.HTML(500, "error", "Database is unavailable.")
		}
	}
}
