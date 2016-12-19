package main

import (
	"photo-cloud/routers"
	"photo-cloud/routers/upload"
	"photo-cloud/models"

	"gopkg.in/macaron.v1"
)

func main() {
	//For example
	//user := new(models.User)
	//user.SetLogin("Alexander")
	//fmt.Println(user.GetLogin())

	dbConfig := models.DBconfig{
		ConnectString: "goimg:123@/goimg",
	}

	db := dbConfig.Connect()
	defer db.Close();

	m := macaron.Classic()

	m.Use(macaron.Static("data/img"))
	m.Use(macaron.Renderer())
	m.Map(db)

	m.Get("/", routers.GetHome)
	m.Post("/u", upload.PostUpload)

	m.Run()
}
