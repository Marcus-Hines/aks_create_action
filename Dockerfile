FROM pulumi/pulumi-go:latest

COPY . /action
ENV PATH="/root/.pulumi/bin:${PATH}"
ENTRYPOINT ["/action/entrypoint.sh"]

