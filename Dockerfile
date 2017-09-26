FROM golang:1.7.4
MAINTAINER Armen Donigian

ENV SOURCES /go/src/kaggle_submission_app

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /go/bin/kaggle_submission_app
