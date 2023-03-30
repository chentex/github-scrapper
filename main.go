package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
)

func main() {
	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	repos, _, err := client.Repositories.List(context.Background(), "cloud-bulldozer", nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, r := range repos {
		tags, _, err := client.Repositories.ListAllTopics(context.Background(), "cloud-bulldozer", *r.Name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*r.Name)
		for _, v := range tags {
			fmt.Printf("\t%s", v)
		}
	}
}
