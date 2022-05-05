FROM golang:1.18.1-buster AS builder

WORKDIR /app

COPY ./vendor ./vendor
COPY go.*  ./
COPY ./cmd ./cmd
ADD ./pkg ./pkg
ADD ./script ./script

RUN CGO_ENABLED=0 GOOS=linux go build -o jarvis ./cmd

FROM alpine:latest  
COPY --from=builder /app /
COPY --from=builder /app/script /script
RUN chmod +x /script/* && chmod +x ./jarvis

ENTRYPOINT ["/script/entrypoint.sh"]