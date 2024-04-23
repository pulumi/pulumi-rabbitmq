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

// The “FederationUpstream“ resource creates and manages a federation upstream parameter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
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
//			// downstream exchange
//			foo, err := rabbitmq.NewExchange(ctx, "foo", &rabbitmq.ExchangeArgs{
//				Name:  pulumi.String("foo"),
//				Vhost: guest.Vhost,
//				Settings: &rabbitmq.ExchangeSettingsArgs{
//					Type:    pulumi.String("topic"),
//					Durable: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// upstream broker
//			fooFederationUpstream, err := rabbitmq.NewFederationUpstream(ctx, "foo", &rabbitmq.FederationUpstreamArgs{
//				Name:  pulumi.String("foo"),
//				Vhost: guest.Vhost,
//				Definition: &rabbitmq.FederationUpstreamDefinitionArgs{
//					Uri:            pulumi.String("amqp://guest:guest@upstream-server-name:5672/%2f"),
//					PrefetchCount:  pulumi.Int(1000),
//					ReconnectDelay: pulumi.Int(5),
//					AckMode:        pulumi.String("on-confirm"),
//					TrustUserId:    pulumi.Bool(false),
//					MaxHops:        pulumi.Int(1),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rabbitmq.NewPolicy(ctx, "foo", &rabbitmq.PolicyArgs{
//				Name:  pulumi.String("foo"),
//				Vhost: guest.Vhost,
//				Policy: &rabbitmq.PolicyPolicyArgs{
//					Pattern: foo.Name.ApplyT(func(name string) (string, error) {
//						return fmt.Sprintf("(^%v$)", name), nil
//					}).(pulumi.StringOutput),
//					Priority: pulumi.Int(1),
//					ApplyTo:  pulumi.String("exchanges"),
//					Definition: pulumi.Map{
//						"federation-upstream": fooFederationUpstream.Name,
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
// A Federation upstream can be imported using the resource `id` which is composed of `name@vhost`, e.g.
//
// ```sh
// $ pulumi import rabbitmq:index/federationUpstream:FederationUpstream foo foo@test
// ```
type FederationUpstream struct {
	pulumi.CustomResourceState

	// Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
	Component pulumi.StringOutput `pulumi:"component"`
	// The configuration of the federation upstream. The structure is described below.
	Definition FederationUpstreamDefinitionOutput `pulumi:"definition"`
	// The name of the federation upstream.
	Name pulumi.StringOutput `pulumi:"name"`
	// The vhost to create the resource in.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
}

