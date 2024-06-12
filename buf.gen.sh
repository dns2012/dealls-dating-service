#!/bin/bash

rm -rf proto docs

IMAGE_NAME="dealls-dating-service-protobuf"
DOCKERFILE_NAME="Dockerfile.buf"

docker build -t $IMAGE_NAME -f $DOCKERFILE_NAME .

docker run -it --rm -v $(pwd):/workspace $IMAGE_NAME

docker rmi $IMAGE_NAME