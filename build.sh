#!/usr/bin/env bash
set -e

mkdir -p ./output

GOOS=linux go build -v -o ./output/critool ./cmd/cri-tool/critool.go
