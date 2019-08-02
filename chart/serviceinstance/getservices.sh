apk add curl
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/linux/amd64/kubectl
chmod +x ./kubectl
apk add sudo
sudo mv ./kubectl /usr/local/bin/kubectl
kubectl get pods --all-namespaces