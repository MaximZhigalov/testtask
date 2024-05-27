// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"testz/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	usersNS := beego.NewNamespace("/users",
		
		beego.NSRouter("/index", &controllers.UsersController{}, "get:Index"),
		beego.NSRouter("/create", &controllers.UsersController{}, "get,post:Create"),
		
		beego.NSRouter("/delete", &controllers.UsersController{}, "get:Delete"),
	)
	beego.AddNamespace(usersNS)
}


