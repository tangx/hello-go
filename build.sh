#!/bin/bash
#

set -x 


go build -v -o bin2/hello_`uname -s`_`uname -m`_${GITHUB_REF##*/} .
CGO_ENABLED=0 GOOS=linux go build -v -o bin2/hello_Linux_x86_64_${GITHUB_REF##*/} .


ls -al bin2/
