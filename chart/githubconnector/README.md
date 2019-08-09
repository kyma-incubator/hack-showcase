# Github Connector <!-- omit in toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation in Kyma with Helm](#installing-in-kyma-using-helm)
	- [Prerequisites](#prerequisites-1)
	- [Steps](#steps)

## Overview
The chart creates a GitHub Connector Deployment, and a Namespace inside Kyma.
Moreover it creates a service, binds it to the newly created Namespace and exposes its API. Apart from that it creates an application that, for now, the user has to manually bind to the Namespace


## Installation in Kyma with Helm

### Prerequisites

To install GitHub Connector using Helm chart inside Kyma you have to:

- be connected to your Kyma
- have a properly configured chart

### Steps

1. Go to Kyma repository and run script `/installation/scripts/tiller-tls.sh` to get certificates needed for using helm commands. By default they are stored in `~/.helm`. After that add `--tls` flag to every Helm command to authorize and authenticate yourself
2. Install your chart running the command:

   ```shell
   helm install {HELM_CHART_DIRECTORY} --tls
   ```

>**NOTE:** To define Namespace in which chart should be installed add flag `--namespace`. You can also define name of your release with flag `--name`.

>**CAUTION:** Our application is in Beta version. For now you must specify those flags:

- `--name` - name of your release, 
- `--set container.image={VALUE}` - Docker image from which you want to build GitHub Connector
- `--set kymaAddress={VALUE}` - specify your Kyma address (for example `35.195.198.66.xip.io`)
  
  ``` shell
  helm install --set container.image={DOCKER-IMAGE} \
    --set kymaAddress={KYMA_ADDRESS} -n {NAME} \
    --namespace {NAMESPACE} . --tls
  ```
