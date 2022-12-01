#!/bin/bash
version=$(<VERSION)
echo "uploading version: $version"
go mod tidy
git add .
git commit -m "automate: changes for v$version"
git tag v$version
git push origin v$version
echo "version published."