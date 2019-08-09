# Using GitHub Connector <!-- omit in toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Steps](#steps)

## Overview

This purpose of this guide is to show example usage of GitHub Connector, which allows you to to handle GitHub events. With its help, you will create a lamda in Kyma, which reacts to new issues created in the connected repository with a short message accessible in lambda's logs. 

## Prerequisites

- Kyma with GitHub Connector deployed on your cluster (see [installation guide](/chart/githubconnector/README.md)).
- WebHook configured to deliver payload to the Connector's `/webhook` endpoint.

### Steps

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

3. To test if connetion works, create a new issue on connected repository to trigger the event in lambda. Then, run this command to find name of the Pod in which lambda is running:

   `kubectl get pods -n {NAMESPACE}`

   >**NOTE:** Remember, that the Namespace must be the one in which you have deployed your lambda

4. Run this command to get logs from the Pod that runs the lambda function and search for the "Issue opened" phrase to verify if you have configured everything properly and lambda reacts to the event.

   `kubectl logs -n {NAMESPACE} {LAMBDA-NAME} | grep "Issue opened"`
