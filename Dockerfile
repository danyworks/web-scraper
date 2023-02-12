FROM golang:1.19

ENV NAME $NAME

RUN apt-get update -yq && apt-get upgrade -yq

USER root

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./bin/notybot

ENTRYPOINT [ "/app/bin/notybot" ]