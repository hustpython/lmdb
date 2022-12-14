package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Movie struct {
	MId             string `orm:"pk"`
	ReleaseDate     string
	Poster          string
	Vote            float32
	LastWatch       int64
	RecentWatch     int64
	PathValid       bool
	Size            uint64
	ModifyTime      int64
	Duration        string
	VideoUrl        string
	Title           string
	Desc            string
	Cover           string
	Favourite       bool
	FavouriteUpdate string `orm:"-"`
	Coll            *Coll  `orm:"null;rel(fk)"`
	CollStr         string
	Tags            []*Tag     `orm:"rel(m2m)"`
	TagArray        []string   `orm:"-"`
	Comments        []*Comment `orm:"reverse(many)"`
	CutInfos        []*CutInfo `orm:"reverse(many)"`
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
	Title           string `orm:"-"`
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

type MovieSizeCountMap struct {
	Count uint32
	Size  uint64
}

type MovieTable struct {
	Key           string   `json:"key"`
	Title         string   `json:"title"`
	DurationStr   string   `json:"durationStr"`
	Duration      int      `json:"duration"`
	Desc          string   `json:"desc"`
	Tags          []string `json:"tags"`
	RecentStr     string   `json:"recentStr"`
	SizeStr       string   `json:"sizeStr"`
	ModifyTime    int64    `json:"modifyTime"`
	ModifyTimeStr string   `json:"modifyTimeStr"`
	Size          uint64   `json:"size"`
	Recent        int64    `json:"recent"`
	Online        string   `json:"online"`
}

type MovieTableWithChild struct {
	MovieTable
	Children []*MovieTableWithChild `json:"children"`
}

var ormOpr orm.Ormer

func init() {
	logs.SetLogger(logs.AdapterConsole)
	orm.RegisterModel(new(Movie), new(Coll), new(Tag), new(Comment), new(CutInfo))
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

func GetMovieSizeAndCount() map[string]*MovieSizeCountMap {
	var movies []*Movie
	var res = make(map[string]*MovieSizeCountMap)
	_, err := ormOpr.QueryTable("movie").Filter("path_valid", 0).All(&movies)
	if err != nil {
		return res
	}
	for _, m := range movies {
		tempPath := strings.Split(m.VideoUrl, "/")
		if len(tempPath) <= 0 {
			continue
		}
		_, exist := res[tempPath[0]]

		if exist {
			res[tempPath[0]].Size += m.Size
			res[tempPath[0]].Count++
		} else {
			res[tempPath[0]] = &MovieSizeCountMap{
				Size:  m.Size,
				Count: 1,
			}
		}

	}
	return res
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
					_, err := ormOpr.Update(m, []string{"video_url", "title"}...)
					fmt.Printf("update videurl: %s,err:%s", m.VideoUrl, err)
				}
				if movie.PathValid != m.PathValid {
					_, err := ormOpr.Update(m, "path_valid")
					fmt.Printf("update videurl: %s,err:%s", m.VideoUrl, err)
				}
				if movie.Size == 0 {
					_, err := ormOpr.Update(m, "size")
					fmt.Printf("update video size: %d,err:%s", m.Size, err)
				}
				if movie.ModifyTime == 0 {
					_, err := ormOpr.Update(m, "modifyTime")
					fmt.Printf("update video modifyTime: %d,err:%s", m.ModifyTime, err)
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
	if len(m.Poster) != 0 {
		updateCol = append(updateCol, "poster")
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
	if len(m.FavouriteUpdate) != 0 {
		updateCol = append(updateCol, "favourite")
	}
	if len(m.ReleaseDate) != 0 {
		updateCol = append(updateCol, "release_date")
	}
	if m.Vote != 0 {
		updateCol = append(updateCol, "vote")
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
	if err != nil {
		return nil, err
	}
	for _, m := range c.Movies {
		ormOpr.LoadRelated(m, "Tags")
		for _, tag := range m.Tags {
			m.TagArray = append(m.TagArray, tag.TagName)
		}
		m.Tags = nil
	}
	sort.Slice(c.Movies, func(i, j int) bool {
		if len(c.Movies[i].Title) != len(c.Movies[j].Title) {
			return len(c.Movies[i].Title) < len(c.Movies[j].Title)
		} else {
			return c.Movies[i].Title < c.Movies[j].Title
		}
	})
	return c.Movies, err
}

func (c Coll) GetCommentsByColl() ([][]*Comment, error) {
	_, err := ormOpr.LoadRelated(&c, "Movies")
	if err != nil {
		return nil, err
	}
	sort.Slice(c.Movies, func(i, j int) bool {
		if len(c.Movies[i].Title) != len(c.Movies[j].Title) {
			return len(c.Movies[i].Title) < len(c.Movies[j].Title)
		} else {
			return c.Movies[i].Title < c.Movies[j].Title
		}
	})
	var res [][]*Comment
	for _, m := range c.Movies {
		d, _ := m.GetCommentsByMId()
		res = append(res, d)
	}
	return res, nil
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
	for i := range m.Comments {
		m.Comments[i].Title = m.Title
	}
	return m.Comments, nil
}

func (m Comment) AddCommentByMId() error {
	_, e := ormOpr.Insert(&m)
	return e
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

func GetMovieByMId(mid string) *Movie {
	m := Movie{MId: mid}
	err := ormOpr.Read(&m)
	if err != nil {
		return nil
	}
	ormOpr.LoadRelated(&m, "Tags")
	for _, tag := range m.Tags {
		m.TagArray = append(m.TagArray, tag.TagName)
	}
	m.Tags = nil
	return &m
}

func GetMovieByFavourite() ([]*Movie, error) {
	var movies []*Movie
	_, err := ormOpr.QueryTable("movie").Filter("favourite", true).All(&movies)
	for _, m := range movies {
		ormOpr.LoadRelated(m, "Tags")
		for _, tag := range m.Tags {
			m.TagArray = append(m.TagArray, tag.TagName)
		}
		m.Tags = nil
	}
	return movies, err
}

func GetMovieTables() []*MovieTableWithChild {
	var movies []*Movie
	_, err := ormOpr.QueryTable("movie").All(&movies)
	if err != nil {
		return nil
	}
	var titleMap = make(map[string]struct{})
	var res []*MovieTableWithChild
	for _, m := range movies {
		tmp := &MovieTableWithChild{}
		if len(m.CollStr) != 0 {
			if _, ok := titleMap[m.CollStr]; ok {
				continue
			}
			tmp.Key = m.MId + "cj"
			tmp.Title = m.CollStr
			tmp.Tags = []string{}
			titleMap[m.CollStr] = struct{}{}
			coll := Coll{
				CollName: m.CollStr,
			}
			m, err := coll.GetMoviesByColl()
			if err != nil {
				continue
			}
			tmp.Online = "???"
			var tagMap = make(map[string]struct{})
			var tmpRecent, temModifyTime int64
			for _, m1 := range m {
				kk := setTableByMovie(m1)
				if kk.Online == "???" {
					tmp.Online = "???"
				}
				if m1.RecentWatch > tmpRecent {
					tmpRecent = m1.RecentWatch
				}
				if m1.ModifyTime > temModifyTime {
					temModifyTime = m1.ModifyTime
				}
				for _, t := range kk.Tags {
					if _, ok := tagMap[t]; !ok {
						tagMap[t] = struct{}{}
						tmp.Tags = append(tmp.Tags, t)
					}
				}
				tmp.Children = append(tmp.Children, kk)
			}
			tmp.Recent = tmpRecent
			tmp.RecentStr = timeInt2Date(tmpRecent)
			tmp.ModifyTime = temModifyTime
			tmp.ModifyTimeStr = timeInt2Date(temModifyTime)
		} else {
			tmp = setTableByMovie(m)
		}
		res = append(res, tmp)
	}
	return res
}

func setTableByMovie(m *Movie) *MovieTableWithChild {
	var res = MovieTableWithChild{}
	res.Key = m.MId
	res.Title = m.Title
	res.DurationStr = m.Duration
	res.Duration = durationStr2Int(m.Duration)
	res.Desc = m.Desc
	res.RecentStr = timeInt2Date(m.RecentWatch)
	res.Online = "???"
	res.Size = m.Size
	res.ModifyTime = m.ModifyTime
	res.ModifyTimeStr = timeInt2Date(m.ModifyTime)
	res.SizeStr = ByteSizeTransform(m.Size)
	res.Tags = []string{}
	if m.PathValid {
		res.Online = "???"
	}
	res.Recent = m.RecentWatch
	ormOpr.LoadRelated(m, "Tags")
	for _, tag := range m.Tags {
		res.Tags = append(res.Tags, tag.TagName)
	}
	return &res
}

func durationStr2Int(durationStr string) int {
	timeStr := strings.Split(durationStr, ":")
	if len(timeStr) == 3 {
		h, _ := strconv.Atoi(timeStr[0])
		m, _ := strconv.Atoi(timeStr[1])
		s, _ := strconv.Atoi(timeStr[2])
		return h*3600 + m*60 + s
	} else {
		return 0
	}
}

func timeInt2Date(timeInt int64) string {
	if (timeInt) == 0 {
		return ""
	}
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(timeInt, 0).Format(timeLayout)
	return timeStr
}

// ???1024????????????
func ByteSizeTransform(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %c",
		float64(b)/float64(div), "KMGTPE"[exp])
}

type searchMovieRes struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func ShowMoviesSearch() []*searchMovieRes {
	var movies []*Movie
	_, err := ormOpr.QueryTable("movie").All(&movies)
	if err != nil {
		return nil
	}
	var res []*searchMovieRes
	for _, m := range movies {
		res = append(res, &searchMovieRes{
			Label: m.Title,
			Value: m.MId,
		})
	}
	return res
}

func GetMoviesByMId(mid string) ([]*Movie, error) {
	m := GetMovieByMId(mid)
	if m == nil {
		return nil, nil
	}
	if len(m.CollStr) == 0 {
		return []*Movie{m}, nil
	}
	c := Coll{
		CollName: m.CollStr,
	}
	return c.GetMoviesByColl()
}
