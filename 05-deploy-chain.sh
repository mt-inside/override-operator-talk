#!/usr/bin/env bash
source 00-common.sh

pushd ./deploy/service-chains
pulumi refresh
pulumi up
popd

watch -n1 kubectl get pods --namespace default
