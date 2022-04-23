# Get or generate private key
import os, subprocess, shutil
from pathlib import Path
from subprocess import DEVNULL
import logging

from g import DEFAULT_GID, DEFAULT_UID, shared_cert_dir, storage_ca_bundle_filename, internal_tls_dir, internal_ca_filename
from .misc import (
    mark_file,
    generate_random_string,
    check_permission,
    stat_decorator,
    get_realpath)

SSL_CERT_PATH = os.path.join("/etc/cert", "server.crt")
SSL_CERT_KEY_PATH = os.path.join("/etc/cert", "server.key")

def _get_secret(folder, filename, length=16):
    key_file = os.path.join(folder, filename)
    if os.path.isfile(key_file):
        with open(key_file, 'r') as f:
            key = f.read()
            print("loaded secret from file: %s" % key_file)
        mark_file(key_file)
        return key
    if not os.path.isdir(folder):
        os.makedirs(folder)
    key = generate_random_string(length)
    with open(key_file, 'w') as f:
        f.write(key)
        print("Generated and saved secret to file: %s" % key_file)
    mark_file(key_file)
    return key

def get_secret_key(path):
    secret_key = _get_secret(path, "secretkey")
    if len(secret_key) != 16:
        raise Exception("secret key's length has to be 16 chars, current length: %d" % len(secret_key))
    return secret_key


def get_alias(path):
    alias = _get_secret(path, "defaultalias", length=8)
    return alias

