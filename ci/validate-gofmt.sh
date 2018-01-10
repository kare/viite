#!/bin/bash

# $pkg is relative path to package
pkg=.

gofmt -s -l $pkg
if [ $? -ne 0 ]; then
	echo "validate-gofmt.sh: error $pkg" 1>&2
	exit 1
fi
exit 0
