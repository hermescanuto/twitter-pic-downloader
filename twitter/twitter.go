package twitter

import (
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ChimeraCoder/anaconda"
	"github.com/hermescanuto/twitter-pic-downloader/util"
)

type Result struct {
	Screename string
	Total     int
}

func New() *anaconda.TwitterApi {
	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"))
	return api
}

func GetTweeter(api *anaconda.TwitterApi, screenname string, done chan Result) {

	folder := filepath.Join(util.GetFolder(), "img", screenname)
	util.CheckFolder(folder)
	v := url.Values{}
	v.Set("screen_name", screenname)
	v.Set("count", "1000")
	searchResult, _ := api.GetUserTimeline(v)

	a := 0
	for _, tweet := range searchResult {
		for _, v := range tweet.ExtendedEntities.Media {
			file := filepath.Base(v.Media_url)
			if _, err := os.Stat(filepath.Join(folder, file)); os.IsNotExist(err) {
				err := util.DownloadFile(filepath.Join(folder, file), v.Media_url)
				if err == nil {
					log.Println(screenname, file)
					a++
				} else {
					log.Println(err.Error())
				}
			}

		}
	}
	done <- Result{Screename: screenname, Total: a}
}
