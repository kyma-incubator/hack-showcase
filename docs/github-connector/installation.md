# GitHub Connector Installation

- [GitHub Connector Installation](#github-connector-installation)
  - [Installation in Kyma with Helm](#installation-in-kyma-with-helm)
    - [Prerequisites](#prerequisites)
    - [Steps](#steps)

## Installation in Kyma with Helm

### Prerequisites

- Connection to Kyma cluster
- GitHub Connector Docker image

### Steps

1. Go to [Kyma repository](https://github.com/kyma-project/kyma) and run script `/installation/scripts/tiller-tls.sh` to get certificates needed for using Helm commands. By default they are stored in `~/.helm` directory. After that add `--tls` flag to every Helm command to authorize and authenticate yourself
2. Go to `chart/githubconnector` directory. Run the command to install GitHub Connector:

    ``` shell
    helm install --set container.image={DOCKER_IMAGE} --set kymaAddress={KYMA_ADDRESS} -n {RELEASE_NAME} . --tls
    ```

    >**CAUTION:** Kyma address should be in the right format. It must consist of domain name, without dot  character at the beggining, for example `35.187.32.214.xip.io`

    >**NOTE:** To define Namespace in which chart should be installed add flag `--namespace`.
3. For further steps see [configuration page](/docs/github-connector/configuration.md)
