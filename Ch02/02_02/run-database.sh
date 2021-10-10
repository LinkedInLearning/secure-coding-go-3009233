#!/bin/bash

docker run \
    -e POSTGRES_HOST_AUTH_METHOD=trust \
    -p 5432:5432 \
    postgres:13-alpine