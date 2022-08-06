package controllers

import (
	"encoding/json"
	"lmdbapi/models"
)

type FilterController struct {
	ErrorController
}

// /v1/filter { POST
//    "MinSize" :400,
//	"MovieExt" :".mp4|.mkv|.rmvb"
//}
func (f *FilterController) PostFilter() {
	var searchRule models.Filter
	json.Unmarshal(f.Ctx.Input.RequestBody, &searchRule)
	if searchRule.MinSize < 100 || len(searchRule.MovieExt) == 0 {
		f.setParamInvalid("minsize or movieext invalid")
		return
	}
	if err := models.SearchMovie(searchRule); err != nil {
		f.setInternalError(err.Error())
	}
	f.Data["json"] = models.GetAllMovies()
	f.ServeJSON()
}

// /v1/filter GET

func (f *FilterController) GetFilter() {
	f.Data["json"] = models.GetFilter()
	f.ServeJSON()
}
