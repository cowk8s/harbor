
SHELL := /bin/bash
BUILDPATH=$(CURDIR)

TOOLSPATH=$(BUILDPATH)/tools
PORTAL_PATH=$(BUILDPATH)/src/portal



# version prepare
# for docker image tag
VERSIONTAG=dev
BUILDBASETARGET=portal
IMAGENAMESPACE=goharbor

# docker parameters
DOCKERCMD=$(shell which docker)
DOCKERBUILD=$(DOCKERCMD) build
DOCKERIMAGES=$(DOCKERCMD) images

# go parameters
GOBUILDIMAGE=golang:1.17.7

# docker image name
DOCKERIMAGENAME_PORTAL=$(IMAGENAMESPACE)/harbor-portal

# cmds
DOCKERSAVE_PARA=$(DOCKERIMAGENAME_PORTAL):$(VERSIONTAG)

RUNCONTAINER=$(DOCKERCMD) run --rm -u $(shell id -u):$(shell id -g) -v $(BUILDPATH):$(BUILDPATH) -w $(BUILDPATH)

# $1 the name of the docker image
# $2 the tag of the docker image
# $3 the command to build the docker image
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

# $1 the path of swagger spec
# $2 the path of base directory for generating the files
# $3 the name of the application
define swagger_generate_server
	@echo "generate all the files for API from $(1)"
	@rm -rf $(2)/{models,restapi}
	@mkdir -p $(2)
	@$(SWAGGER_GENERATE_SERVER) -f $(1) -A $(3) --target $(2)
endef

gen_apis:
	$(call prepare_docker_image,${SWAGGER_IMAGENAME},${SWAGGER_VERSION},${SWAGGER_IMAGE_BUILD_CMD})
	$(call swagger_generate_server,api/v2.0/swagger.yaml,src/server/v2.0,harbor)

build_base_docker:
	@for name in $(BUILDBASETARGET); do \
		echo $$name ; \
		sleep 30 ; \
		
		$(DOCKERBUILD) --build-arg BUILD

clean:
	@echo "  make cleanall:		remove binary, Harbor images, specific version docker-compose"
