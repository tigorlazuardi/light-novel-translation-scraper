#!/bin/bash
set -e

go install github.com/evilmartians/lefthook@latest

go install github.com/vektra/mockery/v2@latest

go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lefthook install
