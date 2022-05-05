FROM golang:1.18.1-buster AS builder

WORKDIR /app

COPY ./vendor ./vendor
COPY go.*  ./
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY entrypoint.sh entrypoint.sh


RUN CGO_ENABLED=0 GOOS=linux go build -o jarvis ./cmd

FROM alpine:latest  
COPY --from=builder /app /
RUN chmod +x /entrypoint.sh && chmod +x ./jarvis

ENTRYPOINT ["/entrypoint.sh"]