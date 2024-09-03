# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, admin=None, disable_ui_access=None, email=None, groups=None, id=None, internal_password_disabled=None, name=None, profile_updatable=None):
        if admin and not isinstance(admin, bool):
            raise TypeError("Expected argument 'admin' to be a bool")
        pulumi.set(__self__, "admin", admin)
        if disable_ui_access and not isinstance(disable_ui_access, bool):
            raise TypeError("Expected argument 'disable_ui_access' to be a bool")
        pulumi.set(__self__, "disable_ui_access", disable_ui_access)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_password_disabled and not isinstance(internal_password_disabled, bool):
            raise TypeError("Expected argument 'internal_password_disabled' to be a bool")
        pulumi.set(__self__, "internal_password_disabled", internal_password_disabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if profile_updatable and not isinstance(profile_updatable, bool):
            raise TypeError("Expected argument 'profile_updatable' to be a bool")
        pulumi.set(__self__, "profile_updatable", profile_updatable)

    @property
    @pulumi.getter
    def admin(self) -> Optional[bool]:
        """
        When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        """
        return pulumi.get(self, "admin")

    @property
    @pulumi.getter(name="disableUiAccess")
    def disable_ui_access(self) -> Optional[bool]:
        """
        When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        """
        return pulumi.get(self, "disable_ui_access")

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        """
        Email for user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence[str]]:
        """
        List of groups this user is a part of.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internalPasswordDisabled")
    def internal_password_disabled(self) -> Optional[bool]:
        """
        When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        """
        return pulumi.get(self, "internal_password_disabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="profileUpdatable")
    def profile_updatable(self) -> Optional[bool]:
        """
        When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        return pulumi.get(self, "profile_updatable")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            admin=self.admin,
            disable_ui_access=self.disable_ui_access,
            email=self.email,
            groups=self.groups,
            id=self.id,
            internal_password_disabled=self.internal_password_disabled,
            name=self.name,
            profile_updatable=self.profile_updatable)


def get_user(admin: Optional[bool] = None,
             disable_ui_access: Optional[bool] = None,
             email: Optional[str] = None,
             groups: Optional[Sequence[str]] = None,
             internal_password_disabled: Optional[bool] = None,
             name: Optional[str] = None,
             profile_updatable: Optional[bool] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    # Artifactory User Data Source

    Provides an Artifactory user data source. This can be used to read the configuration of users in artifactory.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    #
    user1 = artifactory.get_user(name="user1")
    ```


    :param bool admin: When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
    :param bool disable_ui_access: When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
    :param str email: Email for user.
    :param Sequence[str] groups: List of groups this user is a part of.
    :param bool internal_password_disabled: When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
    :param str name: Name of the user.
    :param bool profile_updatable: When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
    """
    __args__ = dict()
    __args__['admin'] = admin
    __args__['disableUiAccess'] = disable_ui_access
    __args__['email'] = email
    __args__['groups'] = groups
    __args__['internalPasswordDisabled'] = internal_password_disabled
    __args__['name'] = name
    __args__['profileUpdatable'] = profile_updatable
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        admin=pulumi.get(__ret__, 'admin'),
        disable_ui_access=pulumi.get(__ret__, 'disable_ui_access'),
        email=pulumi.get(__ret__, 'email'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        internal_password_disabled=pulumi.get(__ret__, 'internal_password_disabled'),
        name=pulumi.get(__ret__, 'name'),
        profile_updatable=pulumi.get(__ret__, 'profile_updatable'))


@_utilities.lift_output_func(get_user)
def get_user_output(admin: Optional[pulumi.Input[Optional[bool]]] = None,
                    disable_ui_access: Optional[pulumi.Input[Optional[bool]]] = None,
                    email: Optional[pulumi.Input[Optional[str]]] = None,
                    groups: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    internal_password_disabled: Optional[pulumi.Input[Optional[bool]]] = None,
                    name: Optional[pulumi.Input[str]] = None,
                    profile_updatable: Optional[pulumi.Input[Optional[bool]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    # Artifactory User Data Source

    Provides an Artifactory user data source. This can be used to read the configuration of users in artifactory.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    #
    user1 = artifactory.get_user(name="user1")
    ```


    :param bool admin: When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
    :param bool disable_ui_access: When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
    :param str email: Email for user.
    :param Sequence[str] groups: List of groups this user is a part of.
    :param bool internal_password_disabled: When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
    :param str name: Name of the user.
    :param bool profile_updatable: When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
    """
    ...
