module github.com/pulumi/pulumi-rabbitmq/provider

go 1.13

require (
	github.com/pulumi/pulumi v1.13.1
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/terraform-providers/terraform-provider-rabbitmq v1.3.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
