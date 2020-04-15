#!/bin/bash
#

set -x 

TAG=${GITHUB_REF##*/}
Machine=`uname -m`

BINARY=hello-go

CGO_ENABLED=0 GOOS=darwin go build -v -o bin/${BINARY}_Darwin_${Machine}_${TAG} .
CGO_ENABLED=0 GOOS=linux go build -v -o bin/${BINARY}_Linux_${Machine}_${TAG} .



