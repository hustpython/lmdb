// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"lmdbapi/controllers"
	"net/http"
)

var fileDirs = []string{"C", "D", "E", "F", "G", "H", "I", "J", "K"}

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"POST", "PUT", "GET", "DELETE"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/movie",
			beego.NSRouter("", &controllers.MovieController{}, "POST:RefreshMovies"),
			beego.NSRouter("", &controllers.MovieController{}, "PUT:UpdateMovie"),
			beego.NSRouter("", &controllers.MovieController{}, "GET:ShowMovies"),
		),
		beego.NSNamespace("/tag",
			beego.NSRouter("/?:Tag", &controllers.MovieController{}, "GET:GetMovieByTag"),
			beego.NSRouter("/:AllTag", &controllers.MovieController{}, "GET:GetAllTags"),
		),
		beego.NSNamespace("/coll",
			beego.NSRouter("/?:Coll", &controllers.MovieController{}, "GET:GetMoviesByColl"),
			beego.NSRouter("/:AllColl", &controllers.MovieController{}, "GET:GetAllColl"),
			beego.NSRouter("/?:MId", &controllers.MovieController{}, "DELETE:DelMovieColl"),
		),
		beego.NSNamespace("/recent",
			beego.NSRouter("", &controllers.MovieController{}, "GET:GetMoviesByRecent"),
		),

	)
	beego.AddNamespace(ns)

	for _, fileDir := range fileDirs {
		beego.Handler("/"+fileDir+"/*", http.StripPrefix("/"+fileDir+"/", http.FileServer(http.Dir(""+fileDir+":"))))
	}
}
