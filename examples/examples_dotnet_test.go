// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccUserCsharp(t *testing.T) {
	test := getCsharpBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "user", "csharp"),
		})

	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions() integration.ProgramTestOptions {
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.RabbitMQ",
		},
	})

	return baseCsharp
}
