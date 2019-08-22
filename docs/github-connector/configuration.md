# Configuring GitHub Connector

## Overview

This document describes how to correctly connect GitHub repository to the GitHub Connector installed in Kyma environment. After completion you will be able to handle events incoming from GitHub in lambdas.

## Prerequisites

- Kyma with the GitHub Connector [installed](/docs/github-connector/installation.md)
- Connection to Kyma

## Installation

1. Find the newly created github-connector application and [bind it to the namespace](https://kyma-project.io/docs/components/application-connector/#tutorials-bind-an-application-to-a-namespace) of your choice.
2. Open settings of GitHub repository you want to connect to, go to `Webhooks` page and click `Add webhook`.
3. On the configuration page, fill the field `Payload URL` with exposed service URL (you can find it in `Kyma Console > {NAMESPACE} > APIs`) and add `/webhook` at the end of it.
4. Set other fields as follows:

    - **Content type**: `application/json`
    - **Secret**: `my-secret-key`
    - **SSL verification**: `Disabled`

    >**NOTE:** Secret is temporarily defined statically in code and SSL verification is disabled. It will be changed in the future.

5. Select which events you would like to receive in GitHub Connector.
6. Click `Add webhook`. You will be redirected back to webhooks page. There will be new webhook on list. Successful configuration results with green tick next to it.
