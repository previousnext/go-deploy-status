package create

import (
	"context"
	"github.com/alecthomas/kingpin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// CommonOptions is the common create command options.
type CommonOptions struct {
	Owner       string
	Repo        string
	Token       string
	Description string
}

// NewCommonOptions creates new create options.
func NewCommonOptions(cmd *kingpin.CmdClause) *CommonOptions {
	options := &CommonOptions{}
	cmd.Flag("owner", "The repository Owner or organisation.").Envar("CIRCLE_PROJECT_USERNAME").Required().StringVar(&options.Owner)
	cmd.Flag("repo", "The repository name").Envar("CIRCLE_PROJECT_REPONAME").Required().StringVar(&options.Repo)
	cmd.Flag("token", "The GitHub OAuth access Token").Envar("GITHUB_TOKEN").Required().StringVar(&options.Token)
	cmd.Flag("desc", "The Description").StringVar(&options.Description)
	return options
}

// NewGithubClient creates a new GitHub client.
func NewGithubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
