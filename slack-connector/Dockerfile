FROM golang:1.12.7-alpine3.10 as builder

WORKDIR /go/src/github.com/kyma-incubator/hack-showcase/slack-connector

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/slack-connector
RUN mkdir /app && mv ./main /app/main && mv ./internal/registration/configs/slackasyncapi.json /app/slackasyncapi.json

FROM alpine:3.10
WORKDIR /app

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /app/main /app 
COPY --from=builder /app/slackasyncapi.json /app

CMD ["./main"]