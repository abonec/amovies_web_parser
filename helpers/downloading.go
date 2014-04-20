package helpers

import (
  "os"
  "net/http"
  "strconv"
  "fmt"
  "io"
  "amovies_parser/conf"
)

func StartDownload(link, filename string) {
  progress_chan := make(chan int)
  go func(){
    length_chan := make(chan int64)
    go download(link, filename, progress_chan, length_chan)
    go capture_downloading(link, filename, <- length_chan, progress_chan)
  }()
  <- progress_chan
}

func download(_url, file_name string, progress chan int, length_chan chan int64)(written int64, err error) {
  file_name = conf.DOWNLOAD_DIR + file_name
  temp_name := file_name + ".part"
  out, _ := os.Create(temp_name)
  resp, _ := http.Get(_url)
  defer out.Close()
  defer resp.Body.Close()
  defer close(progress)

  length, _ := strconv.ParseInt(resp.Header["Content-Length"][0], 10, 64)
  length_chan <- length
  every := length / 100
  var percentage int64 = 0


  fmt.Println(length)

  buff := make([]byte, 32*1024)

  for {
    nr, er := resp.Body.Read(buff)
    if nr > 0 {
      nw, ew := out.Write(buff[0:nr])
      if nw > 0 {
        written += int64(nw)
        if written > every * percentage {
          progress <- int(percentage)
          percentage++
        }
      }
      if ew != nil {
        err = ew
        break
      }
      if nr != nw {
        err = io.ErrShortWrite
      }
    }
    if er == io.EOF {
      break
    }
    if er != nil {
      err = er
    }
  }
  os.Rename(temp_name, file_name)
  return written, err
}

func capture_downloading(url, filename string, length int64, progress chan int){
  download := conf.DOWNLOADS.AddDownload(url, filename, length)
  for p := range progress {
    download.Progress = p
    fmt.Println(p)
    /* conf.DOWNLOADS.SaveToFile("downloads.json") */
  }
  conf.DOWNLOADS.Finish(download)
}
