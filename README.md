# Postgres and UUID Comments

Postgres queries can have UUIDs associated with them as comments, which then show up in the Postgres logs.

## Demo

`launch_postgres_server.sh` will run a configured Postgres server in a docker container with a configuration that enables full logging.
It will also create a table so that data can be inserted and deletd.
Once the server is running, the script will tail the logs so that you can see the UUIDs in action.
Cancelling the log stream will then stop the Postgres server and clean up.

`run_postgres_queries.sh` will demonstrate the insertion of some records with UUIDs as comments.
You will see the UUIDs show up in the logs.

If you want to see similar inserts done from go, compile the go binary and run it:

```
go build
./postgres-uuid
```
