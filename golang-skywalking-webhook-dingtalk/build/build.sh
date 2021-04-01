#!/usr/bin/env bash

export CGO_ENABLED=0
export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -mod=vendor -o bin/golang-skywalking-webhook main.go
