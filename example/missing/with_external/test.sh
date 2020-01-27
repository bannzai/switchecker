#! /bin/bash 
set -eu
set -o pipefail

switchecker -source=thirdparty/*.go --verbose
