package create

import (
	"context"
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/google/go-github/github"
)

type cmdDeployment struct {
	ref       string
	autoMerge bool
	env       string
	contexts  []string
	options   *CommonOptions
}

func (v *cmdDeployment) run(c *kingpin.ParseContext) error {
	requiredContexts := make([]string, 0, 5)
	if len(v.contexts) > 0 {
		requiredContexts = v.contexts
	}

	deploymentRequest := &github.DeploymentRequest{
		Ref:              github.String(v.ref),
		Description:      github.String(v.options.Description),
		AutoMerge:        github.Bool(v.autoMerge),
		Environment:      github.String(v.env),
		RequiredContexts: &requiredContexts,
	}
	ctx := context.Background()
	client := NewGithubClient(ctx, v.options.Token)
	deployment, _, err := client.Repositories.CreateDeployment(ctx, v.options.Owner, v.options.Repo, deploymentRequest)
	if err != nil {
		return err
	}

	fmt.Printf(fmt.Sprintf("%d", deployment.GetID()))

	return nil
}

// Deployment is the create deployment command.
func Deployment(c *kingpin.CmdClause) {
	v := cmdDeployment{}
	command := c.Command("deployment", "Create a new deployment").Action(v.run)
	command.Flag("ref", "The Git reference. Can be a branch, tag or commit ID.").Envar("CIRCLE_BRANCH").Required().StringVar(&v.ref)
	command.Flag("auto-merge", "Auto merge the default branch into the requested ref if it is behind the default branch.").BoolVar(&v.autoMerge)
	command.Flag("env", "The environment").Default("dev").StringVar(&v.env)
	command.Flag("contexts", "The required contexts").StringsVar(&v.contexts)
	options := NewCommonOptions(command)
	v.options = options
}
