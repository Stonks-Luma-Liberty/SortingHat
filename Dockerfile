FROM golang:1.20-buster AS build

WORKDIR /app
COPY . .
RUN go mod download

WORKDIR /app/src
RUN go build -o /sortinghat

FROM debian:buster-slim as runtime
RUN apt-get -y update && apt-get -y install libssl-dev ca-certificates
WORKDIR /
COPY --from=build /app/.env .
COPY --from=build /sortinghat /sortinghat
