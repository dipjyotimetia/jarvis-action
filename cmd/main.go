package main

import (
	"os"

	"github.com/jarvis-action/pkg/git"
)

func main() {
	conf := &git.Config{
		Token: os.Getenv("GITHUB_TOKEN"),
		Owner: os.Getenv("OWNER"),
		Repo:  os.Getenv("GITHUB_REPOSITORY"),
	}

	git.MergeCheck(conf)
}
