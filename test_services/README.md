Two additional service for:
a) testing POST data converting service without Frontend and/or Database.
b) testing your Frontend service to Post new employee in Json. 
c) testing your DB service to recieve employee data in Yaml.
___

Additional service for test sending JSON (such as Python Frontend works):

1. Generate data array.
2. Print data on the screen.
3. Encode data array into JSON.
4. Send JSON to MAIN service on port http :8082
___

Additional service for test getting YAML (such as Python DB service works:

1. Listen requests on port http :8083 path /add
2. Decode body from YAML.
3. Print structure on the screen.
___

To start testing:

1. In first commandline console start Yaml listener:
go run getyamlsrv.go
or execute ./getyamlsrv
2. In second commandline console start main Post converting service:
go run main.go
or execute ./main
3. In third commandline console start Json sender:
go run sendjsonsrv.go
or execute ./sendjsonsrv

Check the output in appropriate commandline consoles:
1. Recieved message with employee field values for DB.
2. Responce code and employee field structure in the converter.
3. Source message with employee information, date, time and responce code.

___
