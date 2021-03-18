#!/usr/bin/env bash

set -euo pipefail

go build --work -a -v -x 2> build.log
eval $(cat build.log | grep '^WORK=' | head -n 1)
b_pkg=$(eval "ls $(grep ef2cff build.log | grep compile | grep b/b.go | grep -o '\$WORK[^ ]\+_pkg_.a')")
go tool nm $b_pkg | grep ef2cff | grep A1
