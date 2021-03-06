#!/bin/sh

export TOKEN=${1}
export REPO=${2}
export REPO_OWNER=${3}

echo 'Initializing jarvis action...'
existcode=0
set -o pipefail
GITHUB_TOKEN=$TOKEN GITHUB_REPOSITORY=$REPO OWNER=$REPO_OWNER /jarvis
if [ $? != 0 ]; then
    existcode=1
fi

exit $existcode