#!/bin/bash
set -e
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source lib.sh || { echo "Are you at repo root?"; exit 1; }

tar -czf ./dist/pakkretqc.tgz ./dist/pakkretqc