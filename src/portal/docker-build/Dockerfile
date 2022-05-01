FROM node:16.10.0 as builder

WORKDIR /build_dir

COPY src/portal /build_dir
COPY api/v2.0/legacy_swagger.yaml /build_dir/swagger.yaml
COPY api/v2.0/swagger.yaml /build_dir/swagger2.yaml
COPY api/swagger.yaml /build_dir/swagger3.yaml

RUN apt-get update \
    && apt-get install -y --no-install-recommends python-yaml
RUN npm install --unsafe-perm
RUN npm run postinstall
RUN npm run generate-build-timestamp
RUN node --max_old_space_size=2048 'node_modules/@angular/cli/bin/ng' build --configuration production
RUN python -c 'import sys, yaml, json; y=yaml.load(sys.stdin.read()); print json.dumps(y)' < swagger.yaml > dist/swagger.json
RUN python -c 'import sys, yaml, json; y=yaml.load(sys.stdin.read()); print json.dumps(y)' < swagger2.yaml > dist/swagger2.json
RUN python -c 'import sys, yaml, json; y=yaml.load(sys.stdin.read()); print json.dumps(y)' < swagger3.yaml > dist/swagger3.json
