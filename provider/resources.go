// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rabbitmq

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-rabbitmq/provider/v2/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-rabbitmq/rabbitmq"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "rabbitmq"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

func Provider() tfbridge.ProviderInfo {
	p := shimv1.NewProvider(rabbitmq.Provider().(*schema.Provider))
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "rabbitmq",
		Description: "A Pulumi package for creating and managing RabbitMQ resources.",
		Keywords:    []string{"pulumi", "rabbitmq"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-rabbitmq",
		Config: map[string]*tfbridge.SchemaInfo{
			"insecure": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"RABBITMQ_INSECURE"},
				},
			},
			"cacert_file": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"RABBITMQ_CACERT"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"rabbitmq_binding":  {Tok: makeResource(mainMod, "Binding")},
			"rabbitmq_exchange": {Tok: makeResource(mainMod, "Exchange")},
			"rabbitmq_permissions": {
				Tok: makeResource(mainMod, "Permissions"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"permissions": {
						CSharpName: "PermissionDetails",
					},
				},
			},
			"rabbitmq_policy": {
				Tok: makeResource(mainMod, "Policy"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"policy": {
						CSharpName: "PolicyBlock",
					},
				},
			},
			"rabbitmq_queue": {Tok: makeResource(mainMod, "Queue")},
			"rabbitmq_user":  {Tok: makeResource(mainMod, "User")},
			"rabbitmq_vhost": {Tok: makeResource(mainMod, "VHost")},
			"rabbitmq_topic_permissions": {
				Tok: makeResource(mainMod, "TopicPermissions"),
				Docs: &tfbridge.DocInfo{
					Source: "topic-permissions.html.markdown",
				},
			},
			"rabbitmq_federation_upstream": {Tok: makeResource(mainMod, "FederationUpstream")},
			"rabbitmq_shovel":              {Tok: makeResource(mainMod, "Shovel")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},

		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.15.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.15.0,<3.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "RabbitMQ",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
