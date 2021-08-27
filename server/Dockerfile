## BUILD ##

FROM golang:1.16-alpine AS builder
WORKDIR /nala
COPY . .
# Compile go dependencies
RUN go build -o main main.go
# Install migrate dependencies
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

## RUN ##

FROM scratch
WORKDIR /nala
COPY --from=builder /nala/main .
COPY --from=builder /nala/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY db/migration ./migration

# Export necessary port.
EXPOSE 8080

# Command to run when starting the container.
CMD [ "/nala/main" ]
ENTRYPOINT [ "/nala/start.sh" ]