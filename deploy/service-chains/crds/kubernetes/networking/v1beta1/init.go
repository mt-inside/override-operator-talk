// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:networking.istio.io/v1beta1:DestinationRule":
		r = &DestinationRule{}
	case "kubernetes:networking.istio.io/v1beta1:Gateway":
		r = &Gateway{}
	case "kubernetes:networking.istio.io/v1beta1:ProxyConfig":
		r = &ProxyConfig{}
	case "kubernetes:networking.istio.io/v1beta1:ServiceEntry":
		r = &ServiceEntry{}
	case "kubernetes:networking.istio.io/v1beta1:Sidecar":
		r = &Sidecar{}
	case "kubernetes:networking.istio.io/v1beta1:VirtualService":
		r = &VirtualService{}
	case "kubernetes:networking.istio.io/v1beta1:WorkloadEntry":
		r = &WorkloadEntry{}
	case "kubernetes:networking.istio.io/v1beta1:WorkloadGroup":
		r = &WorkloadGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := kubernetes.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"crds",
		"networking.istio.io/v1beta1",
		&module{version},
	)
}
