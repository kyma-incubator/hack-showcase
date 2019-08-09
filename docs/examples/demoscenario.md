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

   To test if connetion works, create a new issue on connected repository to trigger the event in lambda. Then, run this command to find lambda's name:

   `kubectl get pods -n {NAMESPACE}`

   Copy the lambda's name from the output (it should consist of name you have provided earlier and `-lambda` suffix) and paste it into this command to check it's logs:

   `kubectl logs -n {NAMESPACE} {LAMBDA-NAME}`

   After you run the command you should find a line containg phrase `Issue opened`. This means your lambda reacted to new issue opened event and you configured everything properly. 



