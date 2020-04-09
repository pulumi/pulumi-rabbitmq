// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rabbitmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``.User`` resource creates and manages a user.
//
// > **Note:** All arguments including username and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/user.html.markdown.
type User struct {
	pulumi.CustomResourceState

	// The name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password of the user. The value of this argument
	// is plain-text so make sure to secure where this is defined.
	Password pulumi.StringOutput `pulumi:"password"`
	// Which permission model to apply to the user. Valid
	// options are: management, policymaker, monitoring, and administrator.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("rabbitmq:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("rabbitmq:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password of the user. The value of this argument
	// is plain-text so make sure to secure where this is defined.
	Password *string `pulumi:"password"`
	// Which permission model to apply to the user. Valid
	// options are: management, policymaker, monitoring, and administrator.
	Tags []string `pulumi:"tags"`
}

type UserState struct {
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password of the user. The value of this argument
	// is plain-text so make sure to secure where this is defined.
	Password pulumi.StringPtrInput
	// Which permission model to apply to the user. Valid
	// options are: management, policymaker, monitoring, and administrator.
	Tags pulumi.StringArrayInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password of the user. The value of this argument
	// is plain-text so make sure to secure where this is defined.
	Password string `pulumi:"password"`
	// Which permission model to apply to the user. Valid
	// options are: management, policymaker, monitoring, and administrator.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password of the user. The value of this argument
	// is plain-text so make sure to secure where this is defined.
	Password pulumi.StringInput
	// Which permission model to apply to the user. Valid
	// options are: management, policymaker, monitoring, and administrator.
	Tags pulumi.StringArrayInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}
