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

__all__ = ['FederationUpstreamArgs', 'FederationUpstream']

@pulumi.input_type
class FederationUpstreamArgs:
    def __init__(__self__, *,
                 definition: pulumi.Input['FederationUpstreamDefinitionArgs'],
                 vhost: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FederationUpstream resource.
        :param pulumi.Input['FederationUpstreamDefinitionArgs'] definition: The configuration of the federation upstream. The structure is described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        :param pulumi.Input[str] name: The name of the federation upstream.
        """
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "vhost", vhost)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Input['FederationUpstreamDefinitionArgs']:
        """
        The configuration of the federation upstream. The structure is described below.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: pulumi.Input['FederationUpstreamDefinitionArgs']):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Input[str]:
        """
        The vhost to create the resource in.
        """
        return pulumi.get(self, "vhost")

    @vhost.setter
    def vhost(self, value: pulumi.Input[str]):
        pulumi.set(self, "vhost", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the federation upstream.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FederationUpstreamState:
    def __init__(__self__, *,
                 component: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input['FederationUpstreamDefinitionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vhost: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FederationUpstream resources.
        :param pulumi.Input[str] component: Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
        :param pulumi.Input['FederationUpstreamDefinitionArgs'] definition: The configuration of the federation upstream. The structure is described below.
        :param pulumi.Input[str] name: The name of the federation upstream.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        if component is not None:
            pulumi.set(__self__, "component", component)
        if definition is not None:
            pulumi.set(__self__, "definition", definition)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vhost is not None:
            pulumi.set(__self__, "vhost", vhost)

    @property
    @pulumi.getter
    def component(self) -> Optional[pulumi.Input[str]]:
        """
        Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
        """
        return pulumi.get(self, "component")

    @component.setter
    def component(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "component", value)

    @property
    @pulumi.getter
    def definition(self) -> Optional[pulumi.Input['FederationUpstreamDefinitionArgs']]:
        """
        The configuration of the federation upstream. The structure is described below.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: Optional[pulumi.Input['FederationUpstreamDefinitionArgs']]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the federation upstream.
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


class FederationUpstream(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['FederationUpstreamDefinitionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The ``FederationUpstream`` resource creates and manages a federation upstream parameter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test = rabbitmq.VHost("test")
        guest = rabbitmq.Permissions("guest",
            user="guest",
            vhost=test.name,
            permissions=rabbitmq.PermissionsPermissionsArgs(
                configure=".*",
                write=".*",
                read=".*",
            ))
        # downstream exchange
        foo_exchange = rabbitmq.Exchange("fooExchange",
            vhost=guest.vhost,
            settings=rabbitmq.ExchangeSettingsArgs(
                type="topic",
                durable=True,
            ))
        # upstream broker
        foo_federation_upstream = rabbitmq.FederationUpstream("fooFederationUpstream",
            vhost=guest.vhost,
            definition=rabbitmq.FederationUpstreamDefinitionArgs(
                uri="amqp://guest:guest@upstream-server-name:5672/%2f",
                prefetch_count=1000,
                reconnect_delay=5,
                ack_mode="on-confirm",
                trust_user_id=False,
                max_hops=1,
            ))
        foo_policy = rabbitmq.Policy("fooPolicy",
            vhost=guest.vhost,
            policy=rabbitmq.PolicyPolicyArgs(
                pattern=foo_exchange.name.apply(lambda name: f"(^{name}$)"),
                priority=1,
                apply_to="exchanges",
                definition={
                    "federation-upstream": foo_federation_upstream.name,
                },
            ))
        ```

        ## Import

        A Federation upstream can be imported using the resource `id` which is composed of `name@vhost`, e.g.

        ```sh
        $ pulumi import rabbitmq:index/federationUpstream:FederationUpstream foo foo@test
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['FederationUpstreamDefinitionArgs']] definition: The configuration of the federation upstream. The structure is described below.
        :param pulumi.Input[str] name: The name of the federation upstream.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FederationUpstreamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``FederationUpstream`` resource creates and manages a federation upstream parameter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test = rabbitmq.VHost("test")
        guest = rabbitmq.Permissions("guest",
            user="guest",
            vhost=test.name,
            permissions=rabbitmq.PermissionsPermissionsArgs(
                configure=".*",
                write=".*",
                read=".*",
            ))
        # downstream exchange
        foo_exchange = rabbitmq.Exchange("fooExchange",
            vhost=guest.vhost,
            settings=rabbitmq.ExchangeSettingsArgs(
                type="topic",
                durable=True,
            ))
        # upstream broker
        foo_federation_upstream = rabbitmq.FederationUpstream("fooFederationUpstream",
            vhost=guest.vhost,
            definition=rabbitmq.FederationUpstreamDefinitionArgs(
                uri="amqp://guest:guest@upstream-server-name:5672/%2f",
                prefetch_count=1000,
                reconnect_delay=5,
                ack_mode="on-confirm",
                trust_user_id=False,
                max_hops=1,
            ))
        foo_policy = rabbitmq.Policy("fooPolicy",
            vhost=guest.vhost,
            policy=rabbitmq.PolicyPolicyArgs(
                pattern=foo_exchange.name.apply(lambda name: f"(^{name}$)"),
                priority=1,
                apply_to="exchanges",
                definition={
                    "federation-upstream": foo_federation_upstream.name,
                },
            ))
        ```

        ## Import

        A Federation upstream can be imported using the resource `id` which is composed of `name@vhost`, e.g.

        ```sh
        $ pulumi import rabbitmq:index/federationUpstream:FederationUpstream foo foo@test
        ```

        :param str resource_name: The name of the resource.
        :param FederationUpstreamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FederationUpstreamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['FederationUpstreamDefinitionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FederationUpstreamArgs.__new__(FederationUpstreamArgs)

            if definition is None and not opts.urn:
                raise TypeError("Missing required property 'definition'")
            __props__.__dict__["definition"] = definition
            __props__.__dict__["name"] = name
            if vhost is None and not opts.urn:
                raise TypeError("Missing required property 'vhost'")
            __props__.__dict__["vhost"] = vhost
            __props__.__dict__["component"] = None
        super(FederationUpstream, __self__).__init__(
            'rabbitmq:index/federationUpstream:FederationUpstream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            component: Optional[pulumi.Input[str]] = None,
            definition: Optional[pulumi.Input[pulumi.InputType['FederationUpstreamDefinitionArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vhost: Optional[pulumi.Input[str]] = None) -> 'FederationUpstream':
        """
        Get an existing FederationUpstream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] component: Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
        :param pulumi.Input[pulumi.InputType['FederationUpstreamDefinitionArgs']] definition: The configuration of the federation upstream. The structure is described below.
        :param pulumi.Input[str] name: The name of the federation upstream.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FederationUpstreamState.__new__(_FederationUpstreamState)

        __props__.__dict__["component"] = component
        __props__.__dict__["definition"] = definition
        __props__.__dict__["name"] = name
        __props__.__dict__["vhost"] = vhost
        return FederationUpstream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def component(self) -> pulumi.Output[str]:
        """
        Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
        """
        return pulumi.get(self, "component")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output['outputs.FederationUpstreamDefinition']:
        """
        The configuration of the federation upstream. The structure is described below.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the federation upstream.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Output[str]:
        """
        The vhost to create the resource in.
        """
        return pulumi.get(self, "vhost")

