#!/bin/sh
set -e -u -x
# Install git for go get

echo ">> Install git"
# apk add --no-cache git
# apk add --no-cache --virtual .build-deps \
#  		bash \
#  		gcc \
#     musl-dev \
#     linux-headers

#apk add --no-cache gcc-arm-none-eabi
apt-get update -y
apt-get install git gcc-arm-linux-gnueabi libc6-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev build-essential -y

# set up directory stuff for golang
echo ">> Setup Directories"
mkdir -p /go/src/github.com/shreddedbacon/
ln -s $PWD/goweather-release /go/src/github.com/shreddedbacon/goweather
cd  /go/src/github.com/shreddedbacon/goweather
echo ">> Get"
go get -v .
cd -
echo ">> Build X86"
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o built-release-x86/goweather github.com/shreddedbacon/goweather
echo ">> Build ARM"
CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=5 go build -a -o built-release-arm/goweather github.com/shreddedbacon/goweather

VERSION=$(cat ${VERSION_FROM})
echo ">> Create X86 artifact"
cd built-release-x86
tar czf goweather-linux-x86-$VERSION.tar.gz goweather

echo ">> Create ARM artifact"
cd ../built-release-arm
tar czf goweather-linux-arm-$VERSION.tar.gz goweather
