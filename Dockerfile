FROM golang:1.18-alpine3.15 AS build-stage

WORKDIR /usr/src/app

COPY . .

RUN go build -o serv


FROM alpine:3.15

WORKDIR /usr/src/app

RUN adduser -D appuser

COPY --from=build-stage /usr/src/app/serv .

COPY data ./data
COPY server/graphql ./server/graphql
COPY server/templates ./server/templates

USER appuser

CMD ["./serv"]
