package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"sort"
	"strings"
)

type Movie struct {
	MId         string `orm:"pk"`
	LastWatch   int64
	RecentWatch int64
	PathValid   bool
	Duration    string
	VideoUrl    string
	Title       string
	Desc        string
	Cover       string
	Coll        *Coll `orm:"null;rel(fk)"`
	CollStr     string
	Tags        []*Tag     `orm:"rel(m2m)"`
	TagArray    []string   `orm:"-"`
	Comments    []*Comment `orm:"reverse(many)"`
}

type Tag struct {
	TagName string   `orm:"pk"`
	Movies  []*Movie `orm:"reverse(many)"`
}

type Coll struct {
	CollName string   `orm:"pk"`
	Movies   []*Movie `orm:"reverse(many)"`
	MIds     []string `orm:"-"`
}

type Comment struct {
	CommentID       string `orm:"pk"` //MID+CommentTime
	CommentFrame    float64
	CommentStr      string
	CommentFrameStr string
	CommentTime     string
	CommentImage    string
	MId             string `orm:"-"`
	Movie           *Movie `orm:"null;rel(fk)"`
}

type CollWithMovie struct {
	Title           string
	DefaultCollName string
	SelectedValues  []string
	TranOptions     []*TranOptions
}

type TranOptions struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
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
	orm.RegisterModel(new(Movie), new(Coll), new(Tag), new(Comment))
	orm.RegisterDataBase("default", "sqlite3", "./lmdb.db")
	orm.RunSyncdb("default", false, true)
	ormOpr = orm.NewOrm()
	ClearInvalidData()
}

func QueryAllMovieData() ([]*Movie, error) {
	var movies []*Movie
	_, err := ormOpr.QueryTable("movie").Filter("path_valid", 0).All(&movies)
	for _, m := range movies {
		ormOpr.LoadRelated(m, "Tags")
		for _, tag := range m.Tags {
			m.TagArray = append(m.TagArray, tag.TagName)
		}
		m.Tags = nil
	}
	return movies, err
}

func QuerySortMovie() ([]*CollWithMovie, error) {
	var d []*Movie
	_, err := ormOpr.QueryTable("movie").All(&d)

	var resMap = make(map[string][]*TranOptions)
	var defaultCollNameMap = make(map[string]string)
	for _, i := range d {
		key := i.VideoUrl[:strings.LastIndex(i.VideoUrl, "/")]
		var label string
		var disabled bool
		if len(i.CollStr) != 0 {
			label = i.Title + " ||| " + i.CollStr
			disabled = true
			defaultCollNameMap[key] = i.CollStr
		} else {
			label = i.Title
		}
		tmp := &TranOptions{
			Label:    label,
			Value:    i.MId,
			Disabled: disabled,
		}
		resMap[key] = append(resMap[key], tmp)
	}
	var resArray []*CollWithMovie
	for k, v := range resMap {
		sort.Slice(v, func(i, j int) bool {
			return v[i].Label < v[j].Label
		})
		tmp := &CollWithMovie{
			Title:           k,
			TranOptions:     v,
			DefaultCollName: defaultCollNameMap[k],
		}
		resArray = append(resArray, tmp)
	}
	sort.Slice(resArray, func(i, j int) bool {
		return resArray[i].Title < resArray[j].Title
	})
	return resArray, err
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
				if movie.PathValid != m.PathValid {
					_, err := ormOpr.Update(m, "path_valid")
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
	_, err := ormOpr.QueryTable("Coll").All(&colls)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, t := range colls {
		res = append(res, t.CollName)
	}
	return res, nil
}

func (c Coll) BatchAddColl() error {
	for _, v := range c.MIds {
		tem := Movie{
			MId:     v,
			CollStr: c.CollName,
		}
		if err := tem.UpdateMovieData(); err != nil {
			fmt.Println("UpdateMovieData get err:", err)
			return err
		}
	}
	return nil
}

func (m Movie) GetCommentsByMId() ([]*Comment, error) {
	_, err := ormOpr.LoadRelated(&m, "Comments")
	if err != nil {
		return nil, err
	}
	return m.Comments, nil
}

func (m Comment) AddCommentByMId() error {
	if _, e := ormOpr.Insert(&m); e != nil {
		return e
	}
	return nil
}

func (m Comment) DelCommentByMId() error {
	qs := ormOpr.QueryTable(new(Comment))
	if qs.Filter("CommentID", m.CommentID).Exist() {
		if _, e := ormOpr.Delete(&m); e != nil {
			return e
		}
	}
	return nil
}

func ClearInvalidData() {
	clearInvalidMovie(false)
	clearInvalidTag()
	clearInvalidColl()
	logs.Info("ClearInvalidData")
}

func clearInvalidMovie(force bool) {
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
			if force {
				_, err = ormOpr.Delete(m, "m_id")
				logs.Info("delete moviedata %v,err:%v", m.VideoUrl, err)
			} else {
				m.PathValid = true
				_, err = ormOpr.Update(m, "path_valid")
				logs.Info("update movie path invalid %v,err:%v", m.VideoUrl, err)
			}
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
	_, err := ormOpr.QueryTable("Coll").All(&colls)
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
			logs.Info("delete Coll %s", coll.CollName)
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
