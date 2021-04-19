# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 password: pulumi.Input[str],
                 username: pulumi.Input[str],
                 cacert_file: Optional[pulumi.Input[str]] = None,
                 clientcert_file: Optional[pulumi.Input[str]] = None,
                 clientkey_file: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        if cacert_file is None:
            cacert_file = _utilities.get_env('RABBITMQ_CACERT')
        if cacert_file is not None:
            pulumi.set(__self__, "cacert_file", cacert_file)
        if clientcert_file is not None:
            pulumi.set(__self__, "clientcert_file", clientcert_file)
        if clientkey_file is not None:
            pulumi.set(__self__, "clientkey_file", clientkey_file)
        if insecure is None:
            insecure = _utilities.get_env_bool('RABBITMQ_INSECURE')
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="cacertFile")
    def cacert_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cacert_file")

    @cacert_file.setter
    def cacert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cacert_file", value)

    @property
    @pulumi.getter(name="clientcertFile")
    def clientcert_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "clientcert_file")

    @clientcert_file.setter
    def clientcert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clientcert_file", value)

    @property
    @pulumi.getter(name="clientkeyFile")
    def clientkey_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "clientkey_file")

    @clientkey_file.setter
    def clientkey_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clientkey_file", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cacert_file: Optional[pulumi.Input[str]] = None,
                 clientcert_file: Optional[pulumi.Input[str]] = None,
                 clientkey_file: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the rabbitmq package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the rabbitmq package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cacert_file: Optional[pulumi.Input[str]] = None,
                 clientcert_file: Optional[pulumi.Input[str]] = None,
                 clientkey_file: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if cacert_file is None:
                cacert_file = _utilities.get_env('RABBITMQ_CACERT')
            __props__.__dict__["cacert_file"] = cacert_file
            __props__.__dict__["clientcert_file"] = clientcert_file
            __props__.__dict__["clientkey_file"] = clientkey_file
            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__.__dict__["endpoint"] = endpoint
            if insecure is None:
                insecure = _utilities.get_env_bool('RABBITMQ_INSECURE')
            __props__.__dict__["insecure"] = pulumi.Output.from_input(insecure).apply(pulumi.runtime.to_json) if insecure is not None else None
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(Provider, __self__).__init__(
            'rabbitmq',
            resource_name,
            __props__,
            opts)

