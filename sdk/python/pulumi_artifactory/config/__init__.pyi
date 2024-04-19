# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

accessToken: Optional[str]
"""
This is a access token that can be given to you by your admin under `User Management > Access Tokens`. If not set, the
'api_key' attribute value will be used.
"""

apiKey: Optional[str]
"""
API key. If `access_token` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set,
the provider will ignore this attribute.
"""

checkLicense: bool
"""
Toggle for pre-flight checking of Artifactory Pro and Enterprise license. Default to `true`.
"""

oidcProviderName: Optional[str]
"""
OIDC provider name. See [Configure an OIDC
Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
more details.
"""

url: Optional[str]
"""
Artifactory URL.
"""

