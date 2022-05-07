package git

import (
	"context"
	"fmt"
	"log"
)

func updateBranches(conf *Config) {
	ctx := context.Background()
	client := AuthGithubAPI(ctx, conf.Token)
	openPr, _, err := client.PullRequests.List(ctx, conf.Owner, conf.Repo, nil)
	if err != nil {
		log.Println(err)
	}
	for _, number := range openPr {
		upBranch, _, err := client.PullRequests.UpdateBranch(ctx, conf.Owner, conf.Repo, number.GetNumber(), nil)
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println(upBranch.GetMessage())
	}
}

//TODO: Add a workflow to approve pr
func checkWorkflows(conf *Config) bool {
	completed := false
	ctx := context.Background()
	client := AuthGithubAPI(ctx, conf.Token)
	wfList, _, err := client.Actions.ListWorkflows(ctx, conf.Owner, conf.Repo, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if !completed {
		for _, wf := range wfList.Workflows {
			wfStatus, _, err := client.Actions.GetWorkflowRunByID(ctx, conf.Owner, conf.Repo, wf.GetID())
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

func approvePr(conf *Config) {
	ctx := context.Background()
	client := AuthGithubAPI(ctx, conf.Token)
	pr, _, err := client.PullRequests.List(ctx, conf.Owner, conf.Repo, nil)
	if err != nil {
		log.Println(err)
	}
	for _, v := range pr {
		result, _, err := client.PullRequests.Merge(ctx, conf.Owner, conf.Repo, v.GetNumber(), "Approved Pr", nil)
		if !result.GetMerged() || err != nil {
			log.Println(err)
		}
		fmt.Println(result.GetMessage())
	}
}

func MergeCheck(conf *Config) {
	// updateBranches(conf)
	approvePr(conf)
}