// NewFederationUpstream registers a new resource with the given unique name, arguments, and options.
func NewFederationUpstream(ctx *pulumi.Context,
	name string, args *FederationUpstreamArgs, opts ...pulumi.ResourceOption) (*FederationUpstream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.Vhost == nil {
		return nil, errors.New("invalid value for required argument 'Vhost'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FederationUpstream
	err := ctx.RegisterResource("rabbitmq:index/federationUpstream:FederationUpstream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederationUpstream gets an existing FederationUpstream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederationUpstream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederationUpstreamState, opts ...pulumi.ResourceOption) (*FederationUpstream, error) {
	var resource FederationUpstream
	err := ctx.ReadResource("rabbitmq:index/federationUpstream:FederationUpstream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederationUpstream resources.
type federationUpstreamState struct {
	// Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
	Component *string `pulumi:"component"`
	// The configuration of the federation upstream. The structure is described below.
	Definition *FederationUpstreamDefinition `pulumi:"definition"`
	// The name of the federation upstream.
	Name *string `pulumi:"name"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type FederationUpstreamState struct {
	// Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
	Component pulumi.StringPtrInput
	// The configuration of the federation upstream. The structure is described below.
	Definition FederationUpstreamDefinitionPtrInput
	// The name of the federation upstream.
	Name pulumi.StringPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (FederationUpstreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*federationUpstreamState)(nil)).Elem()
}

type federationUpstreamArgs struct {
	// The configuration of the federation upstream. The structure is described below.
	Definition FederationUpstreamDefinition `pulumi:"definition"`
	// The name of the federation upstream.
	Name *string `pulumi:"name"`
	// The vhost to create the resource in.
	Vhost string `pulumi:"vhost"`
}

// The set of arguments for constructing a FederationUpstream resource.
type FederationUpstreamArgs struct {
	// The configuration of the federation upstream. The structure is described below.
	Definition FederationUpstreamDefinitionInput
	// The name of the federation upstream.
	Name pulumi.StringPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringInput
}

func (FederationUpstreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federationUpstreamArgs)(nil)).Elem()
}

type FederationUpstreamInput interface {
	pulumi.Input

	ToFederationUpstreamOutput() FederationUpstreamOutput
	ToFederationUpstreamOutputWithContext(ctx context.Context) FederationUpstreamOutput
}

func (*FederationUpstream) ElementType() reflect.Type {
	return reflect.TypeOf((**FederationUpstream)(nil)).Elem()
}

func (i *FederationUpstream) ToFederationUpstreamOutput() FederationUpstreamOutput {
	return i.ToFederationUpstreamOutputWithContext(context.Background())
}

func (i *FederationUpstream) ToFederationUpstreamOutputWithContext(ctx context.Context) FederationUpstreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederationUpstreamOutput)
}

// FederationUpstreamArrayInput is an input type that accepts FederationUpstreamArray and FederationUpstreamArrayOutput values.
// You can construct a concrete instance of `FederationUpstreamArrayInput` via:
//
//	FederationUpstreamArray{ FederationUpstreamArgs{...} }
type FederationUpstreamArrayInput interface {
	pulumi.Input

	ToFederationUpstreamArrayOutput() FederationUpstreamArrayOutput
	ToFederationUpstreamArrayOutputWithContext(context.Context) FederationUpstreamArrayOutput
}

type FederationUpstreamArray []FederationUpstreamInput

func (FederationUpstreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederationUpstream)(nil)).Elem()
}

func (i FederationUpstreamArray) ToFederationUpstreamArrayOutput() FederationUpstreamArrayOutput {
	return i.ToFederationUpstreamArrayOutputWithContext(context.Background())
}

func (i FederationUpstreamArray) ToFederationUpstreamArrayOutputWithContext(ctx context.Context) FederationUpstreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederationUpstreamArrayOutput)
}

// FederationUpstreamMapInput is an input type that accepts FederationUpstreamMap and FederationUpstreamMapOutput values.
// You can construct a concrete instance of `FederationUpstreamMapInput` via:
//
//	FederationUpstreamMap{ "key": FederationUpstreamArgs{...} }
type FederationUpstreamMapInput interface {
	pulumi.Input

	ToFederationUpstreamMapOutput() FederationUpstreamMapOutput
	ToFederationUpstreamMapOutputWithContext(context.Context) FederationUpstreamMapOutput
}

type FederationUpstreamMap map[string]FederationUpstreamInput

func (FederationUpstreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederationUpstream)(nil)).Elem()
}

func (i FederationUpstreamMap) ToFederationUpstreamMapOutput() FederationUpstreamMapOutput {
	return i.ToFederationUpstreamMapOutputWithContext(context.Background())
}

func (i FederationUpstreamMap) ToFederationUpstreamMapOutputWithContext(ctx context.Context) FederationUpstreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederationUpstreamMapOutput)
}

type FederationUpstreamOutput struct{ *pulumi.OutputState }

func (FederationUpstreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederationUpstream)(nil)).Elem()
}

func (o FederationUpstreamOutput) ToFederationUpstreamOutput() FederationUpstreamOutput {
	return o
}

func (o FederationUpstreamOutput) ToFederationUpstreamOutputWithContext(ctx context.Context) FederationUpstreamOutput {
	return o
}

// Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
func (o FederationUpstreamOutput) Component() pulumi.StringOutput {
	return o.ApplyT(func(v *FederationUpstream) pulumi.StringOutput { return v.Component }).(pulumi.StringOutput)
}

// The configuration of the federation upstream. The structure is described below.
func (o FederationUpstreamOutput) Definition() FederationUpstreamDefinitionOutput {
	return o.ApplyT(func(v *FederationUpstream) FederationUpstreamDefinitionOutput { return v.Definition }).(FederationUpstreamDefinitionOutput)
}

// The name of the federation upstream.
func (o FederationUpstreamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FederationUpstream) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The vhost to create the resource in.
func (o FederationUpstreamOutput) Vhost() pulumi.StringOutput {
	return o.ApplyT(func(v *FederationUpstream) pulumi.StringOutput { return v.Vhost }).(pulumi.StringOutput)
}

type FederationUpstreamArrayOutput struct{ *pulumi.OutputState }

func (FederationUpstreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederationUpstream)(nil)).Elem()
}

func (o FederationUpstreamArrayOutput) ToFederationUpstreamArrayOutput() FederationUpstreamArrayOutput {
	return o
}

func (o FederationUpstreamArrayOutput) ToFederationUpstreamArrayOutputWithContext(ctx context.Context) FederationUpstreamArrayOutput {
	return o
}

func (o FederationUpstreamArrayOutput) Index(i pulumi.IntInput) FederationUpstreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederationUpstream {
		return vs[0].([]*FederationUpstream)[vs[1].(int)]
	}).(FederationUpstreamOutput)
}

type FederationUpstreamMapOutput struct{ *pulumi.OutputState }

func (FederationUpstreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederationUpstream)(nil)).Elem()
}

func (o FederationUpstreamMapOutput) ToFederationUpstreamMapOutput() FederationUpstreamMapOutput {
	return o
}

func (o FederationUpstreamMapOutput) ToFederationUpstreamMapOutputWithContext(ctx context.Context) FederationUpstreamMapOutput {
	return o
}

func (o FederationUpstreamMapOutput) MapIndex(k pulumi.StringInput) FederationUpstreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederationUpstream {
		return vs[0].(map[string]*FederationUpstream)[vs[1].(string)]
	}).(FederationUpstreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederationUpstreamInput)(nil)).Elem(), &FederationUpstream{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederationUpstreamArrayInput)(nil)).Elem(), FederationUpstreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederationUpstreamMapInput)(nil)).Elem(), FederationUpstreamMap{})
	pulumi.RegisterOutputType(FederationUpstreamOutput{})
	pulumi.RegisterOutputType(FederationUpstreamArrayOutput{})
	pulumi.RegisterOutputType(FederationUpstreamMapOutput{})
}
