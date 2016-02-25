package main

import (
  "net/http"
  "io/ioutil"
  "log"
  "strings"
)

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request){
  path := req.URL.Path[1:]
  log.Println(path)

  data, err := ioutil.ReadFile(string(path))

  if err == nil {
    var contentType string

    if strings.HasSuffix(path, ".css") {
      contentType = "text/css"
    } else if strings.HasSuffix(path, ".html") {
      contentType = "text/html"
    } else if strings.HasSuffix(path, ".js") {
      contentType = "application/javascript"
    } else if strings.HasSuffix(path, ".png") {
      contentType = "image/png"
    } else if strings.HasSuffix(path, ".jpg") {
      contentType = "image/jpg"
    } else {
      contentType = "text/plain"
    }

    rw.Header().Set("Content-Type", contentType)
    rw.Write(data)
  } else {
    rw.WriteHeader(404)
    rw.Write([]byte("404 My Friend - " + http.StatusText(404)))
  }

}

func main() {
  http.Handle("/", new(MyHandler))
  http.ListenAndServe(":8080", nil)
}
