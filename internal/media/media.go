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

func SaveToSeperateFiles(videoName string, video []byte, audioName string, audio []byte) {
	err := os.WriteFile(videoName, video, 0660)
	util.ErrHandler(err, true)
	log.Printf("Saved video as %s\n", videoName)

	err = os.WriteFile(audioName, audio, 0660)
	util.ErrHandler(err, true)
	log.Printf("Saved audio as %s\n", audioName)
}

func GetVideo(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	util.ErrHandler(err, true)
	res, err := client.Do(req)
	util.ErrHandler(err, true)

	body, err := io.ReadAll(res.Body)
	util.ErrHandler(err, true)
	if res.StatusCode != 200 {
		util.ErrHandler(errors.New(res.Status), false)
	}
	return body
}

func GetAudio(baseUrl string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+"DASH_audio.mp4", nil)
	util.ErrHandler(err, true)
	res, err := client.Do(req)
	util.ErrHandler(err, true)

	body, err := io.ReadAll(res.Body)
	util.ErrHandler(err, true)
	if res.StatusCode != 200 {
		util.ErrHandler(errors.New(res.Status), false)
		return nil
	}
	return body
}

func Convert2Gif(videoPath string) {
	err := ffmpeg.Input(videoPath).Output(strings.Split(videoPath, ".mp4")[0] + ".gif").OverWriteOutput().ErrorToStdOut().Run()
	util.ErrHandler(err, true)
}
