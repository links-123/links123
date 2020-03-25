FROM golang:1.11-alpine AS BUILDER

WORKDIR /go/src/github.com/ic2hrmk/links123
COPY . .
RUN apk add make git && \
    make install-dependency-manager && \
    make run-dependency-manager && \
    make build

FROM alpine:3.11.5 AS RUNNER
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=BUILDER /go/src/github.com/ic2hrmk/links123/links123 .
CMD ["./links123"]
