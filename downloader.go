package main
import (
  "net/http"
  "io"
  "os"
  "strconv"
  "fmt"
)


var (
  DOWNLOAD_DIR = ""
  DOWNLOADING = make([]string, 10)
  DOWNLOADED = make([]string, 100)
)

func init() {
  DOWNLOAD_DIR = os.Getenv("DIR")
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
func download_imp(_url, file_name string, progress chan int)(written int64, err error) {
  out, _ := os.Create(file_name)
  resp, _ := http.Get(_url)
  defer out.Close()
  defer resp.Body.Close()
  defer close(progress)

  length, _ := strconv.ParseInt(resp.Header["Content-Length"][0], 10, 64)
  every := length / 100
  percentage := int64(0)


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
  return written, err
}
