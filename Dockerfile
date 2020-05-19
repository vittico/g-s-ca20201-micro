FROM alpine:latest

WORKDIR /opt/gsca20201-micro
COPY  g-s-ca20201-micro .
EXPOSE 10000/tcp
 CMD [ "./g-s-ca20201-micro" ]