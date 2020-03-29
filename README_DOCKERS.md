### First run containers in the different commandline consoles to check how they work:
```
docker run --name="db" --network=host amazon/dynamodb-local
docker run --name="create" --network=host -e Db_url="http://127.0.0.1:8000" -e Region="local" vnikolayenko/db-service:latest-init-create
docker run --name="load" --network=host -e Db_url="http://127.0.0.1:8000" -e Region="local" vnikolayenko/db-service:latest-init-load
docker run --name="back" --network=host -e Server_port=8083 -e Db_url="http://127.0.0.1:8000" -e Region="local" vnikolayenko/db-service:latest-db-service

docker run --name="post" --network=host -e BACKADDR="127.0.0.1" nigth/postconv
docker run --name="get" --network=host -e HOST_PORT=8081 -e ENDPOINT="127.0.0.1:8083" nikitasadok/go-get-service
docker run --name="front" --network=host -e URL_GET="http://127.0.0.1:8081" -e URL_POST="http://127.0.0.1:8082" dimeder13/frontend
```
### Visit web pages:
http://127.0.0.1:8080 _Frontend_  
http://127.0.0.1:8081 _GET service_  
http://127.0.0.1:8083/list _DB service_  
### Stop all running containers:
```
docker stop front
docker stop post
docker stop get
docker stop back
docker stop db
```
Or execute `./test_services/stop_dockers_demo2.sh`
### Run full existing infrastructure again in the one commandline console:
```
docker start db
docker start -i create
docker start -i load
docker start back
docker start get
docker start post
docker start front
```
Or execute `./test_services/run_dockers_demo2.sh`
### Updating image sources
```
docker pull nigth/postconv
docker pull dimeder13/frontend
docker pull amazon/dynamodb-local
docker pull nikitasadok/go-get-service
docker pull vnikolayenko/db-service:latest-init-load
docker pull vnikolayenko/db-service:latest-db-service
docker pull vnikolayenko/db-service:latest-init-create
```
### DockerHub web pages:
https://hub.docker.com/r/nigth/postconv/tags _Size 72 MB (Compressed size 29 MB)_  
https://hub.docker.com/r/dimeder13/frontend/tags _Size 141 MB (Compressed size 46 MB)_  
https://hub.docker.com/r/nikitasadok/go-get-service/tags _Size 71 MB (Compressed size 28 MB)_  
https://hub.docker.com/r/amazon/dynamodb-local/tags _Size 611 MB (Compressed size 227 MB)_  
https://hub.docker.com/r/vnikolayenko/db-service/tags  _Sizes: **latest-db-service** 165 MB (Compressed 50 MB),  
**latest-init-create** Size 159 MB (Compressed Size 48 MB); **latest-init-load** Size 159 MB (Compressed size 48 MB);_
___

