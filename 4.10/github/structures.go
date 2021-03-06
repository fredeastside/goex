package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type SearchIssueResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func (r SearchIssueResult) Len() int {
	return len(r.Items)
}

func (r SearchIssueResult) Swap(i, j int) {
	r.Items[i], r.Items[j] = r.Items[j], r.Items[i]
}

func (r SearchIssueResult) Less(i, j int) bool {
	return r.Items[j].CreatedAt.Before(r.Items[i].CreatedAt)
}
