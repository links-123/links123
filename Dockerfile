FROM golang:1.14.1-alpine3.11 AS BUILDER

WORKDIR /app
COPY go.mod go.sum ./
RUN apk add git make && \
    go mod download
COPY . .
RUN make build

FROM alpine:3.11.5 AS RUNNER
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=BUILDER /app/links123 .
ENTRYPOINT ["./links123"]
CMD ["--version"]
