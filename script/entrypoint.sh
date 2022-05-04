#!/bin/sh

timestamp(){
    date +"%Y-%m-%d %H:%M:%S"
}

export TOKEN=${1}
export REPO=${2}
export REPO_OWNER=${3}

echo 'Initializing...'
GITHUB_TOKEN=$TOKEN GITHUB_REPOSITORY=$REPO OWNER=$REPO_OWNER ./jarves
existcode=0
set -o pipefail
./bin/
if [ $? != 0 ]; then
    existcode=1
fi

exit $existcode