package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Movie struct {
	MId       string `orm:"pk"`
	LastWatch int64
	Duration  int64
	VideoUrl  string
	Title     string
	Desc      string
	Actors    string
	Nation    string
	Cover     string
}

type Filter struct {
	MinSize  int
	Include  string
	MovieExt []string
}

var ormOpr orm.Ormer

func init() {
	orm.RegisterModel(new(Movie))
	orm.RegisterDataBase("default", "sqlite3", "./lmdb.db")
	orm.RunSyncdb("default", false, true)
	ormOpr = orm.NewOrm()
}

func QueryAllMovieData() ([]*Movie, error) {
	var movies []*Movie
	_, err := ormOpr.QueryTable("movie").All(&movies)
	return movies, err
}

func InsertMovieData(movieArray []*Movie) ([]*Movie, error) {
	qs := ormOpr.QueryTable(new(Movie))
	for _, m := range movieArray {
		if !qs.Filter("MId", m.MId).Exist() {
			if _, e := ormOpr.Insert(m); e != nil {
				return nil, e
			}
		}
	}
	return movieArray, nil
}

func UpdateMovieCover(uid, cover string) error {
	temMovieDb := &Movie{MId: uid, Cover: cover}
	_, err := ormOpr.Update(temMovieDb, "cover")
	return err
}
