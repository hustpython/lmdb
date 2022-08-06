package models

import (
	"errors"
	"fmt"
	"lmdbapi/util"
	"strconv"
)

type Movie struct {
	Path string
}

var (
	MovieList  map[string]*Movie
	FilterData Filter
)

func init() {
	MovieList = make(map[string]*Movie)
	util.InitViper()
	util.MovieFilterViper.UnmarshalKey(movieFilterJsonPrefix, &FilterData)
	util.MovieDataViper.UnmarshalKey(movieConfigJsonPrefix, &MovieList)
}

func GetMovieByID(uid string) (*Movie, error) {
	uu, ok := MovieList[uid]
	if ok {
		return uu, nil
	} else {
		return &Movie{}, fmt.Errorf("not found movie %s", uid)
	}
}

func GetAllMovies() map[string]*Movie {
	return MovieList
}

func AddMovie(m Movie) string {
	id := strconv.Itoa(len(MovieList))
	MovieList[id] = &m
	updateMovieConfig()
	return id
}

func DelMovie(uid string) {
	delete(MovieList, uid)
	updateMovieConfig()
}

func UpdateMovie(uid string, uu *Movie) (*Movie, error) {
	if u, ok := MovieList[uid]; ok {
		if uu.Path != "" {
			u.Path = uu.Path
		}
		updateMovieConfig()
		return u, nil
	}
	return nil, errors.New("movie Not exist")
}
