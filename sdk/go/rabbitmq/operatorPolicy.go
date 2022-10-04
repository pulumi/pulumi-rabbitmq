// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “OperatorPolicy“ resource creates and manages operator policies for queues.
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
//				Permissions: &PermissionsPermissionsArgs{
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
//			_, err = rabbitmq.NewOperatorPolicy(ctx, "testOperatorPolicy", &rabbitmq.OperatorPolicyArgs{
//				Policy: &OperatorPolicyPolicyArgs{
//					ApplyTo: pulumi.String("queues"),
//					Definition: pulumi.AnyMap{
//						"expires":     pulumi.Any(1800000),
//						"message-ttl": pulumi.Any(3600000),
//					},
//					Pattern:  pulumi.String(".*"),
//					Priority: pulumi.Int(0),
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
// Operator policies can be imported using the `id` which is composed of `name@vhost`. E.g.
//
// ```sh
//
//	$ pulumi import rabbitmq:index/operatorPolicy:OperatorPolicy test name@vhost
//
// ```
type OperatorPolicy struct {
	pulumi.CustomResourceState

	// The name of the operator policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The settings of the operator policy. The structure is
	// described below.
	Policy OperatorPolicyPolicyOutput `pulumi:"policy"`
	// The vhost to create the resource in.
	Vhost pulumi.StringOutput `pulumi:"vhost"`
}

// NewOperatorPolicy registers a new resource with the given unique name, arguments, and options.
func NewOperatorPolicy(ctx *pulumi.Context,
	name string, args *OperatorPolicyArgs, opts ...pulumi.ResourceOption) (*OperatorPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.Vhost == nil {
		return nil, errors.New("invalid value for required argument 'Vhost'")
	}
	var resource OperatorPolicy
	err := ctx.RegisterResource("rabbitmq:index/operatorPolicy:OperatorPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOperatorPolicy gets an existing OperatorPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOperatorPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OperatorPolicyState, opts ...pulumi.ResourceOption) (*OperatorPolicy, error) {
	var resource OperatorPolicy
	err := ctx.ReadResource("rabbitmq:index/operatorPolicy:OperatorPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OperatorPolicy resources.
type operatorPolicyState struct {
	// The name of the operator policy.
	Name *string `pulumi:"name"`
	// The settings of the operator policy. The structure is
	// described below.
	Policy *OperatorPolicyPolicy `pulumi:"policy"`
	// The vhost to create the resource in.
	Vhost *string `pulumi:"vhost"`
}

type OperatorPolicyState struct {
	// The name of the operator policy.
	Name pulumi.StringPtrInput
	// The settings of the operator policy. The structure is
	// described below.
	Policy OperatorPolicyPolicyPtrInput
	// The vhost to create the resource in.
	Vhost pulumi.StringPtrInput
}

func (OperatorPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*operatorPolicyState)(nil)).Elem()
}

type operatorPolicyArgs struct {
	// The name of the operator policy.
	Name *string `pulumi:"name"`
	// The settings of the operator policy. The structure is
	// described below.
	Policy OperatorPolicyPolicy `pulumi:"policy"`
	// The vhost to create the resource in.
	Vhost string `pulumi:"vhost"`
}

// The set of arguments for constructing a OperatorPolicy resource.
type OperatorPolicyArgs struct {
	// The name of the operator policy.
	Name pulumi.StringPtrInput
	// The settings of the operator policy. The structure is
	// described below.
	Policy OperatorPolicyPolicyInput
	// The vhost to create the resource in.
	Vhost pulumi.StringInput
}

func (OperatorPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*operatorPolicyArgs)(nil)).Elem()
}

type OperatorPolicyInput interface {
	pulumi.Input

	ToOperatorPolicyOutput() OperatorPolicyOutput
	ToOperatorPolicyOutputWithContext(ctx context.Context) OperatorPolicyOutput
}

func (*OperatorPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatorPolicy)(nil)).Elem()
}

func (i *OperatorPolicy) ToOperatorPolicyOutput() OperatorPolicyOutput {
	return i.ToOperatorPolicyOutputWithContext(context.Background())
}

func (i *OperatorPolicy) ToOperatorPolicyOutputWithContext(ctx context.Context) OperatorPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperatorPolicyOutput)
}

// OperatorPolicyArrayInput is an input type that accepts OperatorPolicyArray and OperatorPolicyArrayOutput values.
// You can construct a concrete instance of `OperatorPolicyArrayInput` via:
//
//	OperatorPolicyArray{ OperatorPolicyArgs{...} }
type OperatorPolicyArrayInput interface {
	pulumi.Input

	ToOperatorPolicyArrayOutput() OperatorPolicyArrayOutput
	ToOperatorPolicyArrayOutputWithContext(context.Context) OperatorPolicyArrayOutput
}

