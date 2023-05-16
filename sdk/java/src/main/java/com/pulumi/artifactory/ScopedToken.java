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
 * ~&gt;Token would not be saved by Artifactory if `expires_in` is less than the persistency threshold value (default to 10800 seconds) set in Access configuration. See [Persistency Threshold](https://www.jfrog.com/confluence/display/JFROG/Access+Tokens#AccessTokens-PersistencyThreshold) for details.
 * 
 * ## Example Usage
 * 
 * ### S
 * ### Create a new Artifactory scoped token for an existing user
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
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
 *         var scopedToken = new ScopedToken(&#34;scopedToken&#34;, ScopedTokenArgs.builder()        
 *             .username(&#34;existing-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * **Note:** This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
 * ### Create a new Artifactory user and scoped token
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.User;
 * import com.pulumi.artifactory.UserArgs;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
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
 *         var newUser = new User(&#34;newUser&#34;, UserArgs.builder()        
 *             .email(&#34;new_user@somewhere.com&#34;)
 *             .groups(&#34;readers&#34;)
 *             .build());
 * 
 *         var scopedTokenUser = new ScopedToken(&#34;scopedTokenUser&#34;, ScopedTokenArgs.builder()        
 *             .username(newUser.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Creates a new token for groups
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
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
 *         var scopedTokenGroup = new ScopedToken(&#34;scopedTokenGroup&#34;, ScopedTokenArgs.builder()        
 *             .scopes(&#34;applied-permissions/groups:readers&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create token with expiry
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
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
 *         var scopedTokenNoExpiry = new ScopedToken(&#34;scopedTokenNoExpiry&#34;, ScopedTokenArgs.builder()        
 *             .expiresIn(7200)
 *             .username(&#34;existing-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Creates a refreshable token
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
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
 *         var scopedTokenRefreshable = new ScopedToken(&#34;scopedTokenRefreshable&#34;, ScopedTokenArgs.builder()        
 *             .refreshable(true)
 *             .username(&#34;existing-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Creates an administrator token
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
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
 *         var admin = new ScopedToken(&#34;admin&#34;, ScopedTokenArgs.builder()        
 *             .scopes(&#34;applied-permissions/admin&#34;)
 *             .username(&#34;admin-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Creates a token with an audience
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ScopedToken;
 * import com.pulumi.artifactory.ScopedTokenArgs;
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
 *         var audience = new ScopedToken(&#34;audience&#34;, ScopedTokenArgs.builder()        
 *             .audiences(&#34;jfrt@*&#34;)
 *             .scopes(&#34;applied-permissions/admin&#34;)
 *             .username(&#34;admin-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## References
 * 
 * - https://www.jfrog.com/confluence/display/JFROG/Access+Tokens
 * - https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-AccessTokens
 * 
 * ## Import
 * 
 * Artifactory **does not** retain scoped tokens and cannot be imported into state.
 * 
 */
@ResourceType(type="artifactory:index/scopedToken:ScopedToken")
public class ScopedToken extends com.pulumi.resources.CustomResource {
    /**
     * Returns the access token to authenticate to Artifactory
     * 
     */
    @Export(name="accessToken", type=String.class, parameters={})
    private Output<String> accessToken;

    /**
     * @return Returns the access token to authenticate to Artifactory
     * 
     */
    public Output<String> accessToken() {
        return this.accessToken;
    }
    /**
     * (Optional) A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to `*@*` if not set. Service ID must begin with `jfrt@`. For instructions to retrieve the Artifactory Service ID see this [documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetServiceID).
     * 
     */
    @Export(name="audiences", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> audiences;

    /**
     * @return (Optional) A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to `*@*` if not set. Service ID must begin with `jfrt@`. For instructions to retrieve the Artifactory Service ID see this [documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetServiceID).
     * 
     */
    public Output<Optional<List<String>>> audiences() {
        return Codegen.optional(this.audiences);
    }
    /**
     * (Optional) Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return (Optional) Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * (Optional) The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in `access.config.yaml`. See [API documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeTokenbyIDrevoketokenbyid) for details.
     * 
     */
    @Export(name="expiresIn", type=Integer.class, parameters={})
    private Output<Integer> expiresIn;

    /**
     * @return (Optional) The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in `access.config.yaml`. See [API documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeTokenbyIDrevoketokenbyid) for details.
     * 
     */
    public Output<Integer> expiresIn() {
        return this.expiresIn;
    }
    /**
     * Returns the token expiry
     * 
     */
    @Export(name="expiry", type=Integer.class, parameters={})
    private Output<Integer> expiry;

    /**
     * @return Returns the token expiry
     * 
     */
    public Output<Integer> expiry() {
        return this.expiry;
    }
    /**
     * (Optional) Should a reference token also be created? Defaults to `false`
     * 
     */
    @Export(name="includeReferenceToken", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> includeReferenceToken;

    /**
     * @return (Optional) Should a reference token also be created? Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> includeReferenceToken() {
        return Codegen.optional(this.includeReferenceToken);
    }
    /**
     * Returns the token issued at date/time
     * 
     */
    @Export(name="issuedAt", type=Integer.class, parameters={})
    private Output<Integer> issuedAt;

    /**
     * @return Returns the token issued at date/time
     * 
     */
    public Output<Integer> issuedAt() {
        return this.issuedAt;
    }
    /**
     * Returns the token issuer
     * 
     */
    @Export(name="issuer", type=String.class, parameters={})
    private Output<String> issuer;

    /**
     * @return Returns the token issuer
     * 
     */
    public Output<String> issuer() {
        return this.issuer;
    }
    /**
     * Returns the reference token to authenticate to Artifactory
     * 
     */
    @Export(name="referenceToken", type=String.class, parameters={})
    private Output<String> referenceToken;

    /**
     * @return Returns the reference token to authenticate to Artifactory
     * 
     */
    public Output<String> referenceToken() {
        return this.referenceToken;
    }
    @Export(name="refreshToken", type=String.class, parameters={})
    private Output<String> refreshToken;

    public Output<String> refreshToken() {
        return this.refreshToken;
    }
    /**
     * (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    @Export(name="refreshable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> refreshable;

    /**
     * @return (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> refreshable() {
        return Codegen.optional(this.refreshable);
    }
    /**
     * (Optional) The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
     * 
     */
    @Export(name="scopes", type=List.class, parameters={String.class})
    private Output<List<String>> scopes;

    /**
     * @return (Optional) The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
     * 
     */
    public Output<List<String>> scopes() {
        return this.scopes;
    }
    /**
     * Returns the token type
     * 
     */
    @Export(name="subject", type=String.class, parameters={})
    private Output<String> subject;

    /**
     * @return Returns the token type
     * 
     */
    public Output<String> subject() {
        return this.subject;
    }
    /**
     * Returns the token type
     * 
     */
    @Export(name="tokenType", type=String.class, parameters={})
    private Output<String> tokenType;

    /**
     * @return Returns the token type
     * 
     */
    public Output<String> tokenType() {
        return this.tokenType;
    }
    /**
     * (Optional) The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: `&lt;service-id&gt;/users/&lt;username&gt;`. Limited to 255 characters.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output</* @Nullable */ String> username;

    /**
     * @return (Optional) The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: `&lt;service-id&gt;/users/&lt;username&gt;`. Limited to 255 characters.
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
