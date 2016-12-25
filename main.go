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

	//Db Test
	//	profile := models.Profile{}
	//	profile.Name = "Alexander"
	//	db.Create(&profile)
	//
	//	user := models.User{}
	//	user.Login = "AlexanderLogin4"
	//	user.Password = "123"
	//	user.Email = "alexander@shreyner.ru5"
	//	user.Profile = profile
	//	db.Create(&user)
	//


	m := macaron.Classic()

	m.Use(macaron.Static("data/img"))
	m.Use(macaron.Renderer())
	m.Map(db)

	m.Get("/", routers.GetHome)
	m.Get("/u", upload.GetUpload)
	m.Post("/u", upload.PostUpload)

	m.Run()
}
