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
                 access_token: Optional[pulumi.Input[str]] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] access_token: This is a bearer token that can be given to you by your admin under `Identity and Access`
        :param pulumi.Input[str] password: Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
               bearer token
        """
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if api_key is not None:
            warnings.warn("""Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""", DeprecationWarning)
            pulumi.log.warn("""api_key is deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""")
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if password is not None:
            warnings.warn("""Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""", DeprecationWarning)
            pulumi.log.warn("""password is deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""")
        if password is not None:
            pulumi.set(__self__, "password", password)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if username is not None:
            warnings.warn("""Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""", DeprecationWarning)
            pulumi.log.warn("""username is deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""")
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        This is a bearer token that can be given to you by your admin under `Identity and Access`
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
        bearer token
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the artifactory package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: This is a bearer token that can be given to you by your admin under `Identity and Access`
        :param pulumi.Input[str] password: Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
               bearer token
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the artifactory package. By default, resources use package-wide configuration
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
                 access_token: Optional[pulumi.Input[str]] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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

            __props__.__dict__["access_token"] = access_token
            if api_key is not None and not opts.urn:
                warnings.warn("""Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""", DeprecationWarning)
                pulumi.log.warn("""api_key is deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""")
            __props__.__dict__["api_key"] = api_key
            if password is not None and not opts.urn:
                warnings.warn("""Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""", DeprecationWarning)
                pulumi.log.warn("""password is deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""")
            __props__.__dict__["password"] = password
            __props__.__dict__["url"] = url
            if username is not None and not opts.urn:
                warnings.warn("""Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""", DeprecationWarning)
                pulumi.log.warn("""username is deprecated: Xray and projects functionality will not work with any auth method other than access tokens (Bearer)""")
            __props__.__dict__["username"] = username
        super(Provider, __self__).__init__(
            'artifactory',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> pulumi.Output[Optional[str]]:
        """
        This is a bearer token that can be given to you by your admin under `Identity and Access`
        """
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Insider note: You may actually use an api_key as the password. This will get your around xray limitations instead of a
        bearer token
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "username")

