package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type ErrorController struct {
	beego.Controller
}

func (m *ErrorController) setParamInvalid(errorMsg string) {
	m.Ctx.Output.SetStatus(http.StatusBadRequest)
	m.Data["json"] = errorMsg
	m.ServeJSON()
}

func (m *ErrorController) setNotFound(errorMsg string) {
	m.Ctx.Output.SetStatus(http.StatusNotFound)
	m.Data["json"] = errorMsg
	m.ServeJSON()
}

func (m *ErrorController) setInternalError(errorMsg string) {
	m.Ctx.Output.SetStatus(http.StatusInternalServerError)
	m.Data["json"] = errorMsg
	m.ServeJSON()
}
