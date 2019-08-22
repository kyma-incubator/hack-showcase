# GitHub Connector Installation

- [Installation in Kyma environment with Helm](#installing-in-kyma-using-helm)
	- [Prerequisites](#prerequisites-1)
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
