# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['OauthSettingsArgs', 'OauthSettings']

@pulumi.input_type
class OauthSettingsArgs:
    def __init__(__self__, *,
                 oauth_providers: pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]],
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 persist_users: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a OauthSettings resource.
        :param pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]] oauth_providers: OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `false`.
        :param pulumi.Input[bool] enable: Enable OAuth SSO.  Default value is `true`.
        :param pulumi.Input[bool] persist_users: Enable the creation of local Artifactory users.  Default value is `false`.
        """
        OauthSettingsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            oauth_providers=oauth_providers,
            allow_user_to_access_profile=allow_user_to_access_profile,
            enable=enable,
            persist_users=persist_users,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             oauth_providers: pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]],
             allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
             enable: Optional[pulumi.Input[bool]] = None,
             persist_users: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'oauthProviders' in kwargs:
            oauth_providers = kwargs['oauthProviders']
        if 'allowUserToAccessProfile' in kwargs:
            allow_user_to_access_profile = kwargs['allowUserToAccessProfile']
        if 'persistUsers' in kwargs:
            persist_users = kwargs['persistUsers']

        _setter("oauth_providers", oauth_providers)
        if allow_user_to_access_profile is not None:
            _setter("allow_user_to_access_profile", allow_user_to_access_profile)
        if enable is not None:
            _setter("enable", enable)
        if persist_users is not None:
            _setter("persist_users", persist_users)

    @property
    @pulumi.getter(name="oauthProviders")
    def oauth_providers(self) -> pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]]:
        """
        OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        """
        return pulumi.get(self, "oauth_providers")

    @oauth_providers.setter
    def oauth_providers(self, value: pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]]):
        pulumi.set(self, "oauth_providers", value)

    @property
    @pulumi.getter(name="allowUserToAccessProfile")
    def allow_user_to_access_profile(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow persisted users to access their profile.  Default value is `false`.
        """
        return pulumi.get(self, "allow_user_to_access_profile")

    @allow_user_to_access_profile.setter
    def allow_user_to_access_profile(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_user_to_access_profile", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable OAuth SSO.  Default value is `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="persistUsers")
    def persist_users(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable the creation of local Artifactory users.  Default value is `false`.
        """
        return pulumi.get(self, "persist_users")

    @persist_users.setter
    def persist_users(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "persist_users", value)


@pulumi.input_type
class _OauthSettingsState:
    def __init__(__self__, *,
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 oauth_providers: Optional[pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]]] = None,
                 persist_users: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering OauthSettings resources.
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `false`.
        :param pulumi.Input[bool] enable: Enable OAuth SSO.  Default value is `true`.
        :param pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]] oauth_providers: OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        :param pulumi.Input[bool] persist_users: Enable the creation of local Artifactory users.  Default value is `false`.
        """
        _OauthSettingsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allow_user_to_access_profile=allow_user_to_access_profile,
            enable=enable,
            oauth_providers=oauth_providers,
            persist_users=persist_users,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
             enable: Optional[pulumi.Input[bool]] = None,
             oauth_providers: Optional[pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]]] = None,
             persist_users: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'allowUserToAccessProfile' in kwargs:
            allow_user_to_access_profile = kwargs['allowUserToAccessProfile']
        if 'oauthProviders' in kwargs:
            oauth_providers = kwargs['oauthProviders']
        if 'persistUsers' in kwargs:
            persist_users = kwargs['persistUsers']

        if allow_user_to_access_profile is not None:
            _setter("allow_user_to_access_profile", allow_user_to_access_profile)
        if enable is not None:
            _setter("enable", enable)
        if oauth_providers is not None:
            _setter("oauth_providers", oauth_providers)
        if persist_users is not None:
            _setter("persist_users", persist_users)

    @property
    @pulumi.getter(name="allowUserToAccessProfile")
    def allow_user_to_access_profile(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow persisted users to access their profile.  Default value is `false`.
        """
        return pulumi.get(self, "allow_user_to_access_profile")

    @allow_user_to_access_profile.setter
    def allow_user_to_access_profile(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_user_to_access_profile", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable OAuth SSO.  Default value is `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="oauthProviders")
    def oauth_providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]]]:
        """
        OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        """
        return pulumi.get(self, "oauth_providers")

    @oauth_providers.setter
    def oauth_providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OauthSettingsOauthProviderArgs']]]]):
        pulumi.set(self, "oauth_providers", value)

    @property
    @pulumi.getter(name="persistUsers")
    def persist_users(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable the creation of local Artifactory users.  Default value is `false`.
        """
        return pulumi.get(self, "persist_users")

    @persist_users.setter
    def persist_users(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "persist_users", value)


class OauthSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 oauth_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OauthSettingsOauthProviderArgs']]]]] = None,
                 persist_users: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resource can be used to manage Artifactory's OAuth SSO settings.

        Only a single `OauthSettings` resource is meant to be defined.

        ~>The `OauthSettings` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory OAuth SSO settings
        oauth = artifactory.OauthSettings("oauth",
            allow_user_to_access_profile=True,
            enable=True,
            oauth_providers=[artifactory.OauthSettingsOauthProviderArgs(
                api_url="https://organization.okta.com/oauth2/v1/userinfo",
                auth_url="https://organization.okta.com/oauth2/v1/authorize",
                client_id="foo",
                client_secret="bar",
                enabled=False,
                name="okta",
                token_url="https://organization.okta.com/oauth2/v1/token",
                type="openId",
            )],
            persist_users=True)
        ```

        ## Import

        Current OAuth SSO settings can be imported using `oauth_settings` as the `ID`. If the resource is being imported, there will be a state drift, because `client_secret` can't be known. There are two options on how to approach this1) Don't set `client_secret` initially, import, then update the config with actual secret; 2) Accept that there is a drift initially and run `pulumi up` twice;

        ```sh
         $ pulumi import artifactory:index/oauthSettings:OauthSettings oauth oauth_settings
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `false`.
        :param pulumi.Input[bool] enable: Enable OAuth SSO.  Default value is `true`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OauthSettingsOauthProviderArgs']]]] oauth_providers: OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        :param pulumi.Input[bool] persist_users: Enable the creation of local Artifactory users.  Default value is `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OauthSettingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can be used to manage Artifactory's OAuth SSO settings.

        Only a single `OauthSettings` resource is meant to be defined.

        ~>The `OauthSettings` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory OAuth SSO settings
        oauth = artifactory.OauthSettings("oauth",
            allow_user_to_access_profile=True,
            enable=True,
            oauth_providers=[artifactory.OauthSettingsOauthProviderArgs(
                api_url="https://organization.okta.com/oauth2/v1/userinfo",
                auth_url="https://organization.okta.com/oauth2/v1/authorize",
                client_id="foo",
                client_secret="bar",
                enabled=False,
                name="okta",
                token_url="https://organization.okta.com/oauth2/v1/token",
                type="openId",
            )],
            persist_users=True)
        ```

        ## Import

        Current OAuth SSO settings can be imported using `oauth_settings` as the `ID`. If the resource is being imported, there will be a state drift, because `client_secret` can't be known. There are two options on how to approach this1) Don't set `client_secret` initially, import, then update the config with actual secret; 2) Accept that there is a drift initially and run `pulumi up` twice;

        ```sh
         $ pulumi import artifactory:index/oauthSettings:OauthSettings oauth oauth_settings
        ```

        :param str resource_name: The name of the resource.
        :param OauthSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OauthSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            OauthSettingsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 oauth_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OauthSettingsOauthProviderArgs']]]]] = None,
                 persist_users: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OauthSettingsArgs.__new__(OauthSettingsArgs)

            __props__.__dict__["allow_user_to_access_profile"] = allow_user_to_access_profile
            __props__.__dict__["enable"] = enable
            if oauth_providers is None and not opts.urn:
                raise TypeError("Missing required property 'oauth_providers'")
            __props__.__dict__["oauth_providers"] = oauth_providers
            __props__.__dict__["persist_users"] = persist_users
        super(OauthSettings, __self__).__init__(
            'artifactory:index/oauthSettings:OauthSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            oauth_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OauthSettingsOauthProviderArgs']]]]] = None,
            persist_users: Optional[pulumi.Input[bool]] = None) -> 'OauthSettings':
        """
        Get an existing OauthSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `false`.
        :param pulumi.Input[bool] enable: Enable OAuth SSO.  Default value is `true`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OauthSettingsOauthProviderArgs']]]] oauth_providers: OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        :param pulumi.Input[bool] persist_users: Enable the creation of local Artifactory users.  Default value is `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OauthSettingsState.__new__(_OauthSettingsState)

        __props__.__dict__["allow_user_to_access_profile"] = allow_user_to_access_profile
        __props__.__dict__["enable"] = enable
        __props__.__dict__["oauth_providers"] = oauth_providers
        __props__.__dict__["persist_users"] = persist_users
        return OauthSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowUserToAccessProfile")
    def allow_user_to_access_profile(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow persisted users to access their profile.  Default value is `false`.
        """
        return pulumi.get(self, "allow_user_to_access_profile")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable OAuth SSO.  Default value is `true`.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="oauthProviders")
    def oauth_providers(self) -> pulumi.Output[Sequence['outputs.OauthSettingsOauthProvider']]:
        """
        OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        """
        return pulumi.get(self, "oauth_providers")

    @property
    @pulumi.getter(name="persistUsers")
    def persist_users(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable the creation of local Artifactory users.  Default value is `false`.
        """
        return pulumi.get(self, "persist_users")

