#Running Steps
 ```
put the app in your src directory
make sure you have .env file in the following directorries
...\go\src\halan
```
#Run the following commands
 ```
go mod download
```
#Swagger Documentation
 ```
swag init --parseDependency true
to run the swagger documentation 
ex: http://127.0.0.1:9098/halan/test/swagger/index.html/
```

#Run Test Cases
 ```
 from project directory ex(C:\Users\User\go\src\halan)
 go to test folder  C:\Users\User\go\src\halan\test
 run go test 
 (there are 27 test cases for all 9 problems)
```
#Run the app
```
go run main.go
```
