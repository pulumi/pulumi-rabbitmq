// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "rabbitmq:index/binding:Binding":
		r = &Binding{}
	case "rabbitmq:index/exchange:Exchange":
		r = &Exchange{}
	case "rabbitmq:index/federationUpstream:FederationUpstream":
		r = &FederationUpstream{}
	case "rabbitmq:index/operatorPolicy:OperatorPolicy":
		r = &OperatorPolicy{}
	case "rabbitmq:index/permissions:Permissions":
		r = &Permissions{}
	case "rabbitmq:index/policy:Policy":
		r = &Policy{}
	case "rabbitmq:index/queue:Queue":
		r = &Queue{}
	case "rabbitmq:index/shovel:Shovel":
		r = &Shovel{}
	case "rabbitmq:index/topicPermissions:TopicPermissions":
		r = &TopicPermissions{}
	case "rabbitmq:index/user:User":
		r = &User{}
	case "rabbitmq:index/vHost:VHost":
		r = &VHost{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:rabbitmq" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/binding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/exchange",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/federationUpstream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/operatorPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/permissions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/queue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/shovel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/topicPermissions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"rabbitmq",
		"index/vHost",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"rabbitmq",
		&pkg{version},
	)
}
