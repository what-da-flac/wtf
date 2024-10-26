#!/usr/bin/env bash

set -e

# first parameter is ui version
if [ -z "$1" ]; then
    echo "no version provided"
fi

BUILD_DIR=ui/build/

cd && mkdir -p "${BUILD_DIR}" && cd "${BUILD_DIR}"
rm -rf *
aws --profile wtf s3 sync --delete s3://wtf-ui.what-da-flac.com/"$1" .