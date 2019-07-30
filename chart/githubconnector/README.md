# github-connector

## Overview
The chart creates a github connector deployment, and a namespace inside kyma.
Moreover it creates a service, binds it to the newly created namespace and exposes its API. Apart from that it creates an application that for now, the user has to manually bind to the namespace

## Prerequisites

In order to install the chart inside of kyma you need to:
* be connected to your kyma instance
* have a properly configured chart

## Details

For now, the chart cannot be configured at all, and has to be installed with a preset that can be found in the [installation tutorial](../../docs/githubconnector/helm-installation-tutorial.md)
