// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"github.com/pulumi/pulumi/pkg/testing/integration"
	"os"
	"path"
	"testing"
)

func TestAccUser(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "user"),
		})

	integration.ProgramTest(t, &test)
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		Config: map[string]string{
			"rabbitmq:endpoint": "http://127.0.0.1:15672",
			"rabbitmq:username": "guest",
			"rabbitmq:password": "guest",
		},
		Dependencies: []string{
			"@pulumi/rabbitmq",
		},
	})

	return baseJS
}
