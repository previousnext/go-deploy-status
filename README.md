Deploy Status
=======================

[![CircleCI](https://circleci.com/gh/previousnext/go-deploy-status.svg?style=svg&circle-token=f92a3dbaa70cbf7ee985559bd942a9c8207c11df)](https://circleci.com/gh/previousnext/go-deploy-status)

**Maintainer**: Kim Pepper

Provides a command line tool for creating GitHub deployments and changing deployment status.

## Usage

```
usage: deploy-status create <command> [<args> ...]

Create a new GitHub API object. See sub-commands for options.

Flags:
  --help         Show context-sensitive help (also try --help-long and --help-man).
  --owner=OWNER  The repository owner or organisation.
  --repo=REPO    The repository name
  --token=TOKEN  The GitHub OAuth access token

Subcommands:
  create deployment --ref=REF [<flags>]
    Create a new deployment

  create status --id=ID [<flags>]
    Create a deployment status for a deployment.
```

### Example

First create a deployment for the branch you are working on:

```
deploy-status create deployment \
    --token XXXXXXXXXXXXXXXXXXXXXXXX \
    --owner previousnext \
    --repo pnx-d8 \
    --ref deployments \ # The branch name
    --env pr \ # An environment name.
    --desc "Test deployment"
```

This will return the deployment ID. You will need this as the `--id` flag to set the status.

```
deploy-status create status \
    --token XXXXXXXXXXXXXXXXXXXXXXXX \
    --owner previousnext \
    --repo pnx-d8 \
    --state success \
    --id 1234567890 \
    --auto-inactive \
    --env-url http://pnx-d8-pr46.ci.pnx.com.au/
```

## Resources

* [Dave Cheney - Reproducible Builds](https://www.youtube.com/watch?v=c3dW80eO88I)

## Development

### Principles

* Code lives in the `workspace` directory

### Tools

* **Dependency management** - https://getgb.io
* **Build** - https://github.com/mitchellh/gox
* **Linting** - https://github.com/golang/lint

### Workflow

(While in the `workspace` directory)

**Installing a new dependency**

```bash
gb vendor fetch github.com/foo/bar
```

**Running quality checks**

```bash
make lint test
```

**Building binaries**

```bash
make build
```
