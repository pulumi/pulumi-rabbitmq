// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “Shovel“ resource creates and manages a dynamic shovel.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-rabbitmq/sdk/go/rabbitmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testVHost, err := rabbitmq.NewVHost(ctx, "testVHost", nil)
//			if err != nil {
//				return err
//			}
//			testExchange, err := rabbitmq.NewExchange(ctx, "testExchange", &rabbitmq.ExchangeArgs{
//				Settings: &ExchangeSettingsArgs{
//					AutoDelete: pulumi.Bool(true),
//					Durable:    pulumi.Bool(false),
//					Type:       pulumi.String("fanout"),
//				},
//				Vhost: testVHost.Name,
//			})
//			if err != nil {
//				return err
//			}
//			testQueue, err := rabbitmq.NewQueue(ctx, "testQueue", &rabbitmq.QueueArgs{
//				Settings: &QueueSettingsArgs{
//					AutoDelete: pulumi.Bool(true),
//					Durable:    pulumi.Bool(false),
//				},
//				Vhost: testVHost.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rabbitmq.NewShovel(ctx, "shovelTest", &rabbitmq.ShovelArgs{
//				Info: &ShovelInfoArgs{
//					DestinationQueue:  testQueue.Name,
//					DestinationUri:    pulumi.String("amqp:///test"),
//					SourceExchange:    testExchange.Name,
//					SourceExchangeKey: pulumi.String("test"),
//					SourceUri:         pulumi.String("amqp:///test"),
//				},
//				Vhost: testVHost.Name,
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
// Shovels can be imported using the `name` and `vhost` E.g.
//
// ```sh
//
//	$ pulumi import rabbitmq:index/shovel:Shovel test shovelTest@test
//
// ```
type Shovel struct {
	pulumi.CustomResourceState

	// The settings of the dynamic shovel. The structure is
	// described below.
	Info ShovelInfoOutput `pulumi:"info"`
	// The shovel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The vhost to create the resource in.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
}

