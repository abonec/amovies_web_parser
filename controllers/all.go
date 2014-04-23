package controllers

import (
  "net/http"
  "fmt"
  "net/url"
  "strings"
  "github.com/aquilax/cyrslug"
  "html/template"
  "amovies_parser/models"
  "amovies_parser/helpers"
  "amovies_parser/conf"
  "github.com/go-martini/martini"
  "strconv"
)

var (
  TEMPLATES = template.Must(template.ParseFiles("views/index.tpl", "views/links.tpl", "views/get_serial_form.tpl", "views/downloads.tpl", "views/assets.tpl"))
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
  render(w, "index_page", nil)
}
func LinksPage(w http.ResponseWriter, r *http.Request) {
  serial := &models.Serial{}
  serial.Init(get_param(r, "url"))
  serial.ParseEpisodes()

  render(w, "links_page", serial)
}
func DownloadsPage(w http.ResponseWriter, r *http.Request) {
  // t, _ := template.ParseFiles("views/downloads.tpl")
  // t.Execute(w, DOWNLOADS)
  render(w, "downloads_page", conf.DOWNLOADS)
}

func RemoveDownload(params martini.Params) string {
  id, _ := strconv.Atoi(params["id"])
  conf.DOWNLOADS.Remove(id)
  return ""
}
func get_param(r *http.Request, key string) string {
  uri, _ := url.Parse(r.RequestURI)
  return uri.Query().Get(key)
}

func render(w http.ResponseWriter, template string, data interface{}){
  TEMPLATES.ExecuteTemplate(w, template, data)
}

func AddDownload(w http.ResponseWriter, r *http.Request) {
  filename := get_filename(r)
  link := r.FormValue("link")
  helpers.StartDownload(link, filename)
  fmt.Fprintln(w, "Download started: ", filename)
}

func translite(str string) string {
  return cyrslug.Slug(str)
}

func get_filename(r *http.Request)(result string){
  result = fmt.Sprintf("%s_%s.mp4", 
    translite(r.FormValue("prefix")),
    translite(r.FormValue("episode")))
  result = strings.Replace(result, " ", "_", -1)
  result = strings.TrimSpace(result)
  return result
}
