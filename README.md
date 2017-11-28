# Password Generator

### Generate API Document:
apidoc -i PasswordGenerator/

## Restful API Test

### Set GOPATH:
export GOPATH = .../Password Generator
### Build:
go build main.go
### Run Server:
go run main.go
### Command Test Example:
curl -i "http://localhost:8080/passwords?length=8&specials=2&digits=2&limits=10"

### Using Restful API Testing tool:
Chrome extension: Postman

## Notice:
actucally not send/get json message, just to render simply template parameters
