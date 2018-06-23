FROM gocd/gocd-server:v18.2.0

ENV HOME=/go-working-dir/

# We want to have `htpasswd` available
RUN apk add --update apache2-utils && \
# Openshift
    chgrp -R 0 /go-server/ & chmod -R 775 /go-server/

# entrypoint stuff
COPY --chown=755 entrypoint.sh /
COPY --chown=755 docker-entrypoint.d/* /docker-entrypoint.d/
COPY --chown=755 install-gocd-plugins.sh /usr/local/sbin/install-gocd-plugins

RUN mkdir -p /go-working-dir && \
    chgrp -R 0 /go-working-dir && chmod 775 /go-working-dir

WORKDIR /go-working-dir/