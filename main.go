package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v50/github"
)

type Repo struct {
	Name   string   `json:"name"`
	Topics []string `json:"topics"`
}

func main() {
	var noTopics []string
	var repos []Repo
	b, err := os.ReadFile("organizations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	orgs := strings.Split(string(b), "\n")
	client := github.NewClient(nil)

	for _, o := range orgs {
		ghrepos, _, err := client.Repositories.List(context.Background(), o, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, r := range ghrepos {
			topics := r.Topics
			if err != nil {
				fmt.Println(err)
			}
			if len(topics) > 0 {
				repo := Repo{Name: r.GetName(), Topics: topics}
				repos = append(repos, repo)
			} else {
				noTopics = append(noTopics, r.GetName())
			}
		}
	}
	reposJson, err := json.Marshal(repos)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.WriteFile("repos.json", reposJson, os.ModeAppend)
	fmt.Println("This Repos have no TOPICS")
	for _, v := range noTopics {
		fmt.Println(v)
	}
}
