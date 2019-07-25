---
title: 'Using GitHub Connector chart'
disqus: flying-seals
---

Installation in Kyma using Helm
===


## Requirements
To install Helm chart inside Kyma you have to:
- be connected to your Kyma
- have properly configured chart

## Steps
1. Go to Kyma repository and run script /installation/scripts/tiller-tls.sh to get certificates needed for using helm commands. By default they are stored in ~/.helm. After that add flag --tls to every helm command to authorize and authenticate yourself
2. Install your chart with command:
``` 
helm install {HELM_CHART_DIRECTORY} --tls 
```
**NOTE:** To define namespace in which chart should be installed add flag `--namespace`. You can also define name of your release with flag `--name`.

**ATTENTION:** Our application is in Beta version. For now you HAVE TO specify those flags:
* --name - it has to be "github-connector"
* --set container.image={VALUE} - specify it if you have newer version of docker image than karoljaksik/github-connector:1.0.2
* --set kymaAddress={VALUE} - specify your kyma adddress (for example 35.195.198.66.xip.io)
```
helm install --set container.image=karoljaksik/github-connector:1.0.2 --set kymaAddress=35.195.198.66.xip.io -n github-connector --namespace flying-seals . --tls
```


Uninstalling a chart from Kyma
===


## Requirements
To register a service inside Kyma you have to:
- be connected to your Kyma
- have helm certificate from Kyma (check *Installation with Kyma using Helm*)

## Steps
1. List your Helm charts using
```
helm list --tls
```
and find name of chart you want to delete. Copy it or memorize.

2. Use command below to delete it
```
helm delete {NAME} --purge --tls
```



Registering service in Kyma
===


## Requirements
To register a service inside Kyma you have to:
- be connected to your Kyma
- have registered an application in Kyma

## Steps
1. Create a json covering [this POST schema](https://github.com/kyma-project/kyma/blob/master/components/application-registry/docs/api/api.yaml), like in example below:
```
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
2. Send it as POST to `application-registry-external-api.kyma-integration.svc.cluster.local:8081/{APP-NAME}/v1/metadata/services` from inside of Kyma. There is couple of ways to do so, e.g.:
    * by connecting to pod you want to send it from (e.g. with curl) using command
        ```
        k exec -n {NAMESPACE} {POD-NAME} -c {CONTAINER-NAME} -it -- sh”
        ```
