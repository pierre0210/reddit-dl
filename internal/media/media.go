package media

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pierre0210/reddit-dl/internal/util"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetVideo(client *http.Client, url string, videoName string) {
	req, err := http.NewRequest("GET", url, nil)
	util.ErrHandler(err, true)
	res, err := client.Do(req)
	util.ErrHandler(err, true)
	if res.StatusCode != 200 {
		util.ErrHandler(errors.New(res.Status), false)
		return
	}

	videoBuff, err := io.ReadAll(res.Body)
	util.ErrHandler(err, true)
	res.Body.Close()

	err = os.WriteFile(videoName, videoBuff, 0660)
	util.ErrHandler(err, true)
	log.Printf("Saved video as %s\n", videoName)
}

func GetAudio(client *http.Client, baseUrl string, audioName string) {
	req, err := http.NewRequest("GET", baseUrl+"DASH_audio.mp4", nil)
	util.ErrHandler(err, true)
	res, err := client.Do(req)
	util.ErrHandler(err, true)
	if res.StatusCode != 200 {
		util.ErrHandler(errors.New(res.Status), false)
		return
	}

	audioBuff, err := io.ReadAll(res.Body)
	util.ErrHandler(err, true)
	res.Body.Close()

	err = os.WriteFile(audioName, audioBuff, 0660)
	util.ErrHandler(err, true)
	log.Printf("Saved audio as %s\n", audioName)
}

func Convert2Gif(videoPath string) {
	err := ffmpeg.Input(videoPath).Output(strings.Split(videoPath, ".mp4")[0] + ".gif").OverWriteOutput().ErrorToStdOut().Run()
	util.ErrHandler(err, true)
}
