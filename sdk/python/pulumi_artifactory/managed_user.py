# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ManagedUserArgs', 'ManagedUser']

@pulumi.input_type
class ManagedUserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 password: pulumi.Input[str],
                 admin: Optional[pulumi.Input[bool]] = None,
                 disable_ui_access: Optional[pulumi.Input[bool]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 internal_password_disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile_updatable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ManagedUser resource.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[str] password: Password for the user.
        :param pulumi.Input[bool] admin: When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        :param pulumi.Input[bool] disable_ui_access: When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of.
        :param pulumi.Input[bool] internal_password_disabled: When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user.
        :param pulumi.Input[bool] profile_updatable: When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "password", password)
        if admin is not None:
            pulumi.set(__self__, "admin", admin)
        if disable_ui_access is not None:
            pulumi.set(__self__, "disable_ui_access", disable_ui_access)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if internal_password_disabled is not None:
            pulumi.set(__self__, "internal_password_disabled", internal_password_disabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if profile_updatable is not None:
            pulumi.set(__self__, "profile_updatable", profile_updatable)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        Email for user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password for the user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def admin(self) -> Optional[pulumi.Input[bool]]:
        """
        When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        """
        return pulumi.get(self, "admin")

    @admin.setter
    def admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin", value)

    @property
    @pulumi.getter(name="disableUiAccess")
    def disable_ui_access(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        """
        return pulumi.get(self, "disable_ui_access")

    @disable_ui_access.setter
    def disable_ui_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_ui_access", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of groups this user is a part of.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="internalPasswordDisabled")
    def internal_password_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        """
        return pulumi.get(self, "internal_password_disabled")

    @internal_password_disabled.setter
    def internal_password_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal_password_disabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Username for user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="profileUpdatable")
    def profile_updatable(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        return pulumi.get(self, "profile_updatable")

    @profile_updatable.setter
    def profile_updatable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "profile_updatable", value)


@pulumi.input_type
class _ManagedUserState:
    def __init__(__self__, *,
                 admin: Optional[pulumi.Input[bool]] = None,
                 disable_ui_access: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 internal_password_disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 profile_updatable: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ManagedUser resources.
        :param pulumi.Input[bool] admin: When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        :param pulumi.Input[bool] disable_ui_access: When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of.
        :param pulumi.Input[bool] internal_password_disabled: When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user.
        :param pulumi.Input[str] password: Password for the user.
        :param pulumi.Input[bool] profile_updatable: When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        if admin is not None:
            pulumi.set(__self__, "admin", admin)
        if disable_ui_access is not None:
            pulumi.set(__self__, "disable_ui_access", disable_ui_access)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if internal_password_disabled is not None:
            pulumi.set(__self__, "internal_password_disabled", internal_password_disabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if profile_updatable is not None:
            pulumi.set(__self__, "profile_updatable", profile_updatable)

    @property
    @pulumi.getter
    def admin(self) -> Optional[pulumi.Input[bool]]:
        """
        When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        """
        return pulumi.get(self, "admin")

    @admin.setter
    def admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin", value)

    @property
    @pulumi.getter(name="disableUiAccess")
    def disable_ui_access(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        """
        return pulumi.get(self, "disable_ui_access")

    @disable_ui_access.setter
    def disable_ui_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_ui_access", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email for user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of groups this user is a part of.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="internalPasswordDisabled")
    def internal_password_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        """
        return pulumi.get(self, "internal_password_disabled")

    @internal_password_disabled.setter
    def internal_password_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal_password_disabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Username for user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for the user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="profileUpdatable")
    def profile_updatable(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        return pulumi.get(self, "profile_updatable")

    @profile_updatable.setter
    def profile_updatable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "profile_updatable", value)


class ManagedUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin: Optional[pulumi.Input[bool]] = None,
                 disable_ui_access: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 internal_password_disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 profile_updatable: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Artifactory Managed User Resource

        Provides an Artifactory managed user resource. This can be used to create and maintain Artifactory users. For example, service account where password is known and managed externally.

        Unlike `UnmanagedUser` and `User`, the `password` attribute is required and cannot be empty.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory user called terraform
        test_user = artifactory.ManagedUser("test-user",
            email="test-user@artifactory-terraform.com",
            groups=[
                "logged-in-users",
                "readers",
            ],
            password="my super secret password")
        ```

        ## Import

        Users can be imported using their name, e.g.

        ```sh
         $ pulumi import artifactory:index/managedUser:ManagedUser test-user myusername
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin: When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        :param pulumi.Input[bool] disable_ui_access: When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of.
        :param pulumi.Input[bool] internal_password_disabled: When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user.
        :param pulumi.Input[str] password: Password for the user.
        :param pulumi.Input[bool] profile_updatable: When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManagedUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Artifactory Managed User Resource

        Provides an Artifactory managed user resource. This can be used to create and maintain Artifactory users. For example, service account where password is known and managed externally.

        Unlike `UnmanagedUser` and `User`, the `password` attribute is required and cannot be empty.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory user called terraform
        test_user = artifactory.ManagedUser("test-user",
            email="test-user@artifactory-terraform.com",
            groups=[
                "logged-in-users",
                "readers",
            ],
            password="my super secret password")
        ```

        ## Import

        Users can be imported using their name, e.g.

        ```sh
         $ pulumi import artifactory:index/managedUser:ManagedUser test-user myusername
        ```

        :param str resource_name: The name of the resource.
        :param ManagedUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagedUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin: Optional[pulumi.Input[bool]] = None,
                 disable_ui_access: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 internal_password_disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 profile_updatable: Optional[pulumi.Input[bool]] = None,
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
            __props__ = ManagedUserArgs.__new__(ManagedUserArgs)

            __props__.__dict__["admin"] = admin
            __props__.__dict__["disable_ui_access"] = disable_ui_access
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["groups"] = groups
            __props__.__dict__["internal_password_disabled"] = internal_password_disabled
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            __props__.__dict__["profile_updatable"] = profile_updatable
        super(ManagedUser, __self__).__init__(
            'artifactory:index/managedUser:ManagedUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin: Optional[pulumi.Input[bool]] = None,
            disable_ui_access: Optional[pulumi.Input[bool]] = None,
            email: Optional[pulumi.Input[str]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            internal_password_disabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            profile_updatable: Optional[pulumi.Input[bool]] = None) -> 'ManagedUser':
        """
        Get an existing ManagedUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin: When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        :param pulumi.Input[bool] disable_ui_access: When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of.
        :param pulumi.Input[bool] internal_password_disabled: When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user.
        :param pulumi.Input[str] password: Password for the user.
        :param pulumi.Input[bool] profile_updatable: When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ManagedUserState.__new__(_ManagedUserState)

        __props__.__dict__["admin"] = admin
        __props__.__dict__["disable_ui_access"] = disable_ui_access
        __props__.__dict__["email"] = email
        __props__.__dict__["groups"] = groups
        __props__.__dict__["internal_password_disabled"] = internal_password_disabled
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["profile_updatable"] = profile_updatable
        return ManagedUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def admin(self) -> pulumi.Output[Optional[bool]]:
        """
        When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        """
        return pulumi.get(self, "admin")

    @property
    @pulumi.getter(name="disableUiAccess")
    def disable_ui_access(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        """
        return pulumi.get(self, "disable_ui_access")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        Email for user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of groups this user is a part of.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="internalPasswordDisabled")
    def internal_password_disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        """
        return pulumi.get(self, "internal_password_disabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Username for user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password for the user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="profileUpdatable")
    def profile_updatable(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        """
        return pulumi.get(self, "profile_updatable")

