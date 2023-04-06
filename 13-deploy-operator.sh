#!/usr/bin/env bash
source 00-common.sh

kubectl apply -f overrides/deploy/override-operator.yaml

watch -n1 kubectl get pods --namespace default
