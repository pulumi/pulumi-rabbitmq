// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rabbitmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``.Permissions`` resource creates and manages a user's set of
// permissions.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/permissions.html.markdown.
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
	if args == nil || args.Permissions == nil {
		return nil, errors.New("missing required argument 'Permissions'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	if args == nil {
		args = &PermissionsArgs{}
	}
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