type OperatorPolicyArray []OperatorPolicyInput

func (OperatorPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OperatorPolicy)(nil)).Elem()
}

func (i OperatorPolicyArray) ToOperatorPolicyArrayOutput() OperatorPolicyArrayOutput {
	return i.ToOperatorPolicyArrayOutputWithContext(context.Background())
}

func (i OperatorPolicyArray) ToOperatorPolicyArrayOutputWithContext(ctx context.Context) OperatorPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperatorPolicyArrayOutput)
}

// OperatorPolicyMapInput is an input type that accepts OperatorPolicyMap and OperatorPolicyMapOutput values.
// You can construct a concrete instance of `OperatorPolicyMapInput` via:
//
//	OperatorPolicyMap{ "key": OperatorPolicyArgs{...} }
type OperatorPolicyMapInput interface {
	pulumi.Input

	ToOperatorPolicyMapOutput() OperatorPolicyMapOutput
	ToOperatorPolicyMapOutputWithContext(context.Context) OperatorPolicyMapOutput
}

type OperatorPolicyMap map[string]OperatorPolicyInput

func (OperatorPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OperatorPolicy)(nil)).Elem()
}

func (i OperatorPolicyMap) ToOperatorPolicyMapOutput() OperatorPolicyMapOutput {
	return i.ToOperatorPolicyMapOutputWithContext(context.Background())
}

func (i OperatorPolicyMap) ToOperatorPolicyMapOutputWithContext(ctx context.Context) OperatorPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperatorPolicyMapOutput)
}

type OperatorPolicyOutput struct{ *pulumi.OutputState }

func (OperatorPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatorPolicy)(nil)).Elem()
}

func (o OperatorPolicyOutput) ToOperatorPolicyOutput() OperatorPolicyOutput {
	return o
}

func (o OperatorPolicyOutput) ToOperatorPolicyOutputWithContext(ctx context.Context) OperatorPolicyOutput {
	return o
}

// The name of the operator policy.
func (o OperatorPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OperatorPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The settings of the operator policy. The structure is
// described below.
func (o OperatorPolicyOutput) Policy() OperatorPolicyPolicyOutput {
	return o.ApplyT(func(v *OperatorPolicy) OperatorPolicyPolicyOutput { return v.Policy }).(OperatorPolicyPolicyOutput)
}

// The vhost to create the resource in.
func (o OperatorPolicyOutput) Vhost() pulumi.StringOutput {
	return o.ApplyT(func(v *OperatorPolicy) pulumi.StringOutput { return v.Vhost }).(pulumi.StringOutput)
}

type OperatorPolicyArrayOutput struct{ *pulumi.OutputState }

func (OperatorPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OperatorPolicy)(nil)).Elem()
}

func (o OperatorPolicyArrayOutput) ToOperatorPolicyArrayOutput() OperatorPolicyArrayOutput {
	return o
}

func (o OperatorPolicyArrayOutput) ToOperatorPolicyArrayOutputWithContext(ctx context.Context) OperatorPolicyArrayOutput {
	return o
}

func (o OperatorPolicyArrayOutput) Index(i pulumi.IntInput) OperatorPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OperatorPolicy {
		return vs[0].([]*OperatorPolicy)[vs[1].(int)]
	}).(OperatorPolicyOutput)
}

type OperatorPolicyMapOutput struct{ *pulumi.OutputState }

func (OperatorPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OperatorPolicy)(nil)).Elem()
}

func (o OperatorPolicyMapOutput) ToOperatorPolicyMapOutput() OperatorPolicyMapOutput {
	return o
}

func (o OperatorPolicyMapOutput) ToOperatorPolicyMapOutputWithContext(ctx context.Context) OperatorPolicyMapOutput {
	return o
}

func (o OperatorPolicyMapOutput) MapIndex(k pulumi.StringInput) OperatorPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OperatorPolicy {
		return vs[0].(map[string]*OperatorPolicy)[vs[1].(string)]
	}).(OperatorPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorPolicyInput)(nil)).Elem(), &OperatorPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorPolicyArrayInput)(nil)).Elem(), OperatorPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorPolicyMapInput)(nil)).Elem(), OperatorPolicyMap{})
	pulumi.RegisterOutputType(OperatorPolicyOutput{})
	pulumi.RegisterOutputType(OperatorPolicyArrayOutput{})
	pulumi.RegisterOutputType(OperatorPolicyMapOutput{})
}