// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the rabbitmq package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.CacertFile == nil {
		args.CacertFile = pulumi.StringPtr(getEnvOrDefault("", nil, "RABBITMQ_CACERT").(string))
	}
	if args.Insecure == nil {
		args.Insecure = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "RABBITMQ_INSECURE").(bool))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:rabbitmq", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	CacertFile     *string `pulumi:"cacertFile"`
	ClientcertFile *string `pulumi:"clientcertFile"`
	ClientkeyFile  *string `pulumi:"clientkeyFile"`
	Endpoint       string  `pulumi:"endpoint"`
	Insecure       *bool   `pulumi:"insecure"`
	Password       string  `pulumi:"password"`
	Username       string  `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	CacertFile     pulumi.StringPtrInput
	ClientcertFile pulumi.StringPtrInput
	ClientkeyFile  pulumi.StringPtrInput
	Endpoint       pulumi.StringInput
	Insecure       pulumi.BoolPtrInput
	Password       pulumi.StringInput
	Username       pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput
}

type providerPtrType ProviderArgs

func (*providerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (i *providerPtrType) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *providerPtrType) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderOutput struct {
	*pulumi.OutputState
}

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyT(func(v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

type ProviderPtrOutput struct {
	*pulumi.OutputState
}

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
}
