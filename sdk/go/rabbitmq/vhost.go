// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ``VHost`` resource creates and manages a vhost.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rabbitmq.NewVHost(ctx, "myVhost", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Vhosts can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import rabbitmq:index/vHost:VHost my_vhost my_vhost
// ```
type VHost struct {
	pulumi.CustomResourceState

	// The name of the vhost.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewVHost registers a new resource with the given unique name, arguments, and options.
func NewVHost(ctx *pulumi.Context,
	name string, args *VHostArgs, opts ...pulumi.ResourceOption) (*VHost, error) {
	if args == nil {
		args = &VHostArgs{}
	}

	var resource VHost
	err := ctx.RegisterResource("rabbitmq:index/vHost:VHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVHost gets an existing VHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VHostState, opts ...pulumi.ResourceOption) (*VHost, error) {
	var resource VHost
	err := ctx.ReadResource("rabbitmq:index/vHost:VHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VHost resources.
type vhostState struct {
	// The name of the vhost.
	Name *string `pulumi:"name"`
}

type VHostState struct {
	// The name of the vhost.
	Name pulumi.StringPtrInput
}

func (VHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*vhostState)(nil)).Elem()
}

type vhostArgs struct {
	// The name of the vhost.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a VHost resource.
type VHostArgs struct {
	// The name of the vhost.
	Name pulumi.StringPtrInput
}

func (VHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vhostArgs)(nil)).Elem()
}

type VHostInput interface {
	pulumi.Input

	ToVHostOutput() VHostOutput
	ToVHostOutputWithContext(ctx context.Context) VHostOutput
}

func (*VHost) ElementType() reflect.Type {
	return reflect.TypeOf((*VHost)(nil))
}

func (i *VHost) ToVHostOutput() VHostOutput {
	return i.ToVHostOutputWithContext(context.Background())
}

func (i *VHost) ToVHostOutputWithContext(ctx context.Context) VHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VHostOutput)
}

func (i *VHost) ToVHostPtrOutput() VHostPtrOutput {
	return i.ToVHostPtrOutputWithContext(context.Background())
}

func (i *VHost) ToVHostPtrOutputWithContext(ctx context.Context) VHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VHostPtrOutput)
}

type VHostPtrInput interface {
	pulumi.Input

	ToVHostPtrOutput() VHostPtrOutput
	ToVHostPtrOutputWithContext(ctx context.Context) VHostPtrOutput
}

type vhostPtrType VHostArgs

func (*vhostPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VHost)(nil))
}

func (i *vhostPtrType) ToVHostPtrOutput() VHostPtrOutput {
	return i.ToVHostPtrOutputWithContext(context.Background())
}

func (i *vhostPtrType) ToVHostPtrOutputWithContext(ctx context.Context) VHostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VHostPtrOutput)
}

// VHostArrayInput is an input type that accepts VHostArray and VHostArrayOutput values.
// You can construct a concrete instance of `VHostArrayInput` via:
//
//          VHostArray{ VHostArgs{...} }
type VHostArrayInput interface {
	pulumi.Input

	ToVHostArrayOutput() VHostArrayOutput
	ToVHostArrayOutputWithContext(context.Context) VHostArrayOutput
}

type VHostArray []VHostInput

func (VHostArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VHost)(nil))
}

func (i VHostArray) ToVHostArrayOutput() VHostArrayOutput {
	return i.ToVHostArrayOutputWithContext(context.Background())
}

func (i VHostArray) ToVHostArrayOutputWithContext(ctx context.Context) VHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VHostArrayOutput)
}

// VHostMapInput is an input type that accepts VHostMap and VHostMapOutput values.
// You can construct a concrete instance of `VHostMapInput` via:
//
//          VHostMap{ "key": VHostArgs{...} }
type VHostMapInput interface {
	pulumi.Input

	ToVHostMapOutput() VHostMapOutput
	ToVHostMapOutputWithContext(context.Context) VHostMapOutput
}

type VHostMap map[string]VHostInput

func (VHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VHost)(nil))
}

func (i VHostMap) ToVHostMapOutput() VHostMapOutput {
	return i.ToVHostMapOutputWithContext(context.Background())
}

func (i VHostMap) ToVHostMapOutputWithContext(ctx context.Context) VHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VHostMapOutput)
}

type VHostOutput struct {
	*pulumi.OutputState
}

func (VHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VHost)(nil))
}

func (o VHostOutput) ToVHostOutput() VHostOutput {
	return o
}

func (o VHostOutput) ToVHostOutputWithContext(ctx context.Context) VHostOutput {
	return o
}

func (o VHostOutput) ToVHostPtrOutput() VHostPtrOutput {
	return o.ToVHostPtrOutputWithContext(context.Background())
}

func (o VHostOutput) ToVHostPtrOutputWithContext(ctx context.Context) VHostPtrOutput {
	return o.ApplyT(func(v VHost) *VHost {
		return &v
	}).(VHostPtrOutput)
}

type VHostPtrOutput struct {
	*pulumi.OutputState
}

func (VHostPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VHost)(nil))
}

func (o VHostPtrOutput) ToVHostPtrOutput() VHostPtrOutput {
	return o
}

func (o VHostPtrOutput) ToVHostPtrOutputWithContext(ctx context.Context) VHostPtrOutput {
	return o
}

type VHostArrayOutput struct{ *pulumi.OutputState }

func (VHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VHost)(nil))
}

func (o VHostArrayOutput) ToVHostArrayOutput() VHostArrayOutput {
	return o
}

func (o VHostArrayOutput) ToVHostArrayOutputWithContext(ctx context.Context) VHostArrayOutput {
	return o
}

func (o VHostArrayOutput) Index(i pulumi.IntInput) VHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VHost {
		return vs[0].([]VHost)[vs[1].(int)]
	}).(VHostOutput)
}

type VHostMapOutput struct{ *pulumi.OutputState }

func (VHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VHost)(nil))
}

func (o VHostMapOutput) ToVHostMapOutput() VHostMapOutput {
	return o
}

func (o VHostMapOutput) ToVHostMapOutputWithContext(ctx context.Context) VHostMapOutput {
	return o
}

func (o VHostMapOutput) MapIndex(k pulumi.StringInput) VHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VHost {
		return vs[0].(map[string]VHost)[vs[1].(string)]
	}).(VHostOutput)
}

func init() {
	pulumi.RegisterOutputType(VHostOutput{})
	pulumi.RegisterOutputType(VHostPtrOutput{})
	pulumi.RegisterOutputType(VHostArrayOutput{})
	pulumi.RegisterOutputType(VHostMapOutput{})
}
