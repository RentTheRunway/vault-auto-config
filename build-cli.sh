#!/usr/bin/env bash

EXT=""

if [[ ${GOOS} == "windows" ]]; then
  EXT=".exe"
fi

docker run -t --rm \
  -v ${PWD}:/src \
  --env CGO_ENABLED=0 \
  --env GOOS=${GOOS:-darwin} \
  --env GOARCH=${GOARCH:-amd64} \
  --workdir /src/cmd \
  golang:1.14-alpine \
  go build -o "/src/vault-auto-config${EXT}" /src/cmd/vault-auto-config
