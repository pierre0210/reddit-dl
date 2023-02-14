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

func getDashXml(url string) {

}

func getVideo(res *http.Response) {
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
	baseUrl := strings.Split(media.(map[string]any)["reddit_video"].(map[string]any)["dash_url"].(string), "DASHPlaylist.mpd")[0]
	log.Println(baseUrl)
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
	res, err := http.Get(*url)
	errHandler(err, true)
	if res.StatusCode != 200 {
		errHandler(errors.New(res.Status), false)
	}
	getVideo(res)
}
