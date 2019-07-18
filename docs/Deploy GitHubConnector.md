Deploying GitHub Connector in Kyma 
=====
### 1. Overview
This document describes how to correctly setup the GitHub connector on Kyma, having provided that the user is logged into GCP cluster with Kyma installed.

### 2. Manual webhook setup
1. Log in to Kyma (***console.${IP}.xip.io***).
2. Create new **`namespace`**.
3. In created namespace click **`Deploy new resource`** button and choose `.yaml` file with deployment (or alternatively do it from terminal with `kubectl apply`).
4. Click on **`...`** button next to newly deployed resource and choose **`Expose API`**, then fill in the required fields.
5. Copy exposed URL and paste it into GitHub webhook settings adding '*`/webhook`*' at the end.

	- **Content type**: 'application/json'
	- **Secret**: for the purpose of this part it is temporarily defined **inside connector code** ('*`my-secret-key`*')
	- **SSL verification**: for the purpose of this part is 'Disabled'

![](https://i.imgur.com/wZB67Gj.png)
<div style="text-align: center"><i> Fig. 1. GitHub view after correct webhook setup </i></div></>

6. With configuration as on (Fig. 1.), click **`Update webhook`**.

### 3. Verification
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
	> 2019/07/17 11:17:05 server started
	> ```

4. Now star the repository on which you set the webhook up and you should see:
	<img src="https://imgur.com/ay7T5Qc.png" width="150"/>

	In the log You should see such an output:
	> ```
	> 2019/07/17 11:19:17 0xc00010e040
	> 2019/07/17 11:19:17 SampleUser has started watching your repo
	> ```

What means that **the webhook is properly setup now! ᶘ ᵒᴥᵒᶅ**