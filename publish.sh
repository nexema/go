#!/bin/bash
version=`cat VERSION`
echo "Creating git tag for version $version"
git tag $version
git push origin $version
GOPROXY=proxy.golang.org go list -m github.com/nexema/go@$version