# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 admin_privileges: Optional[pulumi.Input[bool]] = None,
                 auto_join: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detach_all_users: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_manager: Optional[pulumi.Input[bool]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 realm_attributes: Optional[pulumi.Input[str]] = None,
                 reports_manager: Optional[pulumi.Input[bool]] = None,
                 users_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 watch_manager: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input[bool] policy_manager: (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
               'false'.
        :param pulumi.Input[bool] reports_manager: (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        :param pulumi.Input[bool] watch_manager: (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
               'false'.
        """
        if admin_privileges is not None:
            pulumi.set(__self__, "admin_privileges", admin_privileges)
        if auto_join is not None:
            pulumi.set(__self__, "auto_join", auto_join)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if detach_all_users is not None:
            pulumi.set(__self__, "detach_all_users", detach_all_users)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_manager is not None:
            pulumi.set(__self__, "policy_manager", policy_manager)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if realm_attributes is not None:
            pulumi.set(__self__, "realm_attributes", realm_attributes)
        if reports_manager is not None:
            pulumi.set(__self__, "reports_manager", reports_manager)
        if users_names is not None:
            pulumi.set(__self__, "users_names", users_names)
        if watch_manager is not None:
            pulumi.set(__self__, "watch_manager", watch_manager)

    @property
    @pulumi.getter(name="adminPrivileges")
    def admin_privileges(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_privileges")

    @admin_privileges.setter
    def admin_privileges(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_privileges", value)

    @property
    @pulumi.getter(name="autoJoin")
    def auto_join(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_join")

    @auto_join.setter
    def auto_join(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_join", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="detachAllUsers")
    def detach_all_users(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "detach_all_users")

    @detach_all_users.setter
    def detach_all_users(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "detach_all_users", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyManager")
    def policy_manager(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
        'false'.
        """
        return pulumi.get(self, "policy_manager")

    @policy_manager.setter
    def policy_manager(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "policy_manager", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter(name="realmAttributes")
    def realm_attributes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm_attributes")

    @realm_attributes.setter
    def realm_attributes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_attributes", value)

    @property
    @pulumi.getter(name="reportsManager")
    def reports_manager(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        """
        return pulumi.get(self, "reports_manager")

    @reports_manager.setter
    def reports_manager(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reports_manager", value)

    @property
    @pulumi.getter(name="usersNames")
    def users_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "users_names")

    @users_names.setter
    def users_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users_names", value)

    @property
    @pulumi.getter(name="watchManager")
    def watch_manager(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
        'false'.
        """
        return pulumi.get(self, "watch_manager")

    @watch_manager.setter
    def watch_manager(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "watch_manager", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 admin_privileges: Optional[pulumi.Input[bool]] = None,
                 auto_join: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detach_all_users: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_manager: Optional[pulumi.Input[bool]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 realm_attributes: Optional[pulumi.Input[str]] = None,
                 reports_manager: Optional[pulumi.Input[bool]] = None,
                 users_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 watch_manager: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[bool] policy_manager: (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
               'false'.
        :param pulumi.Input[bool] reports_manager: (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        :param pulumi.Input[bool] watch_manager: (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
               'false'.
        """
        if admin_privileges is not None:
            pulumi.set(__self__, "admin_privileges", admin_privileges)
        if auto_join is not None:
            pulumi.set(__self__, "auto_join", auto_join)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if detach_all_users is not None:
            pulumi.set(__self__, "detach_all_users", detach_all_users)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_manager is not None:
            pulumi.set(__self__, "policy_manager", policy_manager)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if realm_attributes is not None:
            pulumi.set(__self__, "realm_attributes", realm_attributes)
        if reports_manager is not None:
            pulumi.set(__self__, "reports_manager", reports_manager)
        if users_names is not None:
            pulumi.set(__self__, "users_names", users_names)
        if watch_manager is not None:
            pulumi.set(__self__, "watch_manager", watch_manager)

    @property
    @pulumi.getter(name="adminPrivileges")
    def admin_privileges(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_privileges")

    @admin_privileges.setter
    def admin_privileges(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_privileges", value)

    @property
    @pulumi.getter(name="autoJoin")
    def auto_join(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_join")

    @auto_join.setter
    def auto_join(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_join", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="detachAllUsers")
    def detach_all_users(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "detach_all_users")

    @detach_all_users.setter
    def detach_all_users(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "detach_all_users", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyManager")
    def policy_manager(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
        'false'.
        """
        return pulumi.get(self, "policy_manager")

    @policy_manager.setter
    def policy_manager(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "policy_manager", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter(name="realmAttributes")
    def realm_attributes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm_attributes")

    @realm_attributes.setter
    def realm_attributes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_attributes", value)

    @property
    @pulumi.getter(name="reportsManager")
    def reports_manager(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        """
        return pulumi.get(self, "reports_manager")

    @reports_manager.setter
    def reports_manager(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reports_manager", value)

    @property
    @pulumi.getter(name="usersNames")
    def users_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "users_names")

    @users_names.setter
    def users_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users_names", value)

    @property
    @pulumi.getter(name="watchManager")
    def watch_manager(self) -> Optional[pulumi.Input[bool]]:
        """
        (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
        'false'.
        """
        return pulumi.get(self, "watch_manager")

    @watch_manager.setter
    def watch_manager(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "watch_manager", value)


class Group(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_privileges: Optional[pulumi.Input[bool]] = None,
                 auto_join: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detach_all_users: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_manager: Optional[pulumi.Input[bool]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 realm_attributes: Optional[pulumi.Input[str]] = None,
                 reports_manager: Optional[pulumi.Input[bool]] = None,
                 users_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 watch_manager: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a Group resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] policy_manager: (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
               'false'.
        :param pulumi.Input[bool] reports_manager: (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        :param pulumi.Input[bool] watch_manager: (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
               'false'.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Group resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_privileges: Optional[pulumi.Input[bool]] = None,
                 auto_join: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 detach_all_users: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_manager: Optional[pulumi.Input[bool]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 realm_attributes: Optional[pulumi.Input[str]] = None,
                 reports_manager: Optional[pulumi.Input[bool]] = None,
                 users_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 watch_manager: Optional[pulumi.Input[bool]] = None,
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
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["admin_privileges"] = admin_privileges
            __props__.__dict__["auto_join"] = auto_join
            __props__.__dict__["description"] = description
            __props__.__dict__["detach_all_users"] = detach_all_users
            __props__.__dict__["name"] = name
            __props__.__dict__["policy_manager"] = policy_manager
            __props__.__dict__["realm"] = realm
            __props__.__dict__["realm_attributes"] = realm_attributes
            __props__.__dict__["reports_manager"] = reports_manager
            __props__.__dict__["users_names"] = users_names
            __props__.__dict__["watch_manager"] = watch_manager
        super(Group, __self__).__init__(
            'artifactory:index/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_privileges: Optional[pulumi.Input[bool]] = None,
            auto_join: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            detach_all_users: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_manager: Optional[pulumi.Input[bool]] = None,
            realm: Optional[pulumi.Input[str]] = None,
            realm_attributes: Optional[pulumi.Input[str]] = None,
            reports_manager: Optional[pulumi.Input[bool]] = None,
            users_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            watch_manager: Optional[pulumi.Input[bool]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] policy_manager: (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
               'false'.
        :param pulumi.Input[bool] reports_manager: (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        :param pulumi.Input[bool] watch_manager: (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
               'false'.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["admin_privileges"] = admin_privileges
        __props__.__dict__["auto_join"] = auto_join
        __props__.__dict__["description"] = description
        __props__.__dict__["detach_all_users"] = detach_all_users
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_manager"] = policy_manager
        __props__.__dict__["realm"] = realm
        __props__.__dict__["realm_attributes"] = realm_attributes
        __props__.__dict__["reports_manager"] = reports_manager
        __props__.__dict__["users_names"] = users_names
        __props__.__dict__["watch_manager"] = watch_manager
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminPrivileges")
    def admin_privileges(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "admin_privileges")

    @property
    @pulumi.getter(name="autoJoin")
    def auto_join(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "auto_join")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="detachAllUsers")
    def detach_all_users(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "detach_all_users")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyManager")
    def policy_manager(self) -> pulumi.Output[Optional[bool]]:
        """
        (Optional) When this override is set, User in the group can set Xray security and compliance policies. Default value is
        'false'.
        """
        return pulumi.get(self, "policy_manager")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter(name="realmAttributes")
    def realm_attributes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "realm_attributes")

    @property
    @pulumi.getter(name="reportsManager")
    def reports_manager(self) -> pulumi.Output[Optional[bool]]:
        """
        (Optional) When this override is set, User in the group can manage Xray Reports. Default value is 'false'.
        """
        return pulumi.get(self, "reports_manager")

    @property
    @pulumi.getter(name="usersNames")
    def users_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "users_names")

    @property
    @pulumi.getter(name="watchManager")
    def watch_manager(self) -> pulumi.Output[Optional[bool]]:
        """
        (Optional) When this override is set, User in the group can manage Xray Watches on any resource type. Default value is
        'false'.
        """
        return pulumi.get(self, "watch_manager")

