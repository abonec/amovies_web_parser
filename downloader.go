package main
import (
  "net/http"
  "html/template"
  "net/url"
  "io"
  "io/ioutil"
  "os"
  "github.com/gorilla/mux"
  "github.com/PuerkitoBio/goquery"
  "fmt"
  "strings"
  "regexp"
  "github.com/aquilax/cyrslug"
  "code.google.com/p/go-charset/charset"
)

import _ "code.google.com/p/go-charset/data"

var (
  DOWNLOAD_DIR = ""
  QUALITY_REGEXP = regexp.MustCompile(`url(?P<quality>\d\d\d)=(?P<url>.*?mp4)`)
  TEMPLATES = template.Must(template.ParseFiles("views/index.tpl", "views/links.tpl", "views/get_serial_form.tpl"))
)

type Serial struct {
  Title string
  Link string
  Prefix string
  Episodes []*Episode

}
type Episode struct {
  Title string
  Link string
  VideoLinks map[string]string
}

func(serial *Serial) Init(url string){
  serial.Link = url
  doc, _ := goquery.NewDocument(url)
  prefix, _ := doc.Find(".prev_img img").First().Attr("title")
  serial.Prefix = strings.TrimSpace(convert_string(prefix))
  nodes := doc.Find("#series option")
  serial.Title = convert_string(doc.Find(".title_d_dot span").First().Text())
  serial.Episodes = make([]*Episode, nodes.Length(), nodes.Length())
  nodes.Each(func(i int, s *goquery.Selection){
    title := convert_string(s.Text())
    link, _ := s.Attr("value")
    episode := &Episode{Title: title, Link: link }
    serial.Episodes[i] = episode
  })
}

func(serial *Serial) ParseEpisodes(){
  c := make(chan bool)
  episodes_count := len(serial.Episodes)
  finished := 0
  for _, episode := range serial.Episodes {
    go parse_episode(episode, c)
  }
  for completed := range c {
    if completed {
      finished++
    }
    if finished >= episodes_count {
      close(c)
    }
  }
}

func parse_episode(episode *Episode, channel chan bool) {
  episode.VideoLinks = make(map[string]string)
  doc, _ := goquery.NewDocument(episode.Link)
  vars, _ := doc.Find("object embed").First().Attr("flashvars")
  for _, r := range QUALITY_REGEXP.FindAllStringSubmatch(vars, -1) {
    episode.VideoLinks[r[1]+ "p"] = r[2]
  }
  channel <- true
}

func convert_string(str string) string {
  reader, err := charset.NewReader("windows-1251", strings.NewReader(str))
  if err != nil {
    fmt.Println(err)
  }
  result, err := ioutil.ReadAll(reader)
  if err != nil {
    fmt.Println(err)
  }
  return string(result)
}

func translite(str string) string {
  return cyrslug.Slug(str)
}

func init() {
  DOWNLOAD_DIR = os.Getenv("DIR")
}
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

func add_download(w http.ResponseWriter, r *http.Request) {
  filename := fmt.Sprintf("%s_%s.mp4", 
    translite(r.FormValue("prefix")),
    translite(r.FormValue("episode")))
  filename = strings.Replace(filename, " ", "_", -1)
  filename = strings.TrimSpace(filename)
  go download(r.FormValue("link"), filename)
  fmt.Fprintln(w, "Queued:", filename)
}
func index_page(w http.ResponseWriter, r *http.Request) {
  TEMPLATES.ExecuteTemplate(w, "index_page", nil)
}
func links_page(w http.ResponseWriter, r *http.Request) {
  serial := &Serial{}
  serial.Init(get_param(r, "url"))
  serial.ParseEpisodes()

  TEMPLATES.ExecuteTemplate(w, "links_page", serial)
}
func get_param(r *http.Request, key string) string {
  uri, _ := url.Parse(r.RequestURI)
  return uri.Query().Get(key)
}

func download(_url string, file_name string){
  file_name = DOWNLOAD_DIR + file_name
  tmp_name := file_name + ".part"
  out, err := os.Create(tmp_name)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer out.Close()

  resp, _ := http.Get(_url)
  defer resp.Body.Close()

  fmt.Println("Downloading started ", tmp_name)
  io.Copy(out, resp.Body)
  os.Rename(tmp_name, file_name)
  fmt.Println("Downloaded ", file_name)
}
