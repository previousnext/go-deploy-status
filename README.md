Deploy Status
=======================

[![CircleCI](https://circleci.com/gh/previousnext/go-deploy-status.svg?style=svg)](https://circleci.com/gh/previousnext/go-deploy-status)

**Maintainer**: Kim Pepper

This is a brief description on what the project does.

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
