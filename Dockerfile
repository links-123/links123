# Build:            docker build -t ship_links .
# Run gateway:      docker run --rm -p=8081:8081 ship_links
# Run service:      docker run --rm -p=10001:10001 ship_links

FROM golang:1.10

WORKDIR /go/src/github.com/ic2hrmk/ship_links
COPY . .

RUN go build -o link-service entry/entry.go && mv link-service /go/bin/

CMD ["link-service","--env=.env","--kind=link-gtw","--address=:8081"]
# CMD ["link-service","--env=.env","--kind=link-srv","--address=:10001"]
