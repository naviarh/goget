#!/bin/bash

go get -u github.com/naviarh/goget
cd $GOPATH/src/github.com/naviarh/goget
go build --ldflags="-s -w"
strip -s -x -v goget
upx -9v goget
cp ./goget $GOPATH/bin
echo
echo Congratulations, the utility is installed!
echo

