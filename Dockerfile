FROM golang:1.18.1-buster AS builder

WORKDIR /app

COPY ./vendor ./vendor
COPY go.*  ./
COPY ./cmd ./cmd
ADD ./pkg ./pkg
ADD ./script ./script

RUN CGO_ENABLED=0 GOOS=linux go build -o jarvis ./cmd

FROM alpine:latest  
WORKDIR /app
COPY --from=builder /app /app
COPY --from=builder /app/script /app/script
RUN chmod +x /app/script/* && chmod +x ./jarvis

ENTRYPOINT ["/app/script/entrypoint.sh"]