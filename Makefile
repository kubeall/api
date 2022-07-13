all: build
.PHONY: all

build:
	go build github.com/kubeall/api/...
.PHONY: build


verify:
	hack/verify-deepcopy.sh
	hack/verify-protobuf.sh
	hack/verify-swagger-docs.sh
	hack/verify-codegen.sh
.PHONY: verify


.PHONY: generate
generate:
	hack/update-deepcopy.sh
	hack/update-protobuf.sh
	hack/update-swagger-docs.sh
	hack/update-codegen.sh


CONTROLLER_GEN = $(GOPATH)/bin/controller-gen

.PHONY: manifests
manifests:
     controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=manifests/crd

