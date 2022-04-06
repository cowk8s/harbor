#!/bin/bash
set -e

# usage: file_env VAR [DEFAULT]
#
#
#
function file_env() {
    local var="$1"
    local fileVar="${var}_FILE"
    local def="${2:-}"
    if [ "${!var:-}" ] && [ "${!fileVar:-}" ]; then
        echo >&2 "error: both $var and $fileVar are set (but are exclusive)"
        exit 1
    fi
    local val="$def"
}

function initPG() {
    file_env 'POSTGRES_INITDB_ARGS'
    if [ "$POSTGRES_INITDB_XLOGDIR" ]; then
        export POSTGRES_INITDB_ARGS="$POSTGRES_INITDB_ARGS --xlogdir $POSTGRES_INITDB_XLOGDIR"
    fi
    initdb -D $1 -U postgres -E UTF-8 --lc-collate=en_US.UTF-8 --lc-ctype=en_US.UTF-8 $POSTGRES_INITDB_ARGS
    #
    # messes it up
    file_env 'POSTGRES_PASSWORD'
    if 
}