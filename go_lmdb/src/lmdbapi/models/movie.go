package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Movie struct {
	MId       string `orm:"pk"`
	LastWatch int64
	Duration  string
	VideoUrl  string
	Title     string
	Desc      string
	Cover     string
	Coll      *Coll `orm:"null;rel(fk)"`
	CollStr   string
	Tags      []*Tag   `orm:"rel(m2m)"`
	TagArray  []string `orm:"-"`
}

type Tag struct {
	TagName string   `orm:"pk"`
	Movies  []*Movie `orm:"reverse(many)"`
}

type Coll struct {
	CollName string   `orm:"pk"`
	Movies   []*Movie `orm:"reverse(many)"`
}

type RecentWatch struct {
	MId  string `orm:"pk"`
	Time string
}

type Filter struct {
	MinSize  int
	Include  string
	MovieExt []string
}

var ormOpr orm.Ormer

func init() {
	orm.RegisterModel(new(Movie), new(Coll), new(Tag))
	orm.RegisterDataBase("default", "sqlite3", "./lmdb.db")
	orm.RunSyncdb("default", false, true)
	ormOpr = orm.NewOrm()
}

func QueryAllMovieData() ([]*Movie, error) {
	var movies []*Movie
	_, err := ormOpr.QueryTable("movie").All(&movies)
	for _, m := range movies {
		ormOpr.LoadRelated(m, "Tags")
		for _, tag := range m.Tags {
			m.TagArray = append(m.TagArray, tag.TagName)
		}
		m.Tags = nil
	}
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

func (m Movie) UpdateMovieData() error {
	var updateCol []string
	if len(m.Cover) != 0 {
		updateCol = append(updateCol, "cover")
	}
	if len(m.Duration) != 0 {
		updateCol = append(updateCol, "duration")
	}
	if len(m.Title) != 0 {
		updateCol = append(updateCol, "title")
	}
	if len(m.Desc) != 0 {
		updateCol = append(updateCol, "desc")
	}
	if len(m.CollStr) != 0 {
		updateCol = append(updateCol, "coll_str")
		m.Coll = &Coll{
			CollName: m.CollStr,
		}
		updateCol = append(updateCol, "coll_id")
	}
	if len(updateCol) != 0 {
		_, err := ormOpr.Update(&m, updateCol...)
		if err != nil {
			return err
		}
	}
	if len(m.TagArray) != 0 {
		if err := m.UpdateTags(); err != nil {
			return err
		}
	}
	if len(m.CollStr) != 0 {
		return m.UpdateColl()
	}
	return nil
}

func (m Movie) UpdateColl() error {
	qs := ormOpr.QueryTable(new(Coll))
	if !qs.Filter("CollName", m.CollStr).Exist() {
		if _, e := ormOpr.Insert(&Coll{CollName: m.CollStr}); e != nil {
			return e
		}
	}
	return nil
}

func (m Movie) UpdateTags() error {
	qs := ormOpr.QueryTable(new(Tag))
	for _, tagStr := range m.TagArray {
		if !qs.Filter("TagName", tagStr).Exist() {
			if _, e := ormOpr.Insert(&Tag{TagName: tagStr}); e != nil {
				return e
			}
		}
	}
	newMovie := &Movie{MId: m.MId}
	mTag := ormOpr.QueryM2M(newMovie, "Tags")
	for _, tag := range m.TagArray {
		newTag := &Tag{TagName: tag}
		if !mTag.Exist(newTag) {
			mTag.Add(newTag)
		}
	}
	var tmpMovieTag = &Movie{MId: m.MId}
	ormOpr.LoadRelated(tmpMovieTag, "Tags")
	for _, tag1 := range tmpMovieTag.Tags {
		var find = false
		for _, tag2 := range m.TagArray {
			if tag1.TagName == tag2 {
				find = true
				continue
			}
		}
		if !find {
			fmt.Println("delete not used tag", tag1.TagName)
			mTag.Remove(tag1)
		}
	}
	return nil
}

func (t Tag) GetMoviesByTag() ([]*Movie, error) {
	_, err := ormOpr.LoadRelated(&t, "Movies")
	return t.Movies, err
}

func (c Coll) GetMoviesByColl() ([]*Movie, error) {
	_, err := ormOpr.LoadRelated(&c, "Movies")
	return c.Movies, err
}
