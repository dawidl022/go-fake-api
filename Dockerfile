FROM alpine:latest

COPY bin/server /serv
COPY data /data
COPY server/graphql /server/graphql
COPY server/templates /server/templates

CMD ["/serv"]
