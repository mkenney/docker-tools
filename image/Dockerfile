
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

    # the base less package needs to be upgraded to support color output
    && apk add less \

    # install 3rd-party packages
    && apk add git \
    && cd /go/src \

    # glog
    && git clone https://github.com/golang/glog.git \

    # cleanup
    && apk del git \
    && rm -rf /var/cache/apk/*

# Build and install the docker-tools source code
COPY ./ /go/src/docker-tools
RUN set -x \
    # Add the tool sources to gopath
    && export GOPATH=$GOPATH:/go/src/docker-tools \
    # Cleanup and build
    && cd /go/src/docker-tools \
    && rm Dockerfile \
    && go build \
    && ln -s /go/src/docker-tools/docker-tools /go/bin/docker-tools

##############################################################################
# ~ fin ~
##############################################################################

ENTRYPOINT ["/go/bin/docker-tools"]
