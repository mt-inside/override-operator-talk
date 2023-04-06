declare -r NAME=demo

declare -r KUBERNETES_VERSION=v1.25.3

declare -rx ISTIO_VERSION=1.15.3
declare -r ISTIOCTL=./istio-${ISTIO_VERSION}/bin/istioctl

declare -r CERT_MANAGER_VERSION=1.9.1

# TODO: not acutally used atm, cause needed in YAMLs
declare -r HTTP_LOG_VERSION=0.7.9
