package xkcd

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const ApiUrl = "https://xkcd.com/"

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func FindByNumber(number int) (*Comic, error) {
	response, err := http.Get(ApiUrl + strconv.Itoa(number) + "/info.0.json")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var comic Comic
	if err := json.NewDecoder(response.Body).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}
