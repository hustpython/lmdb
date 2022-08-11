package main

import (
	"net/http"

	"github.com/astaxie/beego"
	_ "lmdbapi/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	go httpServerHandle()
	beego.Run()
}

func httpServerHandle() {
	mux := http.NewServeMux()
	mux.Handle("/C/", http.StripPrefix("/C/", http.FileServer(http.Dir("C:"))))
	mux.Handle("/D/", http.StripPrefix("/D/", http.FileServer(http.Dir("D:"))))
	mux.Handle("/E/", http.StripPrefix("/E/", http.FileServer(http.Dir("E:"))))
	mux.Handle("/F/", http.StripPrefix("/F/", http.FileServer(http.Dir("F:"))))
	mux.Handle("/G/", http.StripPrefix("/G/", http.FileServer(http.Dir("G:"))))
	mux.Handle("/H/", http.StripPrefix("/H/", http.FileServer(http.Dir("H:"))))
	mux.Handle("/I/", http.StripPrefix("/I/", http.FileServer(http.Dir("I:"))))
	mux.Handle("/J/", http.StripPrefix("/J/", http.FileServer(http.Dir("J:"))))
	mux.Handle("/K/", http.StripPrefix("/K/", http.FileServer(http.Dir("K:"))))
	http.ListenAndServe(":9091", mux)
}
