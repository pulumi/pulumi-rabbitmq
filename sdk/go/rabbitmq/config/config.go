// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GetCacertFile(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "rabbitmq:cacertFile")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "RABBITMQ_CACERT").(string)
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
	return getEnvOrDefault(false, parseEnvBool, "RABBITMQ_INSECURE").(bool)
}
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:password")
}
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "rabbitmq:username")
}
