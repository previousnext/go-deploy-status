package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/google/go-github/github"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	token        = kingpin.Arg("token", "Your GitHub OAuth access token").String()
	deploymentID = kingpin.Arg("id", "The Deployment ID").String()
	ref          = kingpin.Arg("ref", "The Git reference.").String()
	desc         = kingpin.Arg("desc", "The description").String()
	owner        = kingpin.Arg("owner", "The repository owner").String()
	repo         = kingpin.Arg("repo", "The repository name").String()
	state        = kingpin.Arg("state", "The Deployment state to set").Default("pending").String()
)

func main() {
	kingpin.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	deploymentRequest := &github.DeploymentRequest{
		Ref:         github.String(*ref),
		Description: github.String(*desc),
	}

	deployment, _, err := client.Repositories.CreateDeployment(ctx, *owner, *repo, deploymentRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println(deployment.ID)

	statusRequest := &github.DeploymentStatusRequest{}
	status, _, err := client.Repositories.CreateDeploymentStatus(ctx, *owner, *repo, *deployment.ID, statusRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println(status.State)
}
