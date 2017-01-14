FROM golang:latest

RUN go get -u github.com/antihax/mock-esi
RUN go get -u github.com/antihax/mock-esi/go
RUN go install github.com/antihax/mock-esi

ENTRYPOINT /go/bin/mock-esi

EXPOSE 8080