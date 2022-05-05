package main

import (
	"os"

	"github.com/jarvis-action/pkg/git"
)

func main() {
	git.ApprovePr(os.Getenv("GITHUB_TOKEN"), os.Getenv("OWNER"), os.Getenv("GITHUB_REPOSITORY"))
}
