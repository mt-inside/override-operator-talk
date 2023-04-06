// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProxyConfig struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Provides configuration for individual workloads. See more details at: https://istio.io/docs/reference/config/networking/proxy-config.html
	Spec   ProxyConfigSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput         `pulumi:"status"`
}

// NewProxyConfig registers a new resource with the given unique name, arguments, and options.
func NewProxyConfig(ctx *pulumi.Context,
	name string, args *ProxyConfigArgs, opts ...pulumi.ResourceOption) (*ProxyConfig, error) {
	if args == nil {
		args = &ProxyConfigArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("ProxyConfig")
	var resource ProxyConfig
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1beta1:ProxyConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxyConfig gets an existing ProxyConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxyConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyConfigState, opts ...pulumi.ResourceOption) (*ProxyConfig, error) {
	var resource ProxyConfig
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1beta1:ProxyConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProxyConfig resources.
type proxyConfigState struct {
}

type ProxyConfigState struct {
}

func (ProxyConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyConfigState)(nil)).Elem()
}

type proxyConfigArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Provides configuration for individual workloads. See more details at: https://istio.io/docs/reference/config/networking/proxy-config.html
	Spec   *ProxyConfigSpec       `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a ProxyConfig resource.
type ProxyConfigArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Provides configuration for individual workloads. See more details at: https://istio.io/docs/reference/config/networking/proxy-config.html
	Spec   ProxyConfigSpecPtrInput
	Status pulumi.MapInput
}

func (ProxyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyConfigArgs)(nil)).Elem()
}

type ProxyConfigInput interface {
	pulumi.Input

	ToProxyConfigOutput() ProxyConfigOutput
	ToProxyConfigOutputWithContext(ctx context.Context) ProxyConfigOutput
}

func (*ProxyConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyConfig)(nil)).Elem()
}

func (i *ProxyConfig) ToProxyConfigOutput() ProxyConfigOutput {
	return i.ToProxyConfigOutputWithContext(context.Background())
}

func (i *ProxyConfig) ToProxyConfigOutputWithContext(ctx context.Context) ProxyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyConfigOutput)
}

type ProxyConfigOutput struct{ *pulumi.OutputState }

func (ProxyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyConfig)(nil)).Elem()
}

func (o ProxyConfigOutput) ToProxyConfigOutput() ProxyConfigOutput {
	return o
}

func (o ProxyConfigOutput) ToProxyConfigOutputWithContext(ctx context.Context) ProxyConfigOutput {
	return o
}

func (o ProxyConfigOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConfig) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o ProxyConfigOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyConfig) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ProxyConfigOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *ProxyConfig) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Provides configuration for individual workloads. See more details at: https://istio.io/docs/reference/config/networking/proxy-config.html
func (o ProxyConfigOutput) Spec() ProxyConfigSpecPtrOutput {
	return o.ApplyT(func(v *ProxyConfig) ProxyConfigSpecPtrOutput { return v.Spec }).(ProxyConfigSpecPtrOutput)
}

func (o ProxyConfigOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *ProxyConfig) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyConfigInput)(nil)).Elem(), &ProxyConfig{})
	pulumi.RegisterOutputType(ProxyConfigOutput{})
}