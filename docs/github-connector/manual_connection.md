# Connecting a new GitHub repository manually

## Overview

This document describes how to manually connect a GitHub repository to the existing GitHub Connector installed in the Kyma environment.

## Prerequisites

- Kyma with the GitHub Connector [installed](/docs/github-connector/installation.md)
- Connection to the Kyma Console

## Installation

1. Open the settings of the GitHub repository you want to connect to, go to the `Webhooks` page and click `Add webhook`.
2. On the configuration page, fill the field `Payload URL` with the exposed service URL and add `/webhook` at the end of the URL.
   
  > **NOTE:**  You can find the exposed service URL in the Kyma Console in your Namespace's APIs.

3. To get the Secret, which is required during the webhook setup, use this command:

```shell
kubectl get deployments -n {NAMESPACE} {DEPLOYMENT_NAME} -o jsonpath='{.spec.template.spec.containers[0].env[3].value}'
```

4. Set the other fields as follows:

    - **Content type**: `application/json`
    - **Secret**: {OBTAINED_SECRET}
    - **SSL verification**: `Disabled`

5. Select the Events to receive in the GitHub Connector.
6. Click `Add webhook`. This redirects you back to the webhooks page. You can see a new webhook in the list. A successful configuration results in a green tick next to the new webhook.
