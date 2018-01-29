package github

import (
	"bytes"
	"fmt"
	"net/http"
)

const APIURL = "https://api.github.com/"
const (
	LOGIN    = "fredeastside"
	PASSWORD = "SK^&49ws"
)

func Get(url string) (*http.Response, error) {
	response, err := http.Get(APIURL + url)
	if err != nil {
		return nil, fmt.Errorf("Can not get %s. Status: %s", url, response.Status)
	}

	return response, nil
}

func Post(url, json string) (*http.Response, error) {
	response, err := http.Post(APIURL+url, "application/json", bytes.NewBuffer([]byte(json)))
	if err != nil {
		return nil, fmt.Errorf("Can not post %s. Status: %s", url, response.Status)
	}

	return response, nil
}

func PostWithAuth(url, json string) (*http.Response, error) {
	request, _ := http.NewRequest("POST", APIURL+url, bytes.NewBuffer([]byte(json)))
	request.Header.Add("Content-type", "application/json;charset=utf-8")
	request.SetBasicAuth(LOGIN, PASSWORD)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Can not post %s. Status: %s", url, response.Status)
	}

	return response, nil
}

func PatchWithAuth(url, json string) (*http.Response, error) {
	request, _ := http.NewRequest("PATCH", APIURL+url, bytes.NewBuffer([]byte(json)))
	request.Header.Add("Content-type", "application/json;charset=utf-8")
	request.SetBasicAuth(LOGIN, PASSWORD)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Can not patch %s. Status: %s", url, response.Status)
	}

	return response, nil
}
