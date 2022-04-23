import os
from pathlib import Path

## Const
DEFAULT_UID = 10000
DEFAULT_GID = 10000

PG_UID = 999
PG_GID = 999

REDIS_UIS = 999
REDIS_GID = 999

## Global variable
templates_dir = Path("/usr/src/app/templates")

host_root_dir = Path("/hostfs")

base_dir = '/harbor_make'
config_dir = Path('/config')
data_dir = Path('/data')

secret_dir = data_dir.joinpath('secret')
secret_key_dir = secret_dir.joinpath('keys');
trust_ca_dir = secret_dir.joinpath('keys', 'trust_ca')
internal_tls_dir = secret_dir.joinpath('tls')

INTERNAL_NO_PROXY_DN = {
    '127.0.0.1',
    'localhost',
    '.local',
    '.internal',
    'log',
    'db',
    'redis',
    'nginx',
    'core',
    'portal'
}
