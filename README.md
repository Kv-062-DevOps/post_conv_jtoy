# POST Data Converting Service from JSON to YAML
## How it works:
1. Listen requests from frontend on port http :8082.
2. Decode body from JSON.
3. Encode data into YAML.
4. Print structure on the screen.
5. Send YAML to DB service on port http :8083 path /add.  
*For testing without Frontend/Backend/Database, see the **"test_services"** folder:*  
https://github.com/Kv-062-DevOps/post_conv_jtoy/tree/master/test_services
### To start service from source code:
At first set the environment variables:  
- POSTPORT - "Post port" - on which port this service will listen for JSON data from Frontend;
- BACKPORT - "Backend port" - on which port the YAML converted data will be send to the Backend;
- BACKADDR - "Backend address" - IP or DNS name of the Backend **without** `http://` and path (`/add`or `/list`).  

For example:
```
POSTPORT=8082
export POSTPORT=8082
BACKPORT=8083
export BACKPORT=8083
BACKADDR="127.0.0.1"
export BACKADDR="127.0.0.1"
echo $POSTPORT
echo $BACKPORT
echo $BACKADDR
 
go run main.go 
```
_(or execute `./main`)_  
Now you can post requests from Frontend and put them into Database.
## Containerization
### It is recommended to use image from DockerHub:
https://hub.docker.com/r/nigth/postconv
```
docker run --rm --network=host -e BACKADDR="127.0.0.1" nigth/postconv
```
_Or if you want to customize ports:_
```
docker run --rm --network=host -e POSTPORT=8082 -e BACKPORT=8083 -e BACKADDR="127.0.0.1" nigth/postconv
```
Update to the latest image:
```
docker pull nigth/postconv
```
Also you can create and use local image:
```
docker build -t postlocal .
docker run --rm --network=host -e BACKADDR="127.0.0.1" postlocal
```
_Or if you want to customize ports:_
```
docker run --rm --network=host -e POSTPORT=8082 -e BACKPORT=8083 -e BACKADDR="127.0.0.1" postlocal
```
How to run the whole project with all services see **README_DOCKERS.md**:  
https://github.com/Kv-062-DevOps/post_conv_jtoy/blob/master/README_DOCKERS.md
## Additional information
### Namespace for the Kubernetes:
* front
* get
* post
* back
* load
* create
* db  
For Minikube services add "**-srv**" (example "_front-srv_"), for deployments add "**-dep**" (example: "_front-dep_").
### Project DataBase Structure:
- emp_id
- first_name
- second_name
- types
- experience
- default_salary
### Project Ports Explication:
+ **Python frontend**: Listen on 8081 JSON from Go GET service;      Send JSON to 8082 Go POST service.
+ **Go POST service**: Listen on 8082 JSON from Python frontend;     Send YAML to 8083/add Python DB.
+ **Python Backend**:  Listen on 8083/add YAML from Go POST service; Send YAML to 8084 Go GET service.
+ **Go GET service**:  Listen on 8084 YAML Python DB;                Send JSON to 8081 Python frontend. 
___
