package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Output struct {
	Data struct {
		DisplayUrl string `json:"display_url"`
		Expiration int    `json:"expiration"`
		Thumb      struct {
			Filename string `json:"filename"`
			Mime     string `json:"mime"`
			Url      string `json:"url"`
		} `json:"thumb"`
		DeleteUrl string `json:"delete_url"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

func throwError(msg string) {
	fmt.Printf("[error]: %s\n", msg)
}

func main() {
	// creating the uploader
	var uploader Uploader

	// checkign the args
	if len(os.Args) < 2 {
		throwError("you must select your image")
	}

	// reading the filename as first argument
	fileName := os.Args[1]

	// currentPath
	currentPath, _ := os.Getwd()

	// setting the imagePath
	imagePath := filepath.Join(currentPath, fileName)

	// initializing the uploader
	uploader.init(imagePath)

	// uploading the img
	output, err := uploader.UploadImage()
	if err != nil {
		msg := fmt.Sprintf("uploading image: %v", err)
		throwError(msg)
	}

	fmt.Println(output)

}
