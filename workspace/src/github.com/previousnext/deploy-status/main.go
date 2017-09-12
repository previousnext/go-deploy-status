package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/google/go-github/github"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	token        = kingpin.Flag("token", "Your GitHub OAuth access token").String()
	deploymentId = kingpin.Arg("id", "The Deployment ID").String()
)

func main() {
	kingpin.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

}
