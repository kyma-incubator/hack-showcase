# Slack Connector Installation<!-- omit in toc -->

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

The Slack Connector is a component which allows interaction with Slack API from inside of Kyma environment. The simplest way to install Slack Connector in Kyma is to install it as an Add-On.

## Installation in Kyma as an Add-On

### Prerequisites

- Slack Bot with desired privileges installed to the destination workspace. See the tutorial provided by Slack on how to setup an application [here](https://api.slack.com/bot-users#getting-started). After you have created the Signing Secret, it appears in the application's **Settings** in the **Basic Information** section.
- Access to Kyma Console

> **OPTIONAL:** Follow these steps to install the default application. Be aware that it has **full permissions** in the in the workspace.
>
> 1. Go to the [authentication page](https://auth-slack.herokuapp.com/). Click the **Add to Slack** button, which redirects you to another page. Select the desired workspace and click **Allow**.
>       - **NOTE:** If the link does not work, see [this](https://api.slack.com/docs/oauth#flow) tutorial in the Slack API documentation to create your own application.
> 2. Copy the Bot Authentication Token and/or Slack Signing Secret. You will need it later in the installation process.

### Steps

1. In Kyma console, access the **Add-Ons Config** menu.
2. Click **Add New Configuration** and fill in the **Urls\*** field with this URL:

   ```http
   github.com/kyma-incubator/hack-showcase//addons
   ```

3. Go to the Namespace in which you install the Connector.
4. Find the Add-On in the Service Catalog and click it.
5. Click **Add** and select the installation plan. Fill in all required fields and click **Create Instance**.
6. Go to the **Services** tab in the Service Catalog. After provisioning and automatic registration of application's resources, the Service Class of Slack Connector appears here.
7. Click the Service Class to enter its specification screen, click **Add once**, and then **Create Instance**.

After you have created the service, you can easily bind it to the Lambda Function to allow the use of Slack Events.

### Verification

To verify correct configuration, check if Add-Ons and Service instances in the **Instances** area of the Service Catalog have status **RUNNING**.

### Removal

To correctly remove all Slack Connector resources, you must delete them in order reverse to the installation steps.
> **NOTE:** Wait until deprovisioning and removing of all elements is complete before proceeding to next step to avoid possible errors. For example, after removing ServiceClass, the removal of ServiceInstance is impossible.

1. Delete all service bindings from Lambda Functions and other bindings connected with your Slack Connector Service Instance.
2. Delete the Slack Connector Service Instance found under the **Services** tab in the **Instances** area.
3. Delete the Slack Connector Add-On Instance found in the **Add-Ons** tab.
4. To remove the Add-On Configuration, find it in the global **Add-Ons Config** catalog and remove it.
   > **CAUTION**: This step also removes the GitHub Connector template.

## Installation in Kyma with Helm

### Prerequisites

- Connection to Kyma cluster
- Slack Connector Docker image

### Steps

1. Go to the [authentication page](https://auth-slack.herokuapp.com/). Click **Add to Slack**. This redirects you to another page. Select the desired workspace and click **Allow**.
    >**NOTE:** If the link does not work, it means that the application that authenticates the connector in your workspace does not exist and you have to create it yourself. To create such an application, see [this tutorial](https://api.slack.com/docs/oauth#flow) in the Slack API documentation.

2. Copy the authentication token. You will need it later in the Helm command.
3. Go to [Kyma repository](https://github.com/kyma-project/kyma) and run the script `/installation/scripts/tiller-tls.sh` to get certificates needed for using Helm commands. By default they are stored in `~/.helm` directory. After that add `--tls` flag to every Helm command to authorize and authenticate the user.
4. Go to the `chart/slackconnector` directory. Run this command to install the Slack Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} --set slackBotToken={SLACK_TOKEN} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Make sure that the Kyma address is in the correct format. It consists of the domain name and omits the dot at the beginning. For example, `35.187.32.214.xip.io`.
    >**NOTE:** To define Namespace in which chart should be installed add flag `--namespace`.
