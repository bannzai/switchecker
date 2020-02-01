#! /bin/bash 
set -u
set -o pipefail

switchecker -source=*.go -target=*.go --verbose
if [ $? -ne 0 ]; then
  echo "Test Failed"
  exit 1
fi
echo "Test Success"
