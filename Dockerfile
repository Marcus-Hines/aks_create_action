
FROM golang:1.15 AS builder
COPY . /action
ENV GOFLAGS=-mod=vendor
ENV CGO_ENABLED=0
RUN go build -o /action ./cmd


# Package stage
#
FROM pulumi/pulumi:latest

COPY --from=build /action /action

ENTRYPOINT ["/action/entrypoint.sh"]

