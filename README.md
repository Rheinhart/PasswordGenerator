#Password Generator

## Restful API Test
### Build:
go build main.go
### Run Server:
go run main.go
### Command Test Example:
curl -i "http://localhost:8080/passwords?length=8&specials=2&digits=2&limits=10"
### Generate API Document:
apidoc -i PasswordGenerator/
### Using Restful API Testing tool:
Chrome extension: Postman

## Notice:
actucally not send/get json message, just to render simply template parameters