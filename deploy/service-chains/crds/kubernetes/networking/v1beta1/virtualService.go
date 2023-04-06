// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualService struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html
	Spec   VirtualServiceSpecPtrOutput `pulumi:"spec"`
	Status pulumi.MapOutput            `pulumi:"status"`
}

// NewVirtualService registers a new resource with the given unique name, arguments, and options.
func NewVirtualService(ctx *pulumi.Context,
	name string, args *VirtualServiceArgs, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	if args == nil {
		args = &VirtualServiceArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.istio.io/v1beta1")
	args.Kind = pulumi.StringPtr("VirtualService")
	var resource VirtualService
	err := ctx.RegisterResource("kubernetes:networking.istio.io/v1beta1:VirtualService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualService gets an existing VirtualService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualServiceState, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	var resource VirtualService
	err := ctx.ReadResource("kubernetes:networking.istio.io/v1beta1:VirtualService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualService resources.
type virtualServiceState struct {
}

type VirtualServiceState struct {
}

func (VirtualServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceState)(nil)).Elem()
}

type virtualServiceArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html
	Spec   *VirtualServiceSpec    `pulumi:"spec"`
	Status map[string]interface{} `pulumi:"status"`
}

// The set of arguments for constructing a VirtualService resource.
type VirtualServiceArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html
	Spec   VirtualServiceSpecPtrInput
	Status pulumi.MapInput
}

func (VirtualServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceArgs)(nil)).Elem()
}

type VirtualServiceInput interface {
	pulumi.Input

	ToVirtualServiceOutput() VirtualServiceOutput
	ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput
}

func (*VirtualService) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualService)(nil)).Elem()
}

func (i *VirtualService) ToVirtualServiceOutput() VirtualServiceOutput {
	return i.ToVirtualServiceOutputWithContext(context.Background())
}

func (i *VirtualService) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceOutput)
}

type VirtualServiceOutput struct{ *pulumi.OutputState }

func (VirtualServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualService)(nil)).Elem()
}

func (o VirtualServiceOutput) ToVirtualServiceOutput() VirtualServiceOutput {
	return o
}

func (o VirtualServiceOutput) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return o
}

func (o VirtualServiceOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualServiceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VirtualServiceOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *VirtualService) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Configuration affecting label/content routing, sni routing, etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html
func (o VirtualServiceOutput) Spec() VirtualServiceSpecPtrOutput {
	return o.ApplyT(func(v *VirtualService) VirtualServiceSpecPtrOutput { return v.Spec }).(VirtualServiceSpecPtrOutput)
}

func (o VirtualServiceOutput) Status() pulumi.MapOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.MapOutput { return v.Status }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualServiceInput)(nil)).Elem(), &VirtualService{})
	pulumi.RegisterOutputType(VirtualServiceOutput{})
}