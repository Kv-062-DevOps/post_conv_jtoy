POST Data Converting Service JSON to YAML:

1. Listen requests from frontend on port http :8082
2. Decode body from JSON.
3. Encode data into YAML.
4. Print structure on the screen.
5. Send YAML to DB service on port http :8083 path /add.
___

To start service:

go run main.go
or execute ./main

Now you can post requests from Frontend and put them into Database.
For testing without Frontend or/and Database, see "test_services" folder.
___

* Database structure:
emp_id;
first_name;
second_name;
types;
experience;
default_salary;
___

** Ports explication:
1) Python frontend: Listen on 8081 JSON from Go GET service;      Send JSON to 8082 Go POST service.
2) Go POST service: Listen on 8082 JSON from Python frontend;     Send YAML to 8083/add Python DB.
3) Python DB:       Listen on 8083/add YAML from Go POST service; Send YAML to 8084 Go GET service.
4) Go GET service:  Listen on 8084 YAML Python DB;                Send JSON to 8081 Python frontend. 
___
