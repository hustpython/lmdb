// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"lmdbapi/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/movie",
			beego.NSRouter("", &controllers.MovieController{}, "GET:ShowMovies"),
			beego.NSRouter("", &controllers.MovieController{}, "POST:AddMovie"),
			beego.NSRouter("", &controllers.MovieController{}, "DELETE:DelMovie"),
			beego.NSRouter("", &controllers.MovieController{}, "PUT:UpdateMovie"),
		),
		beego.NSNamespace("/filter",
			beego.NSRouter("", &controllers.FilterController{}, "GET:GetFilter"),
			beego.NSRouter("", &controllers.FilterController{}, "POST:PostFilter"),
		),
	)
	beego.AddNamespace(ns)
}
