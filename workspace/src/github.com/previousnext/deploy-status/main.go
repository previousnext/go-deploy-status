package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/google/go-github/github"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	app   = kingpin.New("deploy-status", "A command-line tool for interacting with GitHub deployment API")
	owner = app.Flag("owner", "The repository owner").Required().String()
	repo  = app.Flag("repo", "The repository name").Required().String()
	token = app.Flag("token", "Your GitHub OAuth access token").Required().String()

	deployment = app.Command("deployment", "Create a new deployment")
	ref        = deployment.Flag("ref", "The Git reference. Can be a branch, tag or commit ID.").Required().String()
	desc       = deployment.Flag("desc", "The description").String()
	autoMerge  = deployment.Flag("auto-merge", "Auto merge the default branch into the requested ref if it is behind the default branch.").Bool()
	env        = deployment.Flag("env", "The environment").Default("dev").String()

	status       = app.Command("status", "")
	state        = status.Arg("state", "The Deployment state to set").Default("pending").String()
	autoInactive = status.Flag("auto-inactive", "Add a new inactive status to all non-transient, non-production environment deployments with the same repository and environment name as the created status's deployment.").Bool()
	deploymentID = status.Flag("id", "The Deployment ID").Required().Int()
	envURL       = status.Flag("env-url", "The environment URL").String()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case deployment.FullCommand():
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: *token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)
		deploymentCommand(ctx, *client)

	case status.FullCommand():
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: *token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		statusCommand(ctx, *client)
	}
}

func deploymentCommand(ctx context.Context, client github.Client) {

	deploymentRequest := &github.DeploymentRequest{
		Ref:         github.String(*ref),
		Description: github.String(*desc),
		AutoMerge:   github.Bool(*autoMerge),
		Environment: github.String(*env),
	}

	deployment, _, err := client.Repositories.CreateDeployment(ctx, *owner, *repo, deploymentRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println(deployment.ID)
}

func statusCommand(ctx context.Context, client github.Client) {

	statusRequest := &github.DeploymentStatusRequest{
		AutoInactive:   github.Bool(*autoInactive),
		EnvironmentURL: github.String(*envURL),
		State:          github.String(*state),
	}
	status, _, err := client.Repositories.CreateDeploymentStatus(ctx, *owner, *repo, *deploymentID, statusRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println(status.State)
}
