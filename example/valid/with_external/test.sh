#! /bin/bash 
set -u
set -o pipefail

switchecker -source=thirdparty/*.go --verbose
if [ $? -eq 0 ]; then
  echo "Test Failed"
  exit 1
fi
echo "Test Success"

