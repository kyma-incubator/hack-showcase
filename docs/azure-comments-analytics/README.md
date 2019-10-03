# GitHub Issue Sentiment Analysis

## Overview

This Add-On allows you to install an example scenario. Azure Text Analytic receives information about GitHub's Issue from the GitHub Connector. Next, the lambda function analyses it using the Azure Text Analytics provided by Azure Broker. If the Issue's sentiment is negative, the lambda will send it to Slack and will label the Issue on GitHub.

## Installation

1. Provision [the GitHub Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/github-connector/README.md).
2. Provision [the Slack Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/slack-connector/README.md).
3. Provision [the Azure Broker](https://github.com/kyma-project/addons/tree/master/addons/azure-service-broker-0.0.1).
4. [Provision](#provisioning) this Add-On.

The GitHub Issue Sentiment Analysis Scenario is now ready to use. Add a new issue or edit one in a given repository or organization.

## Provisioning

### Default plan

In this plan, you only provide the required values.

| PARAMETER NAME  | DISPLAY NAME      | TYPE     | DESCRIPTION                                                                     | REQUIRED |
| --------------- | ----------------- | -------- | ------------------------------------------------------------------------------- | :------: |
| `githubURL`     | GitHub repository | `string` | Link to a GitHub repository in the proper format: repos/:owner/:repo or orgs/:org |   yes    |
| `workspaceName` | Workspace Name    | `string` | The name of the workspace to which to install the Application to.                        |   yes    |
