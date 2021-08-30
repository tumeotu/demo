#!/bin/sh

docker build -t tumeo/demo_app .
docker run --publish 3000:3000 tumeo/demo_app