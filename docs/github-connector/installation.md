# GitHub Connector Installation

- [GitHub Connector Installation](#github-connector-installation)
  - [Installation in Kyma with Helm](#installation-in-kyma-with-helm)
    - [Prerequisites](#prerequisites)
    - [Steps](#steps)

## Installation in Kyma with Helm

### Prerequisites

- Connection to Kyma cluster
- The GitHub Connector Docker image

### Steps

1. Go to [Kyma repository](https://github.com/kyma-project/kyma) and run script `/installation/scripts/tiller-tls.sh` to get certificates needed for using Helm commands. By default they are stored in `~/.helm` directory. After that add the `--tls` flag to every Helm command to authorize and authenticate a user.
2. Go to the `chart/githubconnector` directory. Run this command to install the GitHub Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Make sure the Kyma address is in the correct format. It consists of the domain name and cannot begin with the dot. For example, `35.187.32.214.xip.io`.

    >**NOTE:** To define the Namespace in which to install chart, add the flag `--namespace`.
3. For further steps see [configuration page](/docs/github-connector/configuration.md)
