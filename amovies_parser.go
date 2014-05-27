package main

import (
	"amovies_parser/conf"
	"amovies_parser/controllers"
	"log"
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(martini.Static("public"))
	m.Use(martini.Static(conf.DOWNLOAD_DIR))
	m.Post("/download", controllers.AddDownload)
	m.Delete("/remove/:id", controllers.RemoveDownload)
	m.Get("/links", controllers.LinksPage)
	m.Get("/downloads", controllers.DownloadsPage)
	m.Get("/", controllers.IndexPage)
	log.Fatal(http.ListenAndServe(":6100", m))
}
