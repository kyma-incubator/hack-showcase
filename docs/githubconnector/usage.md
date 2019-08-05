# Usage of GitHub Connector <!-- omit in toc -->

- [Prerequisites](#prerequisites)
- [Configuration of ServiceInstance to use Lambda functions](#configuration-of-serviceinstance-to-use-lambda-functions)
	- [Option 1. Through Kyma user interface](#option-1-through-kyma-user-interface)
		- [Binding application to the namespace](#binding-application-to-the-namespace)
		- [Setting up ServiceInstance and Lambda](#setting-up-serviceinstance-and-lambda)
	- [Option 2. Helm chart](#option-2-helm-chart)
	- [Option 3. Shell script](#option-3-shell-script)

## Prerequisites

* Kyma with GitHub Connector deployed on your cluster (see [installation guide](helm-installation-tutorial.md)).
* WebHook configured to deliver payload to the Connector's ```/webhook``` endpoint.

## Configuration of ServiceInstance to use Lambda functions

To write lambda in Kyma that can utilise the functionalities of GitHub Connector you must configure the application binding and ServiceInstance.

### Option 1. Through Kyma user interface

#### Binding application to the namespace

1. Choose namespace created during installation of GitHub Connector.
	![Choose namespace](./pictures/usage-01-choose-namespace.png)

2. Click **```Show All Applications```**.
	![Show applications](./pictures/usage-02-show-applications.png)

3. Click the application name to access its properties.
	![Click application](./pictures/usage-03-click-application.png)

4. Create namespace binding. During this process **ServiceClass** and **ApplicationMapping** custom resources are created.
	![Create binding](./pictures/usage-04-create-binding.png)
	![Bind namespace](./pictures/usage-05-bind-namespace.png)
	
5. Now you should see this - application is bound to created namespace.
	![See binding](./pictures/usage-06-see-binding.png)

#### Setting up ServiceInstance and Lambda

1. Go to the namespace catalog.
2. Enter Service Catalog.
	![Service catalog](./pictures/usage-07-service-catalog.png)

3. Here, under the **Services** tab you can find defined service classes. Enter **github-connector** ServiceClass. In this place you are able to find defined events and registered GitHub API.
	![Service class](./pictures/usage-08-service-class.png)

4. Click **```Add once```** button to provide ServiceInstance for **github-connector** ServiceClass. Create it with Default Plan and provided name.
	![Service instance](./pictures/usage-09-service-instance.png)

5. In the Instance Catalog under **Services** tab you can find the newly created instance and its status.
	![Service instance](./pictures/usage-10-service-instance-status.png)

6. In the **Lambdas** catalog in namespace click **```+ Add Lambda```** to start creating lambda.
	![Create lambda](./pictures/usage-11-create-lambda.png)

### Option 2. Helm chart

### Option 3. Shell script