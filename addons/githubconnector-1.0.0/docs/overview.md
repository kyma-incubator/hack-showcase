# Overview

Welcome to the GitHub Connector addon!

This addon allows you to communicate with GitHub. You can handle events incoming from GitHub repositories or manage repositories through GitHub API. You must provision an instance for every repository you want to communicate with.

## Installation

1. Provision this addon. Plans' and fields' meaning is explained below.
2. Go to `Service Management > Catalog > Services`. Find a service named `github-{REPOSITORY-NAME}` and add it.

Now you can start using the GitHub Connector. Add new event trigger to react to chosen GitHub notifications or bind this service in lambda to send authorized request to GitHub API.

## Provisioning

### Default plan

This plan allows to both handle events incoming from connected GitHub repositories to an exposed endpoint and POST jsons to GitHub API through Application Gateway, which automatically adds all necessary informations needed to communicate with GitHub.

### Fields

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `githubToken` | Token | `string` | The GitHub repository token, which you can find on this site: <https://auth-github-connector.herokuapp.com/> | yes |
| `githubEndpoint` | GitHub Endpoint (org or repo) | `string` | Link to GitHub repository in proper format: repos/{OWNER}/{REPO} or orgs/{ORG}. For example, "repos/kyma-incubator/hack-showcase". | yes |
| `kymaAddress` | Kyma Domain name | `string` | Kyma domain address in proper format. For example, "domain.sap.com". | yes |
| `image` | Docker image | `string` | The GitHub Connector image on DockerHub | no |
