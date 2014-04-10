package main
import (
  "net/http"
  "net/url"
  "io"
  "os"
  "github.com/gorilla/mux"
  "fmt"
)

var (
  DOWNLOAD_DIR = ""
)


func init() {
  DOWNLOAD_DIR = os.Getenv("DIR")
}
func main() {
  r := mux.NewRouter()
  r.HandleFunc("/down", add_download).Methods("GET")
  http.Handle("/", r)
  http.ListenAndServe(":4545", nil)
}

func add_download(w http.ResponseWriter, r *http.Request) {
  go download(get_param(r, "url"), get_param(r, "filename"))
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
