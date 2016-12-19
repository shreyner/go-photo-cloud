package routers

import (
	"gopkg.in/macaron.v1"
)

//func GetHome(ctx *macaron.Context) {
//	ctx.HTML(200, "home")
//}

func GetHome(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}