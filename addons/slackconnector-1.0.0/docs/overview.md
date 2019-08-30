# Overview

Welcome to the Slack Connector addon!

With the use of token provided during provision of addon, Slack Connector allows sending requests to Slack
Web API, that were specified at the Slack Application installation to the workspace, such as posting a
message to specified channel, getting list of current users, etc.

## Installation

1. [Provision](#provisioning) this addon.
2. Go to `Service Management > Catalog > Services`. Find a service named `slack-connector-{WORKSPACE-NAME}` and add it.
3. Done

## Provisioning

### Minimal plan

| PARAMETER NAME 	| DISPLAY NAME 		| TYPE 		| DESCRIPTION 	| REQUIRED 	|
|-----------------	|-----------------	|---------	|------------	|:---------:|
| `kymaAddress` 	| Kyma Address 	| `string` 	| Kyma domain address. 	| yes 	|
| `slackBotToken` 	| Bot Token 	| `string` 	| The Slack workspace token, which you can find on this site: https://auth-slack.herokuapp.com/ 	| yes 	|
| `workspaceName` 	| Workspace Name 	| `string` 	| The name of workspace application will be installed to. 	| yes 	|
