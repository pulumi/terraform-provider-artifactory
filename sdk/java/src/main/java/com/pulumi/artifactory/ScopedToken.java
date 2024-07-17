// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ScopedTokenArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ScopedTokenState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory Scoped Token resource. This can be used to create and manage Artifactory Scoped Tokens.
 * 
 * !&gt;Scoped Tokens will be stored in the raw state as plain-text. Read more about sensitive data in
 * state.
 * 
 * ~&gt;Token would not be saved by Artifactory if `expires_in` is less than the persistency threshold value (default to 10800 seconds) set in Access configuration. See [Persistency Threshold](https://jfrog.com/help/r/jfrog-platform-administration-documentation/using-the-revocable-and-persistency-thresholds) for details.
 * 
 * ## Example Usage
 * 
 * ### S
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
 * import com.pulumi.artifactory.User;
 * import com.pulumi.artifactory.UserArgs;
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
 *         //## Create a new Artifactory scoped token for an existing user
 *         var scopedToken = new ScopedToken("scopedToken", ScopedTokenArgs.builder()
 *             .username("existing-user")
 *             .build());
 * 
 *         //## **Note:** This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
 *         //## Create a new Artifactory user and scoped token
 *         var newUser = new User("newUser", UserArgs.builder()
 *             .name("new_user")
 *             .email("new_user{@literal @}somewhere.com")
 *             .groups("readers")
 *             .build());
 * 
 *         var scopedTokenUser = new ScopedToken("scopedTokenUser", ScopedTokenArgs.builder()
 *             .username(newUser.name())
 *             .build());
 * 
 *         //## Creates a new token for groups
 *         var scopedTokenGroup = new ScopedToken("scopedTokenGroup", ScopedTokenArgs.builder()
 *             .scopes("applied-permissions/groups:readers")
 *             .build());
 * 
 *         //## Create token with expiry
 *         var scopedTokenNoExpiry = new ScopedToken("scopedTokenNoExpiry", ScopedTokenArgs.builder()
 *             .username("existing-user")
 *             .expiresIn(7200)
 *             .build());
 * 
 *         //## Creates a refreshable token
 *         var scopedTokenRefreshable = new ScopedToken("scopedTokenRefreshable", ScopedTokenArgs.builder()
 *             .username("existing-user")
 *             .refreshable(true)
 *             .build());
 * 
 *         //## Creates an administrator token
 *         var admin = new ScopedToken("admin", ScopedTokenArgs.builder()
 *             .username("admin-user")
 *             .scopes("applied-permissions/admin")
 *             .build());
 * 
 *         //## Creates a token with an audience
 *         var audience = new ScopedToken("audience", ScopedTokenArgs.builder()
 *             .username("admin-user")
 *             .scopes("applied-permissions/admin")
 *             .audiences("jfrt{@literal @}*")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## References
 * 
 * - https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-tokens
 * - https://jfrog.com/help/r/jfrog-rest-apis/access-tokens
 * 
 * ## Import
 * 
 * Artifactory **does not** retain scoped tokens, and they cannot be imported into state.
 * 
 */
@ResourceType(type="artifactory:index/scopedToken:ScopedToken")
public class ScopedToken extends com.pulumi.resources.CustomResource {
    /**
     * Returns the access token to authenticate to Artifactory.
     * 
     */
    @Export(name="accessToken", refs={String.class}, tree="[0]")
    private Output<String> accessToken;

    /**
     * @return Returns the access token to authenticate to Artifactory.
     * 
     */
    public Output<String> accessToken() {
        return this.accessToken;
    }
    /**
     * A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to
     * total 255 characters. Default to &#39;*{@literal @}*&#39; if not set. Service ID must begin with valid JFrog service type. Options: jfrt,
     * jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see
     * this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
     * 
     */
    @Export(name="audiences", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> audiences;

    /**
     * @return A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to
     * total 255 characters. Default to &#39;*{@literal @}*&#39; if not set. Service ID must begin with valid JFrog service type. Options: jfrt,
     * jfxr, jfpip, jfds, jfmc, jfac, jfevt, jfmd, jfcon, or *. For instructions to retrieve the Artifactory Service ID see
     * this [documentation](https://jfrog.com/help/r/jfrog-rest-apis/get-service-id)
     * 
     */
    public Output<Optional<List<String>>> audiences() {
        return Codegen.optional(this.audiences);
    }
    /**
     * Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is
     * mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is
     * based on configuration in &#39;access.config.yaml&#39;. See [API
     * documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved
     * by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access
     * configuration. See [official
     * documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/persistency-threshold) for details.
     * 
     */
    @Export(name="expiresIn", refs={Integer.class}, tree="[0]")
    private Output<Integer> expiresIn;

    /**
     * @return The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is
     * mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is
     * based on configuration in &#39;access.config.yaml&#39;. See [API
     * documentation](https://jfrog.com/help/r/jfrog-rest-apis/revoke-token-by-id) for details. Access Token would not be saved
     * by Artifactory if this is less than the persistence threshold value (default to 10800 seconds) set in Access
     * configuration. See [official
     * documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/persistency-threshold) for details.
     * 
     */
    public Output<Integer> expiresIn() {
        return this.expiresIn;
    }
    /**
     * Returns the token expiry.
     * 
     */
    @Export(name="expiry", refs={Integer.class}, tree="[0]")
    private Output<Integer> expiry;

    /**
     * @return Returns the token expiry.
     * 
     */
    public Output<Integer> expiry() {
        return this.expiry;
    }
    /**
     * The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is
     * also the default value if this parameter is not specified.
     * 
     */
    @Export(name="grantType", refs={String.class}, tree="[0]")
    private Output<String> grantType;

    /**
     * @return The grant type used to authenticate the request. In this case, the only value supported is `client_credentials` which is
     * also the default value if this parameter is not specified.
     * 
     */
    public Output<String> grantType() {
        return this.grantType;
    }
    /**
     * Toggle to ignore warning message when token was missing or not created and stored by Artifactory. Default is `false`.
     * 
     */
    @Export(name="ignoreMissingTokenWarning", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ignoreMissingTokenWarning;

    /**
     * @return Toggle to ignore warning message when token was missing or not created and stored by Artifactory. Default is `false`.
     * 
     */
    public Output<Boolean> ignoreMissingTokenWarning() {
        return this.ignoreMissingTokenWarning;
    }
    /**
     * Also create a reference token which can be used like an API key. Default is `false`.
     * 
     */
    @Export(name="includeReferenceToken", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> includeReferenceToken;

    /**
     * @return Also create a reference token which can be used like an API key. Default is `false`.
     * 
     */
    public Output<Boolean> includeReferenceToken() {
        return this.includeReferenceToken;
    }
    /**
     * Returns the token issued at date/time.
     * 
     */
    @Export(name="issuedAt", refs={Integer.class}, tree="[0]")
    private Output<Integer> issuedAt;

    /**
     * @return Returns the token issued at date/time.
     * 
     */
    public Output<Integer> issuedAt() {
        return this.issuedAt;
    }
    /**
     * Returns the token issuer.
     * 
     */
    @Export(name="issuer", refs={String.class}, tree="[0]")
    private Output<String> issuer;

    /**
     * @return Returns the token issuer.
     * 
     */
    public Output<String> issuer() {
        return this.issuer;
    }
    /**
     * The project for which this token is created. Enter the project name on which you want to apply this token.
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return The project for which this token is created. Enter the project name on which you want to apply this token.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * Reference Token (alias to Access Token).
     * 
     */
    @Export(name="referenceToken", refs={String.class}, tree="[0]")
    private Output<String> referenceToken;

    /**
     * @return Reference Token (alias to Access Token).
     * 
     */
    public Output<String> referenceToken() {
        return this.referenceToken;
    }
    /**
     * Refresh token.
     * 
     */
    @Export(name="refreshToken", refs={String.class}, tree="[0]")
    private Output<String> refreshToken;

    /**
     * @return Refresh token.
     * 
     */
    public Output<String> refreshToken() {
        return this.refreshToken;
    }
    /**
     * Is this token refreshable? Default is `false`.
     * 
     */
    @Export(name="refreshable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> refreshable;

    /**
     * @return Is this token refreshable? Default is `false`.
     * 
     */
    public Output<Boolean> refreshable() {
        return this.refreshable;
    }
    /**
     * The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can
     * set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong. The
     * supported scopes include: - `applied-permissions/user` - provides user access. If left at the default setting, the token
     * will be created with the user-identity scope, which allows users to identify themselves in the Platform but does not
     * grant any specific access permissions. - `applied-permissions/admin` - the scope assigned to admin users. -
     * `applied-permissions/groups` - this scope assigns permissions to groups using the following format:
     * `applied-permissions/groups:&lt;group-name&gt;[,&lt;group-name&gt;...]` - `system:metrics:r` - for getting the service metrics -
     * `system:livelogs:r` - for getting the service livelogs - Resource Permissions: From Artifactory 7.38.x, resource
     * permissions scoped tokens are also supported in the REST API. A permission can be represented as a scope token string in
     * the following format: `&lt;resource-type&gt;:&lt;target&gt;[/&lt;sub-resource&gt;]:&lt;actions&gt;` - Where: - `&lt;resource-type&gt;` - one of the
     * permission resource types, from a predefined closed list. Currently, the only resource type that is supported is the
     * artifact resource type. - `&lt;target&gt;` - the target resource, can be exact name or a pattern - `&lt;sub-resource&gt;` -
     * optional, the target sub-resource, can be exact name or a pattern - `&lt;actions&gt;` - comma-separated list of action
     * acronyms. The actions allowed are `r`, `w`, `d`, `a`, `m`, `x`, `s`, or any combination of these actions. To allow all
     * actions - use `*` - Examples: - `[&#34;applied-permissions/user&#34;, &#34;artifact:generic-local:r&#34;]` -
     * `[&#34;applied-permissions/group&#34;, &#34;artifact:generic-local/path:*&#34;]` - `[&#34;applied-permissions/admin&#34;, &#34;system:metrics:r&#34;,
     * &#34;artifact:generic-local:*&#34;]` - `applied-permissions/roles:project-key` - provides access to elements associated with the
     * project based on the project role. For example, `applied-permissions/roles:project-type:developer,qa`. -&gt;The scope to
     * assign to the token should be provided as a list of scope tokens, limited to 500 characters in total. From Artifactory
     * 7.84.3, [project
     * admins](https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-token-creation-by-project-admins)
     * can create access tokens that are tied to the projects in which they hold administrative privileges.
     * 
     */
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> scopes;

    /**
     * @return The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can
     * set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong. The
     * supported scopes include: - `applied-permissions/user` - provides user access. If left at the default setting, the token
     * will be created with the user-identity scope, which allows users to identify themselves in the Platform but does not
     * grant any specific access permissions. - `applied-permissions/admin` - the scope assigned to admin users. -
     * `applied-permissions/groups` - this scope assigns permissions to groups using the following format:
     * `applied-permissions/groups:&lt;group-name&gt;[,&lt;group-name&gt;...]` - `system:metrics:r` - for getting the service metrics -
     * `system:livelogs:r` - for getting the service livelogs - Resource Permissions: From Artifactory 7.38.x, resource
     * permissions scoped tokens are also supported in the REST API. A permission can be represented as a scope token string in
     * the following format: `&lt;resource-type&gt;:&lt;target&gt;[/&lt;sub-resource&gt;]:&lt;actions&gt;` - Where: - `&lt;resource-type&gt;` - one of the
     * permission resource types, from a predefined closed list. Currently, the only resource type that is supported is the
     * artifact resource type. - `&lt;target&gt;` - the target resource, can be exact name or a pattern - `&lt;sub-resource&gt;` -
     * optional, the target sub-resource, can be exact name or a pattern - `&lt;actions&gt;` - comma-separated list of action
     * acronyms. The actions allowed are `r`, `w`, `d`, `a`, `m`, `x`, `s`, or any combination of these actions. To allow all
     * actions - use `*` - Examples: - `[&#34;applied-permissions/user&#34;, &#34;artifact:generic-local:r&#34;]` -
     * `[&#34;applied-permissions/group&#34;, &#34;artifact:generic-local/path:*&#34;]` - `[&#34;applied-permissions/admin&#34;, &#34;system:metrics:r&#34;,
     * &#34;artifact:generic-local:*&#34;]` - `applied-permissions/roles:project-key` - provides access to elements associated with the
     * project based on the project role. For example, `applied-permissions/roles:project-type:developer,qa`. -&gt;The scope to
     * assign to the token should be provided as a list of scope tokens, limited to 500 characters in total. From Artifactory
     * 7.84.3, [project
     * admins](https://jfrog.com/help/r/jfrog-platform-administration-documentation/access-token-creation-by-project-admins)
     * can create access tokens that are tied to the projects in which they hold administrative privileges.
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * Returns the token type.
     * 
     */
    @Export(name="subject", refs={String.class}, tree="[0]")
    private Output<String> subject;

    /**
     * @return Returns the token type.
     * 
     */
    public Output<String> subject() {
        return this.subject;
    }
    /**
     * Returns the token type.
     * 
     */
    @Export(name="tokenType", refs={String.class}, tree="[0]")
    private Output<String> tokenType;

    /**
     * @return Returns the token type.
     * 
     */
    public Output<String> tokenType() {
        return this.tokenType;
    }
    /**
     * The user name for which this token is created. The username is based on the authenticated user - either from the user of
     * the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject
     * of the token: &lt;service-id&gt;/users/&lt;username&gt;. Limited to 255 characters.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> username;

    /**
     * @return The user name for which this token is created. The username is based on the authenticated user - either from the user of
     * the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject
     * of the token: &lt;service-id&gt;/users/&lt;username&gt;. Limited to 255 characters.
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ScopedToken(String name) {
        this(name, ScopedTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ScopedToken(String name, @Nullable ScopedTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ScopedToken(String name, @Nullable ScopedTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/scopedToken:ScopedToken", name, args == null ? ScopedTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ScopedToken(String name, Output<String> id, @Nullable ScopedTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/scopedToken:ScopedToken", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessToken",
                "referenceToken",
                "refreshToken"
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
    public static ScopedToken get(String name, Output<String> id, @Nullable ScopedTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ScopedToken(name, id, state, options);
    }
}
