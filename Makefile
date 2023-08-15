LOCAL_BIN=$(CURDIR)/bin
DIR=$(shell pwd)
BINDIR=${DIR}/bin
SMARTIMPORTS=${BINDIR}/smartimports
PACKAGE=github.com/pav5000/api-example/cmd/main
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
LINTVER=v1.54.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
IMAGE=carebot
MOCKGEN=${BINDIR}/mockgen

precommit: format build test lint

build: bindir
	GOOS=linux go build -o ${BINDIR}/bot ${PACKAGE}

run:
	go run ${PACKAGE}

test:
	go test ./...

.PHONY: install-smartimports
install-smartimports: bindir
	test -f ${SMARTIMPORTS} || GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest

.PHONY: format
format: install-smartimports
	${SMARTIMPORTS}

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

lint: install-lint
	${LINTBIN} run

bindir:
	mkdir -p ${BINDIR}
