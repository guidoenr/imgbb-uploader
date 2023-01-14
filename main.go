package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Output struct {
	Data struct {
		DisplayUrl string `json:"display_url,omitempty"`
		Expiration int    `json:"expiration,omitempty"`
		Thumb      struct {
			Filename string `json:"filename,omitempty"`
			Mime     string `json:"mime,omitempty"`
			Url      string `json:"url,omitempty"`
		} `json:"thumb,omitempty"`
		DeleteUrl string `json:"delete_url,omitempty"`
	} `json:"data"`
	Success    bool `json:"success,omitempty"`
	Status     int  `json:"status,omitempty"`
	StatusCode int  `json:"status_code,omitempty"`
	Error      struct {
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
	} `json:"error,omitempty"`
	StatusTxt string `json:"status_txt,omitempty"`
}

func throwError(msg string) {
	fmt.Printf("\033[31m[error]\033[0m: %s", msg)
	os.Exit(0)
}

func main() {
	// creating the uploader
	var uploader Uploader

	// checking the args
	if len(os.Args) < 2 {
		throwError("you must provide an image\n(e.g: upload image.jpg)")
	}

	// reading the filename as first argument
	fileName := os.Args[1]

	// get the currentPath
	currentPath, _ := os.Getwd()

	// setting the imagePath
	imagePath := filepath.Join(currentPath, fileName)

	// initializing the uploader
	err := uploader.init(imagePath)
	if err != nil {
		msg := fmt.Sprintf("intializing uploader: %v", err)
		throwError(msg)
	}

	// uploading the img
	output, err := uploader.UploadImage()
	if err != nil {
		msg := fmt.Sprintf("uploading image: %v", err)
		throwError(msg)
	}

	fmt.Println(output)
}
