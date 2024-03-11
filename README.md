# Demo application

A demo K8s controller that listens for the custom resources based on the following CRD. Based on https://itnext.io/how-to-generate-client-codes-for-kubernetes-custom-resource-definitions-crd-b4b9907769ba and https://github.com/kubernetes/sample-controller

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: hellotypes.foo.com
spec:
  group: foo.com
  scope: Namespaced
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              message:
                type: string
          status:
            type: object
            properties:
              availableReplicas:
                type: integer
    subresources:
      status: {}
    additionalPrinterColumns:
    - name: Replicas
      type: string
      jsonPath: .status.availableReplicas
  names:
    kind: HelloType
    shortNames:
    - ht
    plural: hellotypes
    singular: hellotype
```

## To generate your own code

1. Install the following tools to your `GOBIN` directory.
    ``` bash
      $ go install k8s.io/code-generator/cmd/client-gen@latest
      $ go install k8s.io/code-generator/cmd/applyconfiguration-gen@latest
      $ go install k8s.io/code-generator/cmd/conversion-gen@latest
      $ go install k8s.io/code-generator/cmd/deepcopy-gen@latest
      $ go install k8s.io/code-generator/cmd/defaulter-gen@latest
      $ go install k8s.io/code-generator/cmd/go-to-protobuf@latest
      $ go install k8s.io/code-generator/cmd/informer-gen@latest
      $ go install k8s.io/code-generator/cmd/prerelease-lifecycle-gen@latest
      $ go install k8s.io/code-generator/cmd/register-gen@latest
    ```
1. Clone the kubernetes code-generator repository
    1. `cd ~`
    1. `mkdir -p go/src/k8s.io`
    1. `cd go/src/k8s.io`
    1. `git clone https://github.com/kubernetes/code-generator.git`
    1. `cd code-generator`
    1. `git checkout v0.29.2`
1. Follow [this](https://itnext.io/how-to-generate-client-codes-for-kubernetes-custom-resource-definitions-crd-b4b9907769ba) guide to create the superman/demo repo with the golang code in the appropriate files:
    1. `pkg/apis/foo/v1/doc.go`
    1. `pkg/apis/foo/v1/register.go`
    1. `pkg/apis/foo/v1/types.go`
1. In your demo repo
    1. `go mod init`
    1. `git init`
    1. `git add .`
    1. `git commit -m "check in all demo files"`
1.  Run the following to generate the k8s required code:
    ```bash
      /go/src/k8s.io/code-generator/generate-groups.sh all "github.com/superman/demo/pkg/client" "github.com/superman/demo/pkg/apis" foo:v1 --go-header-file ~/go/src/k8s.io/code-generator/examples/hack/boilerplate.go.txt
    ```
1. Copy the generated code that was generated to github.com/ inside of your repo to the top level demo folder
    1. `cp -R github.com/jacob-delgado/demo/pkg .`
    1. `rm -rf github.com`
1. `go mod tidy`
1. `git add .`
1. `git commit -m "check in generated files"`

## Running this repository

1. Install [Skaffold](https://skaffold.dev/).
1. Install [kind](https://github.com/kubernetes-sigs/kind/releases)
1. From your command line
    ```bash
      $ skaffold dev
    ```
