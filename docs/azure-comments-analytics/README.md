# Github Issue Sentiment Analysis

## Overview

This addon allows you to install an example scenario. Azure Text Analytic receives information about Github's Issue from the Github Connector. Next, the lambda function analyses it using the Azure Broker and then, if the Issue's sentiment is negative, the lambda sends it to Slack and labels the Issue on Github.

## Installation

1. Provision [the Github Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/github-connector/README.md).
2. Provision [the Slack Connector](https://github.com/kyma-incubator/github-slack-connectors/blob/master/docs/slack-connector/README.md).
3. Provision [the Azure Broker](https://github.com/kyma-project/addons/tree/master/addons/azure-service-broker-0.0.1).
4. [Provision](#provisioning) this addon.

The Github Issue Sentiment Analysis Scenario is now ready to use. Add a new issue or edit one in a given repository or organization.

## Provisioning

### Default plan

In this plan, you only provide the required values.

| PARAMETER NAME  | DISPLAY NAME      | TYPE     | DESCRIPTION                                                                     | REQUIRED |
| --------------- | ----------------- | -------- | ------------------------------------------------------------------------------- | :------: |
| `githubURL`     | GitHub repository | `string` | Link to a GitHub repository in the proper format: repos/:owner/:repo or orgs/:org |   yes    |
| `workspaceName` | Workspace Name    | `string` | The name of the workspace to which to install the Application to.                        |   yes    |
