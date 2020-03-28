#!/bin/bash

./stop_dockers_demo2.sh

docker start db
docker start -i create
docker start -i load
docker start back
docker start get
docker start post
docker start front

sleep 1s
echo "http://127.0.0.1:8080/home"
curl -X GET http://127.0.0.1:8081
curl -X GET http://127.0.0.1:8083/list
curl -sD - -o /dev/null http://127.0.0.1:8080/home

docker ps


