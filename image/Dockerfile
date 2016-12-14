
#
# docker-tools application service
#
# This is a short-lived service that generates documentation and bash scripts
# for the docker-tools application manager
#

FROM golang:1.7-alpine

MAINTAINER Michael Kenney <mkenney@webbedlam.com>

ENV DOCKER_TOOLS_CONF_DIR /usr/local/docker-tools
ENV TERM xterm-256color
ENV PATH /go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# Setup the image in a separate slice so compiling is fast
RUN set -x \
    && apk update \

    # for my sanity
    && apk add bash \

    # Install golog
    && apk add git \
    && mkdir -p /go/src/github.com/golang \
    && cd /go/src/github.com/golang \
    && git clone https://github.com/golang/glog.git \
    && apk del git \

    # install a proper less package
    && apk add less \

    && rm -rf /var/cache/apk/*

# add sources
COPY src /go/src/docker-tools
COPY lib /go/src/lib

# Build and install the docker-tools source code
RUN set -x \
    && cd /go/src/docker-tools \
    && go build \
    && cp ./docker-tools /go/bin \
    && chmod 0755 /go/bin/docker-tools

##############################################################################
# ~ fin ~
##############################################################################


ENTRYPOINT ["/go/bin/docker-tools"]
