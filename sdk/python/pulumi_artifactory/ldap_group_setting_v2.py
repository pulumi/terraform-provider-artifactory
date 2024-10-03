# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['LdapGroupSettingV2Args', 'LdapGroupSettingV2']

@pulumi.input_type
class LdapGroupSettingV2Args:
    def __init__(__self__, *,
                 description_attribute: pulumi.Input[str],
                 filter: pulumi.Input[str],
                 group_member_attribute: pulumi.Input[str],
                 group_name_attribute: pulumi.Input[str],
                 strategy: pulumi.Input[str],
                 enabled_ldap: Optional[pulumi.Input[str]] = None,
                 force_attribute_search: Optional[pulumi.Input[bool]] = None,
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sub_tree: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a LdapGroupSettingV2 resource.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
        :param pulumi.Input[str] enabled_ldap: The LDAP setting key you want to use for group retrieval.
        :param pulumi.Input[bool] force_attribute_search: This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
        """
        pulumi.set(__self__, "description_attribute", description_attribute)
        pulumi.set(__self__, "filter", filter)
        pulumi.set(__self__, "group_member_attribute", group_member_attribute)
        pulumi.set(__self__, "group_name_attribute", group_name_attribute)
        pulumi.set(__self__, "strategy", strategy)
        if enabled_ldap is not None:
            pulumi.set(__self__, "enabled_ldap", enabled_ldap)
        if force_attribute_search is not None:
            pulumi.set(__self__, "force_attribute_search", force_attribute_search)
        if group_base_dn is not None:
            pulumi.set(__self__, "group_base_dn", group_base_dn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sub_tree is not None:
            pulumi.set(__self__, "sub_tree", sub_tree)

    @property
    @pulumi.getter(name="descriptionAttribute")
    def description_attribute(self) -> pulumi.Input[str]:
        """
        An attribute on the group entry which denoting the group description. Used when importing groups.
        """
        return pulumi.get(self, "description_attribute")

    @description_attribute.setter
    def description_attribute(self, value: pulumi.Input[str]):
        pulumi.set(self, "description_attribute", value)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input[str]:
        """
        The LDAP filter used to search for group entries. Used for importing groups.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="groupMemberAttribute")
    def group_member_attribute(self) -> pulumi.Input[str]:
        """
        A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
        """
        return pulumi.get(self, "group_member_attribute")

    @group_member_attribute.setter
    def group_member_attribute(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_member_attribute", value)

    @property
    @pulumi.getter(name="groupNameAttribute")
    def group_name_attribute(self) -> pulumi.Input[str]:
        """
        Attribute on the group entry denoting the group name. Used when importing groups.
        """
        return pulumi.get(self, "group_name_attribute")

    @group_name_attribute.setter
    def group_name_attribute(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name_attribute", value)

    @property
    @pulumi.getter
    def strategy(self) -> pulumi.Input[str]:
        """
        The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
        """
        return pulumi.get(self, "strategy")

    @strategy.setter
    def strategy(self, value: pulumi.Input[str]):
        pulumi.set(self, "strategy", value)

    @property
    @pulumi.getter(name="enabledLdap")
    def enabled_ldap(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP setting key you want to use for group retrieval.
        """
        return pulumi.get(self, "enabled_ldap")

    @enabled_ldap.setter
    def enabled_ldap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enabled_ldap", value)

    @property
    @pulumi.getter(name="forceAttributeSearch")
    def force_attribute_search(self) -> Optional[pulumi.Input[bool]]:
        """
        This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
        """
        return pulumi.get(self, "force_attribute_search")

    @force_attribute_search.setter
    def force_attribute_search(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_attribute_search", value)

    @property
    @pulumi.getter(name="groupBaseDn")
    def group_base_dn(self) -> Optional[pulumi.Input[str]]:
        """
        A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        """
        return pulumi.get(self, "group_base_dn")

    @group_base_dn.setter
    def group_base_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_base_dn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Ldap group setting name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="subTree")
    def sub_tree(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
        """
        return pulumi.get(self, "sub_tree")

    @sub_tree.setter
    def sub_tree(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sub_tree", value)


@pulumi.input_type
class _LdapGroupSettingV2State:
    def __init__(__self__, *,
                 description_attribute: Optional[pulumi.Input[str]] = None,
                 enabled_ldap: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 force_attribute_search: Optional[pulumi.Input[bool]] = None,
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 group_member_attribute: Optional[pulumi.Input[str]] = None,
                 group_name_attribute: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[str]] = None,
                 sub_tree: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering LdapGroupSettingV2 resources.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] enabled_ldap: The LDAP setting key you want to use for group retrieval.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[bool] force_attribute_search: This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
        """
        if description_attribute is not None:
            pulumi.set(__self__, "description_attribute", description_attribute)
        if enabled_ldap is not None:
            pulumi.set(__self__, "enabled_ldap", enabled_ldap)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if force_attribute_search is not None:
            pulumi.set(__self__, "force_attribute_search", force_attribute_search)
        if group_base_dn is not None:
            pulumi.set(__self__, "group_base_dn", group_base_dn)
        if group_member_attribute is not None:
            pulumi.set(__self__, "group_member_attribute", group_member_attribute)
        if group_name_attribute is not None:
            pulumi.set(__self__, "group_name_attribute", group_name_attribute)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if strategy is not None:
            pulumi.set(__self__, "strategy", strategy)
        if sub_tree is not None:
            pulumi.set(__self__, "sub_tree", sub_tree)

    @property
    @pulumi.getter(name="descriptionAttribute")
    def description_attribute(self) -> Optional[pulumi.Input[str]]:
        """
        An attribute on the group entry which denoting the group description. Used when importing groups.
        """
        return pulumi.get(self, "description_attribute")

    @description_attribute.setter
    def description_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description_attribute", value)

    @property
    @pulumi.getter(name="enabledLdap")
    def enabled_ldap(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP setting key you want to use for group retrieval.
        """
        return pulumi.get(self, "enabled_ldap")

    @enabled_ldap.setter
    def enabled_ldap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enabled_ldap", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP filter used to search for group entries. Used for importing groups.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="forceAttributeSearch")
    def force_attribute_search(self) -> Optional[pulumi.Input[bool]]:
        """
        This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
        """
        return pulumi.get(self, "force_attribute_search")

    @force_attribute_search.setter
    def force_attribute_search(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_attribute_search", value)

    @property
    @pulumi.getter(name="groupBaseDn")
    def group_base_dn(self) -> Optional[pulumi.Input[str]]:
        """
        A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        """
        return pulumi.get(self, "group_base_dn")

    @group_base_dn.setter
    def group_base_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_base_dn", value)

    @property
    @pulumi.getter(name="groupMemberAttribute")
    def group_member_attribute(self) -> Optional[pulumi.Input[str]]:
        """
        A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
        """
        return pulumi.get(self, "group_member_attribute")

    @group_member_attribute.setter
    def group_member_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_member_attribute", value)

    @property
    @pulumi.getter(name="groupNameAttribute")
    def group_name_attribute(self) -> Optional[pulumi.Input[str]]:
        """
        Attribute on the group entry denoting the group name. Used when importing groups.
        """
        return pulumi.get(self, "group_name_attribute")

    @group_name_attribute.setter
    def group_name_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name_attribute", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Ldap group setting name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
        """
        return pulumi.get(self, "strategy")

    @strategy.setter
    def strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "strategy", value)

    @property
    @pulumi.getter(name="subTree")
    def sub_tree(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
        """
        return pulumi.get(self, "sub_tree")

    @sub_tree.setter
    def sub_tree(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sub_tree", value)


class LdapGroupSettingV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description_attribute: Optional[pulumi.Input[str]] = None,
                 enabled_ldap: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 force_attribute_search: Optional[pulumi.Input[bool]] = None,
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 group_member_attribute: Optional[pulumi.Input[str]] = None,
                 group_name_attribute: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[str]] = None,
                 sub_tree: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides an Artifactory LDAP Setting resource.

        This resource can be used to manage Artifactory's LDAP Group settings for user authentication.

        LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
        structure for managing group-based permissions.

        [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/ldap-group-setting), [general documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/ldap).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        ldap_group_name = artifactory.LdapGroupSettingV2("ldap_group_name",
            name="ldap_group_name",
            enabled_ldap="ldap_name",
            group_base_dn="CN=Users,DC=MyDomain,DC=com",
            group_name_attribute="cn",
            group_member_attribute="uniqueMember",
            sub_tree=True,
            force_attribute_search=False,
            filter="(objectClass=groupOfNames)",
            description_attribute="description",
            strategy="STATIC")
        ```

        ## Import

        ```sh
        $ pulumi import artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2 ldap ldapGroup1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] enabled_ldap: The LDAP setting key you want to use for group retrieval.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[bool] force_attribute_search: This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LdapGroupSettingV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory LDAP Setting resource.

        This resource can be used to manage Artifactory's LDAP Group settings for user authentication.

        LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
        structure for managing group-based permissions.

        [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/ldap-group-setting), [general documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/ldap).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        ldap_group_name = artifactory.LdapGroupSettingV2("ldap_group_name",
            name="ldap_group_name",
            enabled_ldap="ldap_name",
            group_base_dn="CN=Users,DC=MyDomain,DC=com",
            group_name_attribute="cn",
            group_member_attribute="uniqueMember",
            sub_tree=True,
            force_attribute_search=False,
            filter="(objectClass=groupOfNames)",
            description_attribute="description",
            strategy="STATIC")
        ```

        ## Import

        ```sh
        $ pulumi import artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2 ldap ldapGroup1
        ```

        :param str resource_name: The name of the resource.
        :param LdapGroupSettingV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LdapGroupSettingV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description_attribute: Optional[pulumi.Input[str]] = None,
                 enabled_ldap: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 force_attribute_search: Optional[pulumi.Input[bool]] = None,
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 group_member_attribute: Optional[pulumi.Input[str]] = None,
                 group_name_attribute: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[str]] = None,
                 sub_tree: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LdapGroupSettingV2Args.__new__(LdapGroupSettingV2Args)

            if description_attribute is None and not opts.urn:
                raise TypeError("Missing required property 'description_attribute'")
            __props__.__dict__["description_attribute"] = description_attribute
            __props__.__dict__["enabled_ldap"] = enabled_ldap
            if filter is None and not opts.urn:
                raise TypeError("Missing required property 'filter'")
            __props__.__dict__["filter"] = filter
            __props__.__dict__["force_attribute_search"] = force_attribute_search
            __props__.__dict__["group_base_dn"] = group_base_dn
            if group_member_attribute is None and not opts.urn:
                raise TypeError("Missing required property 'group_member_attribute'")
            __props__.__dict__["group_member_attribute"] = group_member_attribute
            if group_name_attribute is None and not opts.urn:
                raise TypeError("Missing required property 'group_name_attribute'")
            __props__.__dict__["group_name_attribute"] = group_name_attribute
            __props__.__dict__["name"] = name
            if strategy is None and not opts.urn:
                raise TypeError("Missing required property 'strategy'")
            __props__.__dict__["strategy"] = strategy
            __props__.__dict__["sub_tree"] = sub_tree
        super(LdapGroupSettingV2, __self__).__init__(
            'artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description_attribute: Optional[pulumi.Input[str]] = None,
            enabled_ldap: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            force_attribute_search: Optional[pulumi.Input[bool]] = None,
            group_base_dn: Optional[pulumi.Input[str]] = None,
            group_member_attribute: Optional[pulumi.Input[str]] = None,
            group_name_attribute: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            strategy: Optional[pulumi.Input[str]] = None,
            sub_tree: Optional[pulumi.Input[bool]] = None) -> 'LdapGroupSettingV2':
        """
        Get an existing LdapGroupSettingV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] enabled_ldap: The LDAP setting key you want to use for group retrieval.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[bool] force_attribute_search: This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LdapGroupSettingV2State.__new__(_LdapGroupSettingV2State)

        __props__.__dict__["description_attribute"] = description_attribute
        __props__.__dict__["enabled_ldap"] = enabled_ldap
        __props__.__dict__["filter"] = filter
        __props__.__dict__["force_attribute_search"] = force_attribute_search
        __props__.__dict__["group_base_dn"] = group_base_dn
        __props__.__dict__["group_member_attribute"] = group_member_attribute
        __props__.__dict__["group_name_attribute"] = group_name_attribute
        __props__.__dict__["name"] = name
        __props__.__dict__["strategy"] = strategy
        __props__.__dict__["sub_tree"] = sub_tree
        return LdapGroupSettingV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="descriptionAttribute")
    def description_attribute(self) -> pulumi.Output[str]:
        """
        An attribute on the group entry which denoting the group description. Used when importing groups.
        """
        return pulumi.get(self, "description_attribute")

    @property
    @pulumi.getter(name="enabledLdap")
    def enabled_ldap(self) -> pulumi.Output[str]:
        """
        The LDAP setting key you want to use for group retrieval.
        """
        return pulumi.get(self, "enabled_ldap")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        """
        The LDAP filter used to search for group entries. Used for importing groups.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="forceAttributeSearch")
    def force_attribute_search(self) -> pulumi.Output[bool]:
        """
        This attribute is used in very specific cases of LDAP group settings. Don't switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
        """
        return pulumi.get(self, "force_attribute_search")

    @property
    @pulumi.getter(name="groupBaseDn")
    def group_base_dn(self) -> pulumi.Output[str]:
        """
        A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        """
        return pulumi.get(self, "group_base_dn")

    @property
    @pulumi.getter(name="groupMemberAttribute")
    def group_member_attribute(self) -> pulumi.Output[str]:
        """
        A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
        """
        return pulumi.get(self, "group_member_attribute")

    @property
    @pulumi.getter(name="groupNameAttribute")
    def group_name_attribute(self) -> pulumi.Output[str]:
        """
        Attribute on the group entry denoting the group name. Used when importing groups.
        """
        return pulumi.get(self, "group_name_attribute")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Ldap group setting name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def strategy(self) -> pulumi.Output[str]:
        """
        The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
        """
        return pulumi.get(self, "strategy")

    @property
    @pulumi.getter(name="subTree")
    def sub_tree(self) -> pulumi.Output[bool]:
        """
        When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
        """
        return pulumi.get(self, "sub_tree")

