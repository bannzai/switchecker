#! /bin/bash 
set -eu
set -o pipefail

CWD=$(dirname $0)
echo $CWD
cd $CWD
go run main.go
