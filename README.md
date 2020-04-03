# Links 123 - Saving your links

Get existing data from DB:
~~~bash
curl --request GET \
  --url 'http://localhost:8081/api/links?limit=200&offset=0'
~~~

## How to run

### Via Docker Compose
Easiest way to run whole infrastructure is to use _docker compose_:
~~~bash
docker-compose up
~~~

### Manual
If you need to run services separately:

**Gateway:**
~~~bash
go run ./entry/entry.go --kind=link-gtw --address=:8080
~~~

**Service:**
~~~bash
go run ./entry/entry.go --kind=link-srv --address=:10001
~~~

#### Note: if running manually, it's required to pass env. vars and provide Mongo DB
Example of _.env_ file:
~~~bash
MONGO_URL=:27017

GATEWAY_LINK_ADDRESS=:8080
SERVICE_LINK_ADDRESS=:10001

# Max. file size in bytes
PORT_MAX_IMPORT_FILE_SIZE=209715200
~~~

## Test
To run all available tests:
~~~bash
make run-tests
~~~

## Linting
You need to install linter at first (**ONLY ONCE**):
~~~bash
make install-linter
~~~

To run linter, use next command:
~~~bash
make run-linter
~~~
 