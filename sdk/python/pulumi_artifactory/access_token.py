# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AccessTokenArgs', 'AccessToken']

@pulumi.input_type
class AccessTokenArgs:
    def __init__(__self__, *,
                 username: pulumi.Input[str],
                 admin_token: Optional[pulumi.Input['AccessTokenAdminTokenArgs']] = None,
                 audience: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refreshable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AccessToken resource.
        :param pulumi.Input[str] username: (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        :param pulumi.Input['AccessTokenAdminTokenArgs'] admin_token: (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        :param pulumi.Input[str] audience: (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        :param pulumi.Input[str] end_date: (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        :param pulumi.Input[str] end_date_relative: (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        :param pulumi.Input[bool] refreshable: (Optional) Is this token refreshable? Defaults to `false`
        """
        pulumi.set(__self__, "username", username)
        if admin_token is not None:
            pulumi.set(__self__, "admin_token", admin_token)
        if audience is not None:
            pulumi.set(__self__, "audience", audience)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if end_date_relative is not None:
            pulumi.set(__self__, "end_date_relative", end_date_relative)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if refreshable is not None:
            pulumi.set(__self__, "refreshable", refreshable)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="adminToken")
    def admin_token(self) -> Optional[pulumi.Input['AccessTokenAdminTokenArgs']]:
        """
        (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        """
        return pulumi.get(self, "admin_token")

    @admin_token.setter
    def admin_token(self, value: Optional[pulumi.Input['AccessTokenAdminTokenArgs']]):
        pulumi.set(self, "admin_token", value)

    @property
    @pulumi.getter
    def audience(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        """
        return pulumi.get(self, "audience")

    @audience.setter
    def audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audience", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date_relative", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def refreshable(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) Is this token refreshable? Defaults to `false`
        """
        return pulumi.get(self, "refreshable")

    @refreshable.setter
    def refreshable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "refreshable", value)


@pulumi.input_type
class _AccessTokenState:
    def __init__(__self__, *,
                 access_token: Optional[pulumi.Input[str]] = None,
                 admin_token: Optional[pulumi.Input['AccessTokenAdminTokenArgs']] = None,
                 audience: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refresh_token: Optional[pulumi.Input[str]] = None,
                 refreshable: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccessToken resources.
        :param pulumi.Input[str] access_token: Returns the access token to authenciate to Artifactory
        :param pulumi.Input['AccessTokenAdminTokenArgs'] admin_token: (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        :param pulumi.Input[str] audience: (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        :param pulumi.Input[str] end_date: (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        :param pulumi.Input[str] end_date_relative: (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        :param pulumi.Input[str] refresh_token: Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
        :param pulumi.Input[bool] refreshable: (Optional) Is this token refreshable? Defaults to `false`
        :param pulumi.Input[str] username: (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        """
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if admin_token is not None:
            pulumi.set(__self__, "admin_token", admin_token)
        if audience is not None:
            pulumi.set(__self__, "audience", audience)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if end_date_relative is not None:
            pulumi.set(__self__, "end_date_relative", end_date_relative)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if refresh_token is not None:
            pulumi.set(__self__, "refresh_token", refresh_token)
        if refreshable is not None:
            pulumi.set(__self__, "refreshable", refreshable)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        Returns the access token to authenciate to Artifactory
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter(name="adminToken")
    def admin_token(self) -> Optional[pulumi.Input['AccessTokenAdminTokenArgs']]:
        """
        (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        """
        return pulumi.get(self, "admin_token")

    @admin_token.setter
    def admin_token(self, value: Optional[pulumi.Input['AccessTokenAdminTokenArgs']]):
        pulumi.set(self, "admin_token", value)

    @property
    @pulumi.getter
    def audience(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        """
        return pulumi.get(self, "audience")

    @audience.setter
    def audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audience", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date_relative", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="refreshToken")
    def refresh_token(self) -> Optional[pulumi.Input[str]]:
        """
        Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
        """
        return pulumi.get(self, "refresh_token")

    @refresh_token.setter
    def refresh_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "refresh_token", value)

    @property
    @pulumi.getter
    def refreshable(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) Is this token refreshable? Defaults to `false`
        """
        return pulumi.get(self, "refreshable")

    @refreshable.setter
    def refreshable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "refreshable", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class AccessToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_token: Optional[pulumi.Input[pulumi.InputType['AccessTokenAdminTokenArgs']]] = None,
                 audience: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refreshable: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Artifactory Access Token resource. This can be used to create and manage Artifactory Access Tokens.

        > **Note:** Access Tokens will be stored in the raw state as plain-text. Read more about sensitive data in
        state.

        ## Example Usage

        ### S
        ### Create a new Artifactory Access Token for an existing user

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        exising_user = artifactory.AccessToken("exisingUser",
            end_date_relative="5m",
            username="existing-user")
        ```

        Note: This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
        ### Create a new Artifactory User and Access token
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        new_user_user = artifactory.User("newUserUser",
            email="new_user@somewhere.com",
            groups=["readers"])
        new_user_access_token = artifactory.AccessToken("newUserAccessToken",
            username=new_user_user.name,
            end_date_relative="5m")
        ```
        ### Creates a new token for groups
        This creates a transient user called `temporary-user`.
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        temporary_user = artifactory.AccessToken("temporaryUser",
            end_date_relative="1h",
            groups=["readers"],
            username="temporary-user")
        ```
        ### Create token with no expiry
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        no_expiry = artifactory.AccessToken("noExpiry",
            end_date_relative="0s",
            username="existing-user")
        ```
        ### Creates a refreshable token
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        refreshable = artifactory.AccessToken("refreshable",
            end_date_relative="1m",
            groups=["readers"],
            refreshable=True,
            username="refreshable")
        ```
        ### Creates an administrator token
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        admin = artifactory.AccessToken("admin",
            admin_token=artifactory.AccessTokenAdminTokenArgs(
                instance_id="<instance id>",
            ),
            end_date_relative="1m",
            username="admin")
        ```
        ### Creates a token with an audience
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        audience = artifactory.AccessToken("audience",
            audience="jfrt@*",
            end_date_relative="1m",
            refreshable=True,
            username="audience")
        ```
        ### Creates a token with a fixed end date
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        fixeddate = artifactory.AccessToken("fixeddate",
            end_date="2018-01-01T01:02:03Z",
            groups=["readers"],
            username="fixeddate")
        ```
        ### Rotate token after it expires
        This example will generate a token that will expire in 1 hour.

        If `pulumi up` is run before 1 hour, nothing changes.
        One an hour has passed, `pulumi up` will generate a new token.

        ```python
        import pulumi
        import pulumi_artifactory as artifactory
        import pulumiverse_time as time

        now_plus1_hours = time.Rotating("nowPlus1Hours", rotation_hours=1)
        rotating = artifactory.AccessToken("rotating",
            username="rotating",
            end_date=time_rotating["now_plus_1_hour"]["rotation_rfc3339"],
            groups=["readers"])
        ```
        ## References

        - https://www.jfrog.com/confluence/display/ACC1X/Access+Tokens
        - https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateToken

        ## Import

        Artifactory **does not** retain access tokens and cannot be imported into state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AccessTokenAdminTokenArgs']] admin_token: (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        :param pulumi.Input[str] audience: (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        :param pulumi.Input[str] end_date: (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        :param pulumi.Input[str] end_date_relative: (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        :param pulumi.Input[bool] refreshable: (Optional) Is this token refreshable? Defaults to `false`
        :param pulumi.Input[str] username: (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory Access Token resource. This can be used to create and manage Artifactory Access Tokens.

        > **Note:** Access Tokens will be stored in the raw state as plain-text. Read more about sensitive data in
        state.

        ## Example Usage

        ### S
        ### Create a new Artifactory Access Token for an existing user

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        exising_user = artifactory.AccessToken("exisingUser",
            end_date_relative="5m",
            username="existing-user")
        ```

        Note: This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
        ### Create a new Artifactory User and Access token
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        new_user_user = artifactory.User("newUserUser",
            email="new_user@somewhere.com",
            groups=["readers"])
        new_user_access_token = artifactory.AccessToken("newUserAccessToken",
            username=new_user_user.name,
            end_date_relative="5m")
        ```
        ### Creates a new token for groups
        This creates a transient user called `temporary-user`.
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        temporary_user = artifactory.AccessToken("temporaryUser",
            end_date_relative="1h",
            groups=["readers"],
            username="temporary-user")
        ```
        ### Create token with no expiry
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        no_expiry = artifactory.AccessToken("noExpiry",
            end_date_relative="0s",
            username="existing-user")
        ```
        ### Creates a refreshable token
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        refreshable = artifactory.AccessToken("refreshable",
            end_date_relative="1m",
            groups=["readers"],
            refreshable=True,
            username="refreshable")
        ```
        ### Creates an administrator token
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        admin = artifactory.AccessToken("admin",
            admin_token=artifactory.AccessTokenAdminTokenArgs(
                instance_id="<instance id>",
            ),
            end_date_relative="1m",
            username="admin")
        ```
        ### Creates a token with an audience
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        audience = artifactory.AccessToken("audience",
            audience="jfrt@*",
            end_date_relative="1m",
            refreshable=True,
            username="audience")
        ```
        ### Creates a token with a fixed end date
        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        fixeddate = artifactory.AccessToken("fixeddate",
            end_date="2018-01-01T01:02:03Z",
            groups=["readers"],
            username="fixeddate")
        ```
        ### Rotate token after it expires
        This example will generate a token that will expire in 1 hour.

        If `pulumi up` is run before 1 hour, nothing changes.
        One an hour has passed, `pulumi up` will generate a new token.

        ```python
        import pulumi
        import pulumi_artifactory as artifactory
        import pulumiverse_time as time

        now_plus1_hours = time.Rotating("nowPlus1Hours", rotation_hours=1)
        rotating = artifactory.AccessToken("rotating",
            username="rotating",
            end_date=time_rotating["now_plus_1_hour"]["rotation_rfc3339"],
            groups=["readers"])
        ```
        ## References

        - https://www.jfrog.com/confluence/display/ACC1X/Access+Tokens
        - https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateToken

        ## Import

        Artifactory **does not** retain access tokens and cannot be imported into state.

        :param str resource_name: The name of the resource.
        :param AccessTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_token: Optional[pulumi.Input[pulumi.InputType['AccessTokenAdminTokenArgs']]] = None,
                 audience: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 refreshable: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessTokenArgs.__new__(AccessTokenArgs)

            __props__.__dict__["admin_token"] = admin_token
            __props__.__dict__["audience"] = audience
            __props__.__dict__["end_date"] = end_date
            __props__.__dict__["end_date_relative"] = end_date_relative
            __props__.__dict__["groups"] = groups
            __props__.__dict__["refreshable"] = refreshable
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["access_token"] = None
            __props__.__dict__["refresh_token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessToken", "refreshToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AccessToken, __self__).__init__(
            'artifactory:index/accessToken:AccessToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_token: Optional[pulumi.Input[str]] = None,
            admin_token: Optional[pulumi.Input[pulumi.InputType['AccessTokenAdminTokenArgs']]] = None,
            audience: Optional[pulumi.Input[str]] = None,
            end_date: Optional[pulumi.Input[str]] = None,
            end_date_relative: Optional[pulumi.Input[str]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            refresh_token: Optional[pulumi.Input[str]] = None,
            refreshable: Optional[pulumi.Input[bool]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'AccessToken':
        """
        Get an existing AccessToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: Returns the access token to authenciate to Artifactory
        :param pulumi.Input[pulumi.InputType['AccessTokenAdminTokenArgs']] admin_token: (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        :param pulumi.Input[str] audience: (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        :param pulumi.Input[str] end_date: (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        :param pulumi.Input[str] end_date_relative: (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        :param pulumi.Input[str] refresh_token: Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
        :param pulumi.Input[bool] refreshable: (Optional) Is this token refreshable? Defaults to `false`
        :param pulumi.Input[str] username: (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessTokenState.__new__(_AccessTokenState)

        __props__.__dict__["access_token"] = access_token
        __props__.__dict__["admin_token"] = admin_token
        __props__.__dict__["audience"] = audience
        __props__.__dict__["end_date"] = end_date
        __props__.__dict__["end_date_relative"] = end_date_relative
        __props__.__dict__["groups"] = groups
        __props__.__dict__["refresh_token"] = refresh_token
        __props__.__dict__["refreshable"] = refreshable
        __props__.__dict__["username"] = username
        return AccessToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> pulumi.Output[str]:
        """
        Returns the access token to authenciate to Artifactory
        """
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="adminToken")
    def admin_token(self) -> pulumi.Output[Optional['outputs.AccessTokenAdminToken']]:
        """
        (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
        """
        return pulumi.get(self, "admin_token")

    @property
    @pulumi.getter
    def audience(self) -> pulumi.Output[Optional[str]]:
        """
        (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `"jfrt@*"` so the token to be accepted by all Artifactory instances.
        """
        return pulumi.get(self, "audience")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> pulumi.Output[str]:
        """
        (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> pulumi.Output[Optional[str]]:
        """
        (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "s", "m", "h".
        """
        return pulumi.get(self, "end_date_relative")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `["*"]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="refreshToken")
    def refresh_token(self) -> pulumi.Output[str]:
        """
        Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
        """
        return pulumi.get(self, "refresh_token")

    @property
    @pulumi.getter
    def refreshable(self) -> pulumi.Output[Optional[bool]]:
        """
        (Optional) Is this token refreshable? Defaults to `false`
        """
        return pulumi.get(self, "refreshable")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
        """
        return pulumi.get(self, "username")

