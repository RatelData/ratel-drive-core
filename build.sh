#!/bin/bash

yarn --cwd ./ui build

bash build_docs.sh

# clean up the old build
rm -rf build

# build the application
mkdir -p build
go build -o build/

# COPY STATIC RESOURCES
## copy prodcution config files
cp -rf config/production build/config/
## copy ui resources
cp -rf ui/build build/ui
## copy docs
mkdir -p build/docs && cp -f docs/swagger* build/docs/