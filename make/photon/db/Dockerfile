ARG harbor_base_image_version
ARG harbor_base_namespace
FROM ${harbor_base_namespace}/harbor-db-base:${harbor_base_image_version}

VOLUME /var/lib/postgresql/data

COPY ./make/photon/db/docker-entrypoint.sh /docker-entrypoint.sh

USER postgres
