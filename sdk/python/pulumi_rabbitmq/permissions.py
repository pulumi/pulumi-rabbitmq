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

__all__ = ['PermissionsArgs', 'Permissions']

@pulumi.input_type
class PermissionsArgs:
    def __init__(__self__, *,
                 permissions: pulumi.Input['PermissionsPermissionsArgs'],
                 user: pulumi.Input[str],
                 vhost: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Permissions resource.
        :param pulumi.Input['PermissionsPermissionsArgs'] permissions: The settings of the permissions. The structure is
               described below.
        :param pulumi.Input[str] user: The user to apply the permissions to.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "user", user)
        if vhost is not None:
            pulumi.set(__self__, "vhost", vhost)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input['PermissionsPermissionsArgs']:
        """
        The settings of the permissions. The structure is
        described below.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input['PermissionsPermissionsArgs']):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        The user to apply the permissions to.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

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


class Permissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[pulumi.InputType['PermissionsPermissionsArgs']]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vhost: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The ``Permissions`` resource creates and manages a user's set of
        permissions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test_v_host = rabbitmq.VHost("testVHost")
        test_user = rabbitmq.User("testUser",
            password="foobar",
            tags=["administrator"])
        test_permissions = rabbitmq.Permissions("testPermissions",
            permissions=rabbitmq.PermissionsPermissionsArgs(
                configure=".*",
                read=".*",
                write=".*",
            ),
            user=test_user.name,
            vhost=test_v_host.name)
        ```

        ## Import

        Permissions can be imported using the `id` which is composed of

        `user@vhost`. E.g.

        ```sh
         $ pulumi import rabbitmq:index/permissions:Permissions test user@vhost
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PermissionsPermissionsArgs']] permissions: The settings of the permissions. The structure is
               described below.
        :param pulumi.Input[str] user: The user to apply the permissions to.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``Permissions`` resource creates and manages a user's set of
        permissions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test_v_host = rabbitmq.VHost("testVHost")
        test_user = rabbitmq.User("testUser",
            password="foobar",
            tags=["administrator"])
        test_permissions = rabbitmq.Permissions("testPermissions",
            permissions=rabbitmq.PermissionsPermissionsArgs(
                configure=".*",
                read=".*",
                write=".*",
            ),
            user=test_user.name,
            vhost=test_v_host.name)
        ```

        ## Import

        Permissions can be imported using the `id` which is composed of

        `user@vhost`. E.g.

        ```sh
         $ pulumi import rabbitmq:index/permissions:Permissions test user@vhost
        ```

        :param str resource_name: The name of the resource.
        :param PermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions: Optional[pulumi.Input[pulumi.InputType['PermissionsPermissionsArgs']]] = None,
                 user: Optional[pulumi.Input[str]] = None,
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

            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__['permissions'] = permissions
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
            __props__['vhost'] = vhost
        super(Permissions, __self__).__init__(
            'rabbitmq:index/permissions:Permissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            permissions: Optional[pulumi.Input[pulumi.InputType['PermissionsPermissionsArgs']]] = None,
            user: Optional[pulumi.Input[str]] = None,
            vhost: Optional[pulumi.Input[str]] = None) -> 'Permissions':
        """
        Get an existing Permissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PermissionsPermissionsArgs']] permissions: The settings of the permissions. The structure is
               described below.
        :param pulumi.Input[str] user: The user to apply the permissions to.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["permissions"] = permissions
        __props__["user"] = user
        __props__["vhost"] = vhost
        return Permissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output['outputs.PermissionsPermissions']:
        """
        The settings of the permissions. The structure is
        described below.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        The user to apply the permissions to.
        """
        return pulumi.get(self, "user")

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

