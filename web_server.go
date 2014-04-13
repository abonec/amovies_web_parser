package main

import (
  "net/http"
  "fmt"
  "github.com/gorilla/mux"
  "net/url"
  "strings"
  "github.com/aquilax/cyrslug"
  "html/template"
)


var (
  TEMPLATES = template.Must(template.ParseFiles("views/index.tpl", "views/links.tpl", "views/get_serial_form.tpl"))
)

func main() {
  fmt.Println("Download dir: ",DOWNLOAD_DIR)
  r := mux.NewRouter()
  r.HandleFunc("/download", add_download).Methods("POST")
  r.HandleFunc("/links", links_page).Methods("GET")
  r.HandleFunc("/", index_page).Methods("GET")
  http.Handle("/", r)
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
  http.ListenAndServe(":6100", nil)
}

func index_page(w http.ResponseWriter, r *http.Request) {
  render(w, "index_page", nil)
}
func links_page(w http.ResponseWriter, r *http.Request) {
  serial := &Serial{}
  serial.Init(get_param(r, "url"))
  serial.ParseEpisodes()

  render(w, "links_page", serial)
}
func get_param(r *http.Request, key string) string {
  uri, _ := url.Parse(r.RequestURI)
  return uri.Query().Get(key)
}

func render(w http.ResponseWriter, template string, data interface{}){
  TEMPLATES.ExecuteTemplate(w, template, data)
}

func add_download(w http.ResponseWriter, r *http.Request) {
  filename := fmt.Sprintf("%s_%s.mp4", 
    translite(r.FormValue("prefix")),
    translite(r.FormValue("episode")))
  filename = strings.Replace(filename, " ", "_", -1)
  filename = strings.TrimSpace(filename)
  go download(r.FormValue("link"), filename)
  fmt.Fprintln(w, "Queued:", filename)
}

func translite(str string) string {
  return cyrslug.Slug(str)
}

