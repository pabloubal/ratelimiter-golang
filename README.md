# Rate Limiter
HTTP Rate limiter. Sits in front of http endpoints and limits the amount of requests to them.

## Target
- Be able to learn GoLang by implementing the rate limiter, as it contains many interesting aspects such as:
    * Rest api definition
    * Grpc api definition
    * Concurrency
    * How to structure a clean architecture in Go

## Requirements
- Add a remote endpoint with a global limit of requests per second
- Receive a request for any endpoint. If it's present in the rate limiter config, limit it, if not respond properly that it's not a valid endpoint.


## Things to improve
- Domain shouldn't be pegged to RestAPIErrors. They should be generic errors and get their translation to rest api errors in the infra layer

## Useful commands
- Generate proto API and structures: 
***$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ratelimit.proto***