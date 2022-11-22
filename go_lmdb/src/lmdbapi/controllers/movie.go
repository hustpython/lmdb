package controllers

import (
	"github.com/jan-bar/es"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"lmdbapi/models"
)

const (
	maxMovieNum      = 1000
	minMovieSizeBase = 1000000 // M
)

type MovieController struct {
	ErrorController
}

func init() {
	go loopClearInvalidData()
}

// searchExt ".mp4|.mkv|.rmvb"
func searchMovie(rule models.Filter) ([]*models.Movie, error) {
	searchExt := ""
	for i := range rule.MovieExt {
		searchExt += "."
		searchExt += rule.MovieExt[i]
		if i != len(rule.MovieExt)-1 {
			searchExt += "|"
		}
	}
	err := es.EverythingSetSearch(searchExt)
	if err != nil {
		return nil, err
	}
	err = es.EverythingSetMax(maxMovieNum)
	if err != nil {
		return nil, err
	}
	// 设置好需要查询的内容,不然后续遍历时可能报错
	err = es.EverythingSetRequestFlags(es.EverythingRequestFileName | es.EverythingRequestPath |
		es.EverythingRequestDateCreated | es.EverythingRequestDateModified | es.EverythingRequestDateAccessed |
		es.EverythingRequestSize)
	if err != nil {
		return nil, err
	}
	// 定好排序规则
	err = es.EverythingSetSort(es.EverythingSortDateModifiedAscending)
	if err != nil {
		return nil, err
	}
	// 开始查询
	es.EverythingQuery(true)
	// 得到查询结果个数
	return procQueryData(rule, searchExt)
}

func procQueryData(rule models.Filter, searchExt string) ([]*models.Movie, error) {
	var movieArray []*models.Movie
	num, err := es.EverythingGetNumResults()
	if err != nil {
		return nil, err
	}
	minMovieSize := rule.MinSize * minMovieSizeBase
	for i := uint32(0); i < num; i++ {
		s, err := es.EverythingGetResultSize(i)
		if err != nil {
			continue
		}
		if s < int64(minMovieSize) {
			continue
		}
		p, err := es.EverythingGetResultFullPathName(i)
		if err != nil {
			continue
		}
		k, err := es.EverythingGetResultDateCreated(i)
		if err != nil {
			continue
		}
		base := filepath.Ext(p)
		if strings.Contains(searchExt, base) &&
			strings.Contains(p, rule.Include) {
			p = strings.Replace(p, "\\", "/", -1)
			p = strings.Replace(p, ":", "", -1)
			tmpMovie := &models.Movie{
				MId:      strconv.Itoa(k.Second()) + strconv.Itoa(int(s)),
				Title:    strings.TrimSuffix(filepath.Base(p), base),
				VideoUrl: p,
				Size:     uint64(s),
			}
			movieArray = append(movieArray, tmpMovie)
		}
	}
	return models.InsertOrUpdateMovieData(rule.Force, movieArray)
}

// GET /v1/movie
func (m *MovieController) ShowMovies() {
	d, e := models.QueryAllMovieData()
	if e != nil {
		m.setInternalError(e.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

// /v1/filter { POST
//    "MinSize" :400,
//	"MovieExt" :["mp4","rmvb","avi,"mkv"]
//}
func (m *MovieController) RefreshMovies() {
	var searchRule models.Filter
	err := m.BindJSON(&searchRule)
	if err != nil {
		m.setInternalError(err.Error())
	}
	if searchRule.MinSize < 0 || len(searchRule.MovieExt) == 0 {
		m.setParamInvalid("minsize or movieext invalid")
		return
	}
	d, err := searchMovie(searchRule)
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

type PutCover struct {
	MId   string
	Cover string
}

func (m *MovieController) UpdateMovie() {
	var c = &models.Movie{}
	err := m.BindJSON(c)
	if err != nil {
		m.setInternalError(err.Error())
		return
	}
	err = c.UpdateMovieData()
	if err != nil {
		m.setInternalError(err.Error())
	}
}

func (m *MovieController) GetMovieByTag() {
	tag := m.GetString("Tag")
	if len(tag) == 0 {
		m.setParamInvalid("Tag is null")
		return
	}
	var mTag = models.Tag{TagName: tag}
	d, err := mTag.GetMoviesByTag()
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) GetMoviesByColl() {
	coll := m.GetString("Coll")
	if len(coll) == 0 {
		m.setParamInvalid("Coll is null")
		return
	}
	var collModel = models.Coll{CollName: coll}
	d, err := collModel.GetMoviesByColl()
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) GetAllTags() {
	d, err := models.GetAllTags()
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) GetAllColl() {
	d, err := models.GetAllColl()
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) DelMovieColl() {
	var tmpM = models.Movie{MId: m.GetString("MId")}
	err := tmpM.DeleteColl()
	if err != nil {
		m.setInternalError(err.Error())
	}
	m.ServeJSON()
}

func (m *MovieController) GetMoviesByRecent() {
	d, err := models.GetMovieByRecent()
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) GetCommentsByMId() {
	mid := m.GetString("MId")
	if len(mid) == 0 {
		m.setParamInvalid("MId is null")
		return
	}
	var movieModel = models.Movie{MId: mid}
	d, err := movieModel.GetCommentsByMId()
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) GetCommentsByColl() {
	coll := m.GetString("Coll")
	if len(coll) == 0 {
		m.setParamInvalid("Coll is null")
		return
	}
	var movieModel = models.Coll{CollName: coll}
	d, err := movieModel.GetCommentsByColl()
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) AddCommentByMId() {
	var comment models.Comment
	err := m.BindJSON(&comment)
	if err != nil || len(comment.MId) == 0 {
		m.setParamInvalid("MId is null")
		return
	}
	comment.CommentID = comment.MId + strconv.Itoa(int(comment.CommentFrame))
	comment.Movie = &models.Movie{
		MId: comment.MId,
	}
	err = comment.AddCommentByMId()
	if err != nil {
		m.setInternalError(err.Error())
	}
}

func (m *MovieController) DelCommentByCId() {
	var comment models.Comment
	err := m.BindJSON(&comment)
	if err != nil || len(comment.MId) == 0 {
		m.setParamInvalid("MId is null")
		return
	}
	comment.CommentID = comment.MId + strconv.Itoa(int(comment.CommentFrame))
	err = comment.DelCommentByMId()
	if err != nil {
		m.setInternalError(err.Error())
	}
}

func (m *MovieController) GetMoviesAndColl() {
	d, e := models.QuerySortMovie()
	if e != nil {
		m.setInternalError(e.Error())
	} else {
		m.Data["json"] = d
	}
	m.ServeJSON()
}

func (m *MovieController) BatchAddColl() {
	var c models.Coll
	err := m.BindJSON(&c)
	if err != nil {
		m.setInternalError(err.Error())
	} else {
		c.BatchAddColl()
	}
}

func (m *MovieController) GetMoviesByFavourite() {
	d, err := models.GetMovieByFavourite()
	if err != nil {
		m.setInternalError(err.Error())
		return
	}
	m.Data["json"] = d
	m.ServeJSON()
}

func loopClearInvalidData() {
	for {
		time.Sleep(time.Hour)
		models.ClearInvalidData()
	}
}
