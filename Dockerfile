FROM scratch
ADD bin/static-sally /opt/static-sally
ENTRYPOINT ["/opt/static-sally"]
