
SHELL := /bin/bash
BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make
TOOLSPATH=$(BUILDPATH)/tools

IMAGENAMESPACE=goharbor

# docker parameters
DOCKERCMD=$(shell which docker)
DOCKERBUILD=$(DOCKERCMD) build
DOCKERIMAGES=$(DOCKERCMD) images

# go parameters
GOCMD=$(shell which go)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
GOBUILDIMAGE=golang:1.17.7
GOBUILDPATHINCONTAINER=/harbor

RUNCONTAINER=$(DOCKERCMD) run --rm -u $(shell id -u):$(shell id -g) -v $(BUILDPATH):$(BUILDPATH) -w $(BUILDPATH)

define prepare_docker_image
	@if [ "$(shell ${DOCKERIMAGES} -q $(1):$(2) 2> /dev/null)" == "" ]; then \
		$(3) && echo "build $(1):$(2) successfully" || (echo "build $(1):$(2) failed" && exit 1) ; \
	fi
endef

SWAGGER_IMAGENAME=$(IMAGENAMESPACE)/swagger
SWAGGER_VERSION=v0.25.0
SWAGGER=$(RUNCONTAINER) ${SWAGGER_IMAGENAME}:${SWAGGER_VERSION}
SWAGGER_GENERATE_SERVER=${SWAGGER} generate server --template-dir=$(TOOLSPATH)/swagger/templates --exclude-main --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC
SWAGGER_IMAGE_BUILD_CMD=${DOCKERBUILD} -f ${TOOLSPATH}/swagger/Dockerfile --build-arg GOLANG=${GOBUILDIMAGE} --build-arg SWAGGER_VERSION=${SWAGGER_VERSION} -t ${SWAGGER_IMAGENAME}:$(SWAGGER_VERSION) .

define swagger_generate_server
	@echo "generate all the files for API from $(1)"
	@rm -rf $(2)/{models,restapi}
	@mkdir -p $(2)
	@$(SWAGGER_GENERATE_SERVER) -f $(1) -A $(3) --target $(2)
endef

gen_apis:
	$(call prepare_docker_image,${SWAGGER_IMAGENAME},${SWAGGER_VERSION},${SWAGGER_IMAGE_BUILD_CMD})
	$(call swagger_generate_server,api/v2.0/swagger.yaml,src/server/v2.0,harbor)