# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, cacert_file=None, clientcert_file=None, clientkey_file=None, endpoint=None, insecure=None, password=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the rabbitmq package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if cacert_file is None:
                cacert_file = utilities.get_env('RABBITMQ_CACERT')
            __props__['cacert_file'] = cacert_file
            __props__['clientcert_file'] = clientcert_file
            __props__['clientkey_file'] = clientkey_file
            if endpoint is None:
                endpoint = utilities.get_env('RABBITMQ_ENDPOINT')
            __props__['endpoint'] = endpoint
            if insecure is None:
                insecure = utilities.get_env_bool('RABBITMQ_INSECURE')
            __props__['insecure'] = pulumi.Output.from_input(insecure).apply(json.dumps) if insecure is not None else None
            if password is None:
                password = utilities.get_env('RABBITMQ_PASSWORD')
            __props__['password'] = password
            if username is None:
                username = utilities.get_env('RABBITMQ_USERNAME')
            __props__['username'] = username
        super(Provider, __self__).__init__(
            'rabbitmq',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
