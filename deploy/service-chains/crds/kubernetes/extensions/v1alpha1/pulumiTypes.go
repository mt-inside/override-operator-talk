// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WasmPluginType struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
	Spec   *WasmPluginSpec        `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
type WasmPluginSpec struct {
	ImagePullPolicy *string `pulumi:"imagePullPolicy"`
	// Credentials to use for OCI image pulling.
	ImagePullSecret *string               `pulumi:"imagePullSecret"`
	Match           []WasmPluginSpecMatch `pulumi:"match"`
	// Determines where in the filter chain this `WasmPlugin` is to be injected.
	Phase *string `pulumi:"phase"`
	// The configuration that will be passed on to the plugin.
	PluginConfig map[string]interface{} `pulumi:"pluginConfig"`
	PluginName   *string                `pulumi:"pluginName"`
	// Determines ordering of `WasmPlugins` in the same `phase`.
	Priority *int                    `pulumi:"priority"`
	Selector *WasmPluginSpecSelector `pulumi:"selector"`
	// SHA256 checksum that will be used to verify Wasm module or OCI container.
	Sha256 *string `pulumi:"sha256"`
	// URL of a Wasm module or OCI container.
	Url             *string `pulumi:"url"`
	VerificationKey *string `pulumi:"verificationKey"`
	// Configuration for a Wasm VM.
	VmConfig *WasmPluginSpecVmconfig `pulumi:"vmConfig"`
}

// WasmPluginSpecInput is an input type that accepts WasmPluginSpecArgs and WasmPluginSpecOutput values.
// You can construct a concrete instance of `WasmPluginSpecInput` via:
//
//	WasmPluginSpecArgs{...}
type WasmPluginSpecInput interface {
	pulumi.Input

	ToWasmPluginSpecOutput() WasmPluginSpecOutput
	ToWasmPluginSpecOutputWithContext(context.Context) WasmPluginSpecOutput
}

// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
type WasmPluginSpecArgs struct {
	ImagePullPolicy pulumi.StringPtrInput `pulumi:"imagePullPolicy"`
	// Credentials to use for OCI image pulling.
	ImagePullSecret pulumi.StringPtrInput         `pulumi:"imagePullSecret"`
	Match           WasmPluginSpecMatchArrayInput `pulumi:"match"`
	// Determines where in the filter chain this `WasmPlugin` is to be injected.
	Phase pulumi.StringPtrInput `pulumi:"phase"`
	// The configuration that will be passed on to the plugin.
	PluginConfig pulumi.MapInput       `pulumi:"pluginConfig"`
	PluginName   pulumi.StringPtrInput `pulumi:"pluginName"`
	// Determines ordering of `WasmPlugins` in the same `phase`.
	Priority pulumi.IntPtrInput             `pulumi:"priority"`
	Selector WasmPluginSpecSelectorPtrInput `pulumi:"selector"`
	// SHA256 checksum that will be used to verify Wasm module or OCI container.
	Sha256 pulumi.StringPtrInput `pulumi:"sha256"`
	// URL of a Wasm module or OCI container.
	Url             pulumi.StringPtrInput `pulumi:"url"`
	VerificationKey pulumi.StringPtrInput `pulumi:"verificationKey"`
	// Configuration for a Wasm VM.
	VmConfig WasmPluginSpecVmconfigPtrInput `pulumi:"vmConfig"`
}

func (WasmPluginSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpec)(nil)).Elem()
}

func (i WasmPluginSpecArgs) ToWasmPluginSpecOutput() WasmPluginSpecOutput {
	return i.ToWasmPluginSpecOutputWithContext(context.Background())
}

func (i WasmPluginSpecArgs) ToWasmPluginSpecOutputWithContext(ctx context.Context) WasmPluginSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecOutput)
}

func (i WasmPluginSpecArgs) ToWasmPluginSpecPtrOutput() WasmPluginSpecPtrOutput {
	return i.ToWasmPluginSpecPtrOutputWithContext(context.Background())
}

func (i WasmPluginSpecArgs) ToWasmPluginSpecPtrOutputWithContext(ctx context.Context) WasmPluginSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecOutput).ToWasmPluginSpecPtrOutputWithContext(ctx)
}

// WasmPluginSpecPtrInput is an input type that accepts WasmPluginSpecArgs, WasmPluginSpecPtr and WasmPluginSpecPtrOutput values.
// You can construct a concrete instance of `WasmPluginSpecPtrInput` via:
//
//	        WasmPluginSpecArgs{...}
//
//	or:
//
//	        nil
type WasmPluginSpecPtrInput interface {
	pulumi.Input

	ToWasmPluginSpecPtrOutput() WasmPluginSpecPtrOutput
	ToWasmPluginSpecPtrOutputWithContext(context.Context) WasmPluginSpecPtrOutput
}

type wasmPluginSpecPtrType WasmPluginSpecArgs

func WasmPluginSpecPtr(v *WasmPluginSpecArgs) WasmPluginSpecPtrInput {
	return (*wasmPluginSpecPtrType)(v)
}

func (*wasmPluginSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPluginSpec)(nil)).Elem()
}

func (i *wasmPluginSpecPtrType) ToWasmPluginSpecPtrOutput() WasmPluginSpecPtrOutput {
	return i.ToWasmPluginSpecPtrOutputWithContext(context.Background())
}

func (i *wasmPluginSpecPtrType) ToWasmPluginSpecPtrOutputWithContext(ctx context.Context) WasmPluginSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecPtrOutput)
}

// Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
type WasmPluginSpecOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpec)(nil)).Elem()
}

func (o WasmPluginSpecOutput) ToWasmPluginSpecOutput() WasmPluginSpecOutput {
	return o
}

func (o WasmPluginSpecOutput) ToWasmPluginSpecOutputWithContext(ctx context.Context) WasmPluginSpecOutput {
	return o
}

func (o WasmPluginSpecOutput) ToWasmPluginSpecPtrOutput() WasmPluginSpecPtrOutput {
	return o.ToWasmPluginSpecPtrOutputWithContext(context.Background())
}

func (o WasmPluginSpecOutput) ToWasmPluginSpecPtrOutputWithContext(ctx context.Context) WasmPluginSpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WasmPluginSpec) *WasmPluginSpec {
		return &v
	}).(WasmPluginSpecPtrOutput)
}

func (o WasmPluginSpecOutput) ImagePullPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *string { return v.ImagePullPolicy }).(pulumi.StringPtrOutput)
}

// Credentials to use for OCI image pulling.
func (o WasmPluginSpecOutput) ImagePullSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *string { return v.ImagePullSecret }).(pulumi.StringPtrOutput)
}

func (o WasmPluginSpecOutput) Match() WasmPluginSpecMatchArrayOutput {
	return o.ApplyT(func(v WasmPluginSpec) []WasmPluginSpecMatch { return v.Match }).(WasmPluginSpecMatchArrayOutput)
}

// Determines where in the filter chain this `WasmPlugin` is to be injected.
func (o WasmPluginSpecOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *string { return v.Phase }).(pulumi.StringPtrOutput)
}

// The configuration that will be passed on to the plugin.
func (o WasmPluginSpecOutput) PluginConfig() pulumi.MapOutput {
	return o.ApplyT(func(v WasmPluginSpec) map[string]interface{} { return v.PluginConfig }).(pulumi.MapOutput)
}

func (o WasmPluginSpecOutput) PluginName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *string { return v.PluginName }).(pulumi.StringPtrOutput)
}

// Determines ordering of `WasmPlugins` in the same `phase`.
func (o WasmPluginSpecOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o WasmPluginSpecOutput) Selector() WasmPluginSpecSelectorPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *WasmPluginSpecSelector { return v.Selector }).(WasmPluginSpecSelectorPtrOutput)
}

// SHA256 checksum that will be used to verify Wasm module or OCI container.
func (o WasmPluginSpecOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *string { return v.Sha256 }).(pulumi.StringPtrOutput)
}

// URL of a Wasm module or OCI container.
func (o WasmPluginSpecOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o WasmPluginSpecOutput) VerificationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *string { return v.VerificationKey }).(pulumi.StringPtrOutput)
}

// Configuration for a Wasm VM.
func (o WasmPluginSpecOutput) VmConfig() WasmPluginSpecVmconfigPtrOutput {
	return o.ApplyT(func(v WasmPluginSpec) *WasmPluginSpecVmconfig { return v.VmConfig }).(WasmPluginSpecVmconfigPtrOutput)
}

type WasmPluginSpecPtrOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPluginSpec)(nil)).Elem()
}

func (o WasmPluginSpecPtrOutput) ToWasmPluginSpecPtrOutput() WasmPluginSpecPtrOutput {
	return o
}

func (o WasmPluginSpecPtrOutput) ToWasmPluginSpecPtrOutputWithContext(ctx context.Context) WasmPluginSpecPtrOutput {
	return o
}

func (o WasmPluginSpecPtrOutput) Elem() WasmPluginSpecOutput {
	return o.ApplyT(func(v *WasmPluginSpec) WasmPluginSpec {
		if v != nil {
			return *v
		}
		var ret WasmPluginSpec
		return ret
	}).(WasmPluginSpecOutput)
}

func (o WasmPluginSpecPtrOutput) ImagePullPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *string {
		if v == nil {
			return nil
		}
		return v.ImagePullPolicy
	}).(pulumi.StringPtrOutput)
}

// Credentials to use for OCI image pulling.
func (o WasmPluginSpecPtrOutput) ImagePullSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *string {
		if v == nil {
			return nil
		}
		return v.ImagePullSecret
	}).(pulumi.StringPtrOutput)
}

func (o WasmPluginSpecPtrOutput) Match() WasmPluginSpecMatchArrayOutput {
	return o.ApplyT(func(v *WasmPluginSpec) []WasmPluginSpecMatch {
		if v == nil {
			return nil
		}
		return v.Match
	}).(WasmPluginSpecMatchArrayOutput)
}

// Determines where in the filter chain this `WasmPlugin` is to be injected.
func (o WasmPluginSpecPtrOutput) Phase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *string {
		if v == nil {
			return nil
		}
		return v.Phase
	}).(pulumi.StringPtrOutput)
}

// The configuration that will be passed on to the plugin.
func (o WasmPluginSpecPtrOutput) PluginConfig() pulumi.MapOutput {
	return o.ApplyT(func(v *WasmPluginSpec) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.PluginConfig
	}).(pulumi.MapOutput)
}

func (o WasmPluginSpecPtrOutput) PluginName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *string {
		if v == nil {
			return nil
		}
		return v.PluginName
	}).(pulumi.StringPtrOutput)
}

// Determines ordering of `WasmPlugins` in the same `phase`.
func (o WasmPluginSpecPtrOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *int {
		if v == nil {
			return nil
		}
		return v.Priority
	}).(pulumi.IntPtrOutput)
}

func (o WasmPluginSpecPtrOutput) Selector() WasmPluginSpecSelectorPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *WasmPluginSpecSelector {
		if v == nil {
			return nil
		}
		return v.Selector
	}).(WasmPluginSpecSelectorPtrOutput)
}

// SHA256 checksum that will be used to verify Wasm module or OCI container.
func (o WasmPluginSpecPtrOutput) Sha256() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *string {
		if v == nil {
			return nil
		}
		return v.Sha256
	}).(pulumi.StringPtrOutput)
}

// URL of a Wasm module or OCI container.
func (o WasmPluginSpecPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

func (o WasmPluginSpecPtrOutput) VerificationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *string {
		if v == nil {
			return nil
		}
		return v.VerificationKey
	}).(pulumi.StringPtrOutput)
}

// Configuration for a Wasm VM.
func (o WasmPluginSpecPtrOutput) VmConfig() WasmPluginSpecVmconfigPtrOutput {
	return o.ApplyT(func(v *WasmPluginSpec) *WasmPluginSpecVmconfig {
		if v == nil {
			return nil
		}
		return v.VmConfig
	}).(WasmPluginSpecVmconfigPtrOutput)
}

type WasmPluginSpecMatch struct {
	Mode  *string                    `pulumi:"mode"`
	Ports []WasmPluginSpecMatchPorts `pulumi:"ports"`
}

// WasmPluginSpecMatchInput is an input type that accepts WasmPluginSpecMatchArgs and WasmPluginSpecMatchOutput values.
// You can construct a concrete instance of `WasmPluginSpecMatchInput` via:
//
//	WasmPluginSpecMatchArgs{...}
type WasmPluginSpecMatchInput interface {
	pulumi.Input

	ToWasmPluginSpecMatchOutput() WasmPluginSpecMatchOutput
	ToWasmPluginSpecMatchOutputWithContext(context.Context) WasmPluginSpecMatchOutput
}

type WasmPluginSpecMatchArgs struct {
	Mode  pulumi.StringPtrInput              `pulumi:"mode"`
	Ports WasmPluginSpecMatchPortsArrayInput `pulumi:"ports"`
}

func (WasmPluginSpecMatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecMatch)(nil)).Elem()
}

func (i WasmPluginSpecMatchArgs) ToWasmPluginSpecMatchOutput() WasmPluginSpecMatchOutput {
	return i.ToWasmPluginSpecMatchOutputWithContext(context.Background())
}

func (i WasmPluginSpecMatchArgs) ToWasmPluginSpecMatchOutputWithContext(ctx context.Context) WasmPluginSpecMatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecMatchOutput)
}

// WasmPluginSpecMatchArrayInput is an input type that accepts WasmPluginSpecMatchArray and WasmPluginSpecMatchArrayOutput values.
// You can construct a concrete instance of `WasmPluginSpecMatchArrayInput` via:
//
//	WasmPluginSpecMatchArray{ WasmPluginSpecMatchArgs{...} }
type WasmPluginSpecMatchArrayInput interface {
	pulumi.Input

	ToWasmPluginSpecMatchArrayOutput() WasmPluginSpecMatchArrayOutput
	ToWasmPluginSpecMatchArrayOutputWithContext(context.Context) WasmPluginSpecMatchArrayOutput
}

type WasmPluginSpecMatchArray []WasmPluginSpecMatchInput

func (WasmPluginSpecMatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WasmPluginSpecMatch)(nil)).Elem()
}

func (i WasmPluginSpecMatchArray) ToWasmPluginSpecMatchArrayOutput() WasmPluginSpecMatchArrayOutput {
	return i.ToWasmPluginSpecMatchArrayOutputWithContext(context.Background())
}

func (i WasmPluginSpecMatchArray) ToWasmPluginSpecMatchArrayOutputWithContext(ctx context.Context) WasmPluginSpecMatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecMatchArrayOutput)
}

type WasmPluginSpecMatchOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecMatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecMatch)(nil)).Elem()
}

func (o WasmPluginSpecMatchOutput) ToWasmPluginSpecMatchOutput() WasmPluginSpecMatchOutput {
	return o
}

func (o WasmPluginSpecMatchOutput) ToWasmPluginSpecMatchOutputWithContext(ctx context.Context) WasmPluginSpecMatchOutput {
	return o
}

func (o WasmPluginSpecMatchOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpecMatch) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o WasmPluginSpecMatchOutput) Ports() WasmPluginSpecMatchPortsArrayOutput {
	return o.ApplyT(func(v WasmPluginSpecMatch) []WasmPluginSpecMatchPorts { return v.Ports }).(WasmPluginSpecMatchPortsArrayOutput)
}

type WasmPluginSpecMatchArrayOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecMatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WasmPluginSpecMatch)(nil)).Elem()
}

func (o WasmPluginSpecMatchArrayOutput) ToWasmPluginSpecMatchArrayOutput() WasmPluginSpecMatchArrayOutput {
	return o
}

func (o WasmPluginSpecMatchArrayOutput) ToWasmPluginSpecMatchArrayOutputWithContext(ctx context.Context) WasmPluginSpecMatchArrayOutput {
	return o
}

func (o WasmPluginSpecMatchArrayOutput) Index(i pulumi.IntInput) WasmPluginSpecMatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WasmPluginSpecMatch {
		return vs[0].([]WasmPluginSpecMatch)[vs[1].(int)]
	}).(WasmPluginSpecMatchOutput)
}

type WasmPluginSpecMatchPorts struct {
	Number *int `pulumi:"number"`
}

// WasmPluginSpecMatchPortsInput is an input type that accepts WasmPluginSpecMatchPortsArgs and WasmPluginSpecMatchPortsOutput values.
// You can construct a concrete instance of `WasmPluginSpecMatchPortsInput` via:
//
//	WasmPluginSpecMatchPortsArgs{...}
type WasmPluginSpecMatchPortsInput interface {
	pulumi.Input

	ToWasmPluginSpecMatchPortsOutput() WasmPluginSpecMatchPortsOutput
	ToWasmPluginSpecMatchPortsOutputWithContext(context.Context) WasmPluginSpecMatchPortsOutput
}

type WasmPluginSpecMatchPortsArgs struct {
	Number pulumi.IntPtrInput `pulumi:"number"`
}

func (WasmPluginSpecMatchPortsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecMatchPorts)(nil)).Elem()
}

func (i WasmPluginSpecMatchPortsArgs) ToWasmPluginSpecMatchPortsOutput() WasmPluginSpecMatchPortsOutput {
	return i.ToWasmPluginSpecMatchPortsOutputWithContext(context.Background())
}

func (i WasmPluginSpecMatchPortsArgs) ToWasmPluginSpecMatchPortsOutputWithContext(ctx context.Context) WasmPluginSpecMatchPortsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecMatchPortsOutput)
}

// WasmPluginSpecMatchPortsArrayInput is an input type that accepts WasmPluginSpecMatchPortsArray and WasmPluginSpecMatchPortsArrayOutput values.
// You can construct a concrete instance of `WasmPluginSpecMatchPortsArrayInput` via:
//
//	WasmPluginSpecMatchPortsArray{ WasmPluginSpecMatchPortsArgs{...} }
type WasmPluginSpecMatchPortsArrayInput interface {
	pulumi.Input

	ToWasmPluginSpecMatchPortsArrayOutput() WasmPluginSpecMatchPortsArrayOutput
	ToWasmPluginSpecMatchPortsArrayOutputWithContext(context.Context) WasmPluginSpecMatchPortsArrayOutput
}

type WasmPluginSpecMatchPortsArray []WasmPluginSpecMatchPortsInput

func (WasmPluginSpecMatchPortsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WasmPluginSpecMatchPorts)(nil)).Elem()
}

func (i WasmPluginSpecMatchPortsArray) ToWasmPluginSpecMatchPortsArrayOutput() WasmPluginSpecMatchPortsArrayOutput {
	return i.ToWasmPluginSpecMatchPortsArrayOutputWithContext(context.Background())
}

func (i WasmPluginSpecMatchPortsArray) ToWasmPluginSpecMatchPortsArrayOutputWithContext(ctx context.Context) WasmPluginSpecMatchPortsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecMatchPortsArrayOutput)
}

type WasmPluginSpecMatchPortsOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecMatchPortsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecMatchPorts)(nil)).Elem()
}

func (o WasmPluginSpecMatchPortsOutput) ToWasmPluginSpecMatchPortsOutput() WasmPluginSpecMatchPortsOutput {
	return o
}

func (o WasmPluginSpecMatchPortsOutput) ToWasmPluginSpecMatchPortsOutputWithContext(ctx context.Context) WasmPluginSpecMatchPortsOutput {
	return o
}

func (o WasmPluginSpecMatchPortsOutput) Number() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WasmPluginSpecMatchPorts) *int { return v.Number }).(pulumi.IntPtrOutput)
}

type WasmPluginSpecMatchPortsArrayOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecMatchPortsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WasmPluginSpecMatchPorts)(nil)).Elem()
}

func (o WasmPluginSpecMatchPortsArrayOutput) ToWasmPluginSpecMatchPortsArrayOutput() WasmPluginSpecMatchPortsArrayOutput {
	return o
}

func (o WasmPluginSpecMatchPortsArrayOutput) ToWasmPluginSpecMatchPortsArrayOutputWithContext(ctx context.Context) WasmPluginSpecMatchPortsArrayOutput {
	return o
}

func (o WasmPluginSpecMatchPortsArrayOutput) Index(i pulumi.IntInput) WasmPluginSpecMatchPortsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WasmPluginSpecMatchPorts {
		return vs[0].([]WasmPluginSpecMatchPorts)[vs[1].(int)]
	}).(WasmPluginSpecMatchPortsOutput)
}

type WasmPluginSpecSelector struct {
	MatchLabels map[string]string `pulumi:"matchLabels"`
}

// WasmPluginSpecSelectorInput is an input type that accepts WasmPluginSpecSelectorArgs and WasmPluginSpecSelectorOutput values.
// You can construct a concrete instance of `WasmPluginSpecSelectorInput` via:
//
//	WasmPluginSpecSelectorArgs{...}
type WasmPluginSpecSelectorInput interface {
	pulumi.Input

	ToWasmPluginSpecSelectorOutput() WasmPluginSpecSelectorOutput
	ToWasmPluginSpecSelectorOutputWithContext(context.Context) WasmPluginSpecSelectorOutput
}

type WasmPluginSpecSelectorArgs struct {
	MatchLabels pulumi.StringMapInput `pulumi:"matchLabels"`
}

func (WasmPluginSpecSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecSelector)(nil)).Elem()
}

func (i WasmPluginSpecSelectorArgs) ToWasmPluginSpecSelectorOutput() WasmPluginSpecSelectorOutput {
	return i.ToWasmPluginSpecSelectorOutputWithContext(context.Background())
}

func (i WasmPluginSpecSelectorArgs) ToWasmPluginSpecSelectorOutputWithContext(ctx context.Context) WasmPluginSpecSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecSelectorOutput)
}

func (i WasmPluginSpecSelectorArgs) ToWasmPluginSpecSelectorPtrOutput() WasmPluginSpecSelectorPtrOutput {
	return i.ToWasmPluginSpecSelectorPtrOutputWithContext(context.Background())
}

func (i WasmPluginSpecSelectorArgs) ToWasmPluginSpecSelectorPtrOutputWithContext(ctx context.Context) WasmPluginSpecSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecSelectorOutput).ToWasmPluginSpecSelectorPtrOutputWithContext(ctx)
}

// WasmPluginSpecSelectorPtrInput is an input type that accepts WasmPluginSpecSelectorArgs, WasmPluginSpecSelectorPtr and WasmPluginSpecSelectorPtrOutput values.
// You can construct a concrete instance of `WasmPluginSpecSelectorPtrInput` via:
//
//	        WasmPluginSpecSelectorArgs{...}
//
//	or:
//
//	        nil
type WasmPluginSpecSelectorPtrInput interface {
	pulumi.Input

	ToWasmPluginSpecSelectorPtrOutput() WasmPluginSpecSelectorPtrOutput
	ToWasmPluginSpecSelectorPtrOutputWithContext(context.Context) WasmPluginSpecSelectorPtrOutput
}

type wasmPluginSpecSelectorPtrType WasmPluginSpecSelectorArgs

func WasmPluginSpecSelectorPtr(v *WasmPluginSpecSelectorArgs) WasmPluginSpecSelectorPtrInput {
	return (*wasmPluginSpecSelectorPtrType)(v)
}

func (*wasmPluginSpecSelectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPluginSpecSelector)(nil)).Elem()
}

func (i *wasmPluginSpecSelectorPtrType) ToWasmPluginSpecSelectorPtrOutput() WasmPluginSpecSelectorPtrOutput {
	return i.ToWasmPluginSpecSelectorPtrOutputWithContext(context.Background())
}

func (i *wasmPluginSpecSelectorPtrType) ToWasmPluginSpecSelectorPtrOutputWithContext(ctx context.Context) WasmPluginSpecSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecSelectorPtrOutput)
}

type WasmPluginSpecSelectorOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecSelector)(nil)).Elem()
}

func (o WasmPluginSpecSelectorOutput) ToWasmPluginSpecSelectorOutput() WasmPluginSpecSelectorOutput {
	return o
}

func (o WasmPluginSpecSelectorOutput) ToWasmPluginSpecSelectorOutputWithContext(ctx context.Context) WasmPluginSpecSelectorOutput {
	return o
}

func (o WasmPluginSpecSelectorOutput) ToWasmPluginSpecSelectorPtrOutput() WasmPluginSpecSelectorPtrOutput {
	return o.ToWasmPluginSpecSelectorPtrOutputWithContext(context.Background())
}

func (o WasmPluginSpecSelectorOutput) ToWasmPluginSpecSelectorPtrOutputWithContext(ctx context.Context) WasmPluginSpecSelectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WasmPluginSpecSelector) *WasmPluginSpecSelector {
		return &v
	}).(WasmPluginSpecSelectorPtrOutput)
}

func (o WasmPluginSpecSelectorOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v WasmPluginSpecSelector) map[string]string { return v.MatchLabels }).(pulumi.StringMapOutput)
}

type WasmPluginSpecSelectorPtrOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecSelectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPluginSpecSelector)(nil)).Elem()
}

func (o WasmPluginSpecSelectorPtrOutput) ToWasmPluginSpecSelectorPtrOutput() WasmPluginSpecSelectorPtrOutput {
	return o
}

func (o WasmPluginSpecSelectorPtrOutput) ToWasmPluginSpecSelectorPtrOutputWithContext(ctx context.Context) WasmPluginSpecSelectorPtrOutput {
	return o
}

func (o WasmPluginSpecSelectorPtrOutput) Elem() WasmPluginSpecSelectorOutput {
	return o.ApplyT(func(v *WasmPluginSpecSelector) WasmPluginSpecSelector {
		if v != nil {
			return *v
		}
		var ret WasmPluginSpecSelector
		return ret
	}).(WasmPluginSpecSelectorOutput)
}

func (o WasmPluginSpecSelectorPtrOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WasmPluginSpecSelector) map[string]string {
		if v == nil {
			return nil
		}
		return v.MatchLabels
	}).(pulumi.StringMapOutput)
}

type WasmPluginSpecSelectorMatchlabels struct {
}

// Configuration for a Wasm VM.
type WasmPluginSpecVmconfig struct {
	// Specifies environment variables to be injected to this VM.
	Env []WasmPluginSpecVmconfigEnv `pulumi:"env"`
}

// WasmPluginSpecVmconfigInput is an input type that accepts WasmPluginSpecVmconfigArgs and WasmPluginSpecVmconfigOutput values.
// You can construct a concrete instance of `WasmPluginSpecVmconfigInput` via:
//
//	WasmPluginSpecVmconfigArgs{...}
type WasmPluginSpecVmconfigInput interface {
	pulumi.Input

	ToWasmPluginSpecVmconfigOutput() WasmPluginSpecVmconfigOutput
	ToWasmPluginSpecVmconfigOutputWithContext(context.Context) WasmPluginSpecVmconfigOutput
}

// Configuration for a Wasm VM.
type WasmPluginSpecVmconfigArgs struct {
	// Specifies environment variables to be injected to this VM.
	Env WasmPluginSpecVmconfigEnvArrayInput `pulumi:"env"`
}

func (WasmPluginSpecVmconfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecVmconfig)(nil)).Elem()
}

func (i WasmPluginSpecVmconfigArgs) ToWasmPluginSpecVmconfigOutput() WasmPluginSpecVmconfigOutput {
	return i.ToWasmPluginSpecVmconfigOutputWithContext(context.Background())
}

func (i WasmPluginSpecVmconfigArgs) ToWasmPluginSpecVmconfigOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecVmconfigOutput)
}

func (i WasmPluginSpecVmconfigArgs) ToWasmPluginSpecVmconfigPtrOutput() WasmPluginSpecVmconfigPtrOutput {
	return i.ToWasmPluginSpecVmconfigPtrOutputWithContext(context.Background())
}

func (i WasmPluginSpecVmconfigArgs) ToWasmPluginSpecVmconfigPtrOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecVmconfigOutput).ToWasmPluginSpecVmconfigPtrOutputWithContext(ctx)
}

// WasmPluginSpecVmconfigPtrInput is an input type that accepts WasmPluginSpecVmconfigArgs, WasmPluginSpecVmconfigPtr and WasmPluginSpecVmconfigPtrOutput values.
// You can construct a concrete instance of `WasmPluginSpecVmconfigPtrInput` via:
//
//	        WasmPluginSpecVmconfigArgs{...}
//
//	or:
//
//	        nil
type WasmPluginSpecVmconfigPtrInput interface {
	pulumi.Input

	ToWasmPluginSpecVmconfigPtrOutput() WasmPluginSpecVmconfigPtrOutput
	ToWasmPluginSpecVmconfigPtrOutputWithContext(context.Context) WasmPluginSpecVmconfigPtrOutput
}

type wasmPluginSpecVmconfigPtrType WasmPluginSpecVmconfigArgs

func WasmPluginSpecVmconfigPtr(v *WasmPluginSpecVmconfigArgs) WasmPluginSpecVmconfigPtrInput {
	return (*wasmPluginSpecVmconfigPtrType)(v)
}

func (*wasmPluginSpecVmconfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPluginSpecVmconfig)(nil)).Elem()
}

func (i *wasmPluginSpecVmconfigPtrType) ToWasmPluginSpecVmconfigPtrOutput() WasmPluginSpecVmconfigPtrOutput {
	return i.ToWasmPluginSpecVmconfigPtrOutputWithContext(context.Background())
}

func (i *wasmPluginSpecVmconfigPtrType) ToWasmPluginSpecVmconfigPtrOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecVmconfigPtrOutput)
}

// Configuration for a Wasm VM.
type WasmPluginSpecVmconfigOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecVmconfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecVmconfig)(nil)).Elem()
}

func (o WasmPluginSpecVmconfigOutput) ToWasmPluginSpecVmconfigOutput() WasmPluginSpecVmconfigOutput {
	return o
}

func (o WasmPluginSpecVmconfigOutput) ToWasmPluginSpecVmconfigOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigOutput {
	return o
}

func (o WasmPluginSpecVmconfigOutput) ToWasmPluginSpecVmconfigPtrOutput() WasmPluginSpecVmconfigPtrOutput {
	return o.ToWasmPluginSpecVmconfigPtrOutputWithContext(context.Background())
}

func (o WasmPluginSpecVmconfigOutput) ToWasmPluginSpecVmconfigPtrOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WasmPluginSpecVmconfig) *WasmPluginSpecVmconfig {
		return &v
	}).(WasmPluginSpecVmconfigPtrOutput)
}

// Specifies environment variables to be injected to this VM.
func (o WasmPluginSpecVmconfigOutput) Env() WasmPluginSpecVmconfigEnvArrayOutput {
	return o.ApplyT(func(v WasmPluginSpecVmconfig) []WasmPluginSpecVmconfigEnv { return v.Env }).(WasmPluginSpecVmconfigEnvArrayOutput)
}

type WasmPluginSpecVmconfigPtrOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecVmconfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WasmPluginSpecVmconfig)(nil)).Elem()
}

func (o WasmPluginSpecVmconfigPtrOutput) ToWasmPluginSpecVmconfigPtrOutput() WasmPluginSpecVmconfigPtrOutput {
	return o
}

func (o WasmPluginSpecVmconfigPtrOutput) ToWasmPluginSpecVmconfigPtrOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigPtrOutput {
	return o
}

func (o WasmPluginSpecVmconfigPtrOutput) Elem() WasmPluginSpecVmconfigOutput {
	return o.ApplyT(func(v *WasmPluginSpecVmconfig) WasmPluginSpecVmconfig {
		if v != nil {
			return *v
		}
		var ret WasmPluginSpecVmconfig
		return ret
	}).(WasmPluginSpecVmconfigOutput)
}

// Specifies environment variables to be injected to this VM.
func (o WasmPluginSpecVmconfigPtrOutput) Env() WasmPluginSpecVmconfigEnvArrayOutput {
	return o.ApplyT(func(v *WasmPluginSpecVmconfig) []WasmPluginSpecVmconfigEnv {
		if v == nil {
			return nil
		}
		return v.Env
	}).(WasmPluginSpecVmconfigEnvArrayOutput)
}

type WasmPluginSpecVmconfigEnv struct {
	Name *string `pulumi:"name"`
	// Value for the environment variable.
	Value     *string `pulumi:"value"`
	ValueFrom *string `pulumi:"valueFrom"`
}

// WasmPluginSpecVmconfigEnvInput is an input type that accepts WasmPluginSpecVmconfigEnvArgs and WasmPluginSpecVmconfigEnvOutput values.
// You can construct a concrete instance of `WasmPluginSpecVmconfigEnvInput` via:
//
//	WasmPluginSpecVmconfigEnvArgs{...}
type WasmPluginSpecVmconfigEnvInput interface {
	pulumi.Input

	ToWasmPluginSpecVmconfigEnvOutput() WasmPluginSpecVmconfigEnvOutput
	ToWasmPluginSpecVmconfigEnvOutputWithContext(context.Context) WasmPluginSpecVmconfigEnvOutput
}

type WasmPluginSpecVmconfigEnvArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Value for the environment variable.
	Value     pulumi.StringPtrInput `pulumi:"value"`
	ValueFrom pulumi.StringPtrInput `pulumi:"valueFrom"`
}

func (WasmPluginSpecVmconfigEnvArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecVmconfigEnv)(nil)).Elem()
}

func (i WasmPluginSpecVmconfigEnvArgs) ToWasmPluginSpecVmconfigEnvOutput() WasmPluginSpecVmconfigEnvOutput {
	return i.ToWasmPluginSpecVmconfigEnvOutputWithContext(context.Background())
}

func (i WasmPluginSpecVmconfigEnvArgs) ToWasmPluginSpecVmconfigEnvOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigEnvOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecVmconfigEnvOutput)
}

// WasmPluginSpecVmconfigEnvArrayInput is an input type that accepts WasmPluginSpecVmconfigEnvArray and WasmPluginSpecVmconfigEnvArrayOutput values.
// You can construct a concrete instance of `WasmPluginSpecVmconfigEnvArrayInput` via:
//
//	WasmPluginSpecVmconfigEnvArray{ WasmPluginSpecVmconfigEnvArgs{...} }
type WasmPluginSpecVmconfigEnvArrayInput interface {
	pulumi.Input

	ToWasmPluginSpecVmconfigEnvArrayOutput() WasmPluginSpecVmconfigEnvArrayOutput
	ToWasmPluginSpecVmconfigEnvArrayOutputWithContext(context.Context) WasmPluginSpecVmconfigEnvArrayOutput
}

type WasmPluginSpecVmconfigEnvArray []WasmPluginSpecVmconfigEnvInput

func (WasmPluginSpecVmconfigEnvArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WasmPluginSpecVmconfigEnv)(nil)).Elem()
}

func (i WasmPluginSpecVmconfigEnvArray) ToWasmPluginSpecVmconfigEnvArrayOutput() WasmPluginSpecVmconfigEnvArrayOutput {
	return i.ToWasmPluginSpecVmconfigEnvArrayOutputWithContext(context.Background())
}

func (i WasmPluginSpecVmconfigEnvArray) ToWasmPluginSpecVmconfigEnvArrayOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigEnvArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WasmPluginSpecVmconfigEnvArrayOutput)
}

type WasmPluginSpecVmconfigEnvOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecVmconfigEnvOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WasmPluginSpecVmconfigEnv)(nil)).Elem()
}

func (o WasmPluginSpecVmconfigEnvOutput) ToWasmPluginSpecVmconfigEnvOutput() WasmPluginSpecVmconfigEnvOutput {
	return o
}

func (o WasmPluginSpecVmconfigEnvOutput) ToWasmPluginSpecVmconfigEnvOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigEnvOutput {
	return o
}

func (o WasmPluginSpecVmconfigEnvOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpecVmconfigEnv) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Value for the environment variable.
func (o WasmPluginSpecVmconfigEnvOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpecVmconfigEnv) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func (o WasmPluginSpecVmconfigEnvOutput) ValueFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WasmPluginSpecVmconfigEnv) *string { return v.ValueFrom }).(pulumi.StringPtrOutput)
}

type WasmPluginSpecVmconfigEnvArrayOutput struct{ *pulumi.OutputState }

func (WasmPluginSpecVmconfigEnvArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WasmPluginSpecVmconfigEnv)(nil)).Elem()
}

func (o WasmPluginSpecVmconfigEnvArrayOutput) ToWasmPluginSpecVmconfigEnvArrayOutput() WasmPluginSpecVmconfigEnvArrayOutput {
	return o
}

func (o WasmPluginSpecVmconfigEnvArrayOutput) ToWasmPluginSpecVmconfigEnvArrayOutputWithContext(ctx context.Context) WasmPluginSpecVmconfigEnvArrayOutput {
	return o
}

func (o WasmPluginSpecVmconfigEnvArrayOutput) Index(i pulumi.IntInput) WasmPluginSpecVmconfigEnvOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WasmPluginSpecVmconfigEnv {
		return vs[0].([]WasmPluginSpecVmconfigEnv)[vs[1].(int)]
	}).(WasmPluginSpecVmconfigEnvOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecInput)(nil)).Elem(), WasmPluginSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecPtrInput)(nil)).Elem(), WasmPluginSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecMatchInput)(nil)).Elem(), WasmPluginSpecMatchArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecMatchArrayInput)(nil)).Elem(), WasmPluginSpecMatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecMatchPortsInput)(nil)).Elem(), WasmPluginSpecMatchPortsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecMatchPortsArrayInput)(nil)).Elem(), WasmPluginSpecMatchPortsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecSelectorInput)(nil)).Elem(), WasmPluginSpecSelectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecSelectorPtrInput)(nil)).Elem(), WasmPluginSpecSelectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecVmconfigInput)(nil)).Elem(), WasmPluginSpecVmconfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecVmconfigPtrInput)(nil)).Elem(), WasmPluginSpecVmconfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecVmconfigEnvInput)(nil)).Elem(), WasmPluginSpecVmconfigEnvArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WasmPluginSpecVmconfigEnvArrayInput)(nil)).Elem(), WasmPluginSpecVmconfigEnvArray{})
	pulumi.RegisterOutputType(WasmPluginSpecOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecPtrOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecMatchOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecMatchArrayOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecMatchPortsOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecMatchPortsArrayOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecSelectorOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecSelectorPtrOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecVmconfigOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecVmconfigPtrOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecVmconfigEnvOutput{})
	pulumi.RegisterOutputType(WasmPluginSpecVmconfigEnvArrayOutput{})
}
