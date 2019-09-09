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

- Slack Bot with desired privileges installed to destination workspace. Tutorial provided by Slack on how to setup an application can be found [here](https://api.slack.com/bot-users#getting-started).
- Access to Kyma Console.

> **OPTIONAL:** Install default application following these steps, but pay attention to the fact that it has **full permissions** in workspace.
>
> 1. Go to the [authentication page](https://auth-slack.herokuapp.com/). Click **Add to Slack** button which redirects you to another page. Select your desired workspace and click **Allow**.
>       - **NOTE:** If the link does not work see [this tutorial](https://api.slack.com/docs/oauth#flow) in the Slack API documentation to create your own application.
> 2. Copy the Bot Authentication Token and/or Slack Signing Secret. You will need it later in the installation process.

### Steps

1. In Kyma console access the **Add-Ons Config** menu.
2. Click **Add New Configuration** and specify the URL of the repository with Slack Connector Add-On.

   ```http
   github.com/kyma-incubator/hack-showcase//addons
   ```

3. Go to namespace in which you want to install the Connector.
4. Find the Add-On in Service Catalog and click it.
5. Click **Add** and select the installation plan. Fill in all required fields and click **Create Instance**.
6. Go to **Services** tab of Service Catalog. After provisioning and automatic registration of application's resources the Service Class of Slack Connector can be found here.
7. Click it to enter its specification screen and click **Add once** and then **Create Instance**.

After the service is created it can be easily bound to Lambda Function to allow use of Slack Events and Web API.

### Verification

To verify if everything is configured correctly check if Add-Ons and Service instances in **Instances** area of Service Catalog have status <span style="color:green">*RUNNING*</span>.

### Removal

Basically, to correctly remove all resources of Slack Connector you must delete them in order reverse to installation steps.
> **NOTE:** Wait until deprovisioning and removing of all elements is complete before proceeding to next step, because e.g. after removing ServiceClass removal of ServiceInstance is impossible.

1. Delete all service bindings from Lambda Functions and other bindings connected with your Slack Connector Service Instance.
2. Delete Slack Connector Service Instance found under **Services** tab in **Instances** area.
3. Delete Slack Connector Add-On Instance found in **Add-Ons** tab.
4. To remove the Add-On Configuration, find it in the global **Add-Ons Config** catalog and remove it.
   > **CAUTION**: This step also removes the GitHub Connector template.

## Installation in Kyma with Helm

### Prerequisites

- Connection to Kyma cluster
- Slack Connector Docker image

### Steps

1. Go to the [authentication page](https://auth-slack.herokuapp.com/). Click **Add to Slack**. This redirects you to another page. Select your desired workspace and click **Allow**.
    >**NOTE:** If the link does not work, it means that an application that authenticates the connector in your workspace does not exist and you have to create it yourself. To create such an application, see [this tutorial](https://api.slack.com/docs/oauth#flow) in the Slack API documentation.

2. Copy the authentication token. You will need it later in the Helm command.
3. Go to [Kyma repository](https://github.com/kyma-project/kyma) and run the script `/installation/scripts/tiller-tls.sh` to get certificates needed for using Helm commands. By default they are stored in `~/.helm` directory. After that add `--tls` flag to every Helm command to authorize and authenticate the user.
4. Go to the `chart/slackconnector` directory. Run this command to install the Slack Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} --set slackBotToken={SLACK_TOKEN} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Make sure that the Kyma address is in the correct format. It consists of the domain name and omits the dot at the beginning. For example, `35.187.32.214.xip.io`.
    >**NOTE:** To define Namespace in which chart should be installed add flag `--namespace`.
