package github

import (
	"encoding/json"
	"time"
)

type Issue struct {
	Id        int       `json:"number"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Url       string    `json:"html_url"`
	CreatedAt time.Time `json:"created_at"`
}

func Create(owner, repo string, issue *Issue) (*Issue, error) {
	jsonData, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	response, err := PostWithAuth("repos/"+owner+"/"+repo+"/issues", string(jsonData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(issue); err != nil {
		return nil, err
	}

	return issue, nil
}

func Update(owner, repo, id string, issue *Issue) (*Issue, error) {
	jsonData, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	response, err := PatchWithAuth("repos/"+owner+"/"+repo+"/issues/"+id, string(jsonData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(issue); err != nil {
		return nil, err
	}

	return issue, nil
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
