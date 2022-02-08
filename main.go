package main

import (
	_ "pkmapi/routers"
	"github.com/astaxie/beego"
	"pkmapi/global"	
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	global.MainLog()

	// if beego.BConfig.RunMode == "dev" {
		// beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/"] = "swagger"
	// }
	// beego.BConfig.WebConfig.StaticDir["/images"] = "images"


	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowAllOrigins: true,
	}))	

	beego.Run()
}
