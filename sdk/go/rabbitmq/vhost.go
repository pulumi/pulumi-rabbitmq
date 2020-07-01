// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``VHost`` resource creates and manages a vhost.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-rabbitmq/sdk/v2/go/rabbitmq"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
