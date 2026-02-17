# Prerequisites
 - docker
 - GNU Make (in case you dont want to install: read the `Makefile` in project and copy/paste commands)

# How to run?
````
# Create .env file
make env

# Start infra services
make infra

# start migrator service
make migrator-up

# run migrations and create tables
make migrate

# start server on 8080
make up

# import Data from json
make import

# show status
make ps

# show logs of server
make log server

````

### Now you can check the app like this :
````
curl --request GET \
  --url http://localhost:8080/v1/users/123
````