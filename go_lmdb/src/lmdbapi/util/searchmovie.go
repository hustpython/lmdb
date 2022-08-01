package util

import (
	"github.com/jan-bar/es"
	"lmdbapi/models"
	"path/filepath"
	"strings"
)

const (
	maxMovieNum      = 1000
	minMovieSizeBase = 1000000 // M
)

func SearchMovie(rule models.Rule) []string {
	err := es.EverythingSetSearch(rule.MovieExt)
	if err != nil {
		return nil
	}
	err = es.EverythingSetMax(maxMovieNum)
	if err != nil {
		return nil
	}
	// 设置好需要查询的内容,不然后续遍历时可能报错
	err = es.EverythingSetRequestFlags(es.EverythingRequestFileName | es.EverythingRequestPath |
		es.EverythingRequestDateCreated | es.EverythingRequestDateModified | es.EverythingRequestDateAccessed |
		es.EverythingRequestSize)
	if err != nil {
		return nil
	}
	// 定好排序规则
	err = es.EverythingSetSort(es.EverythingSortDateModifiedAscending)
	if err != nil {
		return nil
	}
	// 开始查询
	es.EverythingQuery(true)
	// 得到查询结果个数
	return procQueryData(rule)
}

func procQueryData(rule models.Rule) []string {
	num, err := es.EverythingGetNumResults()
	if err != nil {
		return nil
	}
	minMovieSize := rule.MinSize * minMovieSizeBase
	var res []string
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
		if strings.Contains(rule.MovieExt, base) {
			res = append(res, p)
		}
	}
	return res
}
