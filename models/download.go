package models

import (
  "time"
  "encoding/json"
  "os"
  "fmt"
  "io/ioutil"
)

type Download struct {
  Filename string
  Link string
  Progress int
  Downloaded bool
  Length int64
  Added time.Time
  Id int
}
type Downloads struct {
  Downloading map[*Download]bool
  Downloaded map[int]*Download
  SaveFile string
  DownloadDir string
  LastId int
}
func(d *Downloads)Init(){
  d.Downloading = make(map[*Download]bool)
  d.Downloaded = make(map[int]*Download)
}
func(d *Downloads)AddDownload(link, filename string, length int64)(download *Download){
  download = &Download{Filename: filename, Link: link, Progress: 0, Downloaded: false, Length: length, Added: time.Now()}
  d.Downloading[download] = true
  return download
}
func(d *Downloads)Finish(download *Download){
  download.Downloaded = true
  d.AddToDownloaded(download)
  delete(d.Downloading, download)
  d.SaveToFile()
}
func(d *Downloads)AddToDownloaded(download *Download){
  d.LastId++
  download.Id = d.LastId

  d.Downloaded[download.Id] = download
}

func(d *Downloads)Remove(id int){
  download := d.Downloaded[id]
  download.RemoveFromDisk(d.DownloadDir)
  delete(d.Downloaded, id)
  d.SaveToFile()
}

func(d *Download)Url() string{
  return "/" + d.Filename
}
func(d *Download)RemoveFromDisk(download_dir string) {
  os.Remove(download_dir + d.Filename)
}
func(d *Downloads)ToJson() []byte{
  downloaded := make([]Download, len(d.Downloaded), len(d.Downloaded))
  i := 0
  for _, download := range d.Downloaded {
    downloaded[i] = *download
    i++
  }

  result, err := json.Marshal(downloaded)
  if err != nil {
    fmt.Println(err)
  }
  return result
}

func(d *Downloads)SaveToFile()(err error){
  err = ioutil.WriteFile(d.SaveFile, d.ToJson(), 0644)
  return
}

func(d *Downloads)FromJson(json_text []byte) {
  d.Init()
  downloaded := make([]Download, 100, 100)
  err := json.Unmarshal(json_text, &downloaded)
  if err == nil {
    for _, download := range downloaded {
      d.AddToDownloaded(&download)
    }
  }

}
func(d *Downloads)RestoreFile(path string){
  fi, err := os.Open(path)
  defer close_file(fi)

  if err != nil {
    d.Init()
  }else{
    json, err := ioutil.ReadAll(fi)
    if err != nil {
      d.Init()
    } else {
      d.FromJson(json)
      fmt.Println("Downloads loaded")
    }
  }
  d.SaveFile = path
}

func close_file(file *os.File){
  if file != nil {
    file.Close()
  }
}
