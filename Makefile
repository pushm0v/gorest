.PHONY: default help build run loadtest

SHELL         = /bin/bash
APP_NAME      = gorest

default: help

help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build                 Compile the project.'
	@echo '    make run                   Build then run the project.'
	@echo '    make loadtest              Run load test using K6.'
	@echo '    make package               Build container.'
	@echo

build:
	@echo "Building ${APP_NAME}"
	go build -o bin/${APP_NAME}

run: build
	@echo "Running ${APP_NAME}"
	bin/${APP_NAME} ${ARGS}

loadtest:
	@echo "Running load test using K6"
	k6 run test/k6.js

package:
	@echo "Build container from Dockerfile"
	docker build -t ${APP_NAME}:latest .