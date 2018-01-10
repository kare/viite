#!/bin/bash

pkg=kkn.fi/viite
go test -coverprofile=coverage.txt -covermode=atomic $pkg
