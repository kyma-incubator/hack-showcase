# GitHub Connector Installation <!-- omit in toc -->


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

The GitHub Connector is a component which allows interaction with the GitHub API from inside of Kyma environment. The simplest way to install the GitHub Connector in Kyma is to install it as an Add-On.

## Installation in Kyma as an Add-On

### Prerequisites

- GitHub App with desired privileges installed to the destination repository or organization. To create a new application, go [here](https://github.com/settings/apps) or access **Github Apps** in the account through **Settings** in **Developer settings**.
- Access to the Kyma Console

> **NOTE**: It is best to create or use an additional service account (e.g. Your-Project-Name-Github-Connector) since any actions that the application performs are signed with the name of the user that the token belongs to.

> **OPTIONAL:** Follow these steps to install the default application. Be aware that it has **full permissions** in the repository/organization.
>
> 1. Go to the [authentication page](https://auth-github-connector.herokuapp.com/). Click the **GitHub** button, which redirects you to another page. Select the repositories or organizations you want to install the application to and click **Install**.
>       - **NOTE:** If the link does not work, see [this tutorial](https://developer.github.com/apps/quickstart-guides/setting-up-your-development-environment/#step-2-register-a-new-github-app) in the GitHub documentation to create your own application.
> 2. Copy the Authentication Token. You will need it later in the installation process.

### Steps

1. In Kyma console, access the **Add-Ons Config** menu.
2. Click **Add New Configuration** and in fill in the **Urls** field with this URL:

   ```http
   github.com/kyma-incubator/hack-showcase//addons
   ```

3. Go to the Namespace in which to install the Connector.
4. Find the Add-On in the Service Catalog and click it.
5. Click **Add** and select the installation plan. Fill in all required fields and click **Create Instance**.
6. Go to the **Services** tab in the Service Catalog. After provisioning and automatic registration of application's resources, the Service Class of the GitHub Connector appears here.
7. Click the Service Class to enter its specification screen, click **Add once**, and then **Create Instance**.

To send requests to the GitHub API, bind the service you created to the Lambda Function.

### Verification

- To verify correct configuration, check if Add-Ons and Service instances in the **Instances** area of the Service Catalog have status **RUNNING**.
- Access **Webhooks**  in your GitHub repository or organization's **Settings** and verify that the webhook is **Active**.

### Removal

To correctly remove all GitHub Connector resources, you must delete them in order reverse to the installation steps.
> **CAUTION:** Wait until deprovisioning and removing of all elements is complete before proceeding to the next step to avoid possible errors. For example, after removing ServiceClass, the removal of ServiceInstance is impossible.

1. Delete all service bindings from Lambda Functions and other bindings connected with your GitHub Connector Service Instance.
2. Delete the GitHub Connector Service Instance found under the **Services** tab in the **Instances** area.
3. Delete the GitHub Connector Add-On Instance found in the **Add-Ons** tab.
4. To remove the Add-On Configuration, find it in the global **Add-Ons Config** catalog and remove it.
   > **CAUTION**: This step also removes the Slack Connector template.

## Installation in Kyma with Helm

### Prerequisites

- Connection to Kyma cluster
- The GitHub Connector Docker image

### Steps

1. Go to the [Kyma repository](https://github.com/kyma-project/kyma) and run the script `/installation/scripts/tiller-tls.sh` to get the certificates needed to use Helm commands. By default, they are stored in the `~/.helm` directory. After that, add the `--tls` flag to every Helm command to authorize and authenticate the user.
2. Go to the `chart/githubconnector` directory. Run this command to install the GitHub Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} --set githubURL={GITHUB_REPO_URL} --set githubToken={GITHUB_TOKEN} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Make sure the Kyma address is in the correct format. It consists of the domain name and cannot begin with the dot. For example, `35.187.32.214.xip.io`.

    >**NOTE:** To define the Namespace in which to install chart, add the `--namespace` flag to the command. To define the GitHub URL, add the `--set githubURL` flag. To crate a webhook on one repository, use the construction `repos/:owner/:repo`. To create a webhook on the whole organization, use `orgs/:org`. To provide the security token, use the `--set githubToken` flag.

3. For further steps, see [configuration page](/docs/github-connector/configuration.md).
