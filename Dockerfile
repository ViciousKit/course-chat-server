FROM golang:1.22-alpine AS builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o ./bin/chat_service cmd/main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/bin/chat_service .
COPY --from=builder /app/${ENV_FILE} .
RUN apk add curl

CMD [ "./chat_service" ]
