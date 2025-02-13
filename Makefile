APP_NAME = api-gateway-controller
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)
CRD_OPTIONS ?= "crd:trivialVersions=true"
SHELL = /bin/bash

.EXPORT_ALL_VARIABLES:
GO111MODULE = on


.PHONY: build
build: generate
	./before-commit.sh ci

.PHONY: pull-licenses
pull-licenses:
ifdef LICENSE_PULLER_PATH
	bash $(LICENSE_PULLER_PATH)
else
	mkdir -p licenses
endif

.PHONY: build-image
build-image: pull-licenses
	docker build -t $(APP_NAME):latest .

.PHONY: push-image
push-image:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)

.PHONY: ci-pr
ci-pr: build build-image push-image

.PHONY: ci-master
ci-master: build build-image push-image

.PHONY: ci-release
ci-release: build build-image push-image

.PHONY: clean
clean:
	rm -rf bin

.PHONY: path-to-referenced-charts
path-to-referenced-charts:
	@echo "resources/core"

# Install CRDs into a cluster
install: manifests
	kustomize build config/crd | kubectl apply -f -
	@if ! kubectl get crd virtualservices.networking.istio.io > /dev/null 2>&1 ; then kubectl apply -f hack/networking.istio.io_virtualservice.yaml; fi;

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths=./api/...

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.0-beta.2
CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

run: build
	go run .

samples-clean:
	kubectl delete -f config/samples/valid.yaml --ignore-not-found=true
	kubectl delete -f config/samples/invalid.yaml --ignore-not-found=true

.PHONY: samples
samples: samples-valid

.PHONY: samples-valid
samples-valid: samples-clean
	kubectl apply -f config/samples/valid.yaml

.PHONY: samples-invalid
samples-invalid: samples-clean
	kubectl apply -f config/samples/invalid.yaml
