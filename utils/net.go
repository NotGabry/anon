package AnonUtils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type ResponseUpload struct {
	Status bool       `json:"status"`
	Data   DataStruct `json:"data"`
}

type DataStruct struct {
	File FileStruct `json:"file"`
}

type FileStruct struct {
	URL URL `json:"url"`
}

type URL struct {
	Full string `json:"full"`
}

func Upload(path string) (bool, ResponseUpload) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(path)
	if err != nil {
		return false, ResponseUpload{}
	}

	name := filepath.Base(path)

	defer file.Close()

	formFile, err := writer.CreateFormFile("file", name)
	if err != nil {
		return false, ResponseUpload{}
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return false, ResponseUpload{}
	}

	writer.Close()

	request, err := http.NewRequest("POST", "https://api.anonfiles.com/upload", body)
	if err != nil {
		return false, ResponseUpload{}
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false, ResponseUpload{}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false, ResponseUpload{}
	}

	bytess, _ := io.ReadAll(response.Body)
	var sent ResponseUpload
	json.Unmarshal(bytess, &sent)

	return true, sent
}
