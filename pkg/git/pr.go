package git

import (
	"context"
	"fmt"
	"log"
)

func updateBranches(token string, owner string, repos string) {
	ctx := context.Background()
	client := AuthGithubAPI(ctx, token)
	openPr, _, err := client.PullRequests.List(ctx, owner, repos, nil)
	if err != nil {
		log.Println(err)
	}
	for _, number := range openPr {
		upBranch, _, err := client.PullRequests.UpdateBranch(ctx, owner, repos, number.GetNumber(), nil)
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println(upBranch.GetMessage())
	}
}

//TODO: Add a workflow to approve pr
func checkWorkflows(token string, owner string, repos string) bool {
	completed := false
	ctx := context.Background()
	client := AuthGithubAPI(ctx, token)
	wfList, _, err := client.Actions.ListWorkflows(ctx, owner, repos, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if !completed {
		for _, wf := range wfList.Workflows {
			wfStatus, _, err := client.Actions.GetWorkflowRunByID(ctx, owner, repos, wf.GetID())
			if err != nil {
				log.Fatalf(err.Error())
			}
			if wfStatus.GetStatus() != "completed" {
				log.Println("Workflow is not completed")
				completed = false
			}
		}
	}
	return completed
}

func ApprovePr(token string, owner string, repos string) {
	ctx := context.Background()
	client := AuthGithubAPI(ctx, token)
	pr, _, err := client.PullRequests.List(ctx, owner, repos, nil)
	if err != nil {
		log.Println(err)
	}
	updateBranches(token, owner, repos)
	for _, v := range pr {
		result, _, err := client.PullRequests.Merge(ctx, owner, repos, v.GetNumber(), "Approved Pr", nil)
		if !result.GetMerged() || err != nil {
			log.Println(err)
		}
		fmt.Println(result.GetMessage())
	}
}
