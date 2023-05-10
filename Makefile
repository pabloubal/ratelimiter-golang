generate_grpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/ratelimit.proto

build-app:
	CGO_ENABLED=0 GOOS=linux go build

build-docker:
	docker build --tag ratelimiter .

build-all: build-app build-docker

run:
	docker run -p 80:80 -p 8080:8080 -p 50050:50050 -p 50051:50051 ratelimiter