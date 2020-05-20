VERSION := $(shell git describe --always --long --dirty)
mod:
	GO111MODULE=on go mod tidy

# checkout to release-1.13 first...
code-gen:
	GO111MODULE=off ${GOPATH}/src/k8s.io/code-generator/generate-groups.sh all "github.com/alauda/helm-crds/pkg/client" "github.com/alauda/helm-crds/pkg/apis" app:v1alpha1,v1beta1


fmt:
	find ./pkg -name \*.go  | xargs goimports -w
