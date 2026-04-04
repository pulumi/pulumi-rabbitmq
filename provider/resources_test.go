package rabbitmq

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"

	"github.com/pulumi/pulumi-rabbitmq/provider/v3/pkg/version"
)

func TestProviderPreservesExistingTokens(t *testing.T) {
	prov := testProvider(t)

	expectedResources := map[string]string{
		"rabbitmq_binding":             "rabbitmq:index/binding:Binding",
		"rabbitmq_exchange":            "rabbitmq:index/exchange:Exchange",
		"rabbitmq_permissions":         "rabbitmq:index/permissions:Permissions",
		"rabbitmq_policy":              "rabbitmq:index/policy:Policy",
		"rabbitmq_queue":               "rabbitmq:index/queue:Queue",
		"rabbitmq_user":                "rabbitmq:index/user:User",
		"rabbitmq_vhost":               "rabbitmq:index/vHost:VHost",
		"rabbitmq_topic_permissions":   "rabbitmq:index/topicPermissions:TopicPermissions",
		"rabbitmq_federation_upstream": "rabbitmq:index/federationUpstream:FederationUpstream",
		"rabbitmq_shovel":              "rabbitmq:index/shovel:Shovel",
		"rabbitmq_operator_policy":     "rabbitmq:index/operatorPolicy:OperatorPolicy",
	}

	expectedDataSources := map[string]string{
		"rabbitmq_exchange":     "rabbitmq:index/getExchange:getExchange",
		"rabbitmq_user":         "rabbitmq:index/getUser:getUser",
		"rabbitmq_default_user": "rabbitmq:index/getDefaultUser:getDefaultUser",
		"rabbitmq_vhost":        "rabbitmq:index/getVHost:getVHost",
	}

	for tfToken, expected := range expectedResources {
		info, ok := prov.Resources[tfToken]
		require.Truef(t, ok, "missing resource mapping for %s", tfToken)
		require.Equal(t, expected, string(info.Tok))
	}

	for tfToken, expected := range expectedDataSources {
		info, ok := prov.DataSources[tfToken]
		require.Truef(t, ok, "missing data source mapping for %s", tfToken)
		require.Equal(t, expected, string(info.Tok))
	}
}

func TestProviderComputesTokensForAllTerraformElements(t *testing.T) {
	prov := testProvider(t)

	prov.P.ResourcesMap().Range(func(tfToken string, _ shim.Resource) bool {
		info, ok := prov.Resources[tfToken]
		require.Truef(t, ok, "missing resource mapping for %s", tfToken)
		require.NotEmptyf(t, info.Tok, "missing resource token for %s", tfToken)
		return true
	})

	prov.P.DataSourcesMap().Range(func(tfToken string, _ shim.Resource) bool {
		info, ok := prov.DataSources[tfToken]
		require.Truef(t, ok, "missing data source mapping for %s", tfToken)
		require.NotEmptyf(t, info.Tok, "missing data source token for %s", tfToken)
		return true
	})
}

func testProvider(t *testing.T) tfbridge.ProviderInfo {
	t.Helper()

	oldVersion := version.Version
	version.Version = "3.0.0"
	t.Cleanup(func() {
		version.Version = oldVersion
	})

	return Provider()
}
