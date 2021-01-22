#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source lib.sh || { echo "Are you at repo root?"; exit 1; }

rm -rf ./dist/pakkretqc
mkdir -p ./dist/pakkretqc/bin
mkdir -p ./dist/pakkretqc/web


rm -rf web/dist
go run devtools/cmd/appbundler/main.go
export GOOS=linux 
export GOARCH=amd64
go build -tags=prod -o ./dist/pakkretqc/bin/pakkretqc ./cmd/pakkretqc/main.go

cp -R web/ ./dist/pakkretqc/web