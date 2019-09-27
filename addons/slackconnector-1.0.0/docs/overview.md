# Overview

With the use of the token provided during the provisioning of the addon, the Slack Connector allows for sending requests to the Slack Web API. The requests are specified during the installation of the Slack Application to the workspace. The example requests are: posting a message to a specified channel, getting the list of the current users, etc.

#### Prerequisites

- Slack App with desired privileges installed to the destination workspace. See the tutorial provided by Slack on how to setup an application [here](https://api.slack.com/bot-users#getting-started).

    After creating the app:
    - The Slack Application's {SIGNING_SECRET} is used for validating requests coming from Slack by verifying its unique signature. Find it in the [Application's](https://api.slack.com/apps) **Settings** in the **Basic Information** section. [Learn more](https://api.slack.com/docs/verifying-requests-from-slack).
    - The Slack Application's {SLACK_TOKEN} is used for requests authorization. Find it in the [Application's](https://api.slack.com/apps) **Features** in the **OAuth & Permissions** section. Depending on the usecase it will be *OAuth Access Token* or *Bot User OAuth Access Token*. [Learn more](https://api.slack.com/docs/oauth).

#### Installation

1. [Provision](#provisioning) this addon.
2. In **Service Management**, go to **Catalog** and choose **Services**. Find the service named `slack-connector-{DESIRED_WORKSPACE_NAME}` and add it.

Now you can start using the Slack Connector. Add the channel name to the lambda environmental variables to send authorized requests to the Slack Web API.

## Provisioning

### Default plan

This plan allows to handle Events incoming from connected Slack workspaces to an exposed endpoint, and POST jsons to the Slack API through the Application Gateway, which automatically adds all the information necessary to communicate with Slack.

### Fields

| PARAMETER NAME | DISPLAY NAME | TYPE | DESCRIPTION | REQUIRED |
|----------------|--------------|------|-------------|:--------:|
| `image` | Docker Image | `string` | The Slack Connector image to use. | no |
| `slackBotToken` | Bot Token | `string` | {SLACK_TOKEN} | yes |
| `workspaceName` | Workspace Name | `string` | The name of the workspace to which to connect the Application. | yes |
| `slackSecret` | Slack Signing Secret | `string` | {SIGNING_SECRET} | yes |