### First run containers in the different commandline consoles to check how they work:
```
docker run --name="dyn" --network=host amazon/dynamodb-local
docker run --name="init" --network=host -e Db_url="http://127.0.0.1:8000" -e Region="local" vnikolayenko/db_service:latest_db_init
docker run --name="db" --network=host -e Server_port="8083" -e Db_url="http://127.0.0.1:8000" -e Region="local" vnikolayenko/db_service:latest_db_service

docker run --name="post" --network=host  -e PORT=":8082" -e DBLINK="http://127.0.0.1:8083/add" nigth/post_conv_jtoy
docker run --name="get" --network=host -e HOST_PORT=":8081" -e ENDPOINT="http://127.0.0.1:8083/list"  nikitasadok/go-get-service
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
docker stop db
docker stop dyn
 
```
Or execute `./test_services/stop_dockers_demo2.sh`

_Wait for a minute while service DB is stopping, it stops more longer then other_
### Run full existing infrastructure again in the one commandline console:
```
docker start dyn
docker start -i init
docker start db
docker start get
docker start post
docker start front
 
```
Or execute `./test_services/run_dockers_demo2.sh`

### Updating image sources
```
docker pull amazon/dynamodb-local
docker pull dimeder13/frontend
docker pull nigth/post_conv_jtoy
docker pull nikitasadok/go-get-service
docker pull vnikolayenko/db_service:latest_db_init
docker pull vnikolayenko/db_service:latest_db_service
 
```
### DockerHub web pages:

https://hub.docker.com/r/dimeder13/frontend/tags _Size 141 MB (Compressed size 46 MB)_

https://hub.docker.com/r/nigth/post_conv_jtoy/tags _Size 72 MB (Compressed size 29 MB)_

https://hub.docker.com/r/amazon/dynamodb-local/tags _Size 611 MB (Compressed size 227 MB)_

https://hub.docker.com/r/nikitasadok/go-get-service/tags _Size 71 MB (Compressed size 28 MB)_

https://hub.docker.com/r/vnikolayenko/db_service/tags 
_(latest_db_init: Size 996 MB (Compressed size 349 MB); latest_db_service: 1.01 GB (Compressed size 353 MB))_
___

