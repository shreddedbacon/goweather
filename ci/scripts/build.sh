#!/bin/sh
set -e -u -x
# Install git for go get

echo ">> Install git"
apk add --no-cache git

# set up directory stuff for golang
echo ">> Setup Directories"
mkdir -p /go/src/github.com/shreddedbacon/
ln -s $PWD/goweather-release /go/src/github.com/shreddedbacon/goweather
cd  /go/src/github.com/shreddedbacon/goweather
echo ">> Get"
go get -v .
cd -
echo ">> Build"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o built-release/goweather github.com/shreddedbacon/goweather

echo ">> Create artifact"
VERSION=$(cat ${VERSION_FROM})
cd built-release
tar czf goweather-linux-$VERSION.tar.gz goweather
