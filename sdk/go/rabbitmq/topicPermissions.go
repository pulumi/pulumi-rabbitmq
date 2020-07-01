// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``TopicPermissions`` resource creates and manages a user's set of
// topic permissions.
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
// 		testVHost, err := rabbitmq.NewVHost(ctx, "testVHost", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testUser, err := rabbitmq.NewUser(ctx, "testUser", &rabbitmq.UserArgs{
// 			Password: pulumi.String("foobar"),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("administrator"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rabbitmq.NewTopicPermissions(ctx, "testTopicPermissions", &rabbitmq.TopicPermissionsArgs{
// 			Permissions: rabbitmq.TopicPermissionsPermissionArray{
// 				&rabbitmq.TopicPermissionsPermissionArgs{
// 					Exchange: pulumi.String("amq.topic"),
// 					Read:     pulumi.String(".*"),
// 					Write:    pulumi.String(".*"),
// 				},
// 			},
// 			User:  testUser.Name,
// 			Vhost: testVHost.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
	if args == nil || args.Permissions == nil {
		return nil, errors.New("missing required argument 'Permissions'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	if args == nil {
		args = &TopicPermissionsArgs{}
	}
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
