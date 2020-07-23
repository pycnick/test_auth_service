FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build ./cmd/api/auth_server.go

CMD ./auth_server