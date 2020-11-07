# IMG is the base path for images..
# Individual images will be
# $(IMG)/$(NAME):$(TAG)
CONTEXT=label-bot-dev

IMG ?= gcr.io/${PROJECT}/vscode

CONTEXT = $(shell yq r Kptfile openAPI.definitions["io.k8s.cli.setters.context"].["x-k8s-cli"].setter.value)
PROJECT = $(shell yq r Kptfile openAPI.definitions["io.k8s.cli.setters.gcloud.core.project"].["x-k8s-cli"].setter.value)
NAME = $(shell yq r Kptfile openAPI.definitions["io.k8s.cli.setters.name"].["x-k8s-cli"].setter.value)
NAMESPACE = $(shell yq r Kptfile openAPI.definitions["io.k8s.cli.setters.namespace"].["x-k8s-cli"].setter.value)

# List any changed  files. We only include files in the notebooks directory.
# because that is the code in the docker image.
# In particular we exclude changes to the ksonnet configs.
CHANGED_FILES := $(shell git diff-files)

# Whether to use cached images with GCB
USE_IMAGE_CACHE ?= true

ifeq ($(strip $(CHANGED_FILES)),)
# Changed files is empty; not dirty
# Don't include --dirty because it could be dirty if files outside the ones we care
# about changed.
GIT_VERSION := $(shell git describe --always)
else
GIT_VERSION := $(shell git describe --always)-dirty-$(shell git diff | shasum -a256 | cut -c -6)
endif

TAG := $(shell date +v%Y%m%d)-$(GIT_VERSION)

all: build

echo:
	echo CONTEXT=$(CONTEXT)
	echo PROJECT=$(PROJECT)
	echo NAME=$(NAME)	
	echo NAMESPACE=$(NAMESPACE)

build-image:    	
	skaffold build --kube-context=${CONTEXT} -v info --file-output=image.json

set-image:
	kpt cfg set . image $(shell yq r image.json builds[0].tag )

apply:
	kpt --context=$(CONTEXT) live apply .	