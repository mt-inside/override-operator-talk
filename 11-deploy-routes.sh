#!/usr/bin/env bash
source 00-common.sh

pushd ./overrides
just run | kubectl apply -f -
popd
