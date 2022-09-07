#!/bin/bash

set -e

oc create secret docker-registry quay \
    --docker-server=quay.io \
    --docker-username=$1 \
    --docker-password=$2


oc secrets link builder quay
oc secrets link default quay --for pull