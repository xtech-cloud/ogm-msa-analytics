APP_NAME := ogm-analytics
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
	rm -rf /tmp/ogm-analytics.db

.PHONY: call
call:
	gomu --registry=etcd --client=grpc call xtc.ogm.analytics Healthy.Echo '{"msg":"hello"}'
	# -------------------------------------------------------------------------
	gomu --registry=etcd --client=grpc call xtc.ogm.analytics Tracker.Record '{"appID":"myapp", "deviceID":"mydevice","userID":"myuser","eventID":"myevent"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.analytics Tracker.Record '{"appID":"myapp", "deviceID":"mydevice","userID":"myuser","eventID":"myevent", "parameter":"myparameter"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.analytics Generator.Record '{"appID":"myapp", "deviceID":"mydevice","userID":"myuser","eventID":"myevent"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.analytics Generator.Record '{"appID":"myapp", "deviceID":"mydevice","userID":"myuser","eventID":"myevent", "eventParameter":"ram"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.analytics Generator.Record '{"appID":"myapp", "deviceID":"mydevice","userID":"myuser","eventID":"myevent", "eventParameter":"333"}'



.PHONY: post
post:
	curl -X POST -d '{"msg":"hello"}' localhost/ogm/analytics/Healthy/Echo

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
docker:
	docker build -t xtechcloud/${APP_NAME}:${BUILD_VERSION} .
	docker rm -f ${APP_NAME}
	docker run --restart=always --name=${APP_NAME} --net=host -v /data/${APP_NAME}:/ogm -e MSA_REGISTRY_ADDRESS='localhost:2379' -e MSA_CONFIG_DEFINE='{"source":"file","prefix":"/ogm/config","key":"${APP_NAME}.yaml"}' -d xtechcloud/${APP_NAME}:${BUILD_VERSION}
	docker logs -f ${APP_NAME}
