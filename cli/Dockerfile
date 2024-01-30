FROM alpine:3.19

RUN apk add --no-cache bash tini

COPY admiral /usr/bin/admiral
COPY cli/scripts/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/sbin/tini", "--", "/entrypoint.sh"]
CMD [ "-h" ]
