#!/bin/bash
set -ex

# make sure we run in our own directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd "${DIR}" || exit 1

# build helper
SOURCE_CHECKER_PATH="/tmp/_tmp_source_checker_build"
export SOURCE_CHECKER_PATH=${SOURCE_CHECKER_PATH}
go build -o "${SOURCE_CHECKER_PATH}"

io --input "data.yml" --template "template.adoc" --output "../README.adoc" --allow-exec && echo "Done."
rm -f "${SOURCE_CHECKER_PATH}"