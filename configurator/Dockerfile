FROM leprechaun/gocd-configurator

RUN apk add --update bash git jq gettext

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["sh", "/entrypoint.sh"]
