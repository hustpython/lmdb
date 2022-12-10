package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"io"
	"lmdbapi/models"
	"strconv"

	"net/http"
)

type TMdbController struct {
	ErrorController
}

type GenresStr struct {
	Name string `json:"name"`
}

type TMdbResult struct {
	Id               int          `json:"id"`
	Backdrop_path    string       `json:"backdrop_path"`
	Title            string       `json:"title"`
	Name             string       `json:"name"`
	Overview         string       `json:"overview"`
	First_air_date   string       `json:"first_air_date"`
	Poster_path      string       `json:"poster_path"`
	Release_date     string       `json:"release_date"`
	Media_type       string       `json:"media_type"`
	Vote_average     float32      `json:"vote_average"`
	Runtime          float32      `json:"runtime"`
	Episode_run_time []float32    `json:"episode_run_time"`
	Genres           []*GenresStr `json:"genres"`
}

type TMdbMultiSearchRsp struct {
	Results []*TMdbResult `json:"results"`
}

const (
	baseTmdbApi    = "https://api.themoviedb.org/3/search/multi"
	baseMovieIDApi = "https://api.themoviedb.org/3/movie/"
	baseTvIDApi    = "https://api.themoviedb.org/3/tv/"
	imgBaseApi     = "https://image.tmdb.org/t/p/w300"
	tmdbKey        = "572d4329ffbfc915eff9d63046a7b42b"
	mediaMovie     = "movie"
)

func (t *TMdbController) GetTMdbDataByKeyWord() {
	searchKey := t.GetString("searchKey")
	if len(searchKey) == 0 {
		t.setParamInvalid("invalid searchKey")
		return
	}
	url := fmt.Sprintf("%s?api_key=%s&query=%s&language=zh-CN", baseTmdbApi,
		tmdbKey, searchKey)
	res, err := http.Get(url)
	if err != nil {
		t.setInternalError(err.Error())
		return
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.setInternalError(err.Error())
		return
	}
	var tmdbRes TMdbMultiSearchRsp
	err = json.Unmarshal(data, &tmdbRes)
	if err != nil {
		t.setInternalError(err.Error())
		return
	}
	var res11 []*models.Movie
	for i := range tmdbRes.Results {
		tmpRes := getTmdbVideoInfoByID(tmdbRes.Results[i].Id, tmdbRes.Results[i].Media_type)
		if tmpRes != nil {
			res11 = append(res11, tmpRes)
		}
	}
	t.Data["json"] = res11
	t.ServeJSON()
}

func getTmdbVideoInfoByID(id int, mediaType string) *models.Movie {
	var searchUrl string
	if mediaType == mediaMovie {
		searchUrl = fmt.Sprintf("%s%d?api_key=%s&language=zh-CN",
			baseMovieIDApi, id, tmdbKey)
	} else {
		searchUrl = fmt.Sprintf("%s%d?api_key=%s&language=zh-CN",
			baseTvIDApi, id, tmdbKey)
	}
	res, err := http.Get(searchUrl)
	if err != nil {
		logs.Error(err)
		return nil
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		logs.Error(err)
		return nil
	}
	var tmdbRes TMdbResult
	err = json.Unmarshal(data, &tmdbRes)
	if err != nil {
		logs.Error(err)
		return nil
	}
	var tmpTitle, tmpReleaseDate string
	var tmpRuntime float32
	if mediaType == mediaMovie {
		tmpTitle = tmdbRes.Title
		tmpReleaseDate = tmdbRes.Release_date
		tmpRuntime = tmdbRes.Runtime
	} else {
		tmpTitle = tmdbRes.Name
		tmpReleaseDate = tmdbRes.First_air_date
		if len(tmdbRes.Episode_run_time) > 0 {
			tmpRuntime = tmdbRes.Episode_run_time[0]
		}
	}
	var tmpTagArr []string
	for _, r := range tmdbRes.Genres {
		tmpTagArr = append(tmpTagArr, r.Name)
	}
	tmpVote, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", tmdbRes.Vote_average), 32)
	return &models.Movie{
		ReleaseDate: tmpReleaseDate,
		Poster:      transImagePathToData64(tmdbRes.Poster_path),
		Vote:        float32(tmpVote),
		Duration:    transRuntime2Duration(tmpRuntime),
		Title:       tmpTitle,
		Cover:       transImagePathToData64(tmdbRes.Backdrop_path),
		TagArray:    tmpTagArr,
		Desc:        tmdbRes.Overview,
	}
}

func transRuntime2Duration(t float32) string {
	intT := int(t)
	if intT == 0 {
		return ""
	}
	h := intT / 60
	m := intT % 60
	return fmt.Sprintf("%02d:%02d:00", h, m)
}
func transImagePathToData64(imgPath string) string {
	res, err := http.Get(imgBaseApi + imgPath)
	if err != nil {
		logs.Error(err)
		return ""
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		logs.Error(err)
		return ""
	}
	base64Data := base64.StdEncoding.EncodeToString(data)
	mimeType := http.DetectContentType(data)
	switch mimeType {
	case "image/jpeg":
		base64Data = "data:image/jpeg;base64," + base64Data
	case "image/png":
		base64Data = "data:image/png;base64," + base64Data
	}
	return base64Data
}
