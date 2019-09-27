# Overview

With the use of the token provided during the provisioning of the addon, the Slack Connector allows for sending requests to the Slack Web API. The requests are specified during the installation of the Slack Application to the workspace. The example requests are: posting a message to a specified channel, getting the list of the current users, etc.

## Installation

1. [Provision](#provisioning) this addon.
2. In **Service Management**, go to **Catalog** and choose **Services**. Find the service named `slack-connector-{DESIRED_WORKSPACE_NAME}` and add it.

Now you can start using the Slack Connector. Add channel name to lambda environmental variables to send authorized requests to the Slack Web API.

## Provisioning

### Default plan

This plan allows to handle events incoming from connected Slack workspaces to an exposed endpoint, and POST jsons to the Slack API through the Application Gateway, which automatically adds all the information necessary to communicate with Slack.

### Fields

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `image` | Docker Image | `string` | The Slack Connector image to use. | no |
| `slackBotToken` | Bot Token | `string` | DO NAPISANIA | yes |
| `workspaceName` | Workspace Name | `string` | The name of the workspace to which to connect the application. | yes |
| `slackSecret` | Slack Signing Secret | `string` | The Slack Application's signing Secret used for validating requests coming from Slack by verifying its unique signature. To get this Secret, create your Slack Application. Then, find the Secret in the [Application's](https://api.slack.com/apps) Basic Information. | yes |
