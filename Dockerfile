FROM golang:latest

RUN go version
ENV GOPATH=/

WORKDIR /root/


CMD ["./reminder-app"]