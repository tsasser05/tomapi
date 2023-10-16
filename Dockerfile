# syntax=docker/dockerfile:1

From golang:1.19-alpine

ENV GO111MODULE=on
#ENV GOFLAGS=-mod=vendor
ENV APP_HOME /Users/tom/tomapi
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

RUN apk add alpine-sdk

#COPY go.mod .
#COPY go.sum .
#RUN go mod download
COPY . .
RUN go get -t -v ./...
COPY *.go .
RUN go build -o /tomapi

EXPOSE 8084

CMD [ "/tomapi" ]

