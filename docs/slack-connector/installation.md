# Slack Connector Installation <!-- omit in toc -->

- [Overview](#overview)
- [Installation in Kyma as an Add-On](#installation-in-kyma-as-an-add-on)
  - [Prerequisites](#prerequisites)
  - [Steps](#steps)
  - [Verification](#verification)
  - [Removal](#removal)
- [Installation in Kyma with Helm](#installation-in-kyma-with-helm)
  - [Prerequisites](#prerequisites-1)
  - [Steps](#steps-1)

## Overview

The Slack Connector is a component which allows interaction with the Slack API from inside of Kyma environment. The simplest way to install the Slack Connector in Kyma is to install it as an Add-On.

## Installation in Kyma as an Add-On

### Prerequisites

- Slack App with desired privileges installed to the destination workspace. See the tutorial provided by Slack on how to setup an application [here](https://api.slack.com/bot-users#getting-started).

    After creating the app:
    - The Slack Application's {SIGNING_SECRET} is used for validating requests coming from Slack by verifying its unique signature. Find it in the [Application's](https://api.slack.com/apps) **Settings** in the **Basic Information** section. [Learn more](https://api.slack.com/docs/verifying-requests-from-slack).
    - The Slack Application's {SLACK_TOKEN} is used for requests authorization. Find it in the [Application's](https://api.slack.com/apps) **Features** in the **OAuth & Permissions** section. Depending on the usecase it will be **OAuth Access Token** or **Bot User OAuth Access Token**. [Learn more](https://api.slack.com/docs/oauth).
- Access to Kyma Console

### Steps
1. Add Add-Ons configuration to Kyma. Run:

    ``` shell
    cat <<EOF | kubectl apply -f -
    apiVersion: addons.kyma-project.io/v1alpha1
    kind: ClusterAddonsConfiguration
    metadata:
      name: addons-slack-github-connectors
      finalizers:
      - addons.kyma-project.io
    spec:
      repositories:
        - url: github.com/kyma-incubator/github-slack-connectors//addons/index.yaml
    EOF
    ```
    
2. Go to the Namespace in which to install the Connector.
3. Find the Add-On in the Service Catalog and click it.
4. Click **Add** and select the installation plan. Fill in all required fields and click **Create Instance**.
5. Go to the **Services** tab in the Service Catalog. After provisioning and automatic registration of Application's resources, the Service Class of the Slack Connector appears here.
6. Click the Service Class to enter its specification screen, click **Add once**, and then **Create Instance**.

To send requests to the Slack API, bind the service you created to the Lambda Function.

### Verification

To verify correct configuration, check if Add-Ons and Service instances in the **Instances** area of the Service Catalog have status **RUNNING**.

### Removal

To correctly remove all Slack Connector resources, you must delete them in order reverse to the installation steps.
> **NOTE:** Wait until deprovisioning and removing of all elements is complete before proceeding to the next step to avoid possible errors. For example, after removing ServiceClass, the removal of ServiceInstance is impossible.

1. Delete all service bindings from Lambda Functions and other bindings connected with your Slack Connector Service Instance.
2. Delete the Slack Connector Service Instance found under the **Services** tab in the **Instances** area.
3. Delete the Slack Connector Add-On Instance found in the **Add-Ons** tab.
4. To remove the Add-On Configuration, find it in the global **Add-Ons Config** catalog and remove it.
   > **CAUTION**: This step also removes the GitHub Connector template.

## Installation in Kyma with Helm

### Prerequisites

- Connection to the Kyma cluster
- Slack Connector Docker image

### Steps

1. Go to the [authentication page](https://auth-slack.herokuapp.com/). Click **Add to Slack**. This redirects you to another page. Select the desired workspace and click **Allow**.
    >**NOTE:** If the link does not work, it means that the application that authenticates the connector in your workspace does not exist and you have to create it yourself. To create such an application, see [this tutorial](https://api.slack.com/docs/oauth#flow) in the Slack API documentation.

2. Copy the authentication token. You will need it later in the Helm command.
3. Go to the [Kyma repository](https://github.com/kyma-project/kyma) and run the script `/installation/scripts/tiller-tls.sh` to get the certificates needed to use Helm commands. By default, they are stored in the `~/.helm` directory. After that, add the `--tls` flag to every Helm command to authorize and authenticate the user.
4. Go to the `chart/slackconnector` directory. Run this command to install the Slack Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} --set slackBotToken={SLACK_TOKEN} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Make sure that the Kyma address is in the correct format. It consists of the domain name and omits the dot at the beginning. For example, `35.187.32.214.xip.io`.

    >**NOTE:** To define Namespace in which to install chart, add the `--namespace` flag. To provide security token, add the `--set slackBotToken` flag.
