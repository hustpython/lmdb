package controllers

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"lmdbapi/models"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type FfmpegController struct {
	ErrorController
}

func transTimeSecFromStr(timeStr string) int {
	timeSplit := strings.Split(timeStr, ":")
	hour, _ := strconv.Atoi(timeSplit[0])
	min, _ := strconv.Atoi(timeSplit[1])
	sec, _ := strconv.Atoi(timeSplit[2])
	return hour*3600 + min*60 + sec
}

// POST /v1/ffmpeg/cut
func (f *FfmpegController) CutVideo() {
	var cutInfo models.CutInfo
	err := f.BindJSON(&cutInfo)
	if err != nil {
		f.setInternalError(err.Error())
	}
	m := models.GetMovieByMId(cutInfo.Mid)
	if m == nil {
		f.setInternalError("not find movie")
		return
	}
	path := m.VideoUrl[0:1] + ":/" + m.VideoUrl[2:]
	startInt := transTimeSecFromStr(cutInfo.Start)
	durInt := transTimeSecFromStr(cutInfo.Duration)
	path2 := strings.Replace(path, filepath.Ext(path),
		fmt.Sprintf("_cut_%d_%d%s", startInt,
			durInt, filepath.Ext(path)), -1)

	if models.GetCutInfoByPath(path2) != nil {
		f.setInternalError("already exist")
		return
	}
	err = ffmpeg.Input(path, ffmpeg.KwArgs{"ss": startInt}).
		Output(path2,
			ffmpeg.KwArgs{"t": durInt, "vcodec": "copy", "acodec": "copy"}).OverWriteOutput().
		ErrorToStdOut().Run()
	if err != nil {
		f.setInternalError(err.Error())
		return
	}
	// 写入数据库
	cutInfo.Path = path2
	cutInfo.Movie = &models.Movie{MId: cutInfo.Mid}
	if err = cutInfo.AddCutInfoByPath(); err != nil {
		f.setInternalError(err.Error())
		os.Remove(path2)
		return
	} else {
		temData := &cutInfo
		temData.Path = path2
		f.Data["json"] = temData
		f.ServeJSON()
	}
}

// Put /v1/ffmpeg/cut
func (f *FfmpegController) OpenCutVideo() {
	exec.Command(`cmd`, `/c`, `explorer`, filepath.Dir(f.GetString("Path"))).Start()
}

// Del /v1/ffmpeg/cut
func (f *FfmpegController) DelCutVideo() {
	path := f.GetString("Path")
	e := os.Remove(path)
	if e != nil {
		f.setInternalError(e.Error())
		return
	}
	c := models.CutInfo{Path: path}
	e = c.DelCutInfoByPath()
	if e != nil {
		f.setInternalError(e.Error())
		return
	}
}

func (f *FfmpegController) GetCutInfosByMId() {
	mid := f.GetString("MId")
	if len(mid) == 0 {
		f.setParamInvalid("MId is null")
		return
	}
	var movieModel = models.Movie{MId: mid}
	d, err := movieModel.GetCutInfosByMId()
	if err != nil {
		f.setInternalError(err.Error())
		return
	} else {
		f.Data["json"] = d
		f.ServeJSON()
	}
}