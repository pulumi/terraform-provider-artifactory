# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

accessToken: Optional[str]
"""
This is a bearer token that can be given to you by your admin under `Identity and Access`
"""

apiKey: Optional[str]

password: Optional[str]
"""
Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
bearer token
"""

url: Optional[str]

username: Optional[str]

