FROM golang:1.19-alpine 
WORKDIR /app/

RUN apk add git
COPY . .

RUN go env -w GOSUMDB=off

RUN go mod tidy
RUN go build -o app .

CMD ["./app"]