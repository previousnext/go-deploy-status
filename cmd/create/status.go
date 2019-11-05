package create

import (
	"context"
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/google/go-github/github"
)

type cmdStatus struct {
	autoInactive bool
	envURL       string
	state        string
	desc         string
	deploymentID int64
	options      *CommonOptions
}

func (v *cmdStatus) run(c *kingpin.ParseContext) error {
	statusRequest := &github.DeploymentStatusRequest{
		AutoInactive:   github.Bool(v.autoInactive),
		EnvironmentURL: github.String(v.envURL),
		State:          github.String(v.state),
		Description:    github.String(v.desc),
	}

	ctx := context.Background()
	client := NewGithubClient(ctx, v.options.Token)
	status, _, err := client.Repositories.CreateDeploymentStatus(ctx, v.options.Owner, v.options.Repo, v.deploymentID, statusRequest)
	if err != nil {
		return err
	}

	fmt.Println(status.GetState())

	return nil
}

// Status is the create deployment status command.
func Status(c *kingpin.CmdClause) {
	v := cmdStatus{}
	command := c.Command("status", "Create a deployment status for a deployment.").Action(v.run)
	command.Flag("state", "The Deployment state to set").Default("pending").StringVar(&v.state)
	command.Flag("auto-inactive", "Add a new inactive status to all non-transient, non-production environment deployments with the same repository and environment name as the created status's deployment.").BoolVar(&v.autoInactive)
	command.Flag("id", "The Deployment ID").Required().Int64Var(&v.deploymentID)
	command.Flag("env-url", "The environment URL").StringVar(&v.envURL)
	options := NewCommonOptions(command)
	v.options = options
}
