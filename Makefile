SHELL := /bin/bash

cmds := $(shell find cmd -name \*main.go | awk -F'/' '{print $$2}')

clean:
	rm -rf .build

prepare_build: clean
	mkdir -p .build && \
	echo "$$(git describe --tags --exact-match 2> /dev/null ||  git symbolic-ref --short -q HEAD)-$$(git rev-parse --short HEAD)" > .build/APP_VERSION && \
	cd .build && \
	rsync -avr --exclude='/.git' \
			   --exclude='/.idea' \
			   --exclude='/.github' \
			   --filter=':- .gitignore' \
			   ../ . && \
	npm ci && \
	env GOARCH=amd64 GOOS=linux go get ./...

build: prepare_build
	@for function in $(cmds) ; do \
		cd .build && \
		env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/$$function cmd/$$function/main.go ; \
	done

deploy_sandbox: build
	cd .build && \
	sls deploy --stage sandbox --aws-profile sandbox && \
	cd ../ && rm -rf .build
