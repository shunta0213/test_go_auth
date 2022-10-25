FROM golang:1.19-alpine AS builder
WORKDIR /build-app

COPY . .

RUN go mod tidy
RUN go build -o app .

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /build-app/app ./

CMD ["./app"]