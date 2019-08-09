# Github Connector <!-- omit in toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installing in Kyma using Helm](#installing-in-kyma-using-helm)
	- [Prerequisites](#prerequisites-1)
	- [Steps](#steps)
- [Registering service in Kyma](#registering-service-in-kyma)
	- [Prerequisites](#prerequisites-3)
	- [Steps](#steps-2)


## Overview
The chart creates a github connector deployment, and a namespace inside kyma.
Moreover it creates a service, binds it to the newly created namespace and exposes its API. Apart from that it creates an application that for now, the user has to manually bind to the namespace

## Prerequisites

In order to install the chart inside of kyma you need to:
* be connected to your kyma instance
* have a properly configured chart

## Installing in Kyma using Helm

### Prerequisites

To install Helm chart inside Kyma you have to:

- be connected to your Kyma
- have a properly configured chart

### Steps

1. Go to Kyma repository and run script /installation/scripts/tiller-tls.sh to get certificates needed for using helm commands. By default they are stored in ~/.helm. After that add --tls flag to every helm command to authorize and authenticate yourself
2. Install your chart with command:

   ```shell
   helm install {HELM_CHART_DIRECTORY} --tls
   ```

**NOTE:** To define namespace in which chart should be installed add flag `--namespace`. You can also define name of your release with flag `--name`.

**ATTENTION:** Our application is in Beta version. For now you HAVE TO specify those flags:

- --name - it has to be "github-connector"
- --set container.image={VALUE} - specify it if you have newer version of docker image than karoljaksik/github-connector:1.0.2 --set kymaAddress={VALUE} - specify your kyma address (for example 35.195.198.66.xip.io)
  
  ``` shell
  helm install --set container.image=karoljaksik/github-connector:1.0.2 \
    --set kymaAddress=35.195.198.66.xip.io -n github-connector \
    --namespace flying-seals . --tls
    ```


## Registering service in Kyma

### Prerequisites
To register a service inside Kyma you have to:

- be connected to your Kyma
- have registered an application in Kyma

### Steps

1. Create a json covering [this POST schema](https://github.com/kyma-project/kyma/blob/master/components/application-registry/docs/api/api.yaml), like in example below:

   ```json
   {
     "provider": "kyma",
     "name": "webhook-app",
     "description": "Boilerplate for GitHub connector",
     "api": {
       "targetUrl": "https://console.35.233.90.87.xip.io/github-api"
      },
     "events": {
       "spec": {}
     }
   }
   ```

2. Send it as POST to `application-registry-external-api.kyma-integration.svc.cluster.local:8081/{APP-NAME}/v1/metadata/services` from inside of Kyma. There is couple of ways to do so, for example:

     - by connecting to pod you want to send it from (for example with curl) using command

      ```shell
      k exec -n {NAMESPACE} {POD-NAME} -c {CONTAINER-NAME} -it -- sh”
      ```
