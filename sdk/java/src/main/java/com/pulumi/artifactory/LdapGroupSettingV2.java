// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LdapGroupSettingV2Args;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LdapGroupSettingV2State;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory LDAP Setting resource.
 * 
 * This resource can be used to manage Artifactory&#39;s LDAP Group settings for user authentication.
 * 
 * LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
 * structure for managing group-based permissions.
 * 
 * [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/ldap-group-setting), [general documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/ldap).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.LdapGroupSettingV2;
 * import com.pulumi.artifactory.LdapGroupSettingV2Args;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var ldapGroupName = new LdapGroupSettingV2(&#34;ldapGroupName&#34;, LdapGroupSettingV2Args.builder()        
 *             .descriptionAttribute(&#34;description&#34;)
 *             .enabledLdap(&#34;ldap_name&#34;)
 *             .filter(&#34;(objectClass=groupOfNames)&#34;)
 *             .forceAttributeSearch(false)
 *             .groupBaseDn(&#34;CN=Users,DC=MyDomain,DC=com&#34;)
 *             .groupMemberAttribute(&#34;uniqueMember&#34;)
 *             .groupNameAttribute(&#34;cn&#34;)
 *             .strategy(&#34;STATIC&#34;)
 *             .subTree(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2 ldap ldapGroup1
 * ```
 * 
 */
@ResourceType(type="artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2")
public class LdapGroupSettingV2 extends com.pulumi.resources.CustomResource {
    /**
     * An attribute on the group entry which denoting the group description. Used when importing groups.
     * 
     */
    @Export(name="descriptionAttribute", type=String.class, parameters={})
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
    @Export(name="enabledLdap", type=String.class, parameters={})
    private Output<String> enabledLdap;

    /**
     * @return The LDAP setting key you want to use for group retrieval.
     * 
     */
    public Output<String> enabledLdap() {
        return this.enabledLdap;
    }
    /**
     * The LDAP filter used to search for group entries. Used for importing groups.
     * 
     */
    @Export(name="filter", type=String.class, parameters={})
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
    @Export(name="forceAttributeSearch", type=Boolean.class, parameters={})
    private Output<Boolean> forceAttributeSearch;

    /**
     * @return This attribute is used in very specific cases of LDAP group settings. Don&#39;t switch it to `false`, unless instructed by the JFrog support team. Default value is `false`.
     * 
     */
    public Output<Boolean> forceAttributeSearch() {
        return this.forceAttributeSearch;
    }
    /**
     * A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    @Export(name="groupBaseDn", type=String.class, parameters={})
    private Output<String> groupBaseDn;

    /**
     * @return A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    public Output<String> groupBaseDn() {
        return this.groupBaseDn;
    }
    /**
     * A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember, member).
     * 
     */
    @Export(name="groupMemberAttribute", type=String.class, parameters={})
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
    @Export(name="groupNameAttribute", type=String.class, parameters={})
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
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Ldap group setting name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas: STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN. DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member. HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, `uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org` indicates that `user1` belongs to two groups: `uk` and `developers`. Valid values are: `STATIC`, `DYNAMIC`, `HIERARCHICAL`, case sensitive, all caps.
     * 
     */
    @Export(name="strategy", type=String.class, parameters={})
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
    @Export(name="subTree", type=Boolean.class, parameters={})
    private Output<Boolean> subTree;

    /**
     * @return When set, enables deep search through the sub-tree of the LDAP URL + Search Base. `true` by default. `sub_tree` can be set to true only with `STATIC` or `DYNAMIC` strategy.
     * 
     */
    public Output<Boolean> subTree() {
        return this.subTree;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LdapGroupSettingV2(String name) {
        this(name, LdapGroupSettingV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LdapGroupSettingV2(String name, LdapGroupSettingV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LdapGroupSettingV2(String name, LdapGroupSettingV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2", name, args == null ? LdapGroupSettingV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LdapGroupSettingV2(String name, Output<String> id, @Nullable LdapGroupSettingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapGroupSettingV2:LdapGroupSettingV2", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static LdapGroupSettingV2 get(String name, Output<String> id, @Nullable LdapGroupSettingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LdapGroupSettingV2(name, id, state, options);
    }
}
