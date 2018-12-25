FROM centos:centos7

ADD ./server/server /usr/local/bin/

RUN mkdir -p /var/webroot

ADD ./index.html /var/webroot/

RUN chmod +x /usr/local/bin/server

ENTRYPOINT ["/usr/local/bin/server"]
