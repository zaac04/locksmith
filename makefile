BINARY_NAME=locksmith
DATE =$(shell date "+%d %b %Y")
Version=v1.1
Maintainer="Issac"

BUILD_FILE := build.txt
ifeq (,$(wildcard $(BUILD_FILE)))
    $(shell echo 1 > $(BUILD_FILE))
endif
BUILD_NUMBER := $(shell cat $(BUILD_FILE))
NEW_BUILD_NUMBER := $(shell expr $(BUILD_NUMBER) + 1)

update_build_number:
	@echo $(NEW_BUILD_NUMBER) > $(BUILD_FILE)

show_build_number:
	@echo "Current build number is $(BUILD_NUMBER)"

build:
	go build -ldflags "-w -s -X 'locksmith/version.Maintainer=${Maintainer}' -X 'locksmith/version.Version=${Version}' -X 'locksmith/version.BuildNo=${NEW_BUILD_NUMBER}' -X 'locksmith/version.Date=${DATE}'" -o ${BINARY_NAME}
	upx --best --lzma ${BINARY_NAME}
	@$(MAKE) update_build_number

run:
	./${BINARY_NAME}

clean:
	go clean