// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ``Binding`` resource creates and manages a binding relationship
// between a queue an exchange.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testVHost, err := rabbitmq.NewVHost(ctx, "testVHost", nil)
// 		if err != nil {
// 			return err
// 		}
// 		guest, err := rabbitmq.NewPermissions(ctx, "guest", &rabbitmq.PermissionsArgs{
// 			Permissions: &PermissionsPermissionsArgs{
// 				Configure: pulumi.String(".*"),
// 				Read:      pulumi.String(".*"),
// 				Write:     pulumi.String(".*"),
// 			},
// 			User:  pulumi.String("guest"),
// 			Vhost: testVHost.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testExchange, err := rabbitmq.NewExchange(ctx, "testExchange", &rabbitmq.ExchangeArgs{
// 			Settings: &ExchangeSettingsArgs{
// 				AutoDelete: pulumi.Bool(true),
// 				Durable:    pulumi.Bool(false),
// 				Type:       pulumi.String("fanout"),
// 			},
// 			Vhost: guest.Vhost,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testQueue, err := rabbitmq.NewQueue(ctx, "testQueue", &rabbitmq.QueueArgs{
// 			Settings: &QueueSettingsArgs{
// 				AutoDelete: pulumi.Bool(false),
// 				Durable:    pulumi.Bool(true),
// 			},
// 			Vhost: guest.Vhost,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rabbitmq.NewBinding(ctx, "testBinding", &rabbitmq.BindingArgs{
// 			Destination:     testQueue.Name,
// 			DestinationType: pulumi.String("queue"),
// 			RoutingKey:      pulumi.String("#"),
// 			Source:          testExchange.Name,
// 			Vhost:           testVHost.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Bindings can be imported using the `id` which is composed of
//
//  `vhost/source/destination/destination_type/properties_key`. E.g.
//
// ```sh
//  $ pulumi import rabbitmq:index/binding:Binding test test/test/test/queue/%23
// ```
type Binding struct {
	pulumi.CustomResourceState

	// Additional key/value arguments for the binding.
	Arguments     pulumi.MapOutput       `pulumi:"arguments"`
	ArgumentsJson pulumi.StringPtrOutput `pulumi:"argumentsJson"`
	// The destination queue or exchange.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// The type of destination (queue or exchange).
	DestinationType pulumi.StringOutput `pulumi:"destinationType"`
	// A unique key to refer to the binding.
	PropertiesKey pulumi.StringOutput `pulumi:"propertiesKey"`
	// A routing key for the binding.
	RoutingKey pulumi.StringPtrOutput `pulumi:"routingKey"`
	// The source exchange.
	Source pulumi.StringOutput `pulumi:"source"`
	// The vhost to create the resource in.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
}

