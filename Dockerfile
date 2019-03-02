FROM golang:alpine as build

ENV GO111MODULE on

RUN apk add --no-cache git \
    && go get github.com/gozap/cclean \
    && go install -ldflags "-w -s" github.com/gozap/cclean

FROM alpine:latest as dist

COPY --from=build /go/bin/cclean /usr/bin/

ENTRYPOINT ["cclean"]
