Deploy Status
=======================

[![CircleCI](https://circleci.com/gh/previousnext/go-deploy-status.svg?style=svg&circle-token=f92a3dbaa70cbf7ee985559bd942a9c8207c11df)](https://circleci.com/gh/previousnext/go-deploy-status)

**Maintainer**: Kim Pepper

Provides a command line tool for creating GitHub deployments and changing deployment status.

## Usage

```
usage: deploy-status --owner=OWNER --repo=REPO --token=TOKEN [<flags>] <command> [<args> ...]

A command-line tool for interacting with GitHub deployment API

Flags:
  --help         Show context-sensitive help (also try --help-long and --help-man).
  --owner=OWNER  The repository owner
  --repo=REPO    The repository name
  --token=TOKEN  Your GitHub OAuth access token

Commands:
  help [<command>...]
    Show help.

  deployment --ref=REF [<flags>]
    Create a new deployment

  status --id=ID [<flags>] [<state>]
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
