// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “Permissions“ resource creates and manages a user's set of
// permissions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := rabbitmq.NewVHost(ctx, "test", &rabbitmq.VHostArgs{
//				Name: pulumi.String("test"),
//			})
//			if err != nil {
//				return err
//			}
//			testUser, err := rabbitmq.NewUser(ctx, "test", &rabbitmq.UserArgs{
//				Name:     pulumi.String("mctest"),
//				Password: pulumi.String("foobar"),
//				Tags: pulumi.StringArray{
//					pulumi.String("administrator"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rabbitmq.NewPermissions(ctx, "test", &rabbitmq.PermissionsArgs{
//				User:  testUser.Name,
//				Vhost: test.Name,
//				Permissions: &rabbitmq.PermissionsPermissionsArgs{
//					Configure: pulumi.String(".*"),
//					Write:     pulumi.String(".*"),
//					Read:      pulumi.String(".*"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Permissions can be imported using the `id` which is composed of  `user@vhost`.
//
// E.g.
//
// ```sh
// $ pulumi import rabbitmq:index/permissions:Permissions test user@vhost
// ```
type Permissions struct {
	pulumi.CustomResourceState

	// The settings of the permissions. The structure is
	// described below.
	Permissions PermissionsPermissionsOutput `pulumi:"permissions"`
	// The user to apply the permissions to.
	User pulumi.StringOutput `pulumi:"user"`
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrOutput `pulumi:"vhost"`
}

// NewPermissions registers a new resource with the given unique name, arguments, and options.
func NewPermissions(ctx *pulumi.Context,
	name string, args *PermissionsArgs, opts ...pulumi.ResourceOption) (*Permissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Permissions
	err := ctx.RegisterResource("rabbitmq:index/permissions:Permissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissions gets an existing Permissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionsState, opts ...pulumi.ResourceOption) (*Permissions, error) {
	var resource Permissions
	err := ctx.ReadResource("rabbitmq:index/permissions:Permissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permissions resources.
type permissionsState struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions *PermissionsPermissions `pulumi:"permissions"`
	// The user to apply the permissions to.
	User *string `pulumi:"user"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type PermissionsState struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions PermissionsPermissionsPtrInput
	// The user to apply the permissions to.
	User pulumi.StringPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (PermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionsState)(nil)).Elem()
}

type permissionsArgs struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions PermissionsPermissions `pulumi:"permissions"`
	// The user to apply the permissions to.
	User string `pulumi:"user"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

// The set of arguments for constructing a Permissions resource.
type PermissionsArgs struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions PermissionsPermissionsInput
	// The user to apply the permissions to.
	User pulumi.StringInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (PermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionsArgs)(nil)).Elem()
}

type PermissionsInput interface {
	pulumi.Input

	ToPermissionsOutput() PermissionsOutput
	ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput
}

func (*Permissions) ElementType() reflect.Type {
	return reflect.TypeOf((**Permissions)(nil)).Elem()
}

func (i *Permissions) ToPermissionsOutput() PermissionsOutput {
	return i.ToPermissionsOutputWithContext(context.Background())
}

func (i *Permissions) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsOutput)
}

// PermissionsArrayInput is an input type that accepts PermissionsArray and PermissionsArrayOutput values.
// You can construct a concrete instance of `PermissionsArrayInput` via:
//
//	PermissionsArray{ PermissionsArgs{...} }
type PermissionsArrayInput interface {
	pulumi.Input

	ToPermissionsArrayOutput() PermissionsArrayOutput
	ToPermissionsArrayOutputWithContext(context.Context) PermissionsArrayOutput
}

type PermissionsArray []PermissionsInput

func (PermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Permissions)(nil)).Elem()
}

func (i PermissionsArray) ToPermissionsArrayOutput() PermissionsArrayOutput {
	return i.ToPermissionsArrayOutputWithContext(context.Background())
}

func (i PermissionsArray) ToPermissionsArrayOutputWithContext(ctx context.Context) PermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsArrayOutput)
}

// PermissionsMapInput is an input type that accepts PermissionsMap and PermissionsMapOutput values.
// You can construct a concrete instance of `PermissionsMapInput` via:
//
//	PermissionsMap{ "key": PermissionsArgs{...} }
type PermissionsMapInput interface {
	pulumi.Input

	ToPermissionsMapOutput() PermissionsMapOutput
	ToPermissionsMapOutputWithContext(context.Context) PermissionsMapOutput
}

type PermissionsMap map[string]PermissionsInput

func (PermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Permissions)(nil)).Elem()
}

func (i PermissionsMap) ToPermissionsMapOutput() PermissionsMapOutput {
	return i.ToPermissionsMapOutputWithContext(context.Background())
}

func (i PermissionsMap) ToPermissionsMapOutputWithContext(ctx context.Context) PermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsMapOutput)
}

type PermissionsOutput struct{ *pulumi.OutputState }

func (PermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permissions)(nil)).Elem()
}

func (o PermissionsOutput) ToPermissionsOutput() PermissionsOutput {
	return o
}

func (o PermissionsOutput) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return o
}

// The settings of the permissions. The structure is
// described below.
func (o PermissionsOutput) Permissions() PermissionsPermissionsOutput {
	return o.ApplyT(func(v *Permissions) PermissionsPermissionsOutput { return v.Permissions }).(PermissionsPermissionsOutput)
}

// The user to apply the permissions to.
func (o PermissionsOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Permissions) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// The vhost to create the resource in.
func (o PermissionsOutput) Vhost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permissions) pulumi.StringPtrOutput { return v.Vhost }).(pulumi.StringPtrOutput)
}

type PermissionsArrayOutput struct{ *pulumi.OutputState }

func (PermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Permissions)(nil)).Elem()
}

func (o PermissionsArrayOutput) ToPermissionsArrayOutput() PermissionsArrayOutput {
	return o
}

func (o PermissionsArrayOutput) ToPermissionsArrayOutputWithContext(ctx context.Context) PermissionsArrayOutput {
	return o
}

func (o PermissionsArrayOutput) Index(i pulumi.IntInput) PermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Permissions {
		return vs[0].([]*Permissions)[vs[1].(int)]
	}).(PermissionsOutput)
}

type PermissionsMapOutput struct{ *pulumi.OutputState }

func (PermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Permissions)(nil)).Elem()
}

func (o PermissionsMapOutput) ToPermissionsMapOutput() PermissionsMapOutput {
	return o
}

func (o PermissionsMapOutput) ToPermissionsMapOutputWithContext(ctx context.Context) PermissionsMapOutput {
	return o
}

func (o PermissionsMapOutput) MapIndex(k pulumi.StringInput) PermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Permissions {
		return vs[0].(map[string]*Permissions)[vs[1].(string)]
	}).(PermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsInput)(nil)).Elem(), &Permissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsArrayInput)(nil)).Elem(), PermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsMapInput)(nil)).Elem(), PermissionsMap{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsArrayOutput{})
	pulumi.RegisterOutputType(PermissionsMapOutput{})
}