// NewShovel registers a new resource with the given unique name, arguments, and options.
func NewShovel(ctx *pulumi.Context,
	name string, args *ShovelArgs, opts ...pulumi.ResourceOption) (*Shovel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Info == nil {
		return nil, errors.New("invalid value for required argument 'Info'")
	}
	if args.Vhost == nil {
		return nil, errors.New("invalid value for required argument 'Vhost'")
	}
	var resource Shovel
	err := ctx.RegisterResource("rabbitmq:index/shovel:Shovel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShovel gets an existing Shovel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShovel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShovelState, opts ...pulumi.ResourceOption) (*Shovel, error) {
	var resource Shovel
	err := ctx.ReadResource("rabbitmq:index/shovel:Shovel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Shovel resources.
type shovelState struct {
	// The settings of the dynamic shovel. The structure is
	// described below.
	Info *ShovelInfo `pulumi:"info"`
	// The shovel name.
	Name *string `pulumi:"name"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type ShovelState struct {
	// The settings of the dynamic shovel. The structure is
	// described below.
	Info ShovelInfoPtrInput
	// The shovel name.
	Name pulumi.StringPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (ShovelState) ElementType() reflect.Type {
	return reflect.TypeOf((*shovelState)(nil)).Elem()
}

type shovelArgs struct {
	// The settings of the dynamic shovel. The structure is
	// described below.
	Info ShovelInfo `pulumi:"info"`
	// The shovel name.
	Name *string `pulumi:"name"`
	// The vhost to create the resource in.
	Vhost string `pulumi:"vhost"`
}

// The set of arguments for constructing a Shovel resource.
type ShovelArgs struct {
	// The settings of the dynamic shovel. The structure is
	// described below.
	Info ShovelInfoInput
	// The shovel name.
	Name pulumi.StringPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringInput
}

func (ShovelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shovelArgs)(nil)).Elem()
}

type ShovelInput interface {
	pulumi.Input

	ToShovelOutput() ShovelOutput
	ToShovelOutputWithContext(ctx context.Context) ShovelOutput
}

func (*Shovel) ElementType() reflect.Type {
	return reflect.TypeOf((**Shovel)(nil)).Elem()
}

func (i *Shovel) ToShovelOutput() ShovelOutput {
	return i.ToShovelOutputWithContext(context.Background())
}

func (i *Shovel) ToShovelOutputWithContext(ctx context.Context) ShovelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShovelOutput)
}

// ShovelArrayInput is an input type that accepts ShovelArray and ShovelArrayOutput values.
// You can construct a concrete instance of `ShovelArrayInput` via:
//
//	ShovelArray{ ShovelArgs{...} }
type ShovelArrayInput interface {
	pulumi.Input

	ToShovelArrayOutput() ShovelArrayOutput
	ToShovelArrayOutputWithContext(context.Context) ShovelArrayOutput
}

type ShovelArray []ShovelInput

func (ShovelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Shovel)(nil)).Elem()
}

func (i ShovelArray) ToShovelArrayOutput() ShovelArrayOutput {
	return i.ToShovelArrayOutputWithContext(context.Background())
}

func (i ShovelArray) ToShovelArrayOutputWithContext(ctx context.Context) ShovelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShovelArrayOutput)
}

// ShovelMapInput is an input type that accepts ShovelMap and ShovelMapOutput values.
// You can construct a concrete instance of `ShovelMapInput` via:
//
//	ShovelMap{ "key": ShovelArgs{...} }
type ShovelMapInput interface {
	pulumi.Input

	ToShovelMapOutput() ShovelMapOutput
	ToShovelMapOutputWithContext(context.Context) ShovelMapOutput
}

type ShovelMap map[string]ShovelInput

func (ShovelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Shovel)(nil)).Elem()
}

func (i ShovelMap) ToShovelMapOutput() ShovelMapOutput {
	return i.ToShovelMapOutputWithContext(context.Background())
}

func (i ShovelMap) ToShovelMapOutputWithContext(ctx context.Context) ShovelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShovelMapOutput)
}

type ShovelOutput struct{ *pulumi.OutputState }

func (ShovelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Shovel)(nil)).Elem()
}

func (o ShovelOutput) ToShovelOutput() ShovelOutput {
	return o
}

func (o ShovelOutput) ToShovelOutputWithContext(ctx context.Context) ShovelOutput {
	return o
}

// The settings of the dynamic shovel. The structure is
// described below.
func (o ShovelOutput) Info() ShovelInfoOutput {
	return o.ApplyT(func(v *Shovel) ShovelInfoOutput { return v.Info }).(ShovelInfoOutput)
}

// The shovel name.
func (o ShovelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Shovel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The vhost to create the resource in.
func (o ShovelOutput) Vhost() pulumi.StringOutput {
	return o.ApplyT(func(v *Shovel) pulumi.StringOutput { return v.Vhost }).(pulumi.StringOutput)
}

type ShovelArrayOutput struct{ *pulumi.OutputState }

func (ShovelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Shovel)(nil)).Elem()
}

func (o ShovelArrayOutput) ToShovelArrayOutput() ShovelArrayOutput {
	return o
}

func (o ShovelArrayOutput) ToShovelArrayOutputWithContext(ctx context.Context) ShovelArrayOutput {
	return o
}

func (o ShovelArrayOutput) Index(i pulumi.IntInput) ShovelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Shovel {
		return vs[0].([]*Shovel)[vs[1].(int)]
	}).(ShovelOutput)
}

type ShovelMapOutput struct{ *pulumi.OutputState }

func (ShovelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Shovel)(nil)).Elem()
}

func (o ShovelMapOutput) ToShovelMapOutput() ShovelMapOutput {
	return o
}

func (o ShovelMapOutput) ToShovelMapOutputWithContext(ctx context.Context) ShovelMapOutput {
	return o
}

func (o ShovelMapOutput) MapIndex(k pulumi.StringInput) ShovelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Shovel {
		return vs[0].(map[string]*Shovel)[vs[1].(string)]
	}).(ShovelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShovelInput)(nil)).Elem(), &Shovel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShovelArrayInput)(nil)).Elem(), ShovelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShovelMapInput)(nil)).Elem(), ShovelMap{})
	pulumi.RegisterOutputType(ShovelOutput{})
	pulumi.RegisterOutputType(ShovelArrayOutput{})
	pulumi.RegisterOutputType(ShovelMapOutput{})
}
