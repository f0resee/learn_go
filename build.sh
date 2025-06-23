#!/usr/bin/env bash
set -e

mkdir -p ./output

GOOS=linux go build -v -o ./output/critool ./cmd/cri-tool/critool.go
GOOS=linux go build -v -o ./output/stress ./cmd/stress-tool/stress.go
GOOS=linux go build -v -o ./output/numa ./cmd/numa
