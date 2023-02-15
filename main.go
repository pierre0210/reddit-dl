package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func saveFiles(videoName string, video []byte, audioName string, audio []byte) {

}

func getVideo(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	errHandler(err, true)
	res, err := client.Do(req)
	errHandler(err, true)

	body, err := io.ReadAll(res.Body)
	errHandler(err, true)
	if res.StatusCode != 200 {
		errHandler(errors.New(res.Status), false)
	}
	return body
}

func getAudio(baseUrl string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+"DASH_audio.mp4", nil)
	errHandler(err, true)
	res, err := client.Do(req)
	errHandler(err, true)

	body, err := io.ReadAll(res.Body)
	errHandler(err, true)
	if res.StatusCode != 200 {
		//errHandler(errors.New(res.Status), false)
		return nil
	}
	return body
}

func processData(res *http.Response) {
	var jsonObj []map[string]any

	defer res.Body.Close()
	jsonStr, err := io.ReadAll(res.Body)
	errHandler(err, true)
	err = json.Unmarshal(jsonStr, &jsonObj)
	errHandler(err, true)

	data := jsonObj[0]["data"].(map[string]any)
	children := data["children"].([]any)
	media := children[0].(map[string]any)["data"].(map[string]any)["media"]
	if media == nil {
		errHandler(errors.New("No video."), false)
	}
	videoUrl := strings.Split(media.(map[string]any)["reddit_video"].(map[string]any)["fallback_url"].(string), "?")[0]
	baseUrl := children[0].(map[string]any)["data"].(map[string]any)["url"].(string) + "/"
	video := getVideo(videoUrl)
	audio := getAudio(baseUrl)
	log.Println(video)
	log.Println(audio)
}

func errHandler(err error, fatal bool) {
	if err == nil {
		return
	} else if fatal {
		log.Fatalf("%s\n", err.Error())
	} else {
		log.Printf("%s\n", err.Error())
		os.Exit(0)
	}
}

func main() {
	url := flag.String("u", "", "Reddit post url.")
	flag.Parse()

	match, _ := regexp.Match("https://www.reddit.com/r/([a-zA-Z_]+)/comments/([a-z0-9]+)/([a-z_]+)/", []byte(*url))
	if !match {
		errHandler(errors.New("Wrong Reddit url format."), false)
	}
	*url += ".json"

	client := &http.Client{}
	req, err := http.NewRequest("GET", *url, nil)
	errHandler(err, true)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	errHandler(err, true)
	if res.StatusCode != 200 {
		errHandler(errors.New(res.Status), false)
	}
	processData(res)
}
