package app

import (
  "net/http"
  "github.com/codegangsta/martini"
)

type App struct {
	name string
}

func (a App) Run() {
	m := martini.Classic()
	m.Get("/", func() string {
   	return "Hello world234778!"
})

http.ListenAndServe("goblog.dev:8080", m)
}
