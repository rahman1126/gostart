# Build Stage
FROM golang:1.13.7 AS builder

WORKDIR $GOPATH/src/gostart

ADD . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app cmd/main.go

# Final Stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /root/

RUN mkdir logs
RUN chmod 777 logs

COPY --from=builder /go/src/gostart/app .
COPY --from=builder /go/src/gostart/config/config.json ./config/

CMD ["./app"]
EXPOSE 8080