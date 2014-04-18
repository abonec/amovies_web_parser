package models

import (
  "time"
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
