package controllers

import (
	"bytes"
	"fmt"
	"github.com/beego/beego/v2/core/logs"

	"lmdbapi/models"

	"os/exec"
)

type FfmpegController struct {
	ErrorController
}

// POST /v1/ffmpeg/cut
func (f *FfmpegController) CutVideo() {
	var cutInfo models.CutInfo
	err := f.BindJSON(&cutInfo)
	logs.Info(cutInfo, err)
	if err != nil {
		f.setInternalError(err.Error())
	}
	cmd := exec.Command("ffmpeg", "-ss 00:00:00", "-t 00:00:30", "-i test.mp4", "-vcodec copy -acodec copy output.mp4", "-")

	buf := new(bytes.Buffer)

	cmd.Stdout = buf

	if err := cmd.Run(); err != nil {
		logs.Info(err, cmd.Stderr)
	}

	fmt.Println(cmd.Stdout)
}
