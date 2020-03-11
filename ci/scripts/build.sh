#!/bin/bash -eux

pushd dp-frontend-homepage-controller
  make build
  cp build/dp-frontend-homepage-controller Dockerfile.concourse ../build
popd
