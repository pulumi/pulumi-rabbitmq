# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ExchangeArgs', 'Exchange']

@pulumi.input_type
class ExchangeArgs:
    def __init__(__self__, *,
                 settings: pulumi.Input['ExchangeSettingsArgs'],
                 name: Optional[pulumi.Input[str]] = None,
                 vhost: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Exchange resource.
        :param pulumi.Input['ExchangeSettingsArgs'] settings: The settings of the exchange. The structure is
               described below.
        :param pulumi.Input[str] name: The name of the exchange.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        pulumi.set(__self__, "settings", settings)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vhost is not None:
            pulumi.set(__self__, "vhost", vhost)

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Input['ExchangeSettingsArgs']:
        """
        The settings of the exchange. The structure is
        described below.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: pulumi.Input['ExchangeSettingsArgs']):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the exchange.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vhost(self) -> Optional[pulumi.Input[str]]:
        """
        The vhost to create the resource in.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vhost", value)


@pulumi.input_type
class _ExchangeState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input['ExchangeSettingsArgs']] = None,
                 vhost: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Exchange resources.
        :param pulumi.Input[str] name: The name of the exchange.
        :param pulumi.Input['ExchangeSettingsArgs'] settings: The settings of the exchange. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)
        if vhost is not None:
            pulumi.set(__self__, "vhost", vhost)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the exchange.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def settings(self) -> Optional[pulumi.Input['ExchangeSettingsArgs']]:
        """
        The settings of the exchange. The structure is
        described below.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input['ExchangeSettingsArgs']]):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter
    def vhost(self) -> Optional[pulumi.Input[str]]:
        """
        The vhost to create the resource in.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vhost", value)


class Exchange(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[pulumi.InputType['ExchangeSettingsArgs']]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The ``Exchange`` resource creates and manages an exchange.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test = rabbitmq.VHost("test", name="test")
        guest = rabbitmq.Permissions("guest",
            user="guest",
            vhost=test.name,
            permissions=rabbitmq.PermissionsPermissionsArgs(
                configure=".*",
                write=".*",
                read=".*",
            ))
        test_exchange = rabbitmq.Exchange("test",
            name="test",
            vhost=guest.vhost,
            settings=rabbitmq.ExchangeSettingsArgs(
                type="fanout",
                durable=False,
                auto_delete=True,
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Exchanges can be imported using the `id` which is composed of  `name@vhost`.

        E.g.

        ```sh
        $ pulumi import rabbitmq:index/exchange:Exchange test test@vhost
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the exchange.
        :param pulumi.Input[pulumi.InputType['ExchangeSettingsArgs']] settings: The settings of the exchange. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExchangeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``Exchange`` resource creates and manages an exchange.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test = rabbitmq.VHost("test", name="test")
        guest = rabbitmq.Permissions("guest",
            user="guest",
            vhost=test.name,
            permissions=rabbitmq.PermissionsPermissionsArgs(
                configure=".*",
                write=".*",
                read=".*",
            ))
        test_exchange = rabbitmq.Exchange("test",
            name="test",
            vhost=guest.vhost,
            settings=rabbitmq.ExchangeSettingsArgs(
                type="fanout",
                durable=False,
                auto_delete=True,
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Exchanges can be imported using the `id` which is composed of  `name@vhost`.

        E.g.

        ```sh
        $ pulumi import rabbitmq:index/exchange:Exchange test test@vhost
        ```

        :param str resource_name: The name of the resource.
        :param ExchangeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExchangeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[pulumi.InputType['ExchangeSettingsArgs']]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExchangeArgs.__new__(ExchangeArgs)

            __props__.__dict__["name"] = name
            if settings is None and not opts.urn:
                raise TypeError("Missing required property 'settings'")
            __props__.__dict__["settings"] = settings
            __props__.__dict__["vhost"] = vhost
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

        __props__ = _ExchangeState.__new__(_ExchangeState)

        __props__.__dict__["name"] = name
        __props__.__dict__["settings"] = settings
        __props__.__dict__["vhost"] = vhost
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

