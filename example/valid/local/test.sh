#! /bin/bash 
set -u
set -o pipefail

switchecker -source=main.go -target=main.go --verbose
if [ $? -ne 0 ]; then
  echo "Test Failed"
  exit 1
fi
echo "Test Success"
