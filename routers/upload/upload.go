package upload

import (
	"gopkg.in/macaron.v1"
	"photo-cloud/models"
)

func GetUpload(ctx *macaron.Context) {
	ctx.Redirect("/")
}

func PostUpload(ctx *macaron.Context) {
	fileimg, fileHandler, error := ctx.Req.FormFile("img")
	defer fileimg.Close();

	if error != nil {
		ctx.Error(403, "Error upload file")
	} else {

		if link, err := models.UploadImage(fileimg, fileHandler); err == false{
			ctx.Error(500, "Error server")
		} else {
			ctx.Redirect("/" + link)
		}

	}


}