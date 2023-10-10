# syntax=docker/dockerfile:1

From golang:1.19-alpine

WORKDIR /app

RUN apk add alpine-sdk

#COPY go.mod .
#COPY go.sum .
#RUN go mod download

COPY . .
RUN go get -t -v ./...

COPY *.go .

RUN go build -o /godocker

EXPOSE 8080

CMD [ "/godocker" ]

