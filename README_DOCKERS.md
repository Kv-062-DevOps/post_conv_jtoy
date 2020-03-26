# First run containers in the different commandline consoles to check how they work:
```
docker run --name="dyn" --network=host amazon/dynamodb-local
docker run --name="init" --network=host -e Db_url="http://127.0.0.1:8000" -e Region="local" vnikolayenko/db_service:latest_db_init
docker run --name="db" --network=host -e Db_url="http://127.0.0.1:8000" -e Server_port="8083" -e Region="local" vnikolayenko/db_service:latest_db_service

docker run --name="post" --network=host nigth/post_conv_jtoy
docker run --name="get" --network=host -e ENDPOINT="http://127.0.0.1:8083/list" -e HOST_PORT=":8081" nikitasadok/go-get-service
docker run --name="front" --network=host dimeder13/frontend:latest
```
# Visit web pages:

http://127.0.0.1:8080

http://127.0.0.1:8081

http://127.0.0.1:8083/list

# Stop all running containers:
```
docker stop front
docker stop post
docker stop get
docker stop db
docker stop dyn
```
_Wait for a minute while service DB is stopping, it stops more longer then other_
# Run full existing infrastructure again in the one commandline console:
```
docker start dyn
docker start -i init
docker start db
docker start get
docker start post
docker start front
```
# Updating image sources
```
docker pull amazon/dynamodb-local
docker pull dimeder13/frontend
docker pull nigth/post_conv_jtoy
docker pull nikitasadok/go-get-service
docker pull vnikolayenko/db_service:latest_db_init
docker pull vnikolayenko/db_service:latest_db_service
```
# DockerHub web pages:

https://hub.docker.com/r/dimeder13/frontend/tags _46 MB_

https://hub.docker.com/r/nigth/post_conv_jtoy/tags _29 MB_

https://hub.docker.com/r/nikitasadok/go-get-service/tags _28 MB_

https://hub.docker.com/r/vnikolayenko/db_service/tags (_349 MB_  :latest_db_init; _353 MB_  :latest_db_service)

https://hub.docker.com/r/amazon/dynamodb-local/tags _227 MB_
___

