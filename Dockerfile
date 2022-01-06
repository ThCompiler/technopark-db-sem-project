FROM golang:latest AS build

ADD . /app
WORKDIR /app/build
RUN ls -l ../
RUN go build -o server.out -v ../cmd/server/main.go

FROM ubuntu:20.04

RUN apt-get -y update && apt-get install -y tzdata
ENV TZ=Russia/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get -y update && apt-get install -y postgresql-12
USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER thecompiler WITH SUPERUSER PASSWORD 'qwerty';" &&\
    createdb -O thecompiler forum_db && /etc/init.d/postgresql stop

EXPOSE 5432
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]
USER root

WORKDIR /usr/src/build

COPY . .
COPY --from=build /app/build/ .

EXPOSE 5000
RUN mkdir -p ./logs/
ENV PGPASSWORD qwerty
CMD service postgresql start && psql -h localhost -d forum_db -U thecompiler -p 5432 -a -q -f ./db/db.sql && ./server.out