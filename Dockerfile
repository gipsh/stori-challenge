# Build stage
FROM golang:1.20.13-alpine3.19 AS builder
WORKDIR /app
COPY . .

RUN apk update
RUN apk add libc-dev
RUN apk add make
RUN apk add gcc
RUN apk add bash

RUN make build

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/stori-challenge .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/domain/template/summary.tmpl ./internal/domain/template/summary.tmpl
COPY --from=builder /app/internal/db/migrations/20240222170759-create-transactions.sql ./internal/db/migrations/20240222170759-create-transactions.sql

CMD [ "/app/stori-challenge" ]
ENTRYPOINT [ "/app/stori-challenge" ]