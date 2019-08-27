# Configuring the GitHub Connector

## Overview

This document describes how to correctly connect GitHub repository to the GitHub Connector installed in Kyma environment. After completion you are able to handle events incoming from GitHub in lambdas.

## Prerequisites

- Kyma with the GitHub Connector [installed](/docs/github-connector/installation.md)
- Connection to Kyma

## Installation

1. Find the newly created GitHub Connector application and [bind it to the namespace](https://kyma-project.io/docs/components/application-connector/#tutorials-bind-an-application-to-a-namespace) of your choice.
2. Open the settings of the GitHub repository you want to connect to, go to `Webhooks` page and click `Add webhook`.
3. On the configuration page, fill the field `Payload URL` with exposed service URL (you can find it in `Kyma Console > {NAMESPACE} > APIs`) and add `/webhook` at the end of the URL.
4. Set other fields as follows:

    - **Content type**: `application/json`
    - **Secret**: `my-secret-key`
    - **SSL verification**: `Disabled`

    >**NOTE:** Secret is defined statically in code and SSL verification is disabled.

5. Select which events you want to receive in the GitHub Connector.
6. Click `Add webhook`. This redirects you back to the webhooks page. You can see a new webhook in the list. A successful configuration results in a green tick next to the new webhook.
