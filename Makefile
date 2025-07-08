# Copyright (C) 2025 Adam B Kaplan
#
# SPDX-License-Identifier: MIT

# Assisted by watsonx Code Assistant 

.PHONY: build, test

build:
	go build -o go-design-patterns main.go

test:
	go test -race -v --coverprofile=coverage.out --coverpkg=./... ./... -ginkgo.v

test-coverage: test
	go tool cover -o coverage.html -html coverage.out
