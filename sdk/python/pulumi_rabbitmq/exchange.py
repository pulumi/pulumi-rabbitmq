# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Exchange']


class Exchange(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[pulumi.InputType['ExchangeSettingsArgs']]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The ``Exchange`` resource creates and manages an exchange.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test_v_host = rabbitmq.VHost("testVHost")
        guest = rabbitmq.Permissions("guest",
            permissions=rabbitmq.PermissionsPermissionsArgs(
                configure=".*",
                read=".*",
                write=".*",
            ),
            user="guest",
            vhost=test_v_host.name)
        test_exchange = rabbitmq.Exchange("testExchange",
            settings=rabbitmq.ExchangeSettingsArgs(
                auto_delete=True,
                durable=False,
                type="fanout",
            ),
            vhost=guest.vhost)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the exchange.
        :param pulumi.Input[pulumi.InputType['ExchangeSettingsArgs']] settings: The settings of the exchange. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            if settings is None:
                raise TypeError("Missing required property 'settings'")
            __props__['settings'] = settings
            __props__['vhost'] = vhost
        super(Exchange, __self__).__init__(
            'rabbitmq:index/exchange:Exchange',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            settings: Optional[pulumi.Input[pulumi.InputType['ExchangeSettingsArgs']]] = None,
            vhost: Optional[pulumi.Input[str]] = None) -> 'Exchange':
        """
        Get an existing Exchange resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the exchange.
        :param pulumi.Input[pulumi.InputType['ExchangeSettingsArgs']] settings: The settings of the exchange. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["settings"] = settings
        __props__["vhost"] = vhost
        return Exchange(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the exchange.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output['outputs.ExchangeSettings']:
        """
        The settings of the exchange. The structure is
        described below.
        """
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Output[Optional[str]]:
        """
        The vhost to create the resource in.
        """
        return pulumi.get(self, "vhost")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

