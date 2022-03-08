FROM golang:1.17-alpine3.13 as builder
WORKDIR /go/src/svcpod
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM alpine
COPY --from=builder /go/src/svcpod/svcpod-client /usr/bin
ENV MyServiceName="svcpodserv"
ENTRYPOINT [ "svcpod-client" ]