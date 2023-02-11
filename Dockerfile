FROM golang:1.19

ENV NAME $NAME

RUN apt-get update -yq && apt-get upgrade -yq

USER root

#RUN gcloud config set account ${SA_EMAIL}
#RUN gcloud auth application-default login --no-browser

WORKDIR /data

COPY . .

RUN go mod download

RUN go build -o ./bin/notybot

EXPOSE 8080

ENTRYPOINT [ "/data/bin/notybot" ]