########################################################################################################################
# Copyright (c) 2020 IoTeX Foundation
# This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
# warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
# permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
# License 2.0 that can be found in the LICENSE file.
########################################################################################################################

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
BUILD_TARGET_IOT=pebble-data-container

# Pkgs
ALL_PKGS := $(shell go list ./... )
PKGS := $(shell go list ./... | grep -v /test/ )
ROOT_PKG := "github.com/iotexproject/pebble-data-container"

# Docker parameters
DOCKERCMD=docker

# Package Info
PACKAGE_VERSION := $(shell git describe --tags)
PACKAGE_COMMIT_ID := $(shell git rev-parse HEAD)
GIT_STATUS := $(shell git status --porcelain)
ifdef GIT_STATUS
	GIT_STATUS := "dirty"
else
	GIT_STATUS := "clean"
endif

GO_VERSION := $(shell go version)
BUILD_TIME=$(shell date +%F-%Z/%T)
VersionImportPath := github.com/iotexproject/iotex-core/pkg/version
PackageFlags += -X '$(VersionImportPath).PackageVersion=$(PACKAGE_VERSION)'
PackageFlags += -X '$(VersionImportPath).PackageCommitID=$(PACKAGE_COMMIT_ID)'
PackageFlags += -X '$(VersionImportPath).GitStatus=$(GIT_STATUS)'
PackageFlags += -X '$(VersionImportPath).GoVersion=$(GO_VERSION)'
PackageFlags += -X '$(VersionImportPath).BuildTime=$(BUILD_TIME)'
PackageFlags += -s -w

default: clean build
all: clean build-all

.PHONY: build
build: 
	$(GOBUILD) -ldflags "$(PackageFlags)" -o ./bin/$(BUILD_TARGET_IOT) -v .

.PHONY: build-all
build-all: build

.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf ./bin/$(BUILD_TARGET_IOT)
