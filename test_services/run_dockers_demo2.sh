#!/bin/bash

./stop_dockers_demo2.sh

docker start dyn
docker start -i init
docker start db
docker start get
docker start post
docker start front

sleep 1s
curl -X GET http://127.0.0.1:8081
curl -X GET http://127.0.0.1:8083/list
curl -sD - -o /dev/null http://127.0.0.1:8080/home

docker ps -s


