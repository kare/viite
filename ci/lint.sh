#!/bin/bash

GO111MODULE=off go get github.com/alecthomas/gometalinter
gometalinter --install
gometalinter ./...
