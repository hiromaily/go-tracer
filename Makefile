
###############################################################################
# Golang formatter and detection
###############################################################################
.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: imports
imports:
	./scripts/imports.sh

###############################################################################
# Build
###############################################################################
.PHONY: build
build:
	go build -i -v -o ${GOPATH}/bin/tracer ./cmd/tracer/

run:
	tracer -t ./configs/settings.toml

###############################################################################
# build container image
###############################################################################
build-image:
	docker image build -t hirokiy/go-tracer .

#run-container:
#	docker container run -p 9999:9000 hirokiy/go-tracer

#change-image-name:
#	docker image tag go-tracer hirokiy/go-tracer

push-image:
	docker image push hirokiy/go-tracer
