// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LdapGroupSettingArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LdapGroupSettingState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can be used to manage Artifactory&#39;s LDAP Group settings for user authentication.
 * 
 * LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
 * structure for managing group-based permissions.
 * 
 * ~&gt;The `artifactory.LdapGroupSetting` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.LdapGroupSetting;
 * import com.pulumi.artifactory.LdapGroupSettingArgs;
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
 *         // Configure Artifactory LDAP setting
 *         var ldapGroupName = new LdapGroupSetting(&#34;ldapGroupName&#34;, LdapGroupSettingArgs.builder()        
 *             .descriptionAttribute(&#34;description&#34;)
 *             .filter(&#34;(objectClass=groupOfNames)&#34;)
 *             .groupBaseDn(&#34;&#34;)
 *             .groupMemberAttribute(&#34;uniqueMember&#34;)
 *             .groupNameAttribute(&#34;cn&#34;)
 *             .ldapSettingKey(&#34;ldap_name&#34;)
 *             .strategy(&#34;STATIC&#34;)
 *             .subTree(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * Note: `Name` argument has to match to the resource name.\
 * Reference Link: [JFrog LDAP](https://www.jfrog.com/confluence/display/JFROG/LDAP)
 * 
 * ## Import
 * 
 * LDAP Group setting can be imported using the key, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/ldapGroupSetting:LdapGroupSetting ldap_group_name ldap_group_name
 * ```
 * 
 */
@ResourceType(type="artifactory:index/ldapGroupSetting:LdapGroupSetting")
public class LdapGroupSetting extends com.pulumi.resources.CustomResource {
    /**
     * An attribute on the group entry which denoting the group description. Used when importing groups.
     * 
     */
    @Export(name="descriptionAttribute", refs={String.class}, tree="[0]")
    private Output<String> descriptionAttribute;

    /**
     * @return An attribute on the group entry which denoting the group description. Used when importing groups.
     * 
     */
    public Output<String> descriptionAttribute() {
        return this.descriptionAttribute;
    }
    /**
     * The LDAP filter used to search for group entries. Used for importing groups.
     * 
     */
    @Export(name="filter", refs={String.class}, tree="[0]")
    private Output<String> filter;

    /**
     * @return The LDAP filter used to search for group entries. Used for importing groups.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }
    /**
     * A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    @Export(name="groupBaseDn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupBaseDn;

    /**
     * @return A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    public Output<Optional<String>> groupBaseDn() {
        return Codegen.optional(this.groupBaseDn);
    }
    /**
     * A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
     * 
     */
    @Export(name="groupMemberAttribute", refs={String.class}, tree="[0]")
    private Output<String> groupMemberAttribute;

    /**
     * @return A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
     * 
     */
    public Output<String> groupMemberAttribute() {
        return this.groupMemberAttribute;
    }
    /**
     * Attribute on the group entry denoting the group name. Used when importing groups.
     * 
     */
    @Export(name="groupNameAttribute", refs={String.class}, tree="[0]")
    private Output<String> groupNameAttribute;

    /**
     * @return Attribute on the group entry denoting the group name. Used when importing groups.
     * 
     */
    public Output<String> groupNameAttribute() {
        return this.groupNameAttribute;
    }
    /**
     * The LDAP setting key you want to use for group retrieval. The value for this field corresponds to &#39;enabledLdap&#39; field of the ldap group setting XML block of system configuration.
     * 
     */
    @Export(name="ldapSettingKey", refs={String.class}, tree="[0]")
    private Output<String> ldapSettingKey;

    /**
     * @return The LDAP setting key you want to use for group retrieval. The value for this field corresponds to &#39;enabledLdap&#39; field of the ldap group setting XML block of system configuration.
     * 
     */
    public Output<String> ldapSettingKey() {
        return this.ldapSettingKey;
    }
    /**
     * Ldap group setting name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Ldap group setting name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
     * - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
     * - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
     * - HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
     * 
     */
    @Export(name="strategy", refs={String.class}, tree="[0]")
    private Output<String> strategy;

    /**
     * @return The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
     * - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
     * - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
     * - HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
     * 
     */
    public Output<String> strategy() {
        return this.strategy;
    }
    /**
     * When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
     * 
     */
    @Export(name="subTree", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> subTree;

    /**
     * @return When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
     * 
     */
    public Output<Optional<Boolean>> subTree() {
        return Codegen.optional(this.subTree);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LdapGroupSetting(String name) {
        this(name, LdapGroupSettingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LdapGroupSetting(String name, LdapGroupSettingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LdapGroupSetting(String name, LdapGroupSettingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapGroupSetting:LdapGroupSetting", name, args == null ? LdapGroupSettingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LdapGroupSetting(String name, Output<String> id, @Nullable LdapGroupSettingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapGroupSetting:LdapGroupSetting", name, state, makeResourceOptions(options, id));
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
    public static LdapGroupSetting get(String name, Output<String> id, @Nullable LdapGroupSettingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LdapGroupSetting(name, id, state, options);
    }
}
