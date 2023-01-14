package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Output struct {
	Data struct {
		DisplayUrl string `json:"display_ur,omitempty"`
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
	fmt.Printf("[error]: %s\n", msg)
	os.Exit(0)
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
	err := uploader.init(imagePath)
	if err != nil {
		msg := fmt.Sprintf("intializing uploader: %v", err)
		throwError(msg)
	}

	// uploading the img
	output, err := uploader.UploadImage()
	fmt.Println(uploader.errorMsg)
	if err != nil {
		msg := fmt.Sprintf("uploading image: %v", err)
		throwError(msg)
	}

	fmt.Println(output)

}
