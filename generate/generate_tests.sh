#!/bin/bash
set -ex

# make sure we run in our own directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd "${DIR}" || exit 1

# build helper
pushd "../documentation" > /dev/null
SOURCE_CHECKER_PATH="/tmp/_tmp_source_checker_build"
export SOURCE_CHECKER_PATH=${SOURCE_CHECKER_PATH}
go build -o "${SOURCE_CHECKER_PATH}"
popd > /dev/null

io --input "{}" --template "tplfuncs_test.tpl" --output "../tplfuncs_test.go" --allow-exec && echo "Done."
gofumpt -w -l ../tplfuncs_test.go
rm -f "${SOURCE_CHECKER_PATH}"