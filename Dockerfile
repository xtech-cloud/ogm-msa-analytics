# *************************************
#
# OpenGM
#
# *************************************

FROM alpine:3.14

MAINTAINER XTech Cloud "xtech.cloud"

ENV container docker
ENV MSA_MODE release

EXPOSE 18803

ADD bin/ogm-analytics /usr/local/bin/
RUN chmod +x /usr/local/bin/ogm-analytics

CMD ["/usr/local/bin/ogm-analytics"]
