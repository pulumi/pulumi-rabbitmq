[![Actions Status](https://github.com/pulumi/pulumi-rabbitmq/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-rabbitmq/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Frabbitmq.svg)](https://www.npmjs.com/package/@pulumi/rabbitmq)
[![Python version](https://badge.fury.io/py/pulumi-rabbitmq.svg)](https://pypi.org/project/pulumi-rabbitmq)
[![NuGet version](https://badge.fury.io/nu/pulumi.rabbitmq.svg)](https://badge.fury.io/nu/pulumi.rabbitmq)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-rabbitmq/sdk/v3/go)](https://pkg.go.dev/github.com/pulumi/pulumi-rabbitmq/sdk/v3/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-rabbitmq/blob/master/LICENSE)

# RabbitMQ Resource Provider

The RabbitMQ resource provider for Pulumi lets you manage RabbitMQ resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/rabbitmq

or `yarn`:

    $ yarn add @pulumi/rabbitmq

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_rabbitmq

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-rabbitmq/sdk/v2

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Rabbitmq

## Configuration

The following configuration points are available:

* `rabbitmq:endpoint` - (Required) The HTTP URL of the management plugin on the RabbitMQ server. The RabbitMQ management 
   plugin must be enabled in order to use this provder. Note: This is not the IP address or hostname of the RabbitMQ server 
   that you would use to access RabbitMQ directly. May be set via the `RABBITMQ_ENDPOINT` environment variable.
* `rabbitmq:username` - (Required) Username to use to authenticate with the server. May be set via the `RABBITMQ_USERNAME`
   environment variable.
* `rabbitmq:password` - (Optional) Password for the given user. May be set via the `RABBITMQ_PASSWORD` environment variable.
* `rabbitmq:insecure` - (Optional) Trust self-signed certificates. May be set via the `RABBITMQ_INSECURE` environment variable.
* `rabbitmq:cacertFile` - (Optional) The path to a custom CA / intermediate certificate. May be set via the `RABBITMQ_CACERT` 
  environment variable.


## Reference

For further information, please visit [the RabbitMQ provider docs](https://www.pulumi.com/docs/intro/cloud-providers/rabbitmq) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/rabbitmq).
