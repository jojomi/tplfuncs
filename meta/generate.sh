#!/bin/sh
set -ex

# generate generic type functions automatically using this script

# make sure we run in our own directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
cd "${DIR}" || exit 1

io --template "templates/math.tpl" --input "data.yml" --overwrite "msg=// AUTOGENERATED FILE. DO NOT EDIT." --output "../math.go"
io --template "templates/default.tpl" --input "data.yml" --overwrite "msg=// AUTOGENERATED FILE. DO NOT EDIT." --output "../default.go"

gofmt -s -w "../math.go" "../default.go"
goimports -w "../math.go" "../default.go"