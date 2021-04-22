package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AssignResp struct {
	Fid       string `json:"fid"`
	URL       string `json:"url"`
	PublicURL string `json:"publicUrl"`
	Count     int    `json:"count"`
	Error     string `json:"error"`
}

// Assign gets the designated url to store the media
func Assign(weedAddr string) (*AssignResp, error) {
	dirurl := fmt.Sprintf("%s/dir/assign", weedAddr)
	resp, err := http.Get(dirurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	assignment := &AssignResp{}
	if err := json.NewDecoder(resp.Body).Decode(assignment); err != nil {
		return nil, err
	}
	if assignment.Error != "" {
		return nil, errors.New(assignment.Error)
	}
	return assignment, nil
}

type UploadResp struct {
	Name  int    `json:"name"`
	Size  int    `json:"size"`
	ETag  string `json:"eTag"`
	Error string `json:"error"`
}

func Upload(publicAddr, contentType string, mf io.Reader) (*UploadResp, error) {
	resp, err := http.Post(publicAddr, contentType, mf)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	upload := &UploadResp{}
	if err = json.NewDecoder(resp.Body).Decode(upload); err != nil {
		return nil, err
	}

	if upload.Error != "" {
		return nil, errors.New(upload.Error)
	}
	return upload, nil
}
