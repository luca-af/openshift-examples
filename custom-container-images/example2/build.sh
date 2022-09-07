#!/bin/bash

set -xe

# Build current image with name "example1" and tag "latest"
podman build -t example:paperino ./