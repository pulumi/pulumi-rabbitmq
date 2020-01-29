module github.com/pulumi/pulumi-rabbitmq

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.6.4
	github.com/terraform-providers/terraform-provider-rabbitmq v1.2.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
