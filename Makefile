export PROJ_PATH=github.com/application-research/delta-metrics-rest

export DATE := $(shell date +%Y.%m.%d-%H%M)
export LATEST_COMMIT := $(shell git log --pretty=format:'%h' -n 1)
export BRANCH := $(shell git branch |grep -v "no branch"| grep \*|cut -d ' ' -f2)
export BUILT_ON_IP := $(shell [ $$(uname) = Linux ] && hostname -i || hostname )
export BIN_DIR=./bin
export PACKR2_EXECUTABLE := $(shell command -v packr2  2> /dev/null)
export SWAG_EXECUTABLE := $(shell command -v swag  2> /dev/null)
export RUNTIME_VER := $(shell go version)

export BUILT_ON_OS=$(shell uname -a)
ifeq ($(BRANCH),)
BRANCH := master
endif

export COMMIT_CNT := $(shell git rev-list HEAD | wc -l | sed 's/ //g' )
export BUILD_NUMBER := ${BRANCH}-${COMMIT_CNT}
export COMPILE_LDFLAGS=-s -X "main.BuildDate=${DATE}" \
                          -X "main.LatestCommit=${LATEST_COMMIT}" \
                          -X "main.BuildNumber=${BUILD_NUMBER}" \
                          -X "main.BuiltOnIP=${BUILT_ON_IP}" \
                          -X "main.BuiltOnOs=${BUILT_ON_OS}" \
						  -X "main.RuntimeVer=${RUNTIME_VER}"

build_info: check_prereq ## Build the container
	@echo ''
	@echo '---------------------------------------------------------'
	@echo 'BUILT_ON_IP       $(BUILT_ON_IP)'
	@echo 'BUILT_ON_OS       $(BUILT_ON_OS)'
	@echo 'DATE              $(DATE)'
	@echo 'LATEST_COMMIT     $(LATEST_COMMIT)'
	@echo 'BRANCH            $(BRANCH)'
	@echo 'COMMIT_CNT        $(COMMIT_CNT)'
	@echo 'BUILD_NUMBER      $(BUILD_NUMBER)'
	@echo 'COMPILE_LDFLAGS   $(COMPILE_LDFLAGS)'
	@echo 'PATH              $(PATH)'
	@echo 'PACKR2_EXECUTABLE $(PACKR2_EXECUTABLE)'
	@echo 'SWAG_EXECUTABLE   $(SWAG_EXECUTABLE)'
	@echo 'RUNTIME_VER       $(RUNTIME_VER)'
	@echo '---------------------------------------------------------'
	@echo ''


####################################################################################################################
##
## help for each task - https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
##
####################################################################################################################
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help



####################################################################################################################
##
## Build of binaries
##
####################################################################################################################
all: dmr test ## build dmr and run tests

binaries: dmr ## build binaries in bin dir

create_dir:
	@mkdir -p $(BIN_DIR)

check_prereq: create_dir
ifndef PACKR2_EXECUTABLE
	go get -u github.com/gobuffalo/packr/v2/packr2
endif
	$(warning "found packr2")

ifndef SWAG_EXECUTABLE
	go get -u github.com/swaggo/swag/cmd/swag
endif
	$(warning "found swag")



build_app: create_dir
		packr2 build -o $(BIN_DIR)/$(BIN_NAME) -a -ldflags '$(COMPILE_LDFLAGS)' $(APP_PATH)


dmr: build_info ## build dmr binary in bin dir
	@echo "build dmr server"
	swag init --dir . --parseDependency --parseInternal -g ./main.go
	make BIN_NAME=dmr APP_PATH=$(PROJ_PATH) build_app
	@echo ''
	@echo ''



####################################################################################################################
##
## Cleanup of binaries
##
####################################################################################################################

clean_binaries: clean_dmr  ## clean all binaries in bin dir


clean_binary: ## clean binary in bin dir
	rm -f $(BIN_DIR)/$(BIN_NAME)

clean_dmr: ## clean dmr
	make BIN_NAME=dmr clean_binary



test: ## run tests
	go test -v $(PROJ_PATH)

fmt: ## run fmt on project
	#go fmt $(PROJ_PATH)/...
	gofmt -s -d -w -l .

doc: ## launch godoc on port 6060
	godoc -http=:6060

deps: ## display deps for project
	go list -f '{{ join .Deps  "\n"}}' . |grep "/" | grep -v $(PROJ_PATH)| grep "\." | sort |uniq

lint: ## run lint on the project
	golint ./...

staticcheck: ## run staticcheck on the project
	staticcheck -ignore "$(shell cat .checkignore)" .

vet: ## run go vet on the project
	go vet .

tools: ## install dependent tools for code analysis
	go get -u github.com/gogo/protobuf
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/gogo/protobuf/jsonpb
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u honnef.co/go/tools
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/fzipp/gocyclo
	go get -u golang.org/x/lint/golint
	go get -u github.com/gobuffalo/packr/v2/packr2



regen: ## regenerate generated code
	gen \
     --sqltype=postgres \
     --connstr='host=kEOpaaIn7ZRDLL9IGrKYOH2MpUEavWWg@dpg-cfto8d9a6gdotcfptsrg-a.oregon-postgres.render.com user=deltadb_metrics_user password=kEOpaaIn7ZRDLL9IGrKYOH2MpUEavWWg dbname=deltadb_metrics port=5432' \
     --database=estuary \
     --templateDir=./templates \
     --model=model \
     --dao=dao \
     --api=api \
     --out=./ \
     --module=github.com/application-research/delta-metrics-rest \
     --gorm \
     --makefile \
     --overwrite \
     --host=localhost \
     --port=8080 \
     --rest \
     --listen=:8080 \
     --scheme=http \
     --generate-dao \
     --file_naming='{{.}}' \
     --model_naming='{{FmtFieldName .}}' \
     --swagger_version=1.0 \
     --swagger_path=/ \
     --swagger_tos= \
     --swagger_contact_name=Me \
     --swagger_contact_url=http://me.com/terms.html \
     --swagger_contact_email=me@me.com


