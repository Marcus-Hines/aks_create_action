#
# Build stage
#
FROM golang:latest AS build

COPY . /action

#
# Package stage
#
FROM pulumi/pulumi-go:latest

COPY --from=build . /action

ENTRYPOINT ["/action/entrypoint.sh"]

