#!/bin/bash

function failIfNonZero {
  if [[ ! 0 -eq $1 ]]; then echo FAIL; else echo PASS; fi
}

PORT=8080
CONTAINER=`docker run --detach -p $PORT:8080 render`
sleep 1 # time for container to get up

ADDR=localhost:$PORT
curl -F "file=@model.blend" $ADDR/render/ > /dev/null

curl -I $ADDR/result/ | grep image/png > /dev/null
failIfNonZero $?

docker rm --force $CONTAINER > /dev/null
