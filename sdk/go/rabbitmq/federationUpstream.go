// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FederationUpstream struct {
	pulumi.CustomResourceState

	Component  pulumi.StringOutput                `pulumi:"component"`
	Definition FederationUpstreamDefinitionOutput `pulumi:"definition"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Vhost      pulumi.StringOutput                `pulumi:"vhost"`
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
	Component  *string                       `pulumi:"component"`
	Definition *FederationUpstreamDefinition `pulumi:"definition"`
	Name       *string                       `pulumi:"name"`
	Vhost      *string                       `pulumi:"vhost"`
}

type FederationUpstreamState struct {
	Component  pulumi.StringPtrInput
	Definition FederationUpstreamDefinitionPtrInput
	Name       pulumi.StringPtrInput
	Vhost      pulumi.StringPtrInput
}

func (FederationUpstreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*federationUpstreamState)(nil)).Elem()
}

type federationUpstreamArgs struct {
	Definition FederationUpstreamDefinition `pulumi:"definition"`
	Name       *string                      `pulumi:"name"`
	Vhost      string                       `pulumi:"vhost"`
}

// The set of arguments for constructing a FederationUpstream resource.
type FederationUpstreamArgs struct {
	Definition FederationUpstreamDefinitionInput
	Name       pulumi.StringPtrInput
	Vhost      pulumi.StringInput
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
	return reflect.TypeOf((*FederationUpstream)(nil))
}

func (i *FederationUpstream) ToFederationUpstreamOutput() FederationUpstreamOutput {
	return i.ToFederationUpstreamOutputWithContext(context.Background())
}

func (i *FederationUpstream) ToFederationUpstreamOutputWithContext(ctx context.Context) FederationUpstreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederationUpstreamOutput)
}

func (i *FederationUpstream) ToFederationUpstreamPtrOutput() FederationUpstreamPtrOutput {
	return i.ToFederationUpstreamPtrOutputWithContext(context.Background())
}

func (i *FederationUpstream) ToFederationUpstreamPtrOutputWithContext(ctx context.Context) FederationUpstreamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederationUpstreamPtrOutput)
}

type FederationUpstreamPtrInput interface {
	pulumi.Input

	ToFederationUpstreamPtrOutput() FederationUpstreamPtrOutput
	ToFederationUpstreamPtrOutputWithContext(ctx context.Context) FederationUpstreamPtrOutput
}

type federationUpstreamPtrType FederationUpstreamArgs

func (*federationUpstreamPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FederationUpstream)(nil))
}

func (i *federationUpstreamPtrType) ToFederationUpstreamPtrOutput() FederationUpstreamPtrOutput {
	return i.ToFederationUpstreamPtrOutputWithContext(context.Background())
}

func (i *federationUpstreamPtrType) ToFederationUpstreamPtrOutputWithContext(ctx context.Context) FederationUpstreamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederationUpstreamPtrOutput)
}

// FederationUpstreamArrayInput is an input type that accepts FederationUpstreamArray and FederationUpstreamArrayOutput values.
// You can construct a concrete instance of `FederationUpstreamArrayInput` via:
//
//          FederationUpstreamArray{ FederationUpstreamArgs{...} }
type FederationUpstreamArrayInput interface {
	pulumi.Input

	ToFederationUpstreamArrayOutput() FederationUpstreamArrayOutput
	ToFederationUpstreamArrayOutputWithContext(context.Context) FederationUpstreamArrayOutput
}

type FederationUpstreamArray []FederationUpstreamInput

func (FederationUpstreamArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FederationUpstream)(nil))
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
//          FederationUpstreamMap{ "key": FederationUpstreamArgs{...} }
type FederationUpstreamMapInput interface {
	pulumi.Input

	ToFederationUpstreamMapOutput() FederationUpstreamMapOutput
	ToFederationUpstreamMapOutputWithContext(context.Context) FederationUpstreamMapOutput
}

type FederationUpstreamMap map[string]FederationUpstreamInput

func (FederationUpstreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FederationUpstream)(nil))
}

func (i FederationUpstreamMap) ToFederationUpstreamMapOutput() FederationUpstreamMapOutput {
	return i.ToFederationUpstreamMapOutputWithContext(context.Background())
}

func (i FederationUpstreamMap) ToFederationUpstreamMapOutputWithContext(ctx context.Context) FederationUpstreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederationUpstreamMapOutput)
}

type FederationUpstreamOutput struct {
	*pulumi.OutputState
}

func (FederationUpstreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FederationUpstream)(nil))
}

func (o FederationUpstreamOutput) ToFederationUpstreamOutput() FederationUpstreamOutput {
	return o
}

func (o FederationUpstreamOutput) ToFederationUpstreamOutputWithContext(ctx context.Context) FederationUpstreamOutput {
	return o
}

func (o FederationUpstreamOutput) ToFederationUpstreamPtrOutput() FederationUpstreamPtrOutput {
	return o.ToFederationUpstreamPtrOutputWithContext(context.Background())
}

func (o FederationUpstreamOutput) ToFederationUpstreamPtrOutputWithContext(ctx context.Context) FederationUpstreamPtrOutput {
	return o.ApplyT(func(v FederationUpstream) *FederationUpstream {
		return &v
	}).(FederationUpstreamPtrOutput)
}

type FederationUpstreamPtrOutput struct {
	*pulumi.OutputState
}

func (FederationUpstreamPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederationUpstream)(nil))
}

func (o FederationUpstreamPtrOutput) ToFederationUpstreamPtrOutput() FederationUpstreamPtrOutput {
	return o
}

func (o FederationUpstreamPtrOutput) ToFederationUpstreamPtrOutputWithContext(ctx context.Context) FederationUpstreamPtrOutput {
	return o
}

type FederationUpstreamArrayOutput struct{ *pulumi.OutputState }

func (FederationUpstreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FederationUpstream)(nil))
}

func (o FederationUpstreamArrayOutput) ToFederationUpstreamArrayOutput() FederationUpstreamArrayOutput {
	return o
}

func (o FederationUpstreamArrayOutput) ToFederationUpstreamArrayOutputWithContext(ctx context.Context) FederationUpstreamArrayOutput {
	return o
}

func (o FederationUpstreamArrayOutput) Index(i pulumi.IntInput) FederationUpstreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FederationUpstream {
		return vs[0].([]FederationUpstream)[vs[1].(int)]
	}).(FederationUpstreamOutput)
}

type FederationUpstreamMapOutput struct{ *pulumi.OutputState }

func (FederationUpstreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FederationUpstream)(nil))
}

func (o FederationUpstreamMapOutput) ToFederationUpstreamMapOutput() FederationUpstreamMapOutput {
	return o
}

func (o FederationUpstreamMapOutput) ToFederationUpstreamMapOutputWithContext(ctx context.Context) FederationUpstreamMapOutput {
	return o
}

func (o FederationUpstreamMapOutput) MapIndex(k pulumi.StringInput) FederationUpstreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FederationUpstream {
		return vs[0].(map[string]FederationUpstream)[vs[1].(string)]
	}).(FederationUpstreamOutput)
}

func init() {
	pulumi.RegisterOutputType(FederationUpstreamOutput{})
	pulumi.RegisterOutputType(FederationUpstreamPtrOutput{})
	pulumi.RegisterOutputType(FederationUpstreamArrayOutput{})
	pulumi.RegisterOutputType(FederationUpstreamMapOutput{})
}
