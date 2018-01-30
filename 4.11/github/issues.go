package github

import (
	"encoding/json"
	"strconv"
	"time"
)

type Issue struct {
	Id        int       `json:"number,omitempty"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Url       string    `json:"html_url,omitempty"`
	State     string    `json:"state,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func Create(owner, repo string, issue *Issue) (*Issue, error) {
	jsonData, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	response, err := Post("repos/"+owner+"/"+repo+"/issues", string(jsonData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(issue); err != nil {
		return nil, err
	}

	return issue, nil
}

func Update(owner, repo string, issue *Issue) (*Issue, error) {
	jsonData, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	response, err := Patch(
		"repos/"+owner+"/"+repo+"/issues/"+strconv.Itoa(issue.Id),
		string(jsonData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(issue); err != nil {
		return nil, err
	}

	return issue, nil
}

func Lock(owner, repo string, id int) (bool, error) {
	response, err := Put("repos/"+owner+"/"+repo+"/issues/"+strconv.Itoa(id)+"/lock", "")
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	return true, nil
}

func ListForRepo(owner, repo string) ([]*Issue, error) {
	response, err := Get("repos/" + owner + "/" + repo + "/issues")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var list []*Issue
	if err := json.NewDecoder(response.Body).Decode(&list); err != nil {
		return nil, err
	}

	return list, nil
}
