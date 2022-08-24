package models

import (
	"github.com/jan-bar/es"
	"lmdbapi/util"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	movieConfigJsonPrefix = "moviepath"
	movieFilterJsonPrefix = "filter"
)

type Filter struct {
	MinSize  int
	Include  string
	MovieExt []string
}

const (
	maxMovieNum      = 1000
	minMovieSizeBase = 1000000 // M
)

func GetFilter() Filter {
	return FilterData
}

// searchExt ".mp4|.mkv|.rmvb"
func SearchMovie(rule Filter) error {
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
		return err
	}
	err = es.EverythingSetMax(maxMovieNum)
	if err != nil {
		return err
	}
	// 设置好需要查询的内容,不然后续遍历时可能报错
	err = es.EverythingSetRequestFlags(es.EverythingRequestFileName | es.EverythingRequestPath |
		es.EverythingRequestDateCreated | es.EverythingRequestDateModified | es.EverythingRequestDateAccessed |
		es.EverythingRequestSize)
	if err != nil {
		return err
	}
	// 定好排序规则
	err = es.EverythingSetSort(es.EverythingSortDateModifiedAscending)
	if err != nil {
		return err
	}
	// 开始查询
	es.EverythingQuery(true)
	// 得到查询结果个数
	return procQueryData(rule, searchExt)
}

func procQueryData(rule Filter, searchExt string) error {
	MovieMap = make(map[string]*Movie, 0)
	num, err := es.EverythingGetNumResults()
	if err != nil {
		return err
	}
	minMovieSize := rule.MinSize * minMovieSizeBase
	var id = 0
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
		base := filepath.Ext(p)
		id++
		if strings.Contains(searchExt, base) &&
			strings.Contains(p, rule.Include) {
			p = strings.Replace(p, "\\", "/", -1)
			p = strings.Replace(p, ":", "", -1)
			MovieMap[strconv.Itoa(id)] = &Movie{VideoUrl: p, Title: strings.Trim(filepath.Base(p), base)}
		}
	}
	FilterData = rule
	updateMovieConfig()
	updateFilterConfig()
	return nil
}

func updateMovieConfig() {
	util.MovieDataViper.Set(movieConfigJsonPrefix, MovieMap)
	util.MovieDataViper.WriteConfig()
}

func updateFilterConfig() {
	util.MovieFilterViper.Set(movieFilterJsonPrefix, FilterData)
	util.MovieFilterViper.WriteConfig()
}
