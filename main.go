package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/polds/MyIP"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		Layout:     "layout",
	}))

	m.Get("", func(out render.Render) {
		ip, _ := myip.GetMyIP()
		out.HTML(200, "index", ip)
	})

	m.Run()
}
