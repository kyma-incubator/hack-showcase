# Overview

Welcome to the GitHub Connector addon!

This addon allows you to communicate with GitHub. You can handle events incoming from GitHub repositories or manage repositories through GitHub API. You must provision an instance for every repository you want to communicate with.

## Installation

1. [Provision](#provisioning) this addon.
2. Go to `Service Management > Catalog > Services`. Find a service named `github-{REPOSITORY-NAME}` and add it.

Now you can start using the GitHub Connector. Add new event trigger to react to chosen GitHub notifications or bind this service in lambda to send authorized request to GitHub API.

## Provisioning

### Minimal plan

In this plan you have to provide only only neccessary values.

| PARAMETER NAME | DISPALAY NAME     | TYPE   | DESCRIPTION                                                                                                | REQUIRED |
|----------------|-------------------|--------|------------------------------------------------------------------------------------------------------------|----------|
| githubToken    | Token             | string | The GitHub repository token, which you can find on this site: https://auth-github-connector.herokuapp.com/ | yes      |
| githubURL      | GitHub repository | string | Link to GitHub repository in proper format: repos/:owner/:repo or orgs/:org                                | yes      |
| kymaAddress    | Kyma address      | string | Kyma domain address                                                                                        | yes      |