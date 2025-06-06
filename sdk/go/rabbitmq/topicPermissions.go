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

// The “TopicPermissions“ resource creates and manages a user's set of
// topic permissions.
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
//			_, err = rabbitmq.NewTopicPermissions(ctx, "test", &rabbitmq.TopicPermissionsArgs{
//				User:  testUser.Name,
//				Vhost: test.Name,
//				Permissions: rabbitmq.TopicPermissionsPermissionArray{
//					&rabbitmq.TopicPermissionsPermissionArgs{
//						Exchange: pulumi.String("amq.topic"),
//						Write:    pulumi.String(".*"),
//						Read:     pulumi.String(".*"),
//					},
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
// $ pulumi import rabbitmq:index/topicPermissions:TopicPermissions test user@vhost
// ```
type TopicPermissions struct {
	pulumi.CustomResourceState

	// The settings of the permissions. The structure is
	// described below.
	Permissions TopicPermissionsPermissionArrayOutput `pulumi:"permissions"`
	// The user to apply the permissions to.
	User pulumi.StringOutput `pulumi:"user"`
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrOutput `pulumi:"vhost"`
}

// NewTopicPermissions registers a new resource with the given unique name, arguments, and options.
func NewTopicPermissions(ctx *pulumi.Context,
	name string, args *TopicPermissionsArgs, opts ...pulumi.ResourceOption) (*TopicPermissions, error) {
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
	var resource TopicPermissions
	err := ctx.RegisterResource("rabbitmq:index/topicPermissions:TopicPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicPermissions gets an existing TopicPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicPermissionsState, opts ...pulumi.ResourceOption) (*TopicPermissions, error) {
	var resource TopicPermissions
	err := ctx.ReadResource("rabbitmq:index/topicPermissions:TopicPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicPermissions resources.
type topicPermissionsState struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions []TopicPermissionsPermission `pulumi:"permissions"`
	// The user to apply the permissions to.
	User *string `pulumi:"user"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type TopicPermissionsState struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions TopicPermissionsPermissionArrayInput
	// The user to apply the permissions to.
	User pulumi.StringPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (TopicPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicPermissionsState)(nil)).Elem()
}

type topicPermissionsArgs struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions []TopicPermissionsPermission `pulumi:"permissions"`
	// The user to apply the permissions to.
	User string `pulumi:"user"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

// The set of arguments for constructing a TopicPermissions resource.
type TopicPermissionsArgs struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions TopicPermissionsPermissionArrayInput
	// The user to apply the permissions to.
	User pulumi.StringInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (TopicPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicPermissionsArgs)(nil)).Elem()
}

type TopicPermissionsInput interface {
	pulumi.Input

	ToTopicPermissionsOutput() TopicPermissionsOutput
	ToTopicPermissionsOutputWithContext(ctx context.Context) TopicPermissionsOutput
}

func (*TopicPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicPermissions)(nil)).Elem()
}

func (i *TopicPermissions) ToTopicPermissionsOutput() TopicPermissionsOutput {
	return i.ToTopicPermissionsOutputWithContext(context.Background())
}

func (i *TopicPermissions) ToTopicPermissionsOutputWithContext(ctx context.Context) TopicPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPermissionsOutput)
}

// TopicPermissionsArrayInput is an input type that accepts TopicPermissionsArray and TopicPermissionsArrayOutput values.
// You can construct a concrete instance of `TopicPermissionsArrayInput` via:
//
//	TopicPermissionsArray{ TopicPermissionsArgs{...} }
type TopicPermissionsArrayInput interface {
	pulumi.Input

	ToTopicPermissionsArrayOutput() TopicPermissionsArrayOutput
	ToTopicPermissionsArrayOutputWithContext(context.Context) TopicPermissionsArrayOutput
}

type TopicPermissionsArray []TopicPermissionsInput

func (TopicPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TopicPermissions)(nil)).Elem()
}

func (i TopicPermissionsArray) ToTopicPermissionsArrayOutput() TopicPermissionsArrayOutput {
	return i.ToTopicPermissionsArrayOutputWithContext(context.Background())
}

