#!/bin/bash

docker ps -s

docker stop front
docker stop post
docker stop get
docker stop db
docker stop dyn

docker ps



