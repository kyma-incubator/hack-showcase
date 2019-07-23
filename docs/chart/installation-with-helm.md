---
title: 'Installation in Kyma using Helm'
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