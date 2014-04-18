package main

import (
  "net/http"
  "github.com/go-martini/martini"
  "log"
  "./controllers"
  "./conf"
)

func main() {
  m := martini.Classic()
  m.Use(martini.Static("public"))
  m.Use(martini.Static(conf.DOWNLOAD_DIR))
  m.Post("/download", controllers.AddDownload)
  m.Get("/links", controllers.LinksPage)
  m.Get("/downloads", controllers.DownloadsPage)
  m.Get("/", controllers.IndexPage)
  log.Fatal(http.ListenAndServe(":6100", m))
}
