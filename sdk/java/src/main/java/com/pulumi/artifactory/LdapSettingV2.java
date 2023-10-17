// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LdapSettingV2Args;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LdapSettingV2State;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory LDAP Setting resource.
 * 
 * This resource can be used to manage Artifactory&#39;s LDAP settings for user authentication.
 * 
 * When specified LDAP setting is active, Artifactory first attempts to authenticate the user against the LDAP server.
 * If LDAP authentication fails, it then tries to authenticate via its internal database.
 * 
 * [API documentation](https://jfrog.com/help/r/jfrog-rest-apis/ldap-setting), [general documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/ldap).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.LdapSettingV2;
 * import com.pulumi.artifactory.LdapSettingV2Args;
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
 *         var ldapName = new LdapSettingV2(&#34;ldapName&#34;, LdapSettingV2Args.builder()        
 *             .allowUserToAccessProfile(false)
 *             .autoCreateUser(true)
 *             .emailAttribute(&#34;mail&#34;)
 *             .enabled(true)
 *             .key(&#34;ldap_name&#34;)
 *             .ldapPoisoningProtection(true)
 *             .ldapUrl(&#34;ldap://ldap_server_url&#34;)
 *             .managerDn(&#34;mgr_dn&#34;)
 *             .managerPassword(&#34;mgr_passwd_random&#34;)
 *             .pagingSupportEnabled(false)
 *             .searchBase(&#34;ou=users&#34;)
 *             .searchFilter(&#34;(uid={0})&#34;)
 *             .searchSubTree(true)
 *             .userDnPattern(&#34;uid={0},ou=People&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import artifactory:index/ldapSettingV2:LdapSettingV2 ldap ldap1
 * ```
 * 
 */
@ResourceType(type="artifactory:index/ldapSettingV2:LdapSettingV2")
public class LdapSettingV2 extends com.pulumi.resources.CustomResource {
    /**
     * Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
     * 
     */
    @Export(name="allowUserToAccessProfile", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowUserToAccessProfile;

    /**
     * @return Auto created users will have access to their profile page and will be able to perform actions such as generating an API key. Default value is `false`.
     * 
     */
    public Output<Boolean> allowUserToAccessProfile() {
        return this.allowUserToAccessProfile;
    }
    /**
     * When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
     * 
     */
    @Export(name="autoCreateUser", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> autoCreateUser;

    /**
     * @return When set, users are automatically created when using LDAP. Otherwise, users are transient and associated with auto-join groups defined in Artifactory. Default value is `true`.
     * 
     */
    public Output<Boolean> autoCreateUser() {
        return this.autoCreateUser;
    }
    /**
     * An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default value is`mail`.
     * 
     */
    @Export(name="emailAttribute", refs={String.class}, tree="[0]")
    private Output<String> emailAttribute;

    /**
     * @return An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default value is`mail`.
     * 
     */
    public Output<String> emailAttribute() {
        return this.emailAttribute;
    }
    /**
     * Flag to enable or disable the ldap setting. Default value is `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Flag to enable or disable the ldap setting. Default value is `true`.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Ldap setting name.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return Ldap setting name.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * When this is set to `true`, an empty or missing usernames array will detach all users from the group.
     * 
     */
    @Export(name="ldapPoisoningProtection", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ldapPoisoningProtection;

    /**
     * @return When this is set to `true`, an empty or missing usernames array will detach all users from the group.
     * 
     */
    public Output<Boolean> ldapPoisoningProtection() {
        return this.ldapPoisoningProtection;
    }
    /**
     * Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
     * 
     */
    @Export(name="ldapUrl", refs={String.class}, tree="[0]")
    private Output<String> ldapUrl;

    /**
     * @return Location of the LDAP server in the following format: `ldap://myldapserver/dc=sampledomain,dc=com`
     * 
     */
    public Output<String> ldapUrl() {
        return this.ldapUrl;
    }
    /**
     * The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
     * 
     */
    @Export(name="managerDn", refs={String.class}, tree="[0]")
    private Output<String> managerDn;

    /**
     * @return The full DN of the user that binds to the LDAP server to perform user searches. Only used with `search` authentication.
     * 
     */
    public Output<String> managerDn() {
        return this.managerDn;
    }
    /**
     * The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
     * 
     */
    @Export(name="managerPassword", refs={String.class}, tree="[0]")
    private Output<String> managerPassword;

    /**
     * @return The password of the user that binds to the LDAP server to perform the search. Only used with `search` authentication.
     * 
     */
    public Output<String> managerPassword() {
        return this.managerPassword;
    }
    /**
     * When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
     * 
     */
    @Export(name="pagingSupportEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> pagingSupportEnabled;

    /**
     * @return When set, supports paging results for the LDAP server. This feature requires that the LDAP server supports a PagedResultsControl configuration. Default value is `true`.
     * 
     */
    public Output<Boolean> pagingSupportEnabled() {
        return this.pagingSupportEnabled;
    }
    /**
     * A context name to search in relative to the base DN of the LDAP URL. For example, &#39;ou=users&#39; With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe (&#39;|&#39;) character.
     * 
     */
    @Export(name="searchBase", refs={String.class}, tree="[0]")
    private Output<String> searchBase;

    /**
     * @return A context name to search in relative to the base DN of the LDAP URL. For example, &#39;ou=users&#39; With the LDAP Group Add-on enabled, it is possible to enter multiple search base entries separated by a pipe (&#39;|&#39;) character.
     * 
     */
    public Output<String> searchBase() {
        return this.searchBase;
    }
    /**
     * A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, and is denoted by &#39;{0}&#39;. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
     * 
     */
    @Export(name="searchFilter", refs={String.class}, tree="[0]")
    private Output<String> searchFilter;

    /**
     * @return A filter expression used to search for the user DN used in LDAP authentication. This is an LDAP search filter (as defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, and is denoted by &#39;{0}&#39;. Possible examples are: (uid={0}) - This searches for a username match on the attribute. Authentication to LDAP is performed from the DN found if successful.
     * 
     */
    public Output<String> searchFilter() {
        return this.searchFilter;
    }
    /**
     * When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
     * 
     */
    @Export(name="searchSubTree", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> searchSubTree;

    /**
     * @return When set, enables deep search through the sub tree of the LDAP URL + search base. Default value is `true`.
     * 
     */
    public Output<Boolean> searchSubTree() {
        return this.searchSubTree;
    }
    /**
     * A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for &#39;direct&#39; user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
     * 
     */
    @Export(name="userDnPattern", refs={String.class}, tree="[0]")
    private Output<String> userDnPattern;

    /**
     * @return A DN pattern that can be used to log users directly in to LDAP. This pattern is used to create a DN string for &#39;direct&#39; user authentication where the pattern is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username. This only works if anonymous binding is allowed and a direct user DN can be used, which is not the default case for Active Directory (use User DN search filter instead). Example: uid={0},ou=People. Default value is blank/empty.
     * 
     */
    public Output<String> userDnPattern() {
        return this.userDnPattern;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LdapSettingV2(String name) {
        this(name, LdapSettingV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LdapSettingV2(String name, LdapSettingV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LdapSettingV2(String name, LdapSettingV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapSettingV2:LdapSettingV2", name, args == null ? LdapSettingV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LdapSettingV2(String name, Output<String> id, @Nullable LdapSettingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapSettingV2:LdapSettingV2", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "managerPassword"
            ))
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
    public static LdapSettingV2 get(String name, Output<String> id, @Nullable LdapSettingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LdapSettingV2(name, id, state, options);
    }
}
