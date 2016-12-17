package main

import (
	"gopkg.in/macaron.v1"
	"httpmacaron/routers/upload"
	"httpmacaron/routers"
)

func main() {
	m := macaron.Classic()

	m.Use(macaron.Static("data/img"))
	m.Use(macaron.Static("public"))

	m.Use(macaron.Renderer())

	m.Get("/", routers.GetHome)
	m.Get("/u", upload.GetUpload)
	m.Post("/u", upload.PostUpload)

	m.Run()
}
