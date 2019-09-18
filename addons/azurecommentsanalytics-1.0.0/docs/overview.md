# Overview

Welcome to the Azure Comments Analytics addon!

This addon allows you to install the scenario provided by Team Flying Seals. Azure Comments Analytics will receive information about Github's Issue from Github Connector, next analyze it using Azure Broker and then if Issue sentiment is too low, sent it to Slack and label Issue on Github.

## Installation

1. Provision [Github Connector](https://github.com/kyma-incubator/hack-showcase/blob/master/docs/github-connector/README.md).
2. Provision [Slack Connector](https://github.com/kyma-incubator/hack-showcase/blob/master/docs/slack-connector/README.md).
3. Provision [Azure Bloker](https://github.com/kyma-project/addons/tree/master/addons/azure-service-broker-0.0.1).
4. Wait for end of provisioning.
5. [Provision](#provisioning) this addon.

Now you can start using the Azure Comments Analytics. Add new or edit Issue on given repo or org.

## Provisioning

### Default plan

In this plan you have to provide only necessary values.

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `githubURL` | GitHub repository | `string` | Link to GitHub repository in proper format: repos/:owner/:repo or orgs/:org | yes |
| `workspaceName` | Workspace Name | `string` | The name of workspace application will be installed to. | yes |
