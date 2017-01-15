FROM golang:latest

RUN openssl genrsa -out server.key 2048
RUN openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 -subj "/C=US/ST=Somewhere/O=personal/L=Cheese/CN=localhost/OU=IT/emailAddress=example@example.org/"
RUN go get github.com/gorilla/mux
RUN go get -u github.com/antihax/mock-esi
RUN go install github.com/antihax/mock-esi

ENTRYPOINT /go/bin/mock-esi

EXPOSE 8080