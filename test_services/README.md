# These are two additional services for:
1. Test POST data converting service without Frontend and/or Database.
2. Test how your Frontend service posts new employee in Json. 
3. Test how your DB service recieves employee data in Yaml.
### Additional service for test sending JSON (such as Python Frontend works):
1. Generate data array.
2. Print data on the screen.
3. Encode data array into JSON.
4. Send JSON to MAIN service on port http :8082/a
### Additional service for test getting YAML (such as Python DB service works:
1. Listen requests on port http :8083 path /add
2. Decode body from YAML.
3. Print structure on the screen.
## To start testing:
1. In the first commandline console start Yaml listener:
```
go run getyamlsrv.go
```
_(or execute `./getyamlsrv`)_

2. In the second commandline console start main Post converting service:
```
go run main.go
```
_(or execute `./main`)_

3. In the third commandline console start Json sender:
```
go run sendjsonsrv.go
```
_(or execute `./sendjsonsrv`)_
### Check the output in appropriate commandline consoles:
1. Recieved message with employee field values for DB.
2. Responce code and employee field structure in the converter.
3. Source message with employee information, date, time and responce code.

If you see the same employee information in all services and responce 200 - it is O'k.
___
