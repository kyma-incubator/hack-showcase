# GitHub Issue Sentiment Analysis

## Overview

Welcome to the GitHub Issue Sentiment Analysis Add-On!

This add-on allows you to install the exemplary scenario. GitHub Issue Sentiment Analysis receives information about GitHub's Issue from the GitHub Connector. Next, the lambda function analyses it using the Azure Text Analytics provided by Azure Broker. If the Issue's sentiment is negative, the lambda will send it to Slack and will label the Issue on GitHub.

## Installation

1. Provision [GitHub Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/github-connector/README.md).
2. Provision [Slack Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/slack-connector/README.md).
3. Provision [Azure Broker](https://github.com/kyma-project/addons/tree/master/addons/azure-service-broker-0.0.1).
4. [Provision](#provisioning) this Add-On.

The GitHub Issue Sentiment Analysis Scenario is now ready to use. Add a new issue or edit one in a given repository or organization.

## Provisioning

### Default plan

In this plan, you only provide the required values.

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `githubURL` | GitHub repository | `string` | Link to GitHub repository in the proper format: repos/:owner/:repo or orgs/:org | yes |
| `workspaceName` | Workspace Name | `string` | The name of the workspace to install the Application to. | yes |
| `channelName` | Slack channel name | `string` | The name of the slack channel where notifications will be sent to. | yes |
