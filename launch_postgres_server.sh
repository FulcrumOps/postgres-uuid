#!/bin/bash


docker network rm db
docker network create db
docker run -d --rm -p 5432:5432 --network db --name postgres -v "$PWD/my-postgres.conf":/etc/postgresql/postgresql.conf -e POSTGRES_PASSWORD=mysecretpassword postgres -c 'config_file=/etc/postgresql/postgresql.conf'
sleep 3
docker run -it --rm -e PGPASSWORD=mysecretpassword --network db postgres psql -h postgres -U postgres -c 'CREATE TABLE users (user_id serial PRIMARY KEY, email VARCHAR (255) UNIQUE NOT NULL);'
docker logs -f postgres
docker stop postgres
sleep 3
docker rm postgres
