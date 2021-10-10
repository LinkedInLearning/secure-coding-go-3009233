#!/bin/bash

docker run \
    --rm -it \
    --network=host \
    postgres:13-alpine \
    psql -U postgres -h localhost