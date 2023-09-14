FROM golang:1.18

WORKDIR /go/src/github.com/Kaparouita/db-manager

COPY . .
COPY .env.docker .env


RUN go mod vendor
RUN go build -ldflags="-s -w" -mod=vendor

CMD "./github.com/Kaparouita/db-manager"

