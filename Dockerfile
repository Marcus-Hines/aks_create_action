
FROM pulumi/pulumi:latest

COPY . /action

ENTRYPOINT ["/action/entrypoint.sh"]

