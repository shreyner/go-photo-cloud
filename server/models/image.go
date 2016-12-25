package models

import (
	"github.com/jinzhu/gorm"
	"mime/multipart"
	"net/http"
	"photo-cloud/server/utils"
	"fmt"
	"os"
	"io"
)

type Image struct {
	gorm.Model
	Name		string	`gorm:"unique_index"`
	TypeImg		string
	Private		bool	`gorm:"default:'true'"`
	User		User
	UserID		uint
}

func (img *Image) GetFullName() string {
	return img.Name + "." + img.TypeImg
}

func (img *Image) generateNameImg(db *gorm.DB) {
	//Add len Name to config
	var name string

	for count := 1; count != 0; db.Model(&Image{}).Where("name = ?", name).Count(&count) {
		name = utils.RandomString(15)
	}

	img.Name = name
};

func (img *Image) getTypeImg(fileimg multipart.File) {
	// Get type image
	fileHeader := make([]byte, 512)
	fileimg.Read(fileHeader)
	fileimg.Seek(0,0)

	img.TypeImg, _ = utils.GetTypeImage(http.DetectContentType(fileHeader))
}

func SaveImage(fileImg multipart.File, db *gorm.DB) (Image, bool) {
	defer fileImg.Close()
	img := Image{}
	img.generateNameImg(db)
	img.getTypeImg(fileImg)

	if len(img.TypeImg) == 0 {
		return img, true
		fmt.Println("No type")
	}

	f, err := os.OpenFile("./data/img/" + img.GetFullName() , os.O_WRONLY | os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		panic("No create file!")
	}

	io.Copy(f, fileImg)
	db.NewRecord(&img)
	db.Create(&img)

	return img, false

}