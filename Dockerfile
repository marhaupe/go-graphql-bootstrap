FROM golang:1.11 as buildjob
RUN mkdir -p /src/bin/ 
WORKDIR /go/src/github.com/marhaupe/go-graphql-bootstrap
COPY . .
RUN CGO_ENABLED=0 GOOS=linux make build output=/src/bin/app

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* 
COPY --from=buildjob /src/bin/app /bin/app
WORKDIR /bin
EXPOSE 3000
ENTRYPOINT /bin/app
