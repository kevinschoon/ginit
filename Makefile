PWD := $(shell pwd)
PACKAGES ?= $(shell go list ./...|grep -v vendor)

.PHONY: all ci dep test

all: test

test:
	go $@ -v $(PACKAGES)
	go vet $(PACKAGES)

dep:
	dep ensure

ci: 
	docker run --rm -ti -v $(PWD):/go/src/github.com/mesanine/ginit -w /go/src/github.com/mesanine/ginit quay.io/vektorcloud/go:1.8 make test
