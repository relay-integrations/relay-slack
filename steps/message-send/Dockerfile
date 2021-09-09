FROM golang:1.16-alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /build
COPY . .
RUN go build -a -o "/usr/local/bin/relay-slack-step-message-send" "./cmd/message-send"

FROM alpine:3
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder "/usr/local/bin/relay-slack-step-message-send" "/usr/local/bin/relay-slack-step-message-send"
ENTRYPOINT []
CMD ["/usr/local/bin/relay-slack-step-message-send"]

LABEL "org.opencontainers.image.title"="Send a message to a Slack channel"
LABEL "org.opencontainers.image.description"="This step sends a message to a Slack channel."
LABEL "sh.relay.sdk.version"="v1"