// NewBinding registers a new resource with the given unique name, arguments, and options.
func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOption) (*Binding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.DestinationType == nil {
		return nil, errors.New("invalid value for required argument 'DestinationType'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.Vhost == nil {
		return nil, errors.New("invalid value for required argument 'Vhost'")
	}
	var resource Binding
	err := ctx.RegisterResource("rabbitmq:index/binding:Binding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinding gets an existing Binding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingState, opts ...pulumi.ResourceOption) (*Binding, error) {
	var resource Binding
	err := ctx.ReadResource("rabbitmq:index/binding:Binding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Binding resources.
type bindingState struct {
	// Additional key/value arguments for the binding.
	Arguments     map[string]interface{} `pulumi:"arguments"`
	ArgumentsJson *string                `pulumi:"argumentsJson"`
	// The destination queue or exchange.
	Destination *string `pulumi:"destination"`
	// The type of destination (queue or exchange).
	DestinationType *string `pulumi:"destinationType"`
	// A unique key to refer to the binding.
	PropertiesKey *string `pulumi:"propertiesKey"`
	// A routing key for the binding.
	RoutingKey *string `pulumi:"routingKey"`
	// The source exchange.
	Source *string `pulumi:"source"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type BindingState struct {
	// Additional key/value arguments for the binding.
	Arguments     pulumi.MapInput
	ArgumentsJson pulumi.StringPtrInput
	// The destination queue or exchange.
	Destination pulumi.StringPtrInput
	// The type of destination (queue or exchange).
	DestinationType pulumi.StringPtrInput
	// A unique key to refer to the binding.
	PropertiesKey pulumi.StringPtrInput
	// A routing key for the binding.
	RoutingKey pulumi.StringPtrInput
	// The source exchange.
	Source pulumi.StringPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (BindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingState)(nil)).Elem()
}

type bindingArgs struct {
	// Additional key/value arguments for the binding.
	Arguments     map[string]interface{} `pulumi:"arguments"`
	ArgumentsJson *string                `pulumi:"argumentsJson"`
	// The destination queue or exchange.
	Destination string `pulumi:"destination"`
	// The type of destination (queue or exchange).
	DestinationType string `pulumi:"destinationType"`
	// A routing key for the binding.
	RoutingKey *string `pulumi:"routingKey"`
	// The source exchange.
	Source string `pulumi:"source"`
	// The vhost to create the resource in.
	Vhost string `pulumi:"vhost"`
}

// The set of arguments for constructing a Binding resource.
type BindingArgs struct {
	// Additional key/value arguments for the binding.
	Arguments     pulumi.MapInput
	ArgumentsJson pulumi.StringPtrInput
	// The destination queue or exchange.
	Destination pulumi.StringInput
	// The type of destination (queue or exchange).
	DestinationType pulumi.StringInput
	// A routing key for the binding.
	RoutingKey pulumi.StringPtrInput
	// The source exchange.
	Source pulumi.StringInput
	// The vhost to create the resource in.
	Vhost pulumi.StringInput
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingArgs)(nil)).Elem()
}

type BindingInput interface {
	pulumi.Input

	ToBindingOutput() BindingOutput
	ToBindingOutputWithContext(ctx context.Context) BindingOutput
}

func (*Binding) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (i *Binding) ToBindingOutput() BindingOutput {
	return i.ToBindingOutputWithContext(context.Background())
}

func (i *Binding) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingOutput)
}

// BindingArrayInput is an input type that accepts BindingArray and BindingArrayOutput values.
// You can construct a concrete instance of `BindingArrayInput` via:
//
//          BindingArray{ BindingArgs{...} }
type BindingArrayInput interface {
	pulumi.Input

	ToBindingArrayOutput() BindingArrayOutput
	ToBindingArrayOutputWithContext(context.Context) BindingArrayOutput
}

type BindingArray []BindingInput

func (BindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (i BindingArray) ToBindingArrayOutput() BindingArrayOutput {
	return i.ToBindingArrayOutputWithContext(context.Background())
}

func (i BindingArray) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingArrayOutput)
}

// BindingMapInput is an input type that accepts BindingMap and BindingMapOutput values.
// You can construct a concrete instance of `BindingMapInput` via:
//
//          BindingMap{ "key": BindingArgs{...} }
type BindingMapInput interface {
	pulumi.Input

	ToBindingMapOutput() BindingMapOutput
	ToBindingMapOutputWithContext(context.Context) BindingMapOutput
}

type BindingMap map[string]BindingInput

func (BindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (i BindingMap) ToBindingMapOutput() BindingMapOutput {
	return i.ToBindingMapOutputWithContext(context.Background())
}

func (i BindingMap) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingMapOutput)
}

type BindingOutput struct{ *pulumi.OutputState }

func (BindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (o BindingOutput) ToBindingOutput() BindingOutput {
	return o
}

func (o BindingOutput) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return o
}

type BindingArrayOutput struct{ *pulumi.OutputState }

func (BindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (o BindingArrayOutput) ToBindingArrayOutput() BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) Index(i pulumi.IntInput) BindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].([]*Binding)[vs[1].(int)]
	}).(BindingOutput)
}

type BindingMapOutput struct{ *pulumi.OutputState }

func (BindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (o BindingMapOutput) ToBindingMapOutput() BindingMapOutput {
	return o
}

func (o BindingMapOutput) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return o
}

func (o BindingMapOutput) MapIndex(k pulumi.StringInput) BindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].(map[string]*Binding)[vs[1].(string)]
	}).(BindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindingInput)(nil)).Elem(), &Binding{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingArrayInput)(nil)).Elem(), BindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingMapInput)(nil)).Elem(), BindingMap{})
	pulumi.RegisterOutputType(BindingOutput{})
	pulumi.RegisterOutputType(BindingArrayOutput{})
	pulumi.RegisterOutputType(BindingMapOutput{})
}
