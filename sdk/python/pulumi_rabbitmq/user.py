# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['User']


class User(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The ``User`` resource creates and manages a user.

        > **Note:** All arguments including username and password will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_rabbitmq as rabbitmq

        test = rabbitmq.User("test",
            password="foobar",
            tags=[
                "administrator",
                "management",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the user.
        :param pulumi.Input[str] password: The password of the user. The value of this argument
               is plain-text so make sure to secure where this is defined.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Which permission model to apply to the user. Valid
               options are: management, policymaker, monitoring, and administrator.
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
            if password is None:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            __props__['tags'] = tags
        super(User, __self__).__init__(
            'rabbitmq:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the user.
        :param pulumi.Input[str] password: The password of the user. The value of this argument
               is plain-text so make sure to secure where this is defined.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Which permission model to apply to the user. Valid
               options are: management, policymaker, monitoring, and administrator.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["password"] = password
        __props__["tags"] = tags
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password of the user. The value of this argument
        is plain-text so make sure to secure where this is defined.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Which permission model to apply to the user. Valid
        options are: management, policymaker, monitoring, and administrator.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

