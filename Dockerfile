FROM ubuntu:latest

LABEL MAINTAINER Dwayne Sykes "https://github.com/dsykes16"

COPY bin/pgfib /usr/bin/pgfib
RUN chmod +x /usr/bin/pgfib

RUN mkdir -p /etc/pgfib/sql
COPY sql/fibonacci.sql /etc/pgfib/sql/fibonacci.sql

RUN mkdir -p /etc/pgfib/config
COPY app.env /etc/pgfib/config/

ENV CGO_ENABLED=0

ENV DB_HOST=db \
    DB_USER=postgres \
    DB_PASS=IncorrectPassword \
    DB_PORT=5423 \
    DB_NAME=pgfib \
    DB_SSL=false \
    SERVER_ADDRESS=0.0.0.0:5000 \
    SQL_BOOTSTRAP_FILE=/etc/pgfib/sql/fibonacci.sql


CMD ["pgfib"]
