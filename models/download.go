package models

import (
  "time"
  "encoding/json"
  "os"
  "io/ioutil"
)

type Download struct {
  Filename string
  Link string
  Progress int
  Downloaded bool
  Length int64
  Added time.Time
}
type Downloads struct {
  Downloading map[*Download]bool
  Downloaded map[*Download]bool
}
func(d *Downloads)Init(){
  d.Downloading = make(map[*Download]bool)
  d.Downloaded = make(map[*Download]bool)
}
func(d *Downloads)AddDownload(link, filename string, length int64)(download *Download){
  download = &Download{Filename: filename, Link: link, Progress: 0, Downloaded: false, Length: length, Added: time.Now()}
  d.Downloading[download] = true
  return download
}
func(d *Downloads)Finish(download *Download){
  download.Downloaded = true
  d.Downloaded[download] = true
  delete(d.Downloading, download)
}

func(d *Download)Url() string{
  return "/" + d.Filename
}
func(d *Downloads)ToJson() []byte{
  result , _ := json.Marshal(&d)
  return result
}

func(d *Downloads)SaveToFile(path string)(err error){
  err = ioutil.WriteFile(path, d.ToJson(), 0644)
  return
}

func(d *Downloads)RestoreJson(json_text []byte) {
  err := json.Unmarshal(json_text, d)
  if err != nil {
    d.Init()
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
      d.RestoreJson(json)
    }
  }
}

func close_file(file *os.File){
  if file != nil {
    file.Close()
  }
}
