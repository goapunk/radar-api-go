# `radar-api-go`

[![Go](https://github.com/goapunk/radar-api-go/actions/workflows/go.yml/badge.svg)](https://github.com/goapunk/radar-api-go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/goapunk/radar-api-go.svg)](https://pkg.go.dev/github.com/goapunk/radar-api-go)

Golang bindings for the [Radar](https://radar.squat.net/en) API.

## Development
Commits should be based on [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

### Tests
Run with:

```
go test
```

## Release
Release are created using commitizen. To create a new release:
1. create a python venv if not existing: `python3 -m venv venv`
2. activate the venv: `source venv/bin/activate`
3. install commitizen: `pip install commitizen`
4. run `cz bump`
5. push to a new branch and make pr
6. once merged, delete the tag created by commitizen, tag the release commit on main and push it
