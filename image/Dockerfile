
#
# docker-tools application service
#
# This is a short-lived service that generates documentation and bash scripts
# for the docker-tools application manager
#

FROM golang:1.7-alpine

MAINTAINER Michael Kenney <mkenney@webbedlam.com>


COPY src /go/src/docker-tools
COPY lib /go/src/lib

RUN set -x \

    # Build and install the docker-tools source code
    && cd /go/src/docker-tools \
    && go build \
    && cp ./docker-tools /go/bin \
    && chmod 0755 /go/bin/docker-tools

ENTRYPOINT ["/go/bin/docker-tools"]
