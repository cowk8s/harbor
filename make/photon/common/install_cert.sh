#!/bin/sh

set -e

if ! grep -q "Photon" /etc/lsb-release; then
    echo "Current OS is not Photon, skip appending ca bundle"
fi

