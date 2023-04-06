// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Sidecar struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html
	Spec   SidecarSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput     `pulumi:"status"`
}

// NewSidecar registers a new resource with the given unique name, arguments, and options.
func NewSidecar(ctx *pulumi.Context,
	name string, args *SidecarArgs, opts ...pulumi.ResourceOption) (*Sidecar, error) {
	if args == nil {
		args = &SidecarArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("Sidecar")
	var resource Sidecar
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1beta1:Sidecar", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSidecar gets an existing Sidecar resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSidecar(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SidecarState, opts ...pulumi.ResourceOption) (*Sidecar, error) {
	var resource Sidecar
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1beta1:Sidecar", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sidecar resources.
type sidecarState struct {
}

type SidecarState struct {
}

func (SidecarState) ElementType() reflect.Type {
	return reflect.TypeOf((*sidecarState)(nil)).Elem()
}

type sidecarArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html
	Spec   *SidecarSpec           `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a Sidecar resource.
type SidecarArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html
	Spec   SidecarSpecPtrInput
	Status pulumi.MapInput
}

func (SidecarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sidecarArgs)(nil)).Elem()
}

type SidecarInput interface {
	pulumi.Input

	ToSidecarOutput() SidecarOutput
	ToSidecarOutputWithContext(ctx context.Context) SidecarOutput
}

func (*Sidecar) ElementType() reflect.Type {
	return reflect.TypeOf((**Sidecar)(nil)).Elem()
}

func (i *Sidecar) ToSidecarOutput() SidecarOutput {
	return i.ToSidecarOutputWithContext(context.Background())
}

func (i *Sidecar) ToSidecarOutputWithContext(ctx context.Context) SidecarOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SidecarOutput)
}

type SidecarOutput struct{ *pulumi.OutputState }

func (SidecarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sidecar)(nil)).Elem()
}

func (o SidecarOutput) ToSidecarOutput() SidecarOutput {
	return o
}

func (o SidecarOutput) ToSidecarOutputWithContext(ctx context.Context) SidecarOutput {
	return o
}

func (o SidecarOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sidecar) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o SidecarOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sidecar) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SidecarOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *Sidecar) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html
func (o SidecarOutput) Spec() SidecarSpecPtrOutput {
	return o.ApplyT(func(v *Sidecar) SidecarSpecPtrOutput { return v.Spec }).(SidecarSpecPtrOutput)
}

func (o SidecarOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *Sidecar) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SidecarInput)(nil)).Elem(), &Sidecar{})
	pulumi.RegisterOutputType(SidecarOutput{})
}
