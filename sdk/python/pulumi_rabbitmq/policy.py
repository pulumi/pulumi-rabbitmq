# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input['PolicyPolicyArgs'],
                 vhost: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input['PolicyPolicyArgs'] policy: The settings of the policy. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        :param pulumi.Input[str] name: The name of the policy.
        """
        pulumi.set(__self__, "policy", policy)
        pulumi.set(__self__, "vhost", vhost)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input['PolicyPolicyArgs']:
        """
        The settings of the policy. The structure is
        described below.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input['PolicyPolicyArgs']):
        pulumi.set(self, "policy", value)

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
        The name of the policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[pulumi.InputType['PolicyPolicyArgs']]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The ``Policy`` resource creates and manages policies for exchanges
        and queues.

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
        test_policy = rabbitmq.Policy("testPolicy",
            policy=rabbitmq.PolicyPolicyArgs(
                apply_to="all",
                definition={
                    "ha-mode": "all",
                },
                pattern=".*",
                priority=0,
            ),
            vhost=guest.vhost)
        ```

        ## Import

        Policies can be imported using the `id` which is composed of `name@vhost`. E.g.

        ```sh
         $ pulumi import rabbitmq:index/policy:Policy test name@vhost
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[pulumi.InputType['PolicyPolicyArgs']] policy: The settings of the policy. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``Policy`` resource creates and manages policies for exchanges
        and queues.

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
        test_policy = rabbitmq.Policy("testPolicy",
            policy=rabbitmq.PolicyPolicyArgs(
                apply_to="all",
                definition={
                    "ha-mode": "all",
                },
                pattern=".*",
                priority=0,
            ),
            vhost=guest.vhost)
        ```

        ## Import

        Policies can be imported using the `id` which is composed of `name@vhost`. E.g.

        ```sh
         $ pulumi import rabbitmq:index/policy:Policy test name@vhost
        ```

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[pulumi.InputType['PolicyPolicyArgs']]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
            if vhost is None and not opts.urn:
                raise TypeError("Missing required property 'vhost'")
            __props__['vhost'] = vhost
        super(Policy, __self__).__init__(
            'rabbitmq:index/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[pulumi.InputType['PolicyPolicyArgs']]] = None,
            vhost: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[pulumi.InputType['PolicyPolicyArgs']] policy: The settings of the policy. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["policy"] = policy
        __props__["vhost"] = vhost
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output['outputs.PolicyPolicy']:
        """
        The settings of the policy. The structure is
        described below.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def vhost(self) -> pulumi.Output[str]:
        """
        The vhost to create the resource in.
        """
        return pulumi.get(self, "vhost")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

