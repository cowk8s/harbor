ARG harbor_base_image_version
ARG harbor_base_namespace
FROM ${harbor_base_namespace}/harbor-core-base:${harbor_base_image_version}

HEALTHCHECK CMD curl --fail -s http || exit 1
COPY ./make/photon/common/install_cert.sh /harbor/