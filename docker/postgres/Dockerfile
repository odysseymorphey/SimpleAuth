FROM postgres:15.3-alpine

COPY ./sql/init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_USER=postgres \
    POSTGRES_PASSWORD=mysecretpassword \
    POSTGRES_DB=postgres