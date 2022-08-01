package models

import (
	"errors"
	"fmt"
	"strconv"
)

type Movie struct {
	Path  string
	Cover string
}

var (
	MovieList map[string]*Movie
)

func init() {
	MovieList = make(map[string]*Movie)
	MovieList["0"] = &Movie{"D://1.mp4", "D:/1.png"}
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
	return id
}

func DelMovie(uid string) {
	delete(MovieList, uid)
}

func UpdateMovie(uid string, uu *Movie) (*Movie, error) {
	if u, ok := MovieList[uid]; ok {
		if u.Cover != "" {
			u.Cover = uu.Cover
		}
		if uu.Path != "" {
			u.Path = uu.Path
		}
		return u, nil
	}
	return nil, errors.New("movie Not exist")
}
