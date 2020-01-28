#! /bin/bash 
set -u
set -o pipefail

switchecker --verbose
if [ $? -ne 0 ]; then
    echo "Test Failed"
    exit 1
fi
echo "Test Success"
