// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rabbitmq

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the rabbitmq package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/index.html.markdown.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.CacertFile == nil {
		args.CacertFile = pulumi.StringPtr(getEnvOrDefault("", nil, "RABBITMQ_CACERT").(string))
	}
	if args.Endpoint == nil {
		args.Endpoint = pulumi.StringPtr(getEnvOrDefault("", nil, "RABBITMQ_ENDPOINT").(string))
	}
	if args.Insecure == nil {
		args.Insecure = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "RABBITMQ_INSECURE").(bool))
	}
	if args.Password == nil {
		args.Password = pulumi.StringPtr(getEnvOrDefault("", nil, "RABBITMQ_PASSWORD").(string))
	}
	if args.Username == nil {
		args.Username = pulumi.StringPtr(getEnvOrDefault("", nil, "RABBITMQ_USERNAME").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:rabbitmq", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	CacertFile *string `pulumi:"cacertFile"`
	Endpoint   *string `pulumi:"endpoint"`
	Insecure   *bool   `pulumi:"insecure"`
	Password   *string `pulumi:"password"`
	Username   *string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	CacertFile pulumi.StringPtrInput
	Endpoint   pulumi.StringPtrInput
	Insecure   pulumi.BoolPtrInput
	Password   pulumi.StringPtrInput
	Username   pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}
