# GitHub Connector Installation

- [GitHub Connector Installation](#github-connector-installation)
  - [Overview](#overview)
  - [Installation in Kyma as Add-On](#installation-in-kyma-as-add-on)
    - [Prerequisites](#prerequisites)
    - [Steps](#steps)
    - [Verification](#verification)
    - [Removal](#removal)
  - [Installation in Kyma with Helm](#installation-in-kyma-with-helm)
    - [Prerequisites](#prerequisites-1)
    - [Steps](#steps-1)

## Overview

The GitHub Connector is a component which allows interaction with GitHub API from inside of Kyma environment. The simplest way to install GitHub Connector in Kyma is to install it as Add-On.

## Installation in Kyma as Add-On

### Prerequisites

- GitHub App with desired privileges installed to destination repository or organization. You can create new application [here](https://github.com/settings/apps) (or in account's *Settings/Developer settings/GitHub Apps*).
  - GitHub Repository Token.
- Access to Kyma Console.

> **NOTE**: It is recommended to create or use additional service account (e.g. Your-Project-Name-Github-Connector) since any actions performed by this application will be signed as user that the token belongs to.

> **OPTIONAL:** If you want to install default application proceed these steps, but pay attention to the fact that provided application has **full permissions** in repository/organization.
>
> 1. Go to the [authentication page](https://auth-github-connector.herokuapp.com/). Click **GitHub** button what will redirect you to another page. Select repositories or organizations you want to install the application and click **Install**.
>       - **NOTE:** If the link does not work, see [this tutorial](https://developer.github.com/apps/quickstart-guides/setting-up-your-development-environment/#step-2-register-a-new-github-app) in the GitHub documentation to create your own application.
> 2. Copy the authentication token. You will need it later in the installation process.

### Steps

1. In Kyma console access the **Add-Ons Config** menu.
2. Click **Add New Configuration** and specify the URL of the repository with GitHub Connector Add-On.

   ```http
   github.com/kyma-incubator/hack-showcase//addons
   ```

3. Go to namespace in which you want to install the Connector.
4. Find the Add-On in Service Catalog and click it.
5. Click **Add** and select the installation plan. Fill in all required fields and click **Create Instance**.
6. Go to **Services** tab of Service Catalog. After provisioning and automatic registration of application's resources the Service Class of GitHub Connector can be found here.
7. Click it to enter its specification screen and click **Add once** and then **Create Instance**.

After the service is created it can be easily bound to Lambda Function to allow use of GitHub Events.

### Verification

- To verify if everything is configured correctly check if Add-Ons and Service instances in **Instances** area of Service Catalog have status <span style="color:green">*RUNNING*</span>.
- Check in your GitHub repository's or organization's *Settings/Webhooks* if the webhook is <span style="color:green">*Active*</span>.

### Removal

Basically, to correctly remove all resources of GitHub Connector you need to delete them in order reverse to installation steps.
> **NOTE:** Wait until deprovisioning and removing of all elements is complete before proceeding to next step, because e.g. after removing ServiceClass removing of ServiceInstance is impossible.

1. Delete all service bindings from Lambda Functions and other bindings connected with your GitHub Connector Service Instance.
2. Delete GitHub Connector Service Instance found under **Services** tab in **Instances** area.
3. Delete GitHub Connector Add-On Instance found in **Add-Ons** tab.
4. If you want to remove add-on configuration, you can remove Add-Ons Configuration from global **Add-Ons Config** Catalog, but this removes *GitHub Connector* alongside.

## Installation in Kyma with Helm

### Prerequisites

- Connection to Kyma cluster
- The GitHub Connector Docker image

### Steps

1. Go to [Kyma repository](https://github.com/kyma-project/kyma) and run script `/installation/scripts/tiller-tls.sh` to get certificates needed for using Helm commands. By default they are stored in `~/.helm` directory. After that add the `--tls` flag to every Helm command to authorize and authenticate a user.
2. Go to the `chart/githubconnector` directory. Run this command to install the GitHub Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} --set githubURL={GITHUB_REPO_URL} --set githubToken={GITHUB_TOKEN} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Make sure the Kyma address is in the correct format. It consists of the domain name and cannot begin with the dot. For example, `35.187.32.214.xip.io`.

    >**NOTE:** To define the Namespace in which to install chart, add the flag `--namespace`. To define the GitHub URL, add the flag `--set githubURL`. If you want to crate webhook on one repository use construction `repos/:owner/:repo`. if you want create webhook on whole organization you have to use `orgs/:org`. To provide security token use flag `--set githubToken`.

3. For further steps see [configuration page](/docs/github-connector/configuration.md)
