#!/bin/bash

# Insert a couple of rows
docker run -it --rm -e PGPASSWORD=mysecretpassword --network db postgres \
  psql -h postgres -U postgres -c "INSERT INTO users (email) VALUES ('pete@theemersons.org') /* UUID:123e4567-e89b-12d3-a456-426614174000 */;"


docker run -it --rm -e PGPASSWORD=mysecretpassword --network db postgres \
  psql -h postgres -U postgres -c "INSERT INTO users (email) VALUES ('pete@fulcrumops.com') /* UUID:23e45678-f89c-22d4-b457-266141740002 */;"


# Select data from the demo table
docker run -it --rm -e PGPASSWORD=mysecretpassword --network db postgres \
  psql -h postgres -U postgres -c "SELECT email FROM users /* UUID:3e456712-89be-2d31-456b-266141740001 */;"

# Delete the data
docker run -it --rm -e PGPASSWORD=mysecretpassword --network db postgres \
  psql -h postgres -U postgres -c "DELETE FROM users /* UUID:e456712b-9bec-d31a-256c-766141740002 */;"
