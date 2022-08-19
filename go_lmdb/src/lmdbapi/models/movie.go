package models

import (
	"errors"
	"fmt"
	"lmdbapi/util"
	"strconv"
)

type Movie struct {
	VideoUrl string
	Title    string
}

var (
	MovieMap   map[string]*Movie
	FilterData Filter
)

func init() {
	MovieMap = make(map[string]*Movie)
	util.InitViper()
	util.MovieFilterViper.UnmarshalKey(movieFilterJsonPrefix, &FilterData)
	util.MovieDataViper.UnmarshalKey(movieConfigJsonPrefix, &MovieMap)
}

func GetMovieByID(uid string) (*Movie, error) {
	uu, ok := MovieMap[uid]
	if ok {
		return uu, nil
	} else {
		return &Movie{}, fmt.Errorf("not found movie %s", uid)
	}
}

func GetAllMovies() []*Movie {
	// TODO 根据条件返回
	var res []*Movie
	for i := range MovieMap {
		res = append(res, MovieMap[i])
	}
	return res
}

func AddMovie(m Movie) string {
	id := strconv.Itoa(len(MovieMap))
	MovieMap[id] = &m
	updateMovieConfig()
	return id
}

func DelMovie(uid string) {
	delete(MovieMap, uid)
	updateMovieConfig()
}

func UpdateMovie(uid string, uu *Movie) (*Movie, error) {
	if u, ok := MovieMap[uid]; ok {
		if uu.VideoUrl != "" {
			u.VideoUrl = uu.VideoUrl
		}
		updateMovieConfig()
		return u, nil
	}
	return nil, errors.New("movie Not exist")
}
