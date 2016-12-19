package models

import (
	"github.com/jinzhu/gorm"
	"mime/multipart"
	"os"
	"io"
	"net/http"
	"photo-cloud/utils"
)

type Images struct {
	gorm.Model
	Name		string	`gorm:"unique_index"`
	TypeImg		string
}

func UploadImage(fileImg multipart.File, fileHandler *multipart.FileHeader, db *gorm.DB) (string, bool) {
	defer fileImg.Close()

	// Get type image
	fileHeader := make([]byte, 512)
	fileImg.Read(fileHeader)
	fileImg.Seek(0,0)

	// Generate name
	fileName := utils.RandomString(25)
	fileType := utils.GetTypeImage(http.DetectContentType(fileHeader))
	imgD := Images{
		Name: fileName,
		//Name: "bbb",
		TypeImg: fileType,
	}

	fullnameFile := imgD.Name + "." + imgD.TypeImg


	f, err := os.OpenFile("./data/img/" + fullnameFile , os.O_WRONLY | os.O_CREATE, 0666)
	defer f.Close()

	if err != nil{
		panic("Don't Save!")
		return "", false
	}

	db.Create(&imgD)

	io.Copy(f, fileImg)

	return fullnameFile, true;
}