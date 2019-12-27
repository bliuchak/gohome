FROM golang:1.13.5 as builder

COPY . /opt/apidocexample
WORKDIR /opt/apidocexample

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "\
    -X main.version=0.1.0 \
    -X main.commitFull=013fe2c2980c1237d8b6c521fc870cf1e0b46ae7 \
    -X main.buildTime=$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
    main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /opt/apidocexample .
CMD ["./main"]
