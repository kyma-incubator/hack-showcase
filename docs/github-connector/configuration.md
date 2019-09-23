# Configuring GitHub Connector

## Overview

This document describes how to manually connect GitHub repository to existing GitHub Connector installed in Kyma enviroment. After completion you are able to handle events incoming from GitHub in lambdas.


## Prerequisites

- Kyma with the GitHub Connector [installed](/docs/github-connector/installation.md)
- Connection to Kyma Console

## Installation

1. Find the newly created GitHub Connector application and [bind it to the namespace](https://kyma-project.io/docs/components/application-connector/#tutorials-bind-an-application-to-a-namespace) of your choice.
2. Open the settings of the GitHub repository you want to connect to, go to `Webhooks` page and click `Add webhook`.
3. On the configuration page, fill the field `Payload URL` with exposed service URL (you can find it in `Kyma Console > {NAMESPACE} > APIs`) and add `/webhook` at the end of the URL.
4. To get Secret which is required during webhook setup, you need to use this comand:

```shell
kubectl get deployments -n {NAMESPACE} {DEPLOYMENT_NAME} -o jsonpath='{.spec.template.spec.containers[0].env[3].value}'
```

5. Set other fields as follows:

    - **Content type**: `application/json`
    - **Secret**: previously obtained secret (nie wiem co z tym zrobic)
    - **SSL verification**: `Disabled`

    >**NOTE:** If your Kyma cluster is not SSL secured, SSL verification will not be working.

6. Select which events you want to receive in the GitHub Connector.
7. Click `Add webhook`. This redirects you back to the webhooks page. You can see a new webhook in the list. A successful configuration results in a green tick next to the new webhook.
