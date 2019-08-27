# Slack Connector Installation<!-- omit in toc -->

- [Overview](#overview)
- [Installation in Kyma with Helm](#installation-in-kyma-with-helm)
  - [Prerequisites](#prerequisites)
  - [Steps](#steps)

## Overview

Slack Connector is a component which allows for communication between the Kyma environment and the Slack API.

## Installation in Kyma with Helm

### Prerequisites

- Connection to Kyma cluster
- Slack Connector Docker image

### Steps

1. Go to the [authentication page](https://auth-slack.herokuapp.com/). Click `Add to Slack`. This redirects you to another page. Select your desired workspace and click `Allow`.
    >**NOTE:** If the link does not work, it means that you have to create an application like that yourself. To create an application, see [this tutorial](https://api.slack.com/docs/oauth#flow) in the Slack API documentation.

2. Copy the authentication token. You will need it later in the Helm command.
3. Go to [Kyma repository](https://github.com/kyma-project/kyma) and run script `/installation/scripts/tiller-tls.sh` to get certificates needed for using Helm commands. By default they are stored in `~/.helm` directory. After that add `--tls` flag to every Helm command to authorize and authenticate yourself
4. Go to the `chart/slackconnector` directory. Run this command to install Slack Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} --set slackBotToken={SLACK_TOKEN} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Kyma address should be in the right format. It must consist of domain name, without dot  character at the beggining, for example `35.187.32.214.xip.io`
    >**NOTE:** To define Namespace in which chart should be installed add flag `--namespace`.
