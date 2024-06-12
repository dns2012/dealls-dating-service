#!/bin/bash

IMAGE_NAME="dealls-dating-service-test"
DOCKERFILE_NAME="Dockerfile.test"

docker build -t $IMAGE_NAME -f $DOCKERFILE_NAME .

docker run -it --rm -v $(pwd):/workspace $IMAGE_NAME

docker rmi $IMAGE_NAME