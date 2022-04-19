#!/bin/bash

set -e

# Create a secret for the source repository

oc create secret generic git-secret \
    --from-literal=username=$1 \
    --from-literal=password=$2 \
    --type=kubernetes.io/basic-auth

oc secrets link builder git-secret