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
	return doAuthRequest("POST", url, json)
}

func Patch(url, json string) (*http.Response, error) {
	return doAuthRequest("PATCH", url, json)
}

func Put(url, json string) (*http.Response, error) {
	return doAuthRequest("PUT", url, json)
}

func doAuthRequest(requestType, url, json string) (*http.Response, error) {
	request, _ := http.NewRequest(requestType, APIURL+url, bytes.NewBuffer([]byte(json)))
	request.Header.Add("Content-type", "application/json;charset=utf-8")
	request.SetBasicAuth(LOGIN, PASSWORD)
	client := &http.Client{}

	return client.Do(request)
}
