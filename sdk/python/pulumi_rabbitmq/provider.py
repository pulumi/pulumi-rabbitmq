# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
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
                 insecure: Optional[pulumi.Input[bool]] = None,
                 proxy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        """
        ProviderArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            endpoint=endpoint,
            password=password,
            username=username,
            cacert_file=cacert_file,
            clientcert_file=clientcert_file,
            clientkey_file=clientkey_file,
            insecure=insecure,
            proxy=proxy,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             endpoint: Optional[pulumi.Input[str]] = None,
             password: Optional[pulumi.Input[str]] = None,
             username: Optional[pulumi.Input[str]] = None,
             cacert_file: Optional[pulumi.Input[str]] = None,
             clientcert_file: Optional[pulumi.Input[str]] = None,
             clientkey_file: Optional[pulumi.Input[str]] = None,
             insecure: Optional[pulumi.Input[bool]] = None,
             proxy: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if endpoint is None:
            raise TypeError("Missing 'endpoint' argument")
        if password is None:
            raise TypeError("Missing 'password' argument")
        if username is None:
            raise TypeError("Missing 'username' argument")
        if cacert_file is None and 'cacertFile' in kwargs:
            cacert_file = kwargs['cacertFile']
        if clientcert_file is None and 'clientcertFile' in kwargs:
            clientcert_file = kwargs['clientcertFile']
        if clientkey_file is None and 'clientkeyFile' in kwargs:
            clientkey_file = kwargs['clientkeyFile']

        _setter("endpoint", endpoint)
        _setter("password", password)
        _setter("username", username)
        if cacert_file is None:
            cacert_file = _utilities.get_env('RABBITMQ_CACERT')
        if cacert_file is not None:
            _setter("cacert_file", cacert_file)
        if clientcert_file is not None:
            _setter("clientcert_file", clientcert_file)
        if clientkey_file is not None:
            _setter("clientkey_file", clientkey_file)
        if insecure is None:
            insecure = _utilities.get_env_bool('RABBITMQ_INSECURE')
        if insecure is not None:
            _setter("insecure", insecure)
        if proxy is not None:
            _setter("proxy", proxy)

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

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy", value)


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
                 proxy: Optional[pulumi.Input[str]] = None,
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProviderArgs._configure(_setter, **kwargs)
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
                 proxy: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
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
            __props__.__dict__["proxy"] = proxy
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(Provider, __self__).__init__(
            'rabbitmq',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="cacertFile")
    def cacert_file(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cacert_file")

    @property
    @pulumi.getter(name="clientcertFile")
    def clientcert_file(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "clientcert_file")

    @property
    @pulumi.getter(name="clientkeyFile")
    def clientkey_file(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "clientkey_file")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def proxy(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "proxy")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        return pulumi.get(self, "username")

