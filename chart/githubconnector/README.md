# Github Connector <!-- omit in toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation in Kyma with Helm](#installing-in-kyma-using-helm)
	- [Prerequisites](#prerequisites-1)
	- [Steps](#steps)

## Overview
The chart creates a GitHub Connector Deployment and a Namespace inside Kyma.
Moreover it creates a service, binds it to the newly created Namespace and exposes its API. Apart from that it creates an application.

## Installation in Kyma with Helm

### Prerequisites

To install GitHub Connector using Helm chart inside Kyma you have to:

- be connected to your Kyma
- have a properly configured chart

### Steps

1. Go to Kyma repository and run script `/installation/scripts/tiller-tls.sh` to get certificates needed for using helm commands. By default they are stored in `~/.helm`. After that add `--tls` flag to every Helm command to authorize and authenticate yourself
2. Install your chart running the command:
``` shell
  helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} -n {NAME} . --tls
  ```
  >**CAUTION:** Kyma address should be in the right format. It must consist of domain name, without dot  character at the beggining, for example `35.187.32.214.xip.io`

>**NOTE:** To define Namespace in which chart should be installed add flag `--namespace`.



