APP_NAME := omo-msa-analytics
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )

.PHONY: build
build: 
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: run
run: 
	./bin/${APP_NAME}

.PHONY: install
install: 
	go install

.PHONY: clean
clean: 
	rm -rf /tmp/msa-analytics.db

.PHONY: call
TOKEN := $(shell cat /tmp/msa-token)
call:
	MICRO_REGISTRY=consul micro call omo.msa.analytics Record.Wake '{"serialNumber":"Z0000001", "softwareFamily":"Test", "softwareVersion": "1.0.0", "systemFamily": "Alpine", "systemVersion": "3.11", "deviceModel":"DELL XPS", "deviceType":"PC", "profile": "myprofile"}'
	MICRO_REGISTRY=consul micro call omo.msa.analytics Query.Agent '{"offset":0, "count":60}'
	MICRO_REGISTRY=consul micro call omo.msa.analytics Record.Wake '{"serialNumber":"Z0000001", "softwareFamily":"Test", "softwareVersion": "1.1.0", "systemFamily": "Alpine", "systemVersion": "3.12", "deviceModel":"DELL XPS", "deviceType":"PC", "profile": "myprofile2"}'
	MICRO_REGISTRY=consul micro call omo.msa.analytics Query.Agent '{"offset":0, "count":60}'

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
docker:
	docker build . -t omo-msa-startkit:latest
