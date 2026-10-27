FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o eventhub ./cmd

FROM alpine:latest

WORKDIR /srv

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/eventhub /srv/eventhub

RUN mkdir -p /srv/uploads

EXPOSE 9000

CMD [ "/srv/eventhub" ]