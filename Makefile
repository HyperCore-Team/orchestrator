.PHONY: all clean version

GO ?= latest

BINARY_NAME=orchestrator
BUILDDIR = $(shell pwd)/build
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_COMMIT_FILE=$(shell pwd)/metadata/git_commit.go

build:
	go build --ldflags '-extldflags "-Wl,--allow-multiple-definition"' -o $(BUILDDIR)/${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

version:
	@echo "package metadata\n" > $(GIT_COMMIT_FILE)
	@echo "const (" >> $(GIT_COMMIT_FILE)
	@echo "\tGitCommit = \"${GIT_COMMIT}\"" >> $(GIT_COMMIT_FILE)
	@echo ")" >> $(GIT_COMMIT_FILE)

clean:
	go clean
	rm $(BINARY_NAME)
