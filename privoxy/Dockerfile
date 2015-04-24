FROM ubuntu:14.04

RUN apt-get update
RUN apt-get install -y privoxy

EXPOSE 8118

ADD docker/config/privoxy.conf /etc/privoxy/config

CMD ["privoxy", "--no-daemon", "/etc/privoxy/config"]
