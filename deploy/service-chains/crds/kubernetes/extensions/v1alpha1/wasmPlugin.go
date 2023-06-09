// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WasmPlugin struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
	Spec   WasmPluginSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput        `pulumi:"status"`
}

// NewWasmPlugin registers a new resource with the given unique name, arguments, and options.
func NewWasmPlugin(ctx *pulumi.Context,
	name string, args *WasmPluginArgs, opts ...pulumi.ResourceOption) (*WasmPlugin, error) {
	if args == nil {
		args = &WasmPluginArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("extensions.istio.io/v1alpha1")
	args.Kind = pulumi.StringPtr("WasmPlugin")
	var resource WasmPlugin
	err := ctx.RegisterResource("kubernetes:extensions.istio.io/v1alpha1:WasmPlugin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWasmPlugin gets an existing WasmPlugin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWasmPlugin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WasmPluginState, opts ...pulumi.ResourceOption) (*WasmPlugin, error) {
	var resource WasmPlugin
	err := ctx.ReadResource("kubernetes:extensions.istio.io/v1alpha1:WasmPlugin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WasmPlugin resources.
type wasmPluginState struct {
}

type WasmPluginState struct {
}

func (WasmPluginState) ElementType() reflect.Type {
	return reflect.TypeOf((*wasmPluginState)(nil)).Elem()
}

type wasmPluginArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
	Spec   *WasmPluginSpec        `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a WasmPlugin resource.
type WasmPluginArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
	Spec   WasmPluginSpecPtrInput
	Status pulumi.MapInput
}

func (WasmPluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wasmPluginArgs)(nil)).Elem()
}

type WasmPluginInput interface {
	pulumi.Input

	ToWasmPluginOutput() WasmPluginOutput
	ToWasmPluginOutputWithContext(ctx context.Context) WasmPluginOutput
}

func (*WasmPlugin) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPlugin)(nil)).Elem()
}

func (i *WasmPlugin) ToWasmPluginOutput() WasmPluginOutput {
	return i.ToWasmPluginOutputWithContext(context.Background())
}

func (i *WasmPlugin) ToWasmPluginOutputWithContext(ctx context.Context) WasmPluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginOutput)
}

type WasmPluginOutput struct{ *pulumi.OutputState }

func (WasmPluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPlugin)(nil)).Elem()
}

func (o WasmPluginOutput) ToWasmPluginOutput() WasmPluginOutput {
	return o
}

func (o WasmPluginOutput) ToWasmPluginOutputWithContext(ctx context.Context) WasmPluginOutput {
	return o
}

func (o WasmPluginOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPlugin) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o WasmPluginOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPlugin) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WasmPluginOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *WasmPlugin) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
func (o WasmPluginOutput) Spec() WasmPluginSpecPtrOutput {
	return o.ApplyT(func(v *WasmPlugin) WasmPluginSpecPtrOutput { return v.Spec }).(WasmPluginSpecPtrOutput)
}

func (o WasmPluginOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *WasmPlugin) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginInput)(nil)).Elem(), &WasmPlugin{})
	pulumi.RegisterOutputType(WasmPluginOutput{})
}
