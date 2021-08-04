# Use docker alpine image
FROM golang:1.16-alpine

ENV GO111MODULE=on

# Create 'code' dir
RUN mkdir /code

# Change to 'code' dir
WORKDIR /code

# Install protobuf, git, and protoc-gen-go
RUN apk add --no-cache protobuf git \
  && go get github.com/golang/protobuf/protoc-gen-go \
  && cp /go/bin/protoc-gen-go /usr/bin/

# Copy go mod and go sum files
COPY go.mod .
COPY go.sum .

# Install the go modules dependencies
RUN go mod download

# Copy the entire dir
COPY . .

# Generate go code from proto files
RUN protoc ./internal/proto/calc/calc.proto --go_out=plugins=grpc:.

# Run tests
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test ./...

# Build server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/client-exec ./cmd/client

ENTRYPOINT ["/bin/client-exec"]
CMD ["-method", "add", "-a", "1", "-b", "2"]