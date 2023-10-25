# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LdapGroupSettingArgs', 'LdapGroupSetting']

@pulumi.input_type
class LdapGroupSettingArgs:
    def __init__(__self__, *,
                 description_attribute: pulumi.Input[str],
                 filter: pulumi.Input[str],
                 group_member_attribute: pulumi.Input[str],
                 group_name_attribute: pulumi.Input[str],
                 ldap_setting_key: pulumi.Input[str],
                 strategy: pulumi.Input[str],
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sub_tree: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a LdapGroupSetting resource.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] ldap_setting_key: The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
               - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
               - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
               - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        """
        LdapGroupSettingArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description_attribute=description_attribute,
            filter=filter,
            group_member_attribute=group_member_attribute,
            group_name_attribute=group_name_attribute,
            ldap_setting_key=ldap_setting_key,
            strategy=strategy,
            group_base_dn=group_base_dn,
            name=name,
            sub_tree=sub_tree,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description_attribute: Optional[pulumi.Input[str]] = None,
             filter: Optional[pulumi.Input[str]] = None,
             group_member_attribute: Optional[pulumi.Input[str]] = None,
             group_name_attribute: Optional[pulumi.Input[str]] = None,
             ldap_setting_key: Optional[pulumi.Input[str]] = None,
             strategy: Optional[pulumi.Input[str]] = None,
             group_base_dn: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             sub_tree: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if description_attribute is None and 'descriptionAttribute' in kwargs:
            description_attribute = kwargs['descriptionAttribute']
        if description_attribute is None:
            raise TypeError("Missing 'description_attribute' argument")
        if filter is None:
            raise TypeError("Missing 'filter' argument")
        if group_member_attribute is None and 'groupMemberAttribute' in kwargs:
            group_member_attribute = kwargs['groupMemberAttribute']
        if group_member_attribute is None:
            raise TypeError("Missing 'group_member_attribute' argument")
        if group_name_attribute is None and 'groupNameAttribute' in kwargs:
            group_name_attribute = kwargs['groupNameAttribute']
        if group_name_attribute is None:
            raise TypeError("Missing 'group_name_attribute' argument")
        if ldap_setting_key is None and 'ldapSettingKey' in kwargs:
            ldap_setting_key = kwargs['ldapSettingKey']
        if ldap_setting_key is None:
            raise TypeError("Missing 'ldap_setting_key' argument")
        if strategy is None:
            raise TypeError("Missing 'strategy' argument")
        if group_base_dn is None and 'groupBaseDn' in kwargs:
            group_base_dn = kwargs['groupBaseDn']
        if sub_tree is None and 'subTree' in kwargs:
            sub_tree = kwargs['subTree']

        _setter("description_attribute", description_attribute)
        _setter("filter", filter)
        _setter("group_member_attribute", group_member_attribute)
        _setter("group_name_attribute", group_name_attribute)
        _setter("ldap_setting_key", ldap_setting_key)
        _setter("strategy", strategy)
        if group_base_dn is not None:
            _setter("group_base_dn", group_base_dn)
        if name is not None:
            _setter("name", name)
        if sub_tree is not None:
            _setter("sub_tree", sub_tree)

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
        A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
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
    @pulumi.getter(name="ldapSettingKey")
    def ldap_setting_key(self) -> pulumi.Input[str]:
        """
        The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        """
        return pulumi.get(self, "ldap_setting_key")

    @ldap_setting_key.setter
    def ldap_setting_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "ldap_setting_key", value)

    @property
    @pulumi.getter
    def strategy(self) -> pulumi.Input[str]:
        """
        The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
        - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
        - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
        - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        """
        return pulumi.get(self, "strategy")

    @strategy.setter
    def strategy(self, value: pulumi.Input[str]):
        pulumi.set(self, "strategy", value)

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
        When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        """
        return pulumi.get(self, "sub_tree")

    @sub_tree.setter
    def sub_tree(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sub_tree", value)


@pulumi.input_type
class _LdapGroupSettingState:
    def __init__(__self__, *,
                 description_attribute: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 group_member_attribute: Optional[pulumi.Input[str]] = None,
                 group_name_attribute: Optional[pulumi.Input[str]] = None,
                 ldap_setting_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[str]] = None,
                 sub_tree: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering LdapGroupSetting resources.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] ldap_setting_key: The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
               - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
               - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
               - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        """
        _LdapGroupSettingState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description_attribute=description_attribute,
            filter=filter,
            group_base_dn=group_base_dn,
            group_member_attribute=group_member_attribute,
            group_name_attribute=group_name_attribute,
            ldap_setting_key=ldap_setting_key,
            name=name,
            strategy=strategy,
            sub_tree=sub_tree,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description_attribute: Optional[pulumi.Input[str]] = None,
             filter: Optional[pulumi.Input[str]] = None,
             group_base_dn: Optional[pulumi.Input[str]] = None,
             group_member_attribute: Optional[pulumi.Input[str]] = None,
             group_name_attribute: Optional[pulumi.Input[str]] = None,
             ldap_setting_key: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             strategy: Optional[pulumi.Input[str]] = None,
             sub_tree: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if description_attribute is None and 'descriptionAttribute' in kwargs:
            description_attribute = kwargs['descriptionAttribute']
        if group_base_dn is None and 'groupBaseDn' in kwargs:
            group_base_dn = kwargs['groupBaseDn']
        if group_member_attribute is None and 'groupMemberAttribute' in kwargs:
            group_member_attribute = kwargs['groupMemberAttribute']
        if group_name_attribute is None and 'groupNameAttribute' in kwargs:
            group_name_attribute = kwargs['groupNameAttribute']
        if ldap_setting_key is None and 'ldapSettingKey' in kwargs:
            ldap_setting_key = kwargs['ldapSettingKey']
        if sub_tree is None and 'subTree' in kwargs:
            sub_tree = kwargs['subTree']

        if description_attribute is not None:
            _setter("description_attribute", description_attribute)
        if filter is not None:
            _setter("filter", filter)
        if group_base_dn is not None:
            _setter("group_base_dn", group_base_dn)
        if group_member_attribute is not None:
            _setter("group_member_attribute", group_member_attribute)
        if group_name_attribute is not None:
            _setter("group_name_attribute", group_name_attribute)
        if ldap_setting_key is not None:
            _setter("ldap_setting_key", ldap_setting_key)
        if name is not None:
            _setter("name", name)
        if strategy is not None:
            _setter("strategy", strategy)
        if sub_tree is not None:
            _setter("sub_tree", sub_tree)

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
        A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
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
    @pulumi.getter(name="ldapSettingKey")
    def ldap_setting_key(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        """
        return pulumi.get(self, "ldap_setting_key")

    @ldap_setting_key.setter
    def ldap_setting_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_setting_key", value)

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
        The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
        - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
        - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
        - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        """
        return pulumi.get(self, "strategy")

    @strategy.setter
    def strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "strategy", value)

    @property
    @pulumi.getter(name="subTree")
    def sub_tree(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        """
        return pulumi.get(self, "sub_tree")

    @sub_tree.setter
    def sub_tree(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sub_tree", value)


class LdapGroupSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description_attribute: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 group_member_attribute: Optional[pulumi.Input[str]] = None,
                 group_name_attribute: Optional[pulumi.Input[str]] = None,
                 ldap_setting_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strategy: Optional[pulumi.Input[str]] = None,
                 sub_tree: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resource can be used to manage Artifactory's LDAP Group settings for user authentication.

        LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
        structure for managing group-based permissions.

        ~>The `LdapGroupSetting` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory LDAP setting
        ldap_group_name = artifactory.LdapGroupSetting("ldapGroupName",
            description_attribute="description",
            filter="(objectClass=groupOfNames)",
            group_base_dn="",
            group_member_attribute="uniqueMember",
            group_name_attribute="cn",
            ldap_setting_key="ldap_name",
            strategy="STATIC",
            sub_tree=True)
        ```
        Note: `Name` argument has to match to the resource name.\\
        Reference Link: [JFrog LDAP](https://www.jfrog.com/confluence/display/JFROG/LDAP)

        ## Import

        LDAP Group setting can be imported using the key, e.g.

        ```sh
         $ pulumi import artifactory:index/ldapGroupSetting:LdapGroupSetting ldap_group_name ldap_group_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] ldap_setting_key: The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
               - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
               - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
               - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LdapGroupSettingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can be used to manage Artifactory's LDAP Group settings for user authentication.

        LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
        structure for managing group-based permissions.

        ~>The `LdapGroupSetting` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory LDAP setting
        ldap_group_name = artifactory.LdapGroupSetting("ldapGroupName",
            description_attribute="description",
            filter="(objectClass=groupOfNames)",
            group_base_dn="",
            group_member_attribute="uniqueMember",
            group_name_attribute="cn",
            ldap_setting_key="ldap_name",
            strategy="STATIC",
            sub_tree=True)
        ```
        Note: `Name` argument has to match to the resource name.\\
        Reference Link: [JFrog LDAP](https://www.jfrog.com/confluence/display/JFROG/LDAP)

        ## Import

        LDAP Group setting can be imported using the key, e.g.

        ```sh
         $ pulumi import artifactory:index/ldapGroupSetting:LdapGroupSetting ldap_group_name ldap_group_name
        ```

        :param str resource_name: The name of the resource.
        :param LdapGroupSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LdapGroupSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            LdapGroupSettingArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description_attribute: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 group_base_dn: Optional[pulumi.Input[str]] = None,
                 group_member_attribute: Optional[pulumi.Input[str]] = None,
                 group_name_attribute: Optional[pulumi.Input[str]] = None,
                 ldap_setting_key: Optional[pulumi.Input[str]] = None,
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
            __props__ = LdapGroupSettingArgs.__new__(LdapGroupSettingArgs)

            if description_attribute is None and not opts.urn:
                raise TypeError("Missing required property 'description_attribute'")
            __props__.__dict__["description_attribute"] = description_attribute
            if filter is None and not opts.urn:
                raise TypeError("Missing required property 'filter'")
            __props__.__dict__["filter"] = filter
            __props__.__dict__["group_base_dn"] = group_base_dn
            if group_member_attribute is None and not opts.urn:
                raise TypeError("Missing required property 'group_member_attribute'")
            __props__.__dict__["group_member_attribute"] = group_member_attribute
            if group_name_attribute is None and not opts.urn:
                raise TypeError("Missing required property 'group_name_attribute'")
            __props__.__dict__["group_name_attribute"] = group_name_attribute
            if ldap_setting_key is None and not opts.urn:
                raise TypeError("Missing required property 'ldap_setting_key'")
            __props__.__dict__["ldap_setting_key"] = ldap_setting_key
            __props__.__dict__["name"] = name
            if strategy is None and not opts.urn:
                raise TypeError("Missing required property 'strategy'")
            __props__.__dict__["strategy"] = strategy
            __props__.__dict__["sub_tree"] = sub_tree
        super(LdapGroupSetting, __self__).__init__(
            'artifactory:index/ldapGroupSetting:LdapGroupSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description_attribute: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            group_base_dn: Optional[pulumi.Input[str]] = None,
            group_member_attribute: Optional[pulumi.Input[str]] = None,
            group_name_attribute: Optional[pulumi.Input[str]] = None,
            ldap_setting_key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            strategy: Optional[pulumi.Input[str]] = None,
            sub_tree: Optional[pulumi.Input[bool]] = None) -> 'LdapGroupSetting':
        """
        Get an existing LdapGroupSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description_attribute: An attribute on the group entry which denoting the group description. Used when importing groups.
        :param pulumi.Input[str] filter: The LDAP filter used to search for group entries. Used for importing groups.
        :param pulumi.Input[str] group_base_dn: A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        :param pulumi.Input[str] group_member_attribute: A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
        :param pulumi.Input[str] group_name_attribute: Attribute on the group entry denoting the group name. Used when importing groups.
        :param pulumi.Input[str] ldap_setting_key: The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        :param pulumi.Input[str] name: Ldap group setting name.
        :param pulumi.Input[str] strategy: The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
               - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
               - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
               - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        :param pulumi.Input[bool] sub_tree: When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LdapGroupSettingState.__new__(_LdapGroupSettingState)

        __props__.__dict__["description_attribute"] = description_attribute
        __props__.__dict__["filter"] = filter
        __props__.__dict__["group_base_dn"] = group_base_dn
        __props__.__dict__["group_member_attribute"] = group_member_attribute
        __props__.__dict__["group_name_attribute"] = group_name_attribute
        __props__.__dict__["ldap_setting_key"] = ldap_setting_key
        __props__.__dict__["name"] = name
        __props__.__dict__["strategy"] = strategy
        __props__.__dict__["sub_tree"] = sub_tree
        return LdapGroupSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="descriptionAttribute")
    def description_attribute(self) -> pulumi.Output[str]:
        """
        An attribute on the group entry which denoting the group description. Used when importing groups.
        """
        return pulumi.get(self, "description_attribute")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        """
        The LDAP filter used to search for group entries. Used for importing groups.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="groupBaseDn")
    def group_base_dn(self) -> pulumi.Output[Optional[str]]:
        """
        A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
        """
        return pulumi.get(self, "group_base_dn")

    @property
    @pulumi.getter(name="groupMemberAttribute")
    def group_member_attribute(self) -> pulumi.Output[str]:
        """
        A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
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
    @pulumi.getter(name="ldapSettingKey")
    def ldap_setting_key(self) -> pulumi.Output[str]:
        """
        The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of the ldap group setting XML block of system configuration.
        """
        return pulumi.get(self, "ldap_setting_key")

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
        The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
        - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
        - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
        - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
        """
        return pulumi.get(self, "strategy")

    @property
    @pulumi.getter(name="subTree")
    def sub_tree(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
        """
        return pulumi.get(self, "sub_tree")

