module github.com/pulumi/pulumi-rabbitmq/provider/v3

go 1.16

replace github.com/terraform-providers/terraform-provider-rabbitmq => github.com/cyrilgdn/terraform-provider-rabbitmq v1.6.0

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.25.1
	github.com/pulumi/pulumi/sdk/v3 v3.35.0
	github.com/terraform-providers/terraform-provider-rabbitmq v1.4.0
)
