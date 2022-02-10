#!/bin/bash

echo "Building application for multi OS"
cd ../app

# build for linux and ubuntu
GOOS=linux GOARCH=amd64 go build -o ../build/app_linux

# build for windows
GOOS=windows GOARCH=amd64 go build -o ../build/app_windows

#build for MacOs
GOOS=darwin GOARCH=amd64 go build -o ../build/app_darwin

echo "Build Atifacts at /build"
echo "Done."