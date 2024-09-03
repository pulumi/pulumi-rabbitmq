// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Queue` resource creates and manages a queue.
//
// ## Example Usage
//
// ### Basic Example
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
//			guest, err := rabbitmq.NewPermissions(ctx, "guest", &rabbitmq.PermissionsArgs{
//				User:  pulumi.String("guest"),
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
//			_, err = rabbitmq.NewQueue(ctx, "test", &rabbitmq.QueueArgs{
//				Name:  pulumi.String("test"),
//				Vhost: guest.Vhost,
//				Settings: &rabbitmq.QueueSettingsArgs{
//					Durable:    pulumi.Bool(false),
//					AutoDelete: pulumi.Bool(true),
//					Arguments: pulumi.StringMap{
//						"x-queue-type": pulumi.String("quorum"),
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
// ### Example With JSON Arguments
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			arguments := "{\n  \"x-message-ttl\": 5000\n}\n"
//			if param := cfg.Get("arguments"); param != "" {
//				arguments = param
//			}
//			test, err := rabbitmq.NewVHost(ctx, "test", &rabbitmq.VHostArgs{
//				Name: pulumi.String("test"),
//			})
//			if err != nil {
//				return err
//			}
//			guest, err := rabbitmq.NewPermissions(ctx, "guest", &rabbitmq.PermissionsArgs{
//				User:  pulumi.String("guest"),
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
//			_, err = rabbitmq.NewQueue(ctx, "test", &rabbitmq.QueueArgs{
//				Name:  pulumi.String("test"),
//				Vhost: guest.Vhost,
//				Settings: &rabbitmq.QueueSettingsArgs{
//					Durable:       pulumi.Bool(false),
//					AutoDelete:    pulumi.Bool(true),
//					ArgumentsJson: pulumi.String(arguments),
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
// Queues can be imported using the `id` which is composed of `name@vhost`. E.g.
//
// ```sh
// $ pulumi import rabbitmq:index/queue:Queue test name@vhost
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//	QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (i QueueArray) ToQueueArrayOutput() QueueArrayOutput {
	return i.ToQueueArrayOutputWithContext(context.Background())
}

func (i QueueArray) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueArrayOutput)
}

// QueueMapInput is an input type that accepts QueueMap and QueueMapOutput values.
// You can construct a concrete instance of `QueueMapInput` via:
//
//	QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

// The name of the queue.
func (o QueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The settings of the queue. The structure is
// described below.
func (o QueueOutput) Settings() QueueSettingsOutput {
	return o.ApplyT(func(v *Queue) QueueSettingsOutput { return v.Settings }).(QueueSettingsOutput)
}

// The vhost to create the resource in.
func (o QueueOutput) Vhost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.Vhost }).(pulumi.StringPtrOutput)
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].([]*Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].(map[string]*Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueInput)(nil)).Elem(), &Queue{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueArrayInput)(nil)).Elem(), QueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueMapInput)(nil)).Elem(), QueueMap{})
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
