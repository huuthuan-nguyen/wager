FROM golang:1.17
WORKDIR /go/src/
COPY . /go/src/
RUN bash "/go/src/build/build.sh"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/app .
CMD ["./app"]