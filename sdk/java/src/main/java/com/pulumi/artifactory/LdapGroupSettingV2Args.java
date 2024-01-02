// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LdapGroupSettingV2Args extends com.pulumi.resources.ResourceArgs {

    public static final LdapGroupSettingV2Args Empty = new LdapGroupSettingV2Args();

    /**
     * An attribute on the group entry which denoting the group description. Used when importing groups.
     * 
     */
    @Import(name="descriptionAttribute", required=true)
    private Output<String> descriptionAttribute;

    /**
     * @return An attribute on the group entry which denoting the group description. Used when importing groups.
     * 
     */
    public Output<String> descriptionAttribute() {
        return this.descriptionAttribute;
    }

    /**
     * The LDAP setting key you want to use for group retrieval.
     * 
     */
    @Import(name="enabledLdap")
    private @Nullable Output<String> enabledLdap;

    /**
     * @return The LDAP setting key you want to use for group retrieval.
     * 
     */
    public Optional<Output<String>> enabledLdap() {
        return Optional.ofNullable(this.enabledLdap);
    }

    /**
     * The LDAP filter used to search for group entries. Used for importing groups.
     * 
     */
    @Import(name="filter", required=true)
    private Output<String> filter;

    /**
     * @return The LDAP filter used to search for group entries. Used for importing groups.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }

    /**
     * This attribute is used in very specific cases of LDAP group settings. Don&#39;t switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
     * 
     */
    @Import(name="forceAttributeSearch")
    private @Nullable Output<Boolean> forceAttributeSearch;

    /**
     * @return This attribute is used in very specific cases of LDAP group settings. Don&#39;t switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> forceAttributeSearch() {
        return Optional.ofNullable(this.forceAttributeSearch);
    }

    /**
     * A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    @Import(name="groupBaseDn")
    private @Nullable Output<String> groupBaseDn;

    /**
     * @return A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    public Optional<Output<String>> groupBaseDn() {
        return Optional.ofNullable(this.groupBaseDn);
    }

    /**
     * A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
     * 
     */
    @Import(name="groupMemberAttribute", required=true)
    private Output<String> groupMemberAttribute;

    /**
     * @return A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
     * 
     */
    public Output<String> groupMemberAttribute() {
        return this.groupMemberAttribute;
    }

    /**
     * Attribute on the group entry denoting the group name. Used when importing groups.
     * 
     */
    @Import(name="groupNameAttribute", required=true)
    private Output<String> groupNameAttribute;

    /**
     * @return Attribute on the group entry denoting the group name. Used when importing groups.
     * 
     */
    public Output<String> groupNameAttribute() {
        return this.groupNameAttribute;
    }

    /**
     * Ldap group setting name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Ldap group setting name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
     * 
     */
    @Import(name="strategy", required=true)
    private Output<String> strategy;

    /**
     * @return The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
     * 
     */
    public Output<String> strategy() {
        return this.strategy;
    }

    /**
     * When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
     * 
     */
    @Import(name="subTree")
    private @Nullable Output<Boolean> subTree;

    /**
     * @return When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
     * 
     */
    public Optional<Output<Boolean>> subTree() {
        return Optional.ofNullable(this.subTree);
    }

    private LdapGroupSettingV2Args() {}

    private LdapGroupSettingV2Args(LdapGroupSettingV2Args $) {
        this.descriptionAttribute = $.descriptionAttribute;
        this.enabledLdap = $.enabledLdap;
        this.filter = $.filter;
        this.forceAttributeSearch = $.forceAttributeSearch;
        this.groupBaseDn = $.groupBaseDn;
        this.groupMemberAttribute = $.groupMemberAttribute;
        this.groupNameAttribute = $.groupNameAttribute;
        this.name = $.name;
        this.strategy = $.strategy;
        this.subTree = $.subTree;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LdapGroupSettingV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LdapGroupSettingV2Args $;

        public Builder() {
            $ = new LdapGroupSettingV2Args();
        }

        public Builder(LdapGroupSettingV2Args defaults) {
            $ = new LdapGroupSettingV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param descriptionAttribute An attribute on the group entry which denoting the group description. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder descriptionAttribute(Output<String> descriptionAttribute) {
            $.descriptionAttribute = descriptionAttribute;
            return this;
        }

        /**
         * @param descriptionAttribute An attribute on the group entry which denoting the group description. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder descriptionAttribute(String descriptionAttribute) {
            return descriptionAttribute(Output.of(descriptionAttribute));
        }

        /**
         * @param enabledLdap The LDAP setting key you want to use for group retrieval.
         * 
         * @return builder
         * 
         */
        public Builder enabledLdap(@Nullable Output<String> enabledLdap) {
            $.enabledLdap = enabledLdap;
            return this;
        }

        /**
         * @param enabledLdap The LDAP setting key you want to use for group retrieval.
         * 
         * @return builder
         * 
         */
        public Builder enabledLdap(String enabledLdap) {
            return enabledLdap(Output.of(enabledLdap));
        }

        /**
         * @param filter The LDAP filter used to search for group entries. Used for importing groups.
         * 
         * @return builder
         * 
         */
        public Builder filter(Output<String> filter) {
            $.filter = filter;
            return this;
        }

        /**
         * @param filter The LDAP filter used to search for group entries. Used for importing groups.
         * 
         * @return builder
         * 
         */
        public Builder filter(String filter) {
            return filter(Output.of(filter));
        }

        /**
         * @param forceAttributeSearch This attribute is used in very specific cases of LDAP group settings. Don&#39;t switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder forceAttributeSearch(@Nullable Output<Boolean> forceAttributeSearch) {
            $.forceAttributeSearch = forceAttributeSearch;
            return this;
        }

        /**
         * @param forceAttributeSearch This attribute is used in very specific cases of LDAP group settings. Don&#39;t switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder forceAttributeSearch(Boolean forceAttributeSearch) {
            return forceAttributeSearch(Output.of(forceAttributeSearch));
        }

        /**
         * @param groupBaseDn A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupBaseDn(@Nullable Output<String> groupBaseDn) {
            $.groupBaseDn = groupBaseDn;
            return this;
        }

        /**
         * @param groupBaseDn A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupBaseDn(String groupBaseDn) {
            return groupBaseDn(Output.of(groupBaseDn));
        }

        /**
         * @param groupMemberAttribute A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
         * 
         * @return builder
         * 
         */
        public Builder groupMemberAttribute(Output<String> groupMemberAttribute) {
            $.groupMemberAttribute = groupMemberAttribute;
            return this;
        }

        /**
         * @param groupMemberAttribute A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
         * 
         * @return builder
         * 
         */
        public Builder groupMemberAttribute(String groupMemberAttribute) {
            return groupMemberAttribute(Output.of(groupMemberAttribute));
        }

        /**
         * @param groupNameAttribute Attribute on the group entry denoting the group name. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupNameAttribute(Output<String> groupNameAttribute) {
            $.groupNameAttribute = groupNameAttribute;
            return this;
        }

        /**
         * @param groupNameAttribute Attribute on the group entry denoting the group name. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupNameAttribute(String groupNameAttribute) {
            return groupNameAttribute(Output.of(groupNameAttribute));
        }

        /**
         * @param name Ldap group setting name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Ldap group setting name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param strategy The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
         * 
         * @return builder
         * 
         */
        public Builder strategy(Output<String> strategy) {
            $.strategy = strategy;
            return this;
        }

        /**
         * @param strategy The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
         * 
         * @return builder
         * 
         */
        public Builder strategy(String strategy) {
            return strategy(Output.of(strategy));
        }

        /**
         * @param subTree When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
         * 
         * @return builder
         * 
         */
        public Builder subTree(@Nullable Output<Boolean> subTree) {
            $.subTree = subTree;
            return this;
        }

        /**
         * @param subTree When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
         * 
         * @return builder
         * 
         */
        public Builder subTree(Boolean subTree) {
            return subTree(Output.of(subTree));
        }

        public LdapGroupSettingV2Args build() {
            if ($.descriptionAttribute == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingV2Args", "descriptionAttribute");
            }
            if ($.filter == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingV2Args", "filter");
            }
            if ($.groupMemberAttribute == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingV2Args", "groupMemberAttribute");
            }
            if ($.groupNameAttribute == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingV2Args", "groupNameAttribute");
            }
            if ($.strategy == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingV2Args", "strategy");
            }
            return $;
        }
    }

}
