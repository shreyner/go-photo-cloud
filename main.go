package main

import (
	"github.com/shreyner/photo-cloud/routers"
	"github.com/shreyner/photo-cloud/routers/upload"

	"gopkg.in/macaron.v1"
)

func main() {
	//For example
	//user := new(models.User)
	//user.SetLogin("Alexander")
	//fmt.Println(user.GetLogin())

	m := macaron.Classic()

	m.Use(macaron.Static("data/img"))
	m.Use(macaron.Renderer())

	m.Get("/", routers.GetHome)
	m.Post("/u", upload.PostUpload)

	m.Run()
}
