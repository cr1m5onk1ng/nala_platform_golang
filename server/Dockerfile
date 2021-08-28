## BUILD ##

FROM golang:1.16-alpine3.13 AS builder

WORKDIR /build/

COPY . .

# Compile go dependencies
RUN go build -o main main.go
# Install migrate dependencies
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

## RUN ##

FROM alpine:3.13 AS runtime

USER 1000

WORKDIR /nala/

COPY --from=builder /build/main .
COPY --from=builder /build/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY db/migration ./migration

EXPOSE 8080

CMD [ "/nala/main" ]
ENTRYPOINT [ "/nala/start.sh" ]