#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <package_directory>"
    exit 1
fi

PACKAGE_DIR="$1"
echo "Running tests in package: $PACKAGE_DIR"

go test -coverprofile=coverage.out "$PACKAGE_DIR"

if [ $? -eq 0 ]; then
    echo "Tests ran successfully. Generating coverage report..."
    go tool cover -html=coverage.out
else
    echo "Tests failed or encountered an error"
fi