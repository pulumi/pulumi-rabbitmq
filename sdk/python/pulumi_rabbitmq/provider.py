# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 cacert_file: Optional[pulumi.Input[_builtins.str]] = None,
                 clientcert_file: Optional[pulumi.Input[_builtins.str]] = None,
                 clientkey_file: Optional[pulumi.Input[_builtins.str]] = None,
                 endpoint: Optional[pulumi.Input[_builtins.str]] = None,
                 insecure: Optional[pulumi.Input[_builtins.bool]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 proxy: Optional[pulumi.Input[_builtins.str]] = None,
                 username: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        """
        if cacert_file is None:
            cacert_file = _utilities.get_env('RABBITMQ_CACERT')
        if cacert_file is not None:
            pulumi.set(__self__, "cacert_file", cacert_file)
        if clientcert_file is not None:
            pulumi.set(__self__, "clientcert_file", clientcert_file)
        if clientkey_file is not None:
            pulumi.set(__self__, "clientkey_file", clientkey_file)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if insecure is None:
            insecure = _utilities.get_env_bool('RABBITMQ_INSECURE')
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @_builtins.property
    @pulumi.getter(name="cacertFile")
    def cacert_file(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "cacert_file")

    @cacert_file.setter
    def cacert_file(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "cacert_file", value)

    @_builtins.property
    @pulumi.getter(name="clientcertFile")
    def clientcert_file(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "clientcert_file")

    @clientcert_file.setter
    def clientcert_file(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "clientcert_file", value)

    @_builtins.property
    @pulumi.getter(name="clientkeyFile")
    def clientkey_file(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "clientkey_file")

    @clientkey_file.setter
    def clientkey_file(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "clientkey_file", value)

    @_builtins.property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "endpoint", value)

    @_builtins.property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[_builtins.bool]]:
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "insecure", value)

    @_builtins.property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "password", value)

    @_builtins.property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "proxy", value)

    @_builtins.property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username", value)


@pulumi.type_token("pulumi:providers:rabbitmq")
class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cacert_file: Optional[pulumi.Input[_builtins.str]] = None,
                 clientcert_file: Optional[pulumi.Input[_builtins.str]] = None,
                 clientkey_file: Optional[pulumi.Input[_builtins.str]] = None,
                 endpoint: Optional[pulumi.Input[_builtins.str]] = None,
                 insecure: Optional[pulumi.Input[_builtins.bool]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 proxy: Optional[pulumi.Input[_builtins.str]] = None,
                 username: Optional[pulumi.Input[_builtins.str]] = None,
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
                 args: Optional[ProviderArgs] = None,
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
                 cacert_file: Optional[pulumi.Input[_builtins.str]] = None,
                 clientcert_file: Optional[pulumi.Input[_builtins.str]] = None,
                 clientkey_file: Optional[pulumi.Input[_builtins.str]] = None,
                 endpoint: Optional[pulumi.Input[_builtins.str]] = None,
                 insecure: Optional[pulumi.Input[_builtins.bool]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 proxy: Optional[pulumi.Input[_builtins.str]] = None,
                 username: Optional[pulumi.Input[_builtins.str]] = None,
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
            __props__.__dict__["endpoint"] = endpoint
            if insecure is None:
                insecure = _utilities.get_env_bool('RABBITMQ_INSECURE')
            __props__.__dict__["insecure"] = pulumi.Output.from_input(insecure).apply(pulumi.runtime.to_json) if insecure is not None else None
            __props__.__dict__["password"] = password
            __props__.__dict__["proxy"] = proxy
            __props__.__dict__["username"] = username
        super(Provider, __self__).__init__(
            'rabbitmq',
            resource_name,
            __props__,
            opts)

    @_builtins.property
    @pulumi.getter(name="cacertFile")
    def cacert_file(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "cacert_file")

    @_builtins.property
    @pulumi.getter(name="clientcertFile")
    def clientcert_file(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "clientcert_file")

    @_builtins.property
    @pulumi.getter(name="clientkeyFile")
    def clientkey_file(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "clientkey_file")

    @_builtins.property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "endpoint")

    @_builtins.property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "password")

    @_builtins.property
    @pulumi.getter
    def proxy(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "proxy")

    @_builtins.property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "username")

    @pulumi.output_type
    class TerraformConfigResult:
        def __init__(__self__, result=None):
            if result and not isinstance(result, dict):
                raise TypeError("Expected argument 'result' to be a dict")
            pulumi.set(__self__, "result", result)

        @_builtins.property
        @pulumi.getter
        def result(self) -> Mapping[str, Any]:
            return pulumi.get(self, "result")

    def terraform_config(__self__) -> pulumi.Output['Provider.TerraformConfigResult']:
        """
        This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        return pulumi.runtime.call('pulumi:providers:rabbitmq/terraformConfig', __args__, res=__self__, typ=Provider.TerraformConfigResult)

