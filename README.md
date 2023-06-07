# thomasscothamilton
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## FluxCD

```shell
flux check --pre
```

```shell
flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=thomasscothamilton-infra \
  --branch=main \
  --path=./clusters/my-cluster \
  --personal
```

```shell
flux create source git thomasscothamilton \
  --url=https://github.com/thomasscothamilton/thomasscothamilton \
  --branch=main \
  --interval=30s \
  --export > ./clusters/my-cluster/thomasscothamilton-source.yaml
```

```shell
flux create kustomization thomasscothamilton \
  --target-namespace=default \
  --source=thomasscothamilton \
  --path="./kustomize" \
  --prune=true \
  --interval=5m \
  --export > ./clusters/my-cluster/thomasscothamilton-kustomization.yaml
```

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/thomasscothamilton:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/thomasscothamilton:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

# azure

```shell

ACR_NAME=thomasscothamilton
REGISTRY=$ACR_NAME.azurecr.io

# Login to Azure
az login

# Login to ACR, using a token based on your Azure identity
USER_NAME="00000000-0000-0000-0000-000000000000"
PASSWORD=$(az acr login --name $ACR_NAME --expose-token --output tsv --query accessToken)
# docker login loginServer -u 00000000-0000-0000-0000-000000000000 -p accessToken

```

```shell
docker build -t thomasscothamilton.azurecr.io .
```

```shell
groupId=$(az group show \
  --name thomasscothamilton \
  --query id --output tsv)
```

```shell
az ad sp create-for-rbac \
  --scope $groupId \
  --role Contributor \
  --sdk-auth
```

```shell
registryId=$(az acr show \
  --name thomasscothamilton.azurecr.io \
  --resource-group thomasscothamilton \
  --query id --output tsv)
```

```shell
az role assignment create \
  --assignee e7fdec1e-2c30-4f55-8270-1d31e2b05527 \
  --scope $registryId \
  --role AcrPush
```