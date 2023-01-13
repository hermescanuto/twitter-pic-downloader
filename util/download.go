package util

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Screename struct {
	Screenname string `json:"screenname"`
}

func CheckUp() []Screename {

	SetFolder()

	var payload []Screename

	CheckFolder(filepath.Join(GetFolder(), "files"))
	CheckFolder(filepath.Join(GetFolder(), "files"))

	jsonfile := filepath.Join(GetFolder(), "files", "twitter.json")
	data, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		var item []Screename
		item = append(item, Screename{Screenname: "hermes"})
		e, err := json.MarshalIndent(item, "", "\t")
		if err != nil {
			log.Println(err)
		}
		contend := string(e)
		log.Println(contend)
		CreateJsonFile(jsonfile, contend)
		log.Println("Create the File")
		payload = item
	} else {
		err = json.Unmarshal(data, &payload)
		CatchGeneralError(&err)
	}
	return payload
}

func CreateJsonFile(file string, content string) {
	f, err := os.Create(file)
	if err != nil {
		log.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		log.Println(err)
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
			return
		}
		return
	}
	log.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func CheckFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		log.Println("Creating folder", folder)
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length")) // get the size of file
	downloadSize := int64(size)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	if downloadSize == 55587 ||
		downloadSize == 52380 ||
		downloadSize == 55846 ||
		downloadSize == 38894 ||
		downloadSize == 44583 ||
		downloadSize == 40907 ||
		downloadSize == 34934 ||
		downloadSize == 32410 {
		os.Remove(filepath)
		return fmt.Errorf("file : %v removed", filepath)
	}

	return err
}
