# Azure Comments Analytics Scenario

## Overview

Welcome to the Azure Comments Analytics addon!

This add-on allows you to install the scenario provided by Team Flying Seals. Azure Comments Analytics will receive information about Github's Issue from the Github Connector, next lambda function analyses it using the Azure Broker and then if the Issue's sentiment is too low, lambda sends it to Slack and labels the Issue on Github.

## Installation

1. Provision [Github Connector](https://github.com/kyma-incubator/hack-showcase/blob/master/docs/github-connector/README.md).
2. Provision [Slack Connector](https://github.com/kyma-incubator/hack-showcase/blob/master/docs/slack-connector/README.md).
3. Provision [Azure Broker](https://github.com/kyma-project/addons/tree/master/addons/azure-service-broker-0.0.1).
4. [Provision](#provisioning) this addon.

Now you can start using the Azure Comments Analytics. Add a new issue or edit one on a given repository or organization.

## Provisioning

### Default plan

In this plan you have to provide only necessary values.

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `githubURL` | GitHub repository | `string` | Link to GitHub repository in proper format: repos/:owner/:repo or orgs/:org | yes |
| `workspaceName` | Workspace Name | `string` | The name of workspace application will be installed to. | yes |
| `image` | Docker image | `string` | Docker image to install scenario from. | no |
