# Links 123 - Saving your links

## How to run

### Via Docker Compose
The Easiest way to run whole infrastructure is to use _docker compose_:
~~~bash
docker-compose up
~~~

### Configuring

Example of `.env` file

~~~bash
# Logging

LINK_LOG_LEVEL=INFO

LINK_LOG_LOKI_ENABLED=false
LINK_LOG_LOKI_URL=unknown

# Persistance
MONGO_URL=:27017

~~~
 