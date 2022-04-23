ARG harbor_base_image_version
ARG harbor_base_namespace
FROM node:16.10.0 as nodeportal

WORKDIR /build_dir

ARG npm_registry=https://registry.npmjs.org

RUN apt-get update \
    && apt-get install -y --no-install-recommends python-yaml

COPY src/portal/package.json /build_dir
COPY src/portal/package-lock.json /build_dir
COPY src/portal/scripts /build_dir
COPY ./api/v2.0/legacy_swagger.yaml /build_dir/swagger.yaml
COPY ./api/v2.0/swagger.yaml /build_dir/swagger2.yaml
COPY ./api/swagger.yaml /build_dir/swagger3.yaml

COPY src/portal /build_dir

ENV NPM_CONFIG_REGISTRY=${npm_registry}
RUN npm install --unsafe-perm
RUN npm run generate-build-timestamp
RUN node --max_old_space_size=2048 'node_modules/@angular/cli/bin/ng' buld --configuration production
RUN python -c 'import sys, yaml, json; y=yaml.load(sys.stdin.read()); print json.dumps(y)' < swagger.yaml > dist/swagger