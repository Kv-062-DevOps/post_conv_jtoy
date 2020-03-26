# POST Data Converting Service from JSON to YAML

## How it works:
1. Listen requests from frontend on port http :8082.
2. Decode body from JSON.
3. Encode data into YAML.
4. Print structure on the screen.
5. Send YAML to DB service on port http :8083 path /add.

*For testing without Frontend or/and Database, see **"test_services"** folder.*

### To start service from source code:
At first set the environment variables POST_SRV_PORT and DB_SRV_LINK , for example:
```
POST_SRV_PORT=":8082"
echo $POST_SRV_PORT
DB_SRV_LINK="http://127.0.0.1:8083/add"
echo $DB_SRV_LINK

go run main.go 
```
_(or execute ./main)_

Now you can post requests from Frontend and put them into Database.

## Containerization

### It is recommended to use image from DockerHub:
https://hub.docker.com/r/nigth/post_conv_jtoy
```
docker run --network=host -e DB_SRV_LINK="http://127.0.0.1:8083/add" -e POST_SRV_PORT=":8082" nigth/post_conv_jtoy
```
To update your image:
```
docker pull nigth/post_conv_jtoy
```
Also you can create and use local image:
1) Create local Docker image:
```
docker build post_conv_jtoy
```
2) Wait for the message: 
**Successfully built IMAGE_NAME**
_(example: Successfully built    76db9a5645e9)_

3) Start the Docker service:
```
docker run --network=host -e DB_SRV_LINK="http://127.0.0.1:8083/add" -e POST_SRV_PORT=":8082" IMAGE_NAME
```
*(example: docker run --network=host -e DB_SRV_LINK="http://127.0.0.1:8083/add" -e POST_SRV_PORT=":8082" 76db9a5645e9)*
## Additional information

### Project DataBase Structure:
- emp_id;
- first_name;
- second_name;
- types;
- experience;
- default_salary;

### Project Ports Explication:
+ **Python frontend**: Listen on 8081 JSON from Go GET service;      Send JSON to 8082 Go POST service.
+ **Go POST service**: Listen on 8082 JSON from Python frontend;     Send YAML to 8083/add Python DB.
+ **Python DB**:       Listen on 8083/add YAML from Go POST service; Send YAML to 8084 Go GET service.
+ **Go GET service**:  Listen on 8084 YAML Python DB;                Send JSON to 8081 Python frontend. 
___
