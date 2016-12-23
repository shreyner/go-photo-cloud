package main

import (
	"photo-cloud/routers"
	"photo-cloud/routers/upload"
	"photo-cloud/models"
	"photo-cloud/module/config"

	"gopkg.in/macaron.v1"
)

func main() {

	// Settings Config
	conf := config.Get()

	// Settings DataBase
	dbConfig := models.DBconfig{
		ConnectString: conf.DbSettings.User +
			":" + conf.DbSettings.Password + "@/" +
			conf.DbSettings.DbName,
	}

	db := dbConfig.Connect()
	defer db.Close();
	// end

	m := macaron.Classic()

	m.Use(macaron.Static("data/img"))
	m.Use(macaron.Renderer())
	m.Map(db)

	m.Get("/", routers.GetHome)
	m.Get("/u", upload.GetUpload)
	m.Post("/u", upload.PostUpload)

	m.Run()
}
