BINARY_NAME=lockenv
EXE_NAME=lockenv.exe
DATE =$(shell date "+%d %b %Y")
Version=v2.0
Maintainer="Zaac04"

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

buildLinux:
	@$(MAKE) clean
	go build -o ${BINARY_NAME} -tags desktop,production -ldflags "-w -s -X 'locksmith/version.Maintainer=${Maintainer}' -X 'locksmith/version.Version=${Version}' -X 'locksmith/version.BuildNo=${NEW_BUILD_NUMBER}' -X 'locksmith/version.Date=${DATE}'" -o ./build/linux/${BINARY_NAME}
	upx --best --lzma ./build/linux/${BINARY_NAME}
	@$(MAKE) update_build_number
	cp ./build/linux/${BINARY_NAME} .

buildWindows:
	@$(MAKE) clean
	GOOS=windows GOARCH=amd64 go build -tags desktop,production -ldflags "-w -s -X 'locksmith/version.Maintainer=${Maintainer}' -X 'locksmith/version.Version=${Version}' -X 'locksmith/version.BuildNo=${NEW_BUILD_NUMBER}' -X 'locksmith/version.Date=${DATE}'" -o ./build/windows/${EXE_NAME}
	upx --best --lzma ./build/windows/${EXE_NAME}
	@$(MAKE) update_build_number
	cp ./build/windows/${EXE_NAME} .

Build: 
	@$(MAKE) clean
	go build -o ${BINARY_NAME} -tags desktop,production -ldflags "-w -s -X 'locksmith/version.Maintainer=${Maintainer}' -X 'locksmith/version.Version=${Version}' -X 'locksmith/version.BuildNo=${NEW_BUILD_NUMBER}' -X 'locksmith/version.Date=${DATE}'" -o ./build/linux/${BINARY_NAME}
	upx --best --lzma ./build/linux/${BINARY_NAME}
	GOOS=windows GOARCH=amd64 go build -tags desktop,production -ldflags "-w -s -X 'locksmith/version.Maintainer=${Maintainer}' -X 'locksmith/version.Version=${Version}' -X 'locksmith/version.BuildNo=${NEW_BUILD_NUMBER}' -X 'locksmith/version.Date=${DATE}'" -o ./build/windows/${EXE_NAME}
	upx --best --lzma ./build/windows/${EXE_NAME}
	@$(MAKE) update_build_number
	cp ./build/linux/${BINARY_NAME} .
	cp ./build/windows/${EXE_NAME} .

run:
	./${BINARY_NAME}

clean:
	go clean