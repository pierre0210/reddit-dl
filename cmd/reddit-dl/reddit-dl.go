package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/pierre0210/reddit-dl/internal/media"
	"github.com/pierre0210/reddit-dl/internal/util"
)

func saveFiles(videoName string, video []byte, audioName string, audio []byte) {
	err := ioutil.WriteFile(videoName, video, 0)
	util.ErrHandler(err, true)
	log.Printf("Saved as %s\n", videoName)
}

func processData(res *http.Response) {
	var jsonObj []map[string]any

	defer res.Body.Close()
	jsonStr, err := io.ReadAll(res.Body)
	util.ErrHandler(err, true)
	err = json.Unmarshal(jsonStr, &jsonObj)
	util.ErrHandler(err, true)

	data := jsonObj[0]["data"].(map[string]any)
	children := data["children"].([]any)
	mediaObj := children[0].(map[string]any)["data"].(map[string]any)["media"]
	if mediaObj == nil {
		util.ErrHandler(errors.New("No video."), false)
	}
	videoUrl := strings.Split(mediaObj.(map[string]any)["reddit_video"].(map[string]any)["fallback_url"].(string), "?")[0]
	baseUrl := children[0].(map[string]any)["data"].(map[string]any)["url"].(string) + "/"
	video := media.GetVideo(videoUrl)
	audio := media.GetAudio(baseUrl)
	media.SaveToSeperateFiles("video.mp4", video, "audio.mp4", audio)
}

func main() {
	url := flag.String("u", "", "Reddit post url.")
	flag.Parse()

	match, _ := regexp.Match("https://www.reddit.com/r/([a-zA-Z_]+)/comments/([a-z0-9]+)/([a-z_]+)/", []byte(*url))
	if !match {
		util.ErrHandler(errors.New("Wrong Reddit url format."), false)
	}
	*url += ".json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", *url, nil)
	util.ErrHandler(err, true)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	util.ErrHandler(err, true)
	if res.StatusCode != 200 {
		util.ErrHandler(errors.New(res.Status), false)
	}
	processData(res)
}
