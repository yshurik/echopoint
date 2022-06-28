FROM golang:1.15.11-alpine as dev

RUN apk add git tree

COPY main.go go.mod go.sum /echopoint/
COPY app /echopoint/app/
RUN tree /echopoint

WORKDIR /echopoint
RUN go get || echo end1
RUN go build

RUN ls -al

FROM alpine:3.12 as base

COPY --from=dev /echopoint/echopoint /usr/bin

RUN echopoint --help

ENTRYPOINT ["echopoint"]

