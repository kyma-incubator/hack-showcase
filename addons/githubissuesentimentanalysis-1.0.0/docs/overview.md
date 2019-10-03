# Github Issue Sentiment Analysis

## Overview

Welcome to the Github Issue Sentiment Analysis addon!

This add-on allows you to install the scenario provided by Team Flying Seals. Github Issue Sentiment Analysis receives information about Github's Issue from the Github Connector. Next, the lambda function analyses it using the Azure Broker and then, if the Issue's sentiment is negative, the lambda sends it to Slack and labels the Issue on Github.

## Installation

1. Provision [Github Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/github-connector/README.md).
2. Provision [Slack Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/slack-connector/README.md).
3. Provision [Azure Broker](https://github.com/kyma-project/addons/tree/master/addons/azure-service-broker-0.0.1).
4. [Provision](#provisioning) this addon.

The Github Issue Sentiment Analysis Scenario is now ready to use. Add a new issue or edit one in a given repository or organization.

## Provisioning

### Default plan

In this plan, you only provide the required values.

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `githubURL` | GitHub repository | `string` | Link to GitHub repository in the proper format: repos/:owner/:repo or orgs/:org | yes |
| `workspaceName` | Workspace Name | `string` | The name of the workspace to install the Application to. | yes |
| `channelName` | Slack channel name | `string` | The name of the slack channel where notifications will be sent to. | yes |
