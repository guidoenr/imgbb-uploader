package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const apiUrl = "https://api.imgbb.com/1/upload?"

type Uploader struct {
	apiKey    string
	apiUrl    string
	imagePath string
	client    *http.Client
	errorMsg  string
}

func (u *Uploader) init(imagePath string) {
	u.apiKey = os.Getenv("IMGBB_API_KEY")
	if u.apiKey == "" {
		log.Fatalf("IMGBB_API_KEY not setted")
	}

	u.apiUrl = fmt.Sprintf("%s&key=%s", apiUrl, u.apiKey)
	u.imagePath = imagePath
	u.client = &http.Client{}
	u.errorMsg = ""
}

func (u *Uploader) UploadImage() (string, error) {
	// pr = pipeReader, pw =pipeWriter
	pr, pw := io.Pipe()

	// creating the form multipart writer
	form := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()

		file, err := os.Open(u.imagePath) // path to image file
		if err != nil {
			msg := fmt.Sprintf("opening image '%s': %v", u.imagePath, err)
			u.errorMsg = msg
			return
		}

		w, err := form.CreateFormFile("image", u.imagePath)
		if err != nil {
			msg := fmt.Sprintf("creating form file '%s': %v", u.imagePath, err)
			u.errorMsg = msg
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			msg := fmt.Sprintf("copying image '%s': %v", u.imagePath, err)
			u.errorMsg = msg
			return
		}

		form.Close()
	}()

	// if no error were produced
	if u.errorMsg != "" {
		return "", errors.New(u.errorMsg)
	}

	// creating the request
	r, err := http.NewRequest(http.MethodPost, u.apiUrl, pr)
	if err != nil {
		msg := fmt.Sprintf("creating request': %v", err)
		return "", errors.New(msg)
	}

	// setting the header
	r.Header.Set("Content-Type", form.FormDataContentType())

	// making the request
	res, err := u.client.Do(r)
	if err != nil {
		msg := fmt.Sprintf("making the request: %v", err)
		return "", errors.New(msg)
	}

	output, err := u.generateOutput(res)

	return output, nil
}

func (u *Uploader) generateOutput(response *http.Response) (string, error) {
	var output Output

	// returning output
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		msg := fmt.Sprintf("reading the response body")
		return "", errors.New(msg)
	}

	err = json.Unmarshal(bytes, &output)
	if err != nil {
		msg := fmt.Sprintf("unmarshalling")
		return "", errors.New(msg)
	}

	indent, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}

	return string(indent), nil
}
