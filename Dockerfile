FROM golang:1.22-alpine AS builder

COPY .. /app
WORKDIR /app

RUN go mod download
RUN go build -o ./bin/chat_service cmd/main.go

FROM alpine:3.20

WORKDIR /root

COPY --from=builder /app/bin/chat_service .
COPY --from=builder /app/local.env .
RUN apk add curl

CMD [ "./chat_service" ]
