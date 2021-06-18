FROM golang:1.16.5-alpine AS build

WORKDIR /go/src/datetime-service

COPY . .

RUN go install ./...

FROM alpine:3.12
WORKDIR /usr/bin
COPY --from=build /go/bin .
CMD cmd

#docker build . -t time
#docker run --rm -p 8080:8080 -d --name tm time cmd
#docker kill tm