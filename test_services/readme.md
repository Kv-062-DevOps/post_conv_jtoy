Service 1. Additional, send JSON.
1. Create data string.
2. Print data on the screen.
2. Encode data string into JSON.
3. Send JSON to service 2 on port http :8082
_____________

Service 2. MAIN SERVICE (NUMBER 3 FOR DEMO LAB). ENCODING JSON->YAML.
1. Listen incoming traffic on port http :8082
2. Decode structure from JSON.
3. Print data on the screen.
4. Encode data into YAML.
5. Send YAML to service 3 on port http :8083
_____________

Service 3. Additional, get YAML.
1. Listen incoming traffic on port http :8083
2. Decode structure from YAML.
3. Print data on the screen.
________

* Database structure:
emp_id
first_name
second_name
types
experience
default_salary
________
