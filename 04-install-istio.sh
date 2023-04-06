#!/usr/bin/env bash
source 00-common.sh

# == Check ==

${ISTIOCTL} validate -f 04/values-custom.yaml 


# == Execute ==

${ISTIOCTL} install -y -f 04/values-custom.yaml

kubectl apply -f istio-${ISTIO_VERSION}/samples/addons
while ! kubectl wait --for=condition=available --timeout=600s deployment/kiali -n istio-system; do sleep 1; done

echo


# == Verify ==

# ${ISTIOCTL} verify-install -f <(${ISTIOCTL} manifest generate --set profile=demo)


# == Post ==

kubectl apply -f 04/ingressgateway-add-hop-headers.yaml

# Make default inject for ad hoc stuff; "real" examples should make their own namespace and set injection on that
kubectl label namespace default --overwrite=true istio-injection=enabled

watch -n1 kubectl get pods --namespace istio-system
