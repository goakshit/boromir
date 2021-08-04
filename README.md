## Boromir
Simple calc client/server program using gRPC

#### Requirements
- A system with working Go environment
   - Go package can be downloaded from this [link](https://golang.org/doc/install)
   - In mac, package will be installed to `/usr/local/go`
- Proto compiler
   - Compiler can installed using homebrew on MacOS using `brew install protobuf`

#### Steps
- Clone the repo: `git clone git@github.com:goakshit/boromir.git`
- Run `go get ./...` to download all dependencies
- Run `protoc ./internal/proto/calc/calc.proto --go_out=plugins=grpc:.` to generate go code from proto files
- Run `go build ./cmd/server && ./server` to build server binary and run it
- Run `go build ./cmd/client` to build client binary
- Run `./client -method <METHOD> -a <NUM1> -b <NUM2>` with arguments

#### Tests
- Tests are available for service under `svc/calc` directory
- Run `go test ./svc/calc` to run tests

#### Documentation
- Docs can be served using `godoc -http=:6060`
- Open `localhost:6060` in browser to view docs


### EDIT - 4th Aug 2021

#### Setup using docker
- Clone the repo: `git clone git@github.com:goakshit/boromir.git`
- Build server `docker build -t boromir-server . -f ./build/docker/server/Dockerfile`
- Run server `docker run -p 50051:50051 boromir-server` in seperate terminal
- Build client `docker build -t boromir-client . -f ./build/docker/client/Dockerfile`
- Run client `docker run -it --rm --network host boromir-client -method mul -a 10 -b 5` in different terminal
