FROM golang:1.17 AS build

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o lfscache

FROM alpine:latest

RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app
COPY --from=build /app/lfscache /app/lfscache
ENTRYPOINT ["/app/lfscache"]