# POST Data Converting Service from JSON to YAML

## How it works:
1. Listen requests from frontend on port http :8082.
2. Decode body from JSON.
3. Encode data into YAML.
4. Print structure on the screen.
5. Send YAML to DB service on port http :8083 path /add.

*For testing without Frontend or/and Database, see **"test_services"** folder.*

### To start service from source code:
```
go run main.go 
```
*(or execute ./main)*
Now you can post requests from Frontend and put them into Database.

## Containerization

### It is recommended to use image from DockerHub:
```
docker run --network=host nigth/post_conv_jtoy
```
Also you can create and use local image:
1) Create local Docker image:
```
docker build post_conv_jtoy
```
2) Wait for the message: 
**Successfully built IMAGE_NAME**
*(example: Successfully built    76db9a5645e9)*
3) Start Docker service:
```
docker run --net=host IMAGE_NAME
```
*(example: docker run --net=host 76db9a5645e9)*
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
