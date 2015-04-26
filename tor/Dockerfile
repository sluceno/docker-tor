FROM ubuntu:14.04

RUN apt-get update
RUN apt-get install -y tor

ADD ./docker/config/torsocks.conf /etc/
ADD ./docker/config/torrc etc/tor/

EXPOSE 9050

CMD ["tor"]