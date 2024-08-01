# Build tiny docker image.
FROM alpine:latest

WORKDIR /app

RUN mkdir /config

COPY ./config/client_secret.json /config

COPY go.mod /
COPY storeManagerBotApp /app

CMD ["/app/storeManagerBotApp"]