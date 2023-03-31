package media

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pierre0210/reddit-dl/internal/util"
)

func SaveToSeperateFiles(videoName string, video []byte, audioName string, audio []byte) {
	err := ioutil.WriteFile(videoName, video, 0)
	util.ErrHandler(err, true)
	log.Printf("Saved video as %s\n", videoName)

	err = ioutil.WriteFile(audioName, audio, 0)
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
