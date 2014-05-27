package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"amovies_parser/conf"
	"amovies_parser/helpers"
	"amovies_parser/models"
	"strconv"

	"github.com/aquilax/cyrslug"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func IndexPage(r render.Render) {
	r.HTML(200, "index", nil)
}
func LinksPage(r render.Render, request *http.Request) {
	serial := &models.Serial{}
	serial.Init(get_param(request, "url"))
	serial.ParseEpisodes()
	r.HTML(200, "links", serial)
}

func DownloadsPage(r render.Render) {
	r.HTML(200, "downloads", conf.DOWNLOADS)
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

func AddDownload(w http.ResponseWriter, r *http.Request) {
	filename := get_filename(r)
	link := r.FormValue("link")
	helpers.StartDownload(link, filename)
	fmt.Fprintln(w, "Download started: ", filename)
}

func translite(str string) string {
	return cyrslug.Slug(str)
}

func get_filename(r *http.Request) (result string) {
	result = fmt.Sprintf("%s_%s.mp4",
		translite(r.FormValue("prefix")),
		translite(r.FormValue("episode")))
	result = strings.Replace(result, " ", "_", -1)
	result = strings.TrimSpace(result)
	return result
}
