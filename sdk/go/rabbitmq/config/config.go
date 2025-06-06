// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

func GetCacertFile(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "rabbitmq:cacertFile")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "RABBITMQ_CACERT"); d != nil {
		value = d.(string)
	}
	return value
}
func GetClientcertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:clientcertFile")
}
func GetClientkeyFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:clientkeyFile")
}
func GetEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:endpoint")
}
func GetInsecure(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "rabbitmq:insecure")
	if err == nil {
		return v
	}
	var value bool
	if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "RABBITMQ_INSECURE"); d != nil {
		value = d.(bool)
	}
	return value
}
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:password")
}
func GetProxy(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:proxy")
}
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:username")
}
