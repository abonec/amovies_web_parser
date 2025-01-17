package models

import (
  "github.com/PuerkitoBio/goquery"
  "strings"
  "regexp"
  "code.google.com/p/go-charset/charset"
  "fmt"
  "io/ioutil"
  "log"
  "time"
  "github.com/franela/goreq"
)

import _ "code.google.com/p/go-charset/data"
var (
  QUALITY_REGEXP = regexp.MustCompile(`url(?P<quality>\d\d\d)=(?P<url>.*?mp4)`)
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
  ThumbUrl string
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
  doc, err := goquery.NewDocumentFromReader(get_html_reader(episode.Link))
  if err != nil {
    log.Print(err)
    return
  }
  vars, _ := doc.Find("object embed").First().Attr("flashvars")
  for _, r := range QUALITY_REGEXP.FindAllStringSubmatch(vars, -1) {
    episode.VideoLinks[r[1]+ "p"] = r[2]
  }
  thumb, _ := doc.Find("#player_thumb").First().Attr("src")
  episode.ThumbUrl = thumb
  channel <- true
}

func get_html_reader(link string) *strings.Reader {
  counter := 0
  var html string
  for ; counter < 5 ; {
    counter++
    resp, err := get_request(link).Do()
    if err == nil {
      html, _= resp.Body.ToString()
      break
    }
  }
  reader := strings.NewReader(html)
  return reader
}

func get_request(link string) *goreq.Request {
  request := goreq.Request{
    Uri: link,
    Timeout: 2 * time.Second,
  }
  return &request
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
