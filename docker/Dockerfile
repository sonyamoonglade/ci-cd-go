FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./cmd/main.go

FROM alpine:3.16 as prod

WORKDIR /app

COPY --from=builder ./app/bin/main ./app/bin/main

CMD ["./app/bin/main"]


