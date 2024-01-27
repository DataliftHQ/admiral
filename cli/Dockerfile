FROM alpine:3.19

RUN apk add --no-cache bash tini

COPY datalift /usr/bin/datalift
COPY cli/scripts/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/sbin/tini", "--", "/entrypoint.sh"]
CMD [ "-h" ]
