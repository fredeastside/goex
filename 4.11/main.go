package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/fredeastside/goex/4.11/github"
)

const TEXT_EDITOR = "vim"

func main() {
	var action, owner, repo, title, id string

	flag.StringVar(&owner, "owner", "", "Owner of github repository.")
	flag.StringVar(&repo, "repo", "", "Name of github repository.")
	flag.StringVar(&action, "action", "", "Action with github issues.")
	flag.StringVar(&title, "title", "", "Title of issue.")
	flag.StringVar(&id, "id", "", "Id of issue.")
	flag.Parse()

	switch action {
	case "create":
		createIssue(owner, repo, title)
	case "update":
		updateIssue(owner, repo, id, title)
	case "list":
		printIssuesList(owner, repo)
	default:
		fmt.Println("You can use actions: create, update, list.")
	}
}

func createIssue(owner, repo, title string) {
	content, err := getIssueBody(title)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	issue, err := github.Create(owner, repo, &github.Issue{Title: title, Body: content})
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Issue created with id %d", issue.Id)
}

func updateIssue(owner, repo, id, title string) {
	content, err := getIssueBody(title)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	issue, err := github.Update(owner, repo, id, &github.Issue{Title: title, Body: content})
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Issue updated with id %d", issue.Id)
}

func getIssueBody(title string) (string, error) {
	tmpFileName := generateTmpFileName(title)
	cmd := exec.Command(TEXT_EDITOR, tmpFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	content, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		return "", err
	}
	if err := os.Remove(tmpFileName); err != nil {
		return "", err
	}

	return string(content), nil
}

func printIssuesList(owner, repo string) {
	list, err := github.ListForRepo(owner, repo)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	for _, item := range list {
		fmt.Println(item.Title)
		fmt.Println(item.Body)
	}
}

func generateTmpFileName(title string) string {
	return title + strconv.Itoa(int(time.Now().Unix()))
}
