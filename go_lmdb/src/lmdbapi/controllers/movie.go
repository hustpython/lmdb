package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lmdbapi/models"
)

type MovieController struct {
	ErrorController
}

// GET /v1/movie?id=movie_12121
func (m *MovieController) ShowMovies() {
	uid := m.GetString("id")
	if len(uid) == 0 {
		m.Data["json"] = models.GetAllMovies()
	} else {
		uu, err := models.GetMovieByID(uid)
		if err == nil {
			m.Data["json"] = uu
		} else {
			m.setNotFound(err.Error())
		}
	}
	m.ServeJSON()
}

// POST /v1/movie {
//        "VideoUrl": "D://2.mp4",
//    }
func (m *MovieController) AddMovie() {
	var movie models.Movie
	json.Unmarshal(m.Ctx.Input.RequestBody, &movie)
	uid := models.AddMovie(movie)
	m.Data["json"] = map[string]string{"uid": uid}
	m.ServeJSON()
}

// DELETE /v1/movie?id=movie_12121
func (m *MovieController) DelMovie() {
	models.DelMovie(m.GetString("id"))
	m.Ctx.Output.SetStatus(http.StatusNoContent)
}

// PUT /v1/movie?id=movie_12121 {
//        "VideoUrl": "D://4.mp4",
//        "Cover": "D:/3.png"
//    }
func (m *MovieController) UpdateMovie() {
	uid := m.GetString("id")
	if len(uid) == 0 {
		m.setParamInvalid("uid is invalid")
		return
	}
	var movie models.Movie
	json.Unmarshal(m.Ctx.Input.RequestBody, &movie)
	uu, err := models.UpdateMovie(uid, &movie)
	if err != nil {
		m.setNotFound(fmt.Sprintf("not found updateID %s", uid))
	} else {
		m.Data["json"] = uu
	}
	m.ServeJSON()
}
