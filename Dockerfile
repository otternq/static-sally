FROM alpine:3.8
ADD bin/static-sally /opt/static-sally
ENTRYPOINT /opt/static-sally
