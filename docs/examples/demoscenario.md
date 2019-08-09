# Using GitHub Connector <!-- omit in toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Configuration of ServiceInstance to use Lambda functions](#configuration-of-serviceinstance-to-use-lambda-functions)
	- [Option 1. Through Kyma user interface](#option-1-through-kyma-user-interface)
		- [Binding application to the Namespace](#binding-application-to-the-namespace)
		- [Setting up ServiceInstance and Lambda](#setting-up-serviceinstance-and-lambda)
		- [Creation of Lambda function](#creation-of-lambda-function)
	- [Option 2. Shell script](#option-2-shell-script)

## Overview

This purpose of this guide is to show example usage of GitHub Connector, which allows you to to handle GitHub events. With it's help you will create a lamda in Kyma, which reacts to new issue on connected repository with a short message accessible in lambda's logs. 


## Prerequisites

- Kyma with GitHub Connector deployed on your cluster (see [installation guide](helm-installation-tutorial.md)).
- WebHook configured to deliver payload to the Connector's `/webhook` endpoint.

## Configuration of ServiceInstance to use Lambda functions

To write lambda in Kyma that can utilise the functionalities of GitHub Connector you must configure the application binding and ServiceInstance.

### Option 1. Through Kyma user interface

#### Binding application to the Namespace

1. Choose Namespace created during installation of GitHub Connector.
	![Choose Namespace](./pictures/demoscenario-01-choose-namespace.png)

2. Click **Show All Applications**.
	![Show applications](./pictures/demoscenario-02-show-applications.png)

3. Click the application name to access its properties.
	![Click application](./pictures/demoscenario-03-click-application.png)

4. Create Namespace binding. During this process ServiceClass and ApplicationMapping custom resources are created.
	![Create binding](./pictures/demoscenario-04-create-binding.png)
	![Bind Namespace](./pictures/demoscenario-05-bind-namespace.png)
	
5. Now you should see this - application is bound to created Namespace.
	![See binding](./pictures/demoscenario-06-see-binding.png)

#### Setting up ServiceInstance and Lambda

1. Go to the Namespace catalog.
2. Enter Service Catalog.
	![Service catalog](./pictures/demoscenario-07-service-catalog.png)

3. Here, under the **Services** tab you can find defined service classes. Enter **github-connector** ServiceClass. In this place you are able to find defined events and registered GitHub API.
	![Service class](./pictures/demoscenario-08-service-class.png)

4. Click **Add once** button to provide ServiceInstance for **github-connector** ServiceClass. Create it with Default Plan and provided name.
	![Service instance](./pictures/demoscenario-09-service-instance.png)

5. In the Instance Catalog under **Services** tab you can find the newly created instance and its status.
	![Service instance](./pictures/demoscenario-10-service-instance-status.png)

6. In the **Lambdas** catalog in Namespace click **+ Add Lambda** to start creating lambda function.
	![Create lambda](./pictures/demoscenario-11-create-lambda.png)

#### Creation of Lambda function

After clicking **Add Lambda** you will be able to select triggering options and write body of your function.

1. Choose function name and click **Select Funtion Trigger** to choose the trigger method.
	![Select function trigger](./pictures/demoscenario-12-select-function-trigger.png)
2. From the drop-down menu that appears choose the trigger option you desire. For the purpose of this example **Event Trigger** is chosen.
	![Choose function trigger](./pictures/demoscenario-13-choose-function-trigger.png)
3. Select events for which this specific lambda you desire to be able to react.
	![Select events](./pictures/demoscenario-14-choose-events.png)
4. Create service binding and choose your service. Click **Create Service Binding** to apply binding.
   	![Write lambda](./pictures/demoscenario-15-create-service-binding.png)

5. Write body of your function and click **Create**.
	![Write lambda](./pictures/demoscenario-16-write-lambda.png)

Everything should be working now.

### Option 2. Shell script

1. Go to the `/chart/demoscenario` directory, where you can find `demoscenario.sh` script.

2. Run the script and supply the name, which has to be the same as the one you used as release name while installing GitHub Connector with helm, and the Namespace in which GitHub Connector application is running.

   ```shell
   sh demoscenario.sh {NAME} {NAMESPACE}
   ```

   After you trigger the script you get the following output that shows the created resources:

   ```
   applicationmapping.applicationconnector.kyma-project.io/gh-connector-example-app created
   serviceinstance.servicecatalog.k8s.io/gh-connector-example created
   function.kubeless.io/gh-connector-example-lambda created
   subscription.eventing.kyma-project.io/gh-connector-example-lambda-issuesevent-opened-v1 created
   Subscribed! Happy GitHub Connecting!
   ```

   Now your GitHub Connector is configured to react to new issues opened on GitHub repository you have connected to during Connector installation.

   To test if connetion works, create a new issue on connected repository to trigger the event in lambda. Then, run this command to find lambda's name (it should consist of name you have provided earlier and *-lambda* suffix):

   `kubectl get pods -n {NAMESPACE}`

   Copy the lambda's name from the output and paste it into this command to check it's logs:

   `kubectl logs -n {NAMESPACE} {LAMBDA-NAME}`

   After you run the command you should find a line containg phrase `Issue opened`. This means your lambda reacted to new issue opened event and you configured everything properly. 