func (i TopicPermissionsArray) ToTopicPermissionsArrayOutputWithContext(ctx context.Context) TopicPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPermissionsArrayOutput)
}

// TopicPermissionsMapInput is an input type that accepts TopicPermissionsMap and TopicPermissionsMapOutput values.
// You can construct a concrete instance of `TopicPermissionsMapInput` via:
//
//	TopicPermissionsMap{ "key": TopicPermissionsArgs{...} }
type TopicPermissionsMapInput interface {
	pulumi.Input

	ToTopicPermissionsMapOutput() TopicPermissionsMapOutput
	ToTopicPermissionsMapOutputWithContext(context.Context) TopicPermissionsMapOutput
}

type TopicPermissionsMap map[string]TopicPermissionsInput

func (TopicPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TopicPermissions)(nil)).Elem()
}

func (i TopicPermissionsMap) ToTopicPermissionsMapOutput() TopicPermissionsMapOutput {
	return i.ToTopicPermissionsMapOutputWithContext(context.Background())
}

func (i TopicPermissionsMap) ToTopicPermissionsMapOutputWithContext(ctx context.Context) TopicPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPermissionsMapOutput)
}

type TopicPermissionsOutput struct{ *pulumi.OutputState }

func (TopicPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicPermissions)(nil)).Elem()
}

func (o TopicPermissionsOutput) ToTopicPermissionsOutput() TopicPermissionsOutput {
	return o
}

func (o TopicPermissionsOutput) ToTopicPermissionsOutputWithContext(ctx context.Context) TopicPermissionsOutput {
	return o
}

// The settings of the permissions. The structure is
// described below.
func (o TopicPermissionsOutput) Permissions() TopicPermissionsPermissionArrayOutput {
	return o.ApplyT(func(v *TopicPermissions) TopicPermissionsPermissionArrayOutput { return v.Permissions }).(TopicPermissionsPermissionArrayOutput)
}

// The user to apply the permissions to.
func (o TopicPermissionsOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicPermissions) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// The vhost to create the resource in.
func (o TopicPermissionsOutput) Vhost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopicPermissions) pulumi.StringPtrOutput { return v.Vhost }).(pulumi.StringPtrOutput)
}

type TopicPermissionsArrayOutput struct{ *pulumi.OutputState }

func (TopicPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TopicPermissions)(nil)).Elem()
}

func (o TopicPermissionsArrayOutput) ToTopicPermissionsArrayOutput() TopicPermissionsArrayOutput {
	return o
}

func (o TopicPermissionsArrayOutput) ToTopicPermissionsArrayOutputWithContext(ctx context.Context) TopicPermissionsArrayOutput {
	return o
}

func (o TopicPermissionsArrayOutput) Index(i pulumi.IntInput) TopicPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TopicPermissions {
		return vs[0].([]*TopicPermissions)[vs[1].(int)]
	}).(TopicPermissionsOutput)
}

type TopicPermissionsMapOutput struct{ *pulumi.OutputState }

func (TopicPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TopicPermissions)(nil)).Elem()
}

func (o TopicPermissionsMapOutput) ToTopicPermissionsMapOutput() TopicPermissionsMapOutput {
	return o
}

func (o TopicPermissionsMapOutput) ToTopicPermissionsMapOutputWithContext(ctx context.Context) TopicPermissionsMapOutput {
	return o
}

func (o TopicPermissionsMapOutput) MapIndex(k pulumi.StringInput) TopicPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TopicPermissions {
		return vs[0].(map[string]*TopicPermissions)[vs[1].(string)]
	}).(TopicPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicPermissionsInput)(nil)).Elem(), &TopicPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicPermissionsArrayInput)(nil)).Elem(), TopicPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicPermissionsMapInput)(nil)).Elem(), TopicPermissionsMap{})
	pulumi.RegisterOutputType(TopicPermissionsOutput{})
	pulumi.RegisterOutputType(TopicPermissionsArrayOutput{})
	pulumi.RegisterOutputType(TopicPermissionsMapOutput{})
}
