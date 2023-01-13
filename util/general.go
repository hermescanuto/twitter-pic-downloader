package util

import (
	"log"
	"os"
)

var Folder *string

// SetFolder setup the folder .
// Return *string
func SetFolder() {
	r, _ := os.Getwd()
	Folder = &r
}

func GetFolder() string {
	return *Folder
}

func CatchGeneralError(err *error) {
	if *err != nil {
		log.Fatalln(*err)
	}
}
