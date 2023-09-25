FROM postgres:12.8-alpine

LABEL author="ali khaledi pour"
LABEL description="postgres db for url shortener"
LABEL version="1.0"

COPY *.sql /docker-entrypoint-initdb.d/