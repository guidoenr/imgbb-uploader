package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

// init to initialize the uploader
func (u *Uploader) init(imagePath string) error {
	u.apiKey = os.Getenv("IMGBB_API_KEY")
	if u.apiKey == "" {
		msg := fmt.Sprintf("api key not setted")
		return errors.New(msg)
	}
	if !u.fileExists(imagePath) {
		msg := fmt.Sprintf("the file '%s' doesn't exist", imagePath)
		return errors.New(msg)
	}

	u.apiUrl = fmt.Sprintf("%s&key=%s", apiUrl, u.apiKey)
	u.imagePath = imagePath
	u.client = &http.Client{}
	u.errorMsg = ""

	return nil
}

// UploadImage just do that
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
			fmt.Println(msg)
			u.errorMsg = msg
			return
		}

		w, err := form.CreateFormFile("image", u.imagePath)
		if err != nil {
			msg := fmt.Sprintf("creating form file '%s': %v", u.imagePath, err)
			fmt.Println(msg)
			u.errorMsg = msg
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			msg := fmt.Sprintf("copying image '%s': %v", u.imagePath, err)
			fmt.Println(msg)
			u.errorMsg = msg
			return
		}

		form.Close()
	}()

	// if error were produced
	if u.errorMsg != "" {
		fmt.Println(u.errorMsg)
		return "", errors.New(u.errorMsg)
	}

	// creating the request
	request, err := http.NewRequest(http.MethodPost, u.apiUrl, pr)
	if err != nil {
		msg := fmt.Sprintf("creating request': %v", err)
		fmt.Println(msg)
		return "", errors.New(msg)
	}

	// setting the header
	request.Header.Set("Content-Type", form.FormDataContentType())

	// making the request
	response, err := u.client.Do(request)
	if err != nil {
		msg := fmt.Sprintf("making the request: %v", err)
		fmt.Println(msg)
		return "", errors.New(msg)
	}

	// generating the output
	output, err := u.generateOutput(response)

	return output, nil
}

// generateOutput will parse the response into a json output to print in the terminal
func (u *Uploader) generateOutput(response *http.Response) (string, error) {
	var output Output

	// returning output
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return "", err
	}

	indent, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}

	return string(indent), nil

}

// fileExists check if the file exists
func (u *Uploader) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
