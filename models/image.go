package models

import (
	"mime/multipart"
	"os"
	"fmt"
	"io"
	"net/http"
	"photo-cloud/utils"
)

func UploadImage(fileImg multipart.File, fileHandler *multipart.FileHeader) (string, bool) {
	defer fileImg.Close()


	fileHeader := make([]byte, 512)
	fileImg.Read(fileHeader)
	fileImg.Seek(0,0)

	fileName := utils.RandomString(25)
	fileType := utils.GetTypeImage(http.DetectContentType(fileHeader))

	fullnameFile := fileName + "." + fileType
	f, err := os.OpenFile("./data/img/" + fullnameFile , os.O_WRONLY | os.O_CREATE, 0666)
	defer f.Close()

	if err != nil{
		return "", false
	}

	fmt.Println("Create img: " + fullnameFile)
	io.Copy(f, fileImg)

	return fullnameFile, true;
}