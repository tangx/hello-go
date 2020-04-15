#!/bin/bash
#

set -x 

TAG=${GITHUB_REF##*/}
Machine=`uname -m`

CGO_ENABLED=0 GOOS=darwin go build -v -o bin/hello_${GOOS}_${Machine}_${TAG} .
CGO_ENABLED=0 GOOS=linux go build -v -o bin/hello_${GOOS}_${Machine}_${TAG} .



