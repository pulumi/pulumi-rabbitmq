# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
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
from . import outputs

__all__ = [
    'GetExchangeResult',
    'AwaitableGetExchangeResult',
    'get_exchange',
    'get_exchange_output',
]

@pulumi.output_type
class GetExchangeResult:
    """
    A collection of values returned by getExchange.
    """
    def __init__(__self__, id=None, name=None, settings=None, vhost=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if settings and not isinstance(settings, list):
            raise TypeError("Expected argument 'settings' to be a list")
        pulumi.set(__self__, "settings", settings)
        if vhost and not isinstance(vhost, str):
            raise TypeError("Expected argument 'vhost' to be a str")
        pulumi.set(__self__, "vhost", vhost)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def settings(self) -> Sequence['outputs.GetExchangeSettingResult']:
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter
    def vhost(self) -> Optional[builtins.str]:
        return pulumi.get(self, "vhost")


class AwaitableGetExchangeResult(GetExchangeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExchangeResult(
            id=self.id,
            name=self.name,
            settings=self.settings,
            vhost=self.vhost)


def get_exchange(name: Optional[builtins.str] = None,
                 vhost: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExchangeResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vhost'] = vhost
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('rabbitmq:index/getExchange:getExchange', __args__, opts=opts, typ=GetExchangeResult).value

    return AwaitableGetExchangeResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        settings=pulumi.get(__ret__, 'settings'),
        vhost=pulumi.get(__ret__, 'vhost'))
def get_exchange_output(name: Optional[pulumi.Input[builtins.str]] = None,
                        vhost: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetExchangeResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vhost'] = vhost
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('rabbitmq:index/getExchange:getExchange', __args__, opts=opts, typ=GetExchangeResult)
    return __ret__.apply(lambda __response__: GetExchangeResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        settings=pulumi.get(__response__, 'settings'),
        vhost=pulumi.get(__response__, 'vhost')))
