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
from .. import _utilities

cacertFile: Optional[str]

clientcertFile: Optional[str]

clientkeyFile: Optional[str]

endpoint: Optional[str]

insecure: Optional[bool]

password: Optional[str]

proxy: Optional[str]

username: Optional[str]

