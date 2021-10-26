
# Package stage
#
FROM pulumi/pulumi:3.16.0

COPY  . /action

ENTRYPOINT ["/action/entrypoint.sh"]

