
# Package stage
#
FROM pulumi/pulumi-go:3.13.2

COPY  . /action

ENTRYPOINT ["/action/entrypoint.sh"]

