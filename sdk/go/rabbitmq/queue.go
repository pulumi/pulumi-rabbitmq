// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``.Queue`` resource creates and manages a queue.
type Queue struct {
	pulumi.CustomResourceState

	// The name of the queue.
	Name pulumi.StringOutput `pulumi:"name"`
	// The settings of the queue. The structure is
	// described below.
	Settings QueueSettingsOutput `pulumi:"settings"`
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrOutput `pulumi:"vhost"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	if args == nil {
		args = &QueueArgs{}
	}
	var resource Queue
	err := ctx.RegisterResource("rabbitmq:index/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("rabbitmq:index/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The name of the queue.
	Name *string `pulumi:"name"`
	// The settings of the queue. The structure is
	// described below.
	Settings *QueueSettings `pulumi:"settings"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type QueueState struct {
	// The name of the queue.
	Name pulumi.StringPtrInput
	// The settings of the queue. The structure is
	// described below.
	Settings QueueSettingsPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// The name of the queue.
	Name *string `pulumi:"name"`
	// The settings of the queue. The structure is
	// described below.
	Settings QueueSettings `pulumi:"settings"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// The name of the queue.
	Name pulumi.StringPtrInput
	// The settings of the queue. The structure is
	// described below.
	Settings QueueSettingsInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}
