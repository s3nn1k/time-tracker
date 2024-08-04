FROM golang:1.23-rc-alpine3.20 AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY internal ./internal
COPY cmd ./cmd
RUN go build -o ./bin/app ./cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY /.env /.env
COPY /schema.sql /schema.sql

CMD ["/app"]