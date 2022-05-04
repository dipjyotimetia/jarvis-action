package git

import (
	"context"
	"fmt"
	"log"
)

func ApprovePr(token string, owner string, repos string) {
	ctx := context.Background()
	client := AuthGithubAPI(ctx, token)
	pr, _, err := client.PullRequests.List(ctx, owner, repos, nil)
	if err != nil {
		log.Println(err)
	}
	for _, v := range pr {
		result, _, err := client.PullRequests.Merge(ctx, owner, repos, v.GetNumber(), "Approved Pr", nil)
		if !result.GetMerged() || err != nil {
			log.Println(err)
		}
		fmt.Println(result.GetMessage())
	}
}
