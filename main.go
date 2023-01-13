package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hermescanuto/twitter-pic-downloader/twitter"
	"github.com/hermescanuto/twitter-pic-downloader/util"
)

func main() {
	fmt.Println("Twitter Downloader")
	start := time.Now()
	util.SetFolder()
	payload := util.CheckUp()

	log.Println("Profiles", payload)
	api := twitter.New()

	var results = make([]chan twitter.Result, len(payload))
	for i, v := range payload {
		results[i] = make(chan twitter.Result)
		go twitter.GetTweeter(api, v.Screenname, results[i]) // run goroutine
	}

	//waiting for the goroutine
	for i := range results {
		r := <-results[i]
		log.Printf("%v - %v", r.Screename, r.Total)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Finished ", elapsed)
}
