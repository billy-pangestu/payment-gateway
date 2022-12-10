FROM golang:1.17.8-alpine3.15 AS build 

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /src

COPY . ./

RUN go mod download

WORKDIR /src/amqp_listener_broadcast/

RUN CGO_ENABLED=0 go build -a -o /src/app

FROM alpine

RUN apk update && apk add --no-cache tzdata

WORKDIR /app

COPY --from=build /src/app /app/
COPY --from=build /src/key /key/
COPY --from=build /src/log /log/
COPY --from=build /src/.env /.env

ENTRYPOINT ["./app"]
