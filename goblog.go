package main

import (
   "net/http"
   "github.com/codegangsta/martini"
)

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  http.ListenAndServe("goblog.dev:8080", m)
}
