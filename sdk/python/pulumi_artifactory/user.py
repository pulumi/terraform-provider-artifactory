# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 admin: Optional[pulumi.Input[bool]] = None,
                 disable_ui_access: Optional[pulumi.Input[bool]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 internal_password_disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 profile_updatable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[bool] admin: (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        :param pulumi.Input[bool] disable_ui_access: (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
        :param pulumi.Input[bool] internal_password_disabled: (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        :param pulumi.Input[str] password: (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        :param pulumi.Input[bool] profile_updatable: (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        """
        UserArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            email=email,
            admin=admin,
            disable_ui_access=disable_ui_access,
            groups=groups,
            internal_password_disabled=internal_password_disabled,
            name=name,
            password=password,
            profile_updatable=profile_updatable,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             email: Optional[pulumi.Input[str]] = None,
             admin: Optional[pulumi.Input[bool]] = None,
             disable_ui_access: Optional[pulumi.Input[bool]] = None,
             groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             internal_password_disabled: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             password: Optional[pulumi.Input[str]] = None,
             profile_updatable: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if email is None:
            raise TypeError("Missing 'email' argument")
        if disable_ui_access is None and 'disableUiAccess' in kwargs:
            disable_ui_access = kwargs['disableUiAccess']
        if internal_password_disabled is None and 'internalPasswordDisabled' in kwargs:
            internal_password_disabled = kwargs['internalPasswordDisabled']
        if profile_updatable is None and 'profileUpdatable' in kwargs:
            profile_updatable = kwargs['profileUpdatable']

        _setter("email", email)
        if admin is not None:
            _setter("admin", admin)
        if disable_ui_access is not None:
            _setter("disable_ui_access", disable_ui_access)
        if groups is not None:
            _setter("groups", groups)
        if internal_password_disabled is not None:
            _setter("internal_password_disabled", internal_password_disabled)
        if name is not None:
            _setter("name", name)
        if password is not None:
            _setter("password", password)
        if profile_updatable is not None:
            _setter("profile_updatable", profile_updatable)

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
    def admin(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        """
        return pulumi.get(self, "admin")

    @admin.setter
    def admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin", value)

    @property
    @pulumi.getter(name="disableUiAccess")
    def disable_ui_access(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        """
        return pulumi.get(self, "disable_ui_access")

    @disable_ui_access.setter
    def disable_ui_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_ui_access", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="internalPasswordDisabled")
    def internal_password_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        """
        return pulumi.get(self, "internal_password_disabled")

    @internal_password_disabled.setter
    def internal_password_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal_password_disabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="profileUpdatable")
    def profile_updatable(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        """
        return pulumi.get(self, "profile_updatable")

    @profile_updatable.setter
    def profile_updatable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "profile_updatable", value)


@pulumi.input_type
class _UserState:
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
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[bool] admin: (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        :param pulumi.Input[bool] disable_ui_access: (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
        :param pulumi.Input[bool] internal_password_disabled: (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        :param pulumi.Input[str] password: (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        :param pulumi.Input[bool] profile_updatable: (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        """
        _UserState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            admin=admin,
            disable_ui_access=disable_ui_access,
            email=email,
            groups=groups,
            internal_password_disabled=internal_password_disabled,
            name=name,
            password=password,
            profile_updatable=profile_updatable,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             admin: Optional[pulumi.Input[bool]] = None,
             disable_ui_access: Optional[pulumi.Input[bool]] = None,
             email: Optional[pulumi.Input[str]] = None,
             groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             internal_password_disabled: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             password: Optional[pulumi.Input[str]] = None,
             profile_updatable: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if disable_ui_access is None and 'disableUiAccess' in kwargs:
            disable_ui_access = kwargs['disableUiAccess']
        if internal_password_disabled is None and 'internalPasswordDisabled' in kwargs:
            internal_password_disabled = kwargs['internalPasswordDisabled']
        if profile_updatable is None and 'profileUpdatable' in kwargs:
            profile_updatable = kwargs['profileUpdatable']

        if admin is not None:
            _setter("admin", admin)
        if disable_ui_access is not None:
            _setter("disable_ui_access", disable_ui_access)
        if email is not None:
            _setter("email", email)
        if groups is not None:
            _setter("groups", groups)
        if internal_password_disabled is not None:
            _setter("internal_password_disabled", internal_password_disabled)
        if name is not None:
            _setter("name", name)
        if password is not None:
            _setter("password", password)
        if profile_updatable is not None:
            _setter("profile_updatable", profile_updatable)

    @property
    @pulumi.getter
    def admin(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        """
        return pulumi.get(self, "admin")

    @admin.setter
    def admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin", value)

    @property
    @pulumi.getter(name="disableUiAccess")
    def disable_ui_access(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
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
        List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="internalPasswordDisabled")
    def internal_password_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        """
        return pulumi.get(self, "internal_password_disabled")

    @internal_password_disabled.setter
    def internal_password_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal_password_disabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="profileUpdatable")
    def profile_updatable(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        """
        return pulumi.get(self, "profile_updatable")

    @profile_updatable.setter
    def profile_updatable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "profile_updatable", value)


class User(pulumi.CustomResource):
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
        Provides an Artifactory user resource. This can be used to create and manage Artifactory users.
        The password is a required field by the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser), but we made it optional in this resource to accommodate the scenario where the password is not needed and will be reset by the actual user later.\\
        When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.

        > The generated password won't be stored in the TF state and can not be recovered. The user must reset the password to be able to log in. An admin can always generate the access key for the user as well. The password change won't trigger state drift. We don't recommend to use this resource unless there is a specific use case for it. Recommended resource is `ManagedUser`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        test_user = artifactory.User("test-user",
            admin=False,
            disable_ui_access=False,
            email="test-user@artifactory-terraform.com",
            groups=[
                "readers",
                "logged-in-users",
            ],
            internal_password_disabled=False,
            password="my super secret password",
            profile_updatable=True)
        ```
        ## Managing groups relationship

        See our recommendation on how to manage user-group relationship.

        ## Import

        ```sh
         $ pulumi import artifactory:index/user:User test-user myusername
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin: (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        :param pulumi.Input[bool] disable_ui_access: (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
        :param pulumi.Input[bool] internal_password_disabled: (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        :param pulumi.Input[str] password: (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        :param pulumi.Input[bool] profile_updatable: (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory user resource. This can be used to create and manage Artifactory users.
        The password is a required field by the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser), but we made it optional in this resource to accommodate the scenario where the password is not needed and will be reset by the actual user later.\\
        When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.

        > The generated password won't be stored in the TF state and can not be recovered. The user must reset the password to be able to log in. An admin can always generate the access key for the user as well. The password change won't trigger state drift. We don't recommend to use this resource unless there is a specific use case for it. Recommended resource is `ManagedUser`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        test_user = artifactory.User("test-user",
            admin=False,
            disable_ui_access=False,
            email="test-user@artifactory-terraform.com",
            groups=[
                "readers",
                "logged-in-users",
            ],
            internal_password_disabled=False,
            password="my super secret password",
            profile_updatable=True)
        ```
        ## Managing groups relationship

        See our recommendation on how to manage user-group relationship.

        ## Import

        ```sh
         $ pulumi import artifactory:index/user:User test-user myusername
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            UserArgs._configure(_setter, **kwargs)
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
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["admin"] = admin
            __props__.__dict__["disable_ui_access"] = disable_ui_access
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["groups"] = groups
            __props__.__dict__["internal_password_disabled"] = internal_password_disabled
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["profile_updatable"] = profile_updatable
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(User, __self__).__init__(
            'artifactory:index/user:User',
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
            profile_updatable: Optional[pulumi.Input[bool]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin: (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        :param pulumi.Input[bool] disable_ui_access: (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
        :param pulumi.Input[str] email: Email for user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
        :param pulumi.Input[bool] internal_password_disabled: (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        :param pulumi.Input[str] name: Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        :param pulumi.Input[str] password: (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        :param pulumi.Input[bool] profile_updatable: (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["admin"] = admin
        __props__.__dict__["disable_ui_access"] = disable_ui_access
        __props__.__dict__["email"] = email
        __props__.__dict__["groups"] = groups
        __props__.__dict__["internal_password_disabled"] = internal_password_disabled
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["profile_updatable"] = profile_updatable
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def admin(self) -> pulumi.Output[bool]:
        """
        (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
        """
        return pulumi.get(self, "admin")

    @property
    @pulumi.getter(name="disableUiAccess")
    def disable_ui_access(self) -> pulumi.Output[bool]:
        """
        (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
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
    def groups(self) -> pulumi.Output[Sequence[str]]:
        """
        List of groups this user is a part of. If no groups set, `readers` group will be added by default. If other groups are assigned, `readers` must be added to the list manually to avoid state drift.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="internalPasswordDisabled")
    def internal_password_disabled(self) -> pulumi.Output[bool]:
        """
        (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
        """
        return pulumi.get(self, "internal_password_disabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Username for user. May contain lowercase letters, numbers and symbols: '.-_@'
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="profileUpdatable")
    def profile_updatable(self) -> pulumi.Output[bool]:
        """
        (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
        """
        return pulumi.get(self, "profile_updatable")

