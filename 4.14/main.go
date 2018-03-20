package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/fredeastside/goex/4.11/github"
)

var list []*github.Issue

func init() {
	var err error
	list, err = github.ListForRepo("symfony", "symfony")
	if err != nil {
		fmt.Printf("Can't load golang/go issues. Error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	issuesList := template.Must(template.New("issueslist").Parse(`
		<h1>Symfony</h1>
		<table>
			<tr>
				<td>#</td>
				<td>Title</td>
			</tr>
			{{range .}}
			<tr>
				<td>{{.Id}}</td>
				<td><a href="{{.Url}}">{{.Title}}</a></td>
			</tr>
			{{end}}
		</table>
	`))
	issuesList.Execute(w, list)
}
