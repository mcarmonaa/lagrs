# ENVIRONMENT VARIABLES TO CONFIG GSERVER
export GSERVER_ADDR="localhost:8080"
export GSERVER_HOST="localhost"
export GSERVER_PORT="8080"

# ENVIRONMENT VARIABLES TO CONNECT TO SERVICEUSERS
export SERVICEUSERS_ADDR=""
export SERVICEUSERS_HOST="localhost"
export SERVICEUSERS_PORT="8081"

# [manuel@manubook gserver]$ docker run --rm -p 8080:8080 --name gserver_1 --env-file ./env.sh gserver
