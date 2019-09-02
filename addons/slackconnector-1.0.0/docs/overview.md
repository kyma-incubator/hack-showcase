# Overview

Welcome to the Slack Connector addon!

With the use of token provided during provision of addon, Slack Connector allows sending requests to Slack
Web API, that were specified at the Slack Application installation to the workspace, such as posting a
message to specified channel, getting list of current users, etc.

## Installation

1. [Provision](#provisioning) this addon.
2. Go to `Service Management > Catalog > Services`. Find a service named `slack-connector-{WORKSPACE-NAME}` and add it.

Now you can start using the Slack Connector. Add channel ID to lambda environmental variables to send authorized request to Slack Web API.

## Provisioning

### Default plan

In this plan you have to provide only only necessary values.

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `slackBotToken` | Bot Token | `string` | The Slack workspace token, which you can find on this site: <https://auth-slack.herokuapp.com/> | yes |
| `workspaceName` | Workspace Name | `string` | The name of workspace application will be installed to. | yes |

### Dev plan

In this plan you can provide your own image of application, e.g. for testing purposes.

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `image` | Docker Image | `string` | Specify the Slack Connector image you want to use. | no |
| `slackBotToken` | Bot Token | `string` | The Slack workspace token, which you can find on this site: <https://auth-slack.herokuapp.com/> | yes |
| `workspaceName` | Workspace Name | `string` | The name of workspace application will be installed to. | yes |
