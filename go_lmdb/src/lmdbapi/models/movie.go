package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"sort"
)

type Movie struct {
	MId         string `orm:"pk"`
	LastWatch   int64
	RecentWatch int64
	Duration    string
	VideoUrl    string
	Title       string
	Desc        string
	Cover       string
	Coll        *Coll `orm:"null;rel(fk)"`
	CollStr     string
	Tags        []*Tag   `orm:"rel(m2m)"`
	TagArray    []string `orm:"-"`
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
	Force    bool
	MovieExt []string
}

var ormOpr orm.Ormer

func init() {
	logs.SetLogger(logs.AdapterConsole)
	orm.RegisterModel(new(Movie), new(Coll), new(Tag))
	orm.RegisterDataBase("default", "sqlite3", "./lmdb.db")
	orm.RunSyncdb("default", false, true)
	ormOpr = orm.NewOrm()
	ClearInvalidData()
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

func InsertOrUpdateMovieData(force bool, movieArray []*Movie) ([]*Movie, error) {
	// TODO  if force, delete all tables
	qs := ormOpr.QueryTable(new(Movie))
	var res []*Movie
	for _, m := range movieArray {
		if !qs.Filter("MId", m.MId).Exist() {
			if _, e := ormOpr.Insert(m); e != nil {
				return nil, e
			}
			res = append(res, m)
		} else {
			var movie Movie
			if qs.Filter("MId", m.MId).One(&movie) == nil {
				if movie.VideoUrl != m.VideoUrl {
					_, err := ormOpr.Update(m, "video_url")
					fmt.Printf("update videurl: %s,err:%s", m.VideoUrl, err)
				}
			}
		}
	}
	return res, nil
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
	if m.RecentWatch != 0 {
		updateCol = append(updateCol, "recent_watch")
	}
	if m.LastWatch != 0 {
		updateCol = append(updateCol, "last_watch")
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
				return nil
			}
		}
		if !find {
			fmt.Println("delete not used tag in rel table", tag1.TagName)
			mTag.Remove(tag1)
			ormOpr.LoadRelated(tag1, "Movies")
			if len(tag1.Movies) == 0 {
				fmt.Println("delete tag table", tag1.TagName)
				ormOpr.Delete(tag1)
			}
		}
	}
	return nil
}

func (t Tag) GetMoviesByTag() ([]*Movie, error) {
	_, err := ormOpr.LoadRelated(&t, "Movies")
	for _, m := range t.Movies {
		ormOpr.LoadRelated(m, "Tags")
		for _, tag := range m.Tags {
			m.TagArray = append(m.TagArray, tag.TagName)
		}
		m.Tags = nil
	}
	return t.Movies, err
}

func GetAllTags() ([]string, error) {
	var tags []*Tag
	_, err := ormOpr.QueryTable("tag").All(&tags)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, t := range tags {
		res = append(res, t.TagName)
	}
	return res, nil
}

func (c Coll) GetMoviesByColl() ([]*Movie, error) {
	_, err := ormOpr.LoadRelated(&c, "Movies")
	for _, m := range c.Movies {
		ormOpr.LoadRelated(m, "Tags")
		for _, tag := range m.Tags {
			m.TagArray = append(m.TagArray, tag.TagName)
		}
		m.Tags = nil
	}
	sort.Slice(c.Movies, func(i, j int) bool {
		return c.Movies[i].Title < c.Movies[j].Title
	})
	return c.Movies, err
}

func (m Movie) DeleteColl() error {
	tmpColl := []string{"coll_id", "coll_str"}
	_, err := ormOpr.Update(&m, tmpColl...)
	clearInvalidColl()
	return err
}

func GetAllColl() ([]string, error) {
	var colls []*Coll
	_, err := ormOpr.QueryTable("coll").All(&colls)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, t := range colls {
		res = append(res, t.CollName)
	}
	return res, nil
}

func ClearInvalidData() {
	clearInvalidMovie()
	clearInvalidTag()
	clearInvalidColl()
	logs.Info("ClearInvalidData")
}

func clearInvalidMovie() {
	var movies []*Movie
	_, err := ormOpr.QueryTable("movie").All(&movies)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, m := range movies {
		path := m.VideoUrl[0:1] + ":/" + m.VideoUrl[2:]
		_, err := os.Stat(path)
		if err == nil {
			continue
		}
		if os.IsNotExist(err) {
			_, err = ormOpr.Delete(m, "m_id")
			logs.Info("delete moviedata %v,err:%v", m.VideoUrl, err)
		}
	}
}

func clearInvalidTag() {
	var tags []*Tag
	_, err := ormOpr.QueryTable("tag").All(&tags)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, tag := range tags {
		_, err := ormOpr.LoadRelated(tag, "Movies")
		if err != nil {
			continue
		}
		if len(tag.Movies) == 0 {
			logs.Info("delete tag %s", tag.TagName)
			ormOpr.Delete(tag)
		}
	}
}

func clearInvalidColl() {
	var colls []*Coll
	_, err := ormOpr.QueryTable("coll").All(&colls)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, coll := range colls {
		_, err := ormOpr.LoadRelated(coll, "Movies")
		if err != nil {
			continue
		}
		if len(coll.Movies) == 0 {
			logs.Info("delete coll %s", coll.CollName)
			ormOpr.Delete(coll)
		}
	}
}

func GetMovieByRecent() ([]*Movie, error) {
	m, err := QueryAllMovieData()
	if err != nil {
		return m, err
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].RecentWatch > m[j].RecentWatch
	})
	return m, nil
}
