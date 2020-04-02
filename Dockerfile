FROM golang:1.11-alpine AS BUILDER

WORKDIR /go/src/github.com/links-123/links123
COPY . .
RUN apk add make git
RUN make install-dependency-manager && \
    make run-dependency-manager && \
    make build

FROM alpine:3.11.5 AS RUNNER
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=BUILDER /go/src/github.com/links-123/links123/links123 .
ENTRYPOINT ["./links123"]
CMD ["--version"]
