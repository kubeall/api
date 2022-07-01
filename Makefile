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



generate:
	hack/update-deepcopy.sh
	hack/update-protobuf.sh
	hack/update-swagger-docs.sh
	hack/update-codegen.sh
.PHONY: generate
