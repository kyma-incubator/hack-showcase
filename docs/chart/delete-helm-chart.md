---
title: 'Uninstalling a chart from Kyma'
disqus: flying-seals
---

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
