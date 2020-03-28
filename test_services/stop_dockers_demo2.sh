#!/bin/bash

docker ps

docker stop front
docker stop post
docker stop get
docker stop back
docker stop db
docker stop load
docker stop create

docker ps



