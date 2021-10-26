
FROM pulumi/pulumi-go:latest

COPY . /action

ENTRYPOINT ["/action/entrypoint.sh"]

