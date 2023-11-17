FROM golang:1.18

WORKDIR /go/src/module

COPY *.go ./
COPY core core
COPY handlers handlers
COPY ports ports
COPY repositories repositories
COPY server server
COPY ./.env.docker .env
COPY go.mod .
COPY go.sum .

COPY vendor/ /go/src/module/vendor
RUN go build -ldflags="-s -w" -mod=vendor
RUN ls -a

FROM debian:12-slim
WORKDIR /bridge
COPY --from=builder /go/src/module/db /bridge/db
COPY --from=builder /go/src/module/.env /bridge/.env
RUN ls
CMD "./db"
