FROM golang:1.13 AS build
WORKDIR /go/src/github.com/antoinedao/aecdelta-go-server
COPY go ./go
COPY pkg ./pkg
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/github.com/antoinedao/aecdelta-go-server/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
