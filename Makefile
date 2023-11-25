# Packages content
PKG_OS = darwin linux
PKG_ARCH = amd64

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install -v
GOGET = $(GOCMD) get -v -t
GOTEST = $(GOCMD) test -v

# Environment
WORKDIR := $(shell pwd)
BUILD_PATH := $(WORKDIR)/build
GOCACHE_PATH = $(WORKDIR)/gocache
DOCKER_IMAGE_BUILD = prusalink-screen-build

DEBIAN_PACKAGES = BULLSEYE
ARCH = armhf
# ARCH = amd64

BULLSEYE_NAME := bullseye
BULLSEYE_IMAGE := golang:1.21-bullseye
BULLSEYE_GO_TAGS := "gtk_3_24 glib_deprecated glib_2_66"

# Build information
#GIT_COMMIT = $(shell git rev-parse HEAD | cut -c1-7)
VERSION := 0.0.1
BUILD_DATE ?= $(shell date --utc +%Y%m%d-%H:%M:%S)
#BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

#ifneq ($(BRANCH), master)
#	VERSION := $(shell echo $(BRANCH)| sed -e 's/v//g')
#endif

# Package information
PACKAGE_NAME = prusalink-screen

# we export the variable to allow envsubst, substitute the vars in the
# Dockerfiles
export

build-environment:
	mkdir -p ${BUILD_PATH}
	mkdir -p ${GOCACHE_PATH}

build: | build-environment $(DEBIAN_PACKAGES)

$(DEBIAN_PACKAGES):
	docker build \
		--build-arg IMAGE=${${@}_IMAGE} \
		--build-arg TARGET_ARCH=${ARCH} \
		--build-arg GO_TAGS=${${@}_GO_TAGS} \
		-t ${DOCKER_IMAGE_BUILD}:${${@}_NAME}-${ARCH} . \
		&& \
	docker run --rm \
		-e TARGET_ARCH=${ARCH} \
		-v ${BUILD_PATH}/${${@}_NAME}-${ARCH}:/build \
		-v ${GOCACHE_PATH}/${${@}_NAME}-${ARCH}:/gocache \
		${DOCKER_IMAGE_BUILD}:${${@}_NAME}-${ARCH} \
		make build-internal

build-internal: prepare-internal
	#go build --tags ${GO_TAGS} -v -o /build/bin/${BINARY_NAME} main.go
	cd $(WORKDIR); \
	GOCACHE=/gocache debuild --prepend-path=/usr/local/go/bin/ --preserve-env -us -uc -a${TARGET_ARCH} \
	&& cp ../*.deb /build/;

prepare-internal:
	dch --create -v $(VERSION)-1 --package $(PACKAGE_NAME) --controlmaint empty; \
	cd $(WORKDIR)/..; \
	tar -czf prusalink-screen_$(VERSION).orig.tar.gz --exclude-vcs prusalink-screen

clean:
	rm -rf ${BUILD_PATH}
