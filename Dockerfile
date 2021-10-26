# Build stage
#
FROM golang:latest AS build

COPY . /action

# Package stage
#
FROM pulumi/pulumi:latest

COPY --from=build /action /action

ENTRYPOINT ["/action/entrypoint.sh"]

