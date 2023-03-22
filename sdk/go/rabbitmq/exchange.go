// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “Exchange“ resource creates and manages an exchange.
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
//			guest, err := rabbitmq.NewPermissions(ctx, "guest", &rabbitmq.PermissionsArgs{
//				Permissions: &rabbitmq.PermissionsPermissionsArgs{
//					Configure: pulumi.String(".*"),
//					Read:      pulumi.String(".*"),
//					Write:     pulumi.String(".*"),
//				},
//				User:  pulumi.String("guest"),
//				Vhost: testVHost.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rabbitmq.NewExchange(ctx, "testExchange", &rabbitmq.ExchangeArgs{
//				Settings: &rabbitmq.ExchangeSettingsArgs{
//					AutoDelete: pulumi.Bool(true),
//					Durable:    pulumi.Bool(false),
//					Type:       pulumi.String("fanout"),
//				},
//				Vhost: guest.Vhost,
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
// # Exchanges can be imported using the `id` which is composed of
//
// `name@vhost`. E.g.
//
// ```sh
//
//	$ pulumi import rabbitmq:index/exchange:Exchange test test@vhost
//
// ```
type Exchange struct {
	pulumi.CustomResourceState

	// The name of the exchange.
	Name pulumi.StringOutput `pulumi:"name"`
	// The settings of the exchange. The structure is
	// described below.
	Settings ExchangeSettingsOutput `pulumi:"settings"`
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrOutput `pulumi:"vhost"`
}

// NewExchange registers a new resource with the given unique name, arguments, and options.
func NewExchange(ctx *pulumi.Context,
	name string, args *ExchangeArgs, opts ...pulumi.ResourceOption) (*Exchange, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	var resource Exchange
	err := ctx.RegisterResource("rabbitmq:index/exchange:Exchange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExchange gets an existing Exchange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExchange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExchangeState, opts ...pulumi.ResourceOption) (*Exchange, error) {
	var resource Exchange
	err := ctx.ReadResource("rabbitmq:index/exchange:Exchange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Exchange resources.
type exchangeState struct {
	// The name of the exchange.
	Name *string `pulumi:"name"`
	// The settings of the exchange. The structure is
	// described below.
	Settings *ExchangeSettings `pulumi:"settings"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type ExchangeState struct {
	// The name of the exchange.
	Name pulumi.StringPtrInput
	// The settings of the exchange. The structure is
	// described below.
	Settings ExchangeSettingsPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (ExchangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*exchangeState)(nil)).Elem()
}

type exchangeArgs struct {
	// The name of the exchange.
	Name *string `pulumi:"name"`
	// The settings of the exchange. The structure is
	// described below.
	Settings ExchangeSettings `pulumi:"settings"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

// The set of arguments for constructing a Exchange resource.
type ExchangeArgs struct {
	// The name of the exchange.
	Name pulumi.StringPtrInput
	// The settings of the exchange. The structure is
	// described below.
	Settings ExchangeSettingsInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (ExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exchangeArgs)(nil)).Elem()
}

type ExchangeInput interface {
	pulumi.Input

	ToExchangeOutput() ExchangeOutput
	ToExchangeOutputWithContext(ctx context.Context) ExchangeOutput
}

func (*Exchange) ElementType() reflect.Type {
	return reflect.TypeOf((**Exchange)(nil)).Elem()
}

func (i *Exchange) ToExchangeOutput() ExchangeOutput {
	return i.ToExchangeOutputWithContext(context.Background())
}

func (i *Exchange) ToExchangeOutputWithContext(ctx context.Context) ExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeOutput)
}

// ExchangeArrayInput is an input type that accepts ExchangeArray and ExchangeArrayOutput values.
// You can construct a concrete instance of `ExchangeArrayInput` via:
//
//	ExchangeArray{ ExchangeArgs{...} }
type ExchangeArrayInput interface {
	pulumi.Input

	ToExchangeArrayOutput() ExchangeArrayOutput
	ToExchangeArrayOutputWithContext(context.Context) ExchangeArrayOutput
}

type ExchangeArray []ExchangeInput

func (ExchangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Exchange)(nil)).Elem()
}

func (i ExchangeArray) ToExchangeArrayOutput() ExchangeArrayOutput {
	return i.ToExchangeArrayOutputWithContext(context.Background())
}

func (i ExchangeArray) ToExchangeArrayOutputWithContext(ctx context.Context) ExchangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeArrayOutput)
}

// ExchangeMapInput is an input type that accepts ExchangeMap and ExchangeMapOutput values.
// You can construct a concrete instance of `ExchangeMapInput` via:
//
//	ExchangeMap{ "key": ExchangeArgs{...} }
type ExchangeMapInput interface {
	pulumi.Input

	ToExchangeMapOutput() ExchangeMapOutput
	ToExchangeMapOutputWithContext(context.Context) ExchangeMapOutput
}

type ExchangeMap map[string]ExchangeInput

func (ExchangeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Exchange)(nil)).Elem()
}

func (i ExchangeMap) ToExchangeMapOutput() ExchangeMapOutput {
	return i.ToExchangeMapOutputWithContext(context.Background())
}

func (i ExchangeMap) ToExchangeMapOutputWithContext(ctx context.Context) ExchangeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeMapOutput)
}

type ExchangeOutput struct{ *pulumi.OutputState }

func (ExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Exchange)(nil)).Elem()
}

func (o ExchangeOutput) ToExchangeOutput() ExchangeOutput {
	return o
}

func (o ExchangeOutput) ToExchangeOutputWithContext(ctx context.Context) ExchangeOutput {
	return o
}

// The name of the exchange.
func (o ExchangeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The settings of the exchange. The structure is
// described below.
func (o ExchangeOutput) Settings() ExchangeSettingsOutput {
	return o.ApplyT(func(v *Exchange) ExchangeSettingsOutput { return v.Settings }).(ExchangeSettingsOutput)
}

// The vhost to create the resource in.
func (o ExchangeOutput) Vhost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Exchange) pulumi.StringPtrOutput { return v.Vhost }).(pulumi.StringPtrOutput)
}

type ExchangeArrayOutput struct{ *pulumi.OutputState }

func (ExchangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Exchange)(nil)).Elem()
}

func (o ExchangeArrayOutput) ToExchangeArrayOutput() ExchangeArrayOutput {
	return o
}

func (o ExchangeArrayOutput) ToExchangeArrayOutputWithContext(ctx context.Context) ExchangeArrayOutput {
	return o
}

func (o ExchangeArrayOutput) Index(i pulumi.IntInput) ExchangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Exchange {
		return vs[0].([]*Exchange)[vs[1].(int)]
	}).(ExchangeOutput)
}

type ExchangeMapOutput struct{ *pulumi.OutputState }

func (ExchangeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Exchange)(nil)).Elem()
}

func (o ExchangeMapOutput) ToExchangeMapOutput() ExchangeMapOutput {
	return o
}

func (o ExchangeMapOutput) ToExchangeMapOutputWithContext(ctx context.Context) ExchangeMapOutput {
	return o
}

func (o ExchangeMapOutput) MapIndex(k pulumi.StringInput) ExchangeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Exchange {
		return vs[0].(map[string]*Exchange)[vs[1].(string)]
	}).(ExchangeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeInput)(nil)).Elem(), &Exchange{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeArrayInput)(nil)).Elem(), ExchangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExchangeMapInput)(nil)).Elem(), ExchangeMap{})
	pulumi.RegisterOutputType(ExchangeOutput{})
	pulumi.RegisterOutputType(ExchangeArrayOutput{})
	pulumi.RegisterOutputType(ExchangeMapOutput{})
}
