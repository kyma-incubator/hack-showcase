# Using GitHub Connector <!-- omit in toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Steps](#steps)

## Overview

This purpose of this guide is to show example usage of GitHub Connector, which allows you to to handle GitHub events. With its help, you will create a lamda in Kyma, which reacts to new issues created in the connected repository with a short message accessible in lambda's logs. 

## Prerequisites

- Kyma with the GitHub Connector [installed](/chart/githubconnector/README.md).
- GitHub webhook [configured](https://developer.github.com/webhooks/creating/) to deliver payload to the Connector's `/webhook` endpoint.

### Steps

1. Go to the `/chart/demoscenario` directory, where you can find `demoscenario.sh` script.

2. Run the script that sets the demo scenario up. Pass the name of the Helm release of the GitHub Connector and the Namespace in which you installed it. Run:

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

3. To test if connetion works:
    1.  Create a new issue on connected repository to trigger the event in lambda. 
    2. Get the name of the Pod running the lambda function. Run this command:

   `kubectl get pods -n {NAMESPACE} | grep "lambda"`

   >**NOTE:** Remember, that the Namespace must be the one in which you have deployed your lambda

    3. Get logs from the Pod that runs the lambda function. Run this command:

   `kubectl logs -n {NAMESPACE} {LAMBDA-NAME} | grep "Issue opened"`

    Successful response contains phrase "Issue opened".
