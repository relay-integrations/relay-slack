FROM golang:1.16-alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /build
COPY . .
RUN go build -a -o "/usr/local/bin/relay-slack-step-channel-create" "./cmd/channel-create"

FROM alpine:3
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder "/usr/local/bin/relay-slack-step-channel-create" "/usr/local/bin/relay-slack-step-channel-create"
ENTRYPOINT []
CMD ["/usr/local/bin/relay-slack-step-channel-create"]

LABEL "org.opencontainers.image.title"="Create a Slack channel"
LABEL "org.opencontainers.image.description"="This step creates a Slack channel."
LABEL "com.puppet.relay.sdk.version"="v1"
