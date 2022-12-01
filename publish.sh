#!/bin/bash
version=$(<VERSION)
echo "uploading version: $version"
go mod tidy
git add .
git commit -m "automate: changes for v$version"
git tag v$version
git push origin v$version
GOPROXY=proxy.golang.org go list -m github.com/messagepack-schema/go@v$version
echo "version published."