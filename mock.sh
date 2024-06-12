#!/bin/bash

rm -rf mocks

IMAGE_NAME="vektra/mockery"

docker pull $IMAGE_NAME

docker run -it --rm -v $(pwd):/workspace -w /workspace vektra/mockery --all