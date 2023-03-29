package media

import (
	"io/ioutil"
	"log"

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
