Deploying GitHub Connector in Kyma 
=====
## Overview
This document describes how to correctly setup the GitHub connector on Kyma, having provided that the user is logged into GCP cluster with Kyma installed.

## Prerequisites
Installation guide can be found [`here`] (chart/github-connector/README.md)
For the purpose of this README we assume that you've already done all of the steps mentioned in that doc.

## Installation
1. Log in to Kyma (***console.${IP}.xip.io***).
2. Find the newly created github-connector application and bind it to the namespace that you've set up during the installation 
3. Copy exposed service URL and paste it into GitHub webhook settings adding '*`/webhook`*' at the end.

	- **Content type**: 'application/json'
	- **Secret**: for the purpose of this part it is temporarily defined **inside connector code** ('*`my-secret-key`*')
	- **SSL verification**: for the purpose of this part is 'Disabled'

![](https://i.imgur.com/wZB67Gj.png)
<div style="text-align: center"><i> Fig. 1. GitHub view after correct webhook setup </i></div></>

6. With configuration as on (Fig. 1.), click **`Update webhook`**.

To verify if it's working correctly you need to know **`POD_NAME`** and **`CONTAINER_NAME`** to get logs from destination container, because the reaction of app to an event is placed in logs.

1. To get **`POD_NAME`** and **`CONTAINER_NAME`** check
	```sh 
	kubectl get pods -n ${NAMESPACE_NAME}
	```

2. Then you can describe the pod to get the container name 
	```sh
	kubectl describe -n ${NAMESPACE_NAME} ${POD_NAME}
	```

3. To get logs run command:
	```sh
	kubectl logs -n ${NAMESPACE_NAME} ${POD_NAME} ${CONTAINER_NAME}
	```

	You should see:
    > ```
	> time="2019-07-29T06:34:14Z" level=info msg="server started"
    > 2019/07/29 06:34:20 Application ID: 179d61b2-0fa0-4000-8927-519ae5127984
	> ```

4. Now star the repository on which you set the webhook up and you should see:
	<img src="https://imgur.com/ay7T5Qc.png" width="150"/>

	In the log You should see a similar output:
	> ```
	> 2019/07/29 08:34:30 repository starred
    > 2019/07/29 08:34:30 SampleUser is watching repo "SampleOrganisation/SampleRepo"
	> ```

What means that **the webhook is properly setup now! ᶘ ᵒᴥᵒᶅ**

## Troubleshooting
The connector currently fully supports those events 
* Issues Event
* PullRequestReview Event
* Star Event
* Watch Event

For the rest of the events, it will log out a message like this
> ```
> 2019/07/29 08:32:04 unknown event type: "fork"
> ```
