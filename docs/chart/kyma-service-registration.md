---
title: 'Registering service in Kyma'
disqus: flying-seals
---

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
