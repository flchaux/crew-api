FROM golang:1.18-alpine
WORKDIR /app

RUN apk add openrc --no-cache

#RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
#RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories
#RUN apk update

#RUN apk add mongodb
#RUN apk add mongodb-tools

#RUN rc-update add mongodb default
#RUN rc-service mongodb start

COPY src/api ./api
COPY src/model ./model
COPY src/dal ./dal

WORKDIR /app/dal

RUN go install

WORKDIR /app/api

RUN go install

EXPOSE 8080


CMD go run main.go