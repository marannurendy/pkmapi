// @APIVersion 1.0.0
// @Title PKM API
// @Description api untuk kebutuhan mobile pkm
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"pkmapi/controllers"
	"pkmapi/middleware"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(middleware.Jwt),
		beego.NSNamespace("/pkm",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
		// beego.NSNamespace("/pkmnew",
		// 	beego.NSInclude(
		// 		&controllers.PkmController{},
		// 	),
		// ),		
		beego.NSNamespace("/inisiasi",
			beego.NSInclude(
				&controllers.MasterDataController{},
			),
		),			
		beego.NSNamespace("/post_inisiasi",
			beego.NSInclude(
				&controllers.PostInisiasiController{},
			),
		),	
		beego.NSNamespace("/pkmb",
			beego.NSInclude(
				&controllers.PkmbController{},
			),
		),			
		beego.NSNamespace("/other",
			beego.NSInclude(
				&controllers.OtherController{},
			),
		),				
		beego.NSNamespace("/images",
			beego.NSInclude(
				&controllers.ImagesController{},
			),
		),				
	)
	beego.AddNamespace(ns)
}