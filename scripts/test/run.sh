#! /bin/bash 
set -eu
set -o pipefail

CWD=$(dirname $0)
cd $CWD
go run main.go
