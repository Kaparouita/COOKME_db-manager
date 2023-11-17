# First stage: Builder
FROM golang:1.20 AS builder

WORKDIR /cook

COPY *.go ./
COPY core core
COPY handlers handlers
COPY ports ports
COPY repositories repositories
COPY server server
COPY ./.env.docker .env
COPY go.mod .
COPY go.sum .

COPY vendor/ /cook/vendor
RUN go build -ldflags="-s -w" -mod=vendor -o db
RUN ls -a

# Second stage: Final image
FROM debian:12-slim
WORKDIR /bridge
COPY --from=builder /cook/db /bridge/db
COPY --from=builder /cook/.env /bridge/.env
RUN ls
CMD "./db"
