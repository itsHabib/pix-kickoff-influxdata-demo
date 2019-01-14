FROM golang:alpine as builder

RUN apk update && apk add git

COPY . $GOPATH/src/github.com/itshabib/demo
WORKDIR $GOPATH/src/github.com/itshabib/demo

# Get Dependencies & Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o demo main.go

# Get executable from builder
FROM alpine
WORKDIR /root
COPY --from=builder /go/src/github.com/itshabib/demo .
EXPOSE 8010
ENTRYPOINT ["./demo"]
