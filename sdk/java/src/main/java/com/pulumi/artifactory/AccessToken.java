// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.AccessTokenArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.AccessTokenState;
import com.pulumi.artifactory.outputs.AccessTokenAdminToken;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### S
 * ### Create a new Artifactory Access Token for an existing user
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var exisingUser = new AccessToken(&#34;exisingUser&#34;, AccessTokenArgs.builder()        
 *             .endDateRelative(&#34;5m&#34;)
 *             .username(&#34;existing-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * Note: This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
 * ### Create a new Artifactory User and Access token
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.User;
 * import com.pulumi.artifactory.UserArgs;
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var newUserUser = new User(&#34;newUserUser&#34;, UserArgs.builder()        
 *             .email(&#34;new_user@somewhere.com&#34;)
 *             .groups(&#34;readers&#34;)
 *             .build());
 * 
 *         var newUserAccessToken = new AccessToken(&#34;newUserAccessToken&#34;, AccessTokenArgs.builder()        
 *             .username(newUserUser.name())
 *             .endDateRelative(&#34;5m&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Creates a new token for groups
 * This creates a transient user called `temporary-user`.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var temporaryUser = new AccessToken(&#34;temporaryUser&#34;, AccessTokenArgs.builder()        
 *             .endDateRelative(&#34;1h&#34;)
 *             .groups(&#34;readers&#34;)
 *             .username(&#34;temporary-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create token with no expiry
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var noExpiry = new AccessToken(&#34;noExpiry&#34;, AccessTokenArgs.builder()        
 *             .endDateRelative(&#34;0s&#34;)
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
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var refreshable = new AccessToken(&#34;refreshable&#34;, AccessTokenArgs.builder()        
 *             .endDateRelative(&#34;1m&#34;)
 *             .groups(&#34;readers&#34;)
 *             .refreshable(true)
 *             .username(&#34;refreshable&#34;)
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
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
 * import com.pulumi.artifactory.inputs.AccessTokenAdminTokenArgs;
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
 *         var admin = new AccessToken(&#34;admin&#34;, AccessTokenArgs.builder()        
 *             .adminToken(AccessTokenAdminTokenArgs.builder()
 *                 .instanceId(&#34;&lt;instance id&gt;&#34;)
 *                 .build())
 *             .endDateRelative(&#34;1m&#34;)
 *             .username(&#34;admin&#34;)
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
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var audience = new AccessToken(&#34;audience&#34;, AccessTokenArgs.builder()        
 *             .audience(&#34;jfrt@*&#34;)
 *             .endDateRelative(&#34;1m&#34;)
 *             .refreshable(true)
 *             .username(&#34;audience&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Creates a token with a fixed end date
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var fixeddate = new AccessToken(&#34;fixeddate&#34;, AccessTokenArgs.builder()        
 *             .endDate(&#34;2018-01-01T01:02:03Z&#34;)
 *             .groups(&#34;readers&#34;)
 *             .username(&#34;fixeddate&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Rotate token after it expires
 * This example will generate a token that will expire in 1 hour.
 * 
 * If `pulumi up` is run before 1 hour, nothing changes.
 * One an hour has passed, `pulumi up` will generate a new token.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.time.Rotating;
 * import com.pulumi.time.RotatingArgs;
 * import com.pulumi.artifactory.AccessToken;
 * import com.pulumi.artifactory.AccessTokenArgs;
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
 *         var nowPlus1Hours = new Rotating(&#34;nowPlus1Hours&#34;, RotatingArgs.builder()        
 *             .rotationHours(&#34;1&#34;)
 *             .build());
 * 
 *         var rotating = new AccessToken(&#34;rotating&#34;, AccessTokenArgs.builder()        
 *             .username(&#34;rotating&#34;)
 *             .endDate(time_rotating.now_plus_1_hour().rotation_rfc3339())
 *             .groups(&#34;readers&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## References
 * 
 * - https://www.jfrog.com/confluence/display/ACC1X/Access+Tokens
 * - https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateToken
 * 
 * ## Import
 * 
 * Artifactory **does not** retain access tokens and cannot be imported into state.
 * 
 */
@ResourceType(type="artifactory:index/accessToken:AccessToken")
public class AccessToken extends com.pulumi.resources.CustomResource {
    /**
     * Returns the access token to authenciate to Artifactory
     * 
     */
    @Export(name="accessToken", refs={String.class}, tree="[0]")
    private Output<String> accessToken;

    /**
     * @return Returns the access token to authenciate to Artifactory
     * 
     */
    public Output<String> accessToken() {
        return this.accessToken;
    }
    /**
     * (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
     * 
     */
    @Export(name="adminToken", refs={AccessTokenAdminToken.class}, tree="[0]")
    private Output</* @Nullable */ AccessTokenAdminToken> adminToken;

    /**
     * @return (Optional) Specify the `instance_id` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. `admin_token` cannot be specified with `groups`.
     * 
     */
    public Output<Optional<AccessTokenAdminToken>> adminToken() {
        return Codegen.optional(this.adminToken);
    }
    /**
     * (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt@*&#34;` so the token to be accepted by all Artifactory instances.
     * 
     */
    @Export(name="audience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> audience;

    /**
     * @return (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt@*&#34;` so the token to be accepted by all Artifactory instances.
     * 
     */
    public Output<Optional<String>> audience() {
        return Codegen.optional(this.audience);
    }
    /**
     * (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    @Export(name="endDate", refs={String.class}, tree="[0]")
    private Output<String> endDate;

    /**
     * @return (Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    public Output<String> endDate() {
        return this.endDate;
    }
    /**
     * (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
     * 
     */
    @Export(name="endDateRelative", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endDateRelative;

    /**
     * @return (Optional) A relative duration for which the token is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
     * 
     */
    public Output<Optional<String>> endDateRelative() {
        return Codegen.optional(this.endDateRelative);
    }
    /**
     * (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
     * 
     */
    @Export(name="groups", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> groups;

    /**
     * @return (Optional) List of groups. The token is granted access based on the permissions of the groups. Specify `[&#34;*&#34;]` for all groups that the user belongs to. `groups` cannot be specified with `admin_token`.
     * 
     */
    public Output<Optional<List<String>>> groups() {
        return Codegen.optional(this.groups);
    }
    /**
     * Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
     * 
     */
    @Export(name="refreshToken", refs={String.class}, tree="[0]")
    private Output<String> refreshToken;

    /**
     * @return Returns the refresh token when `refreshable` is true, or an empty string when `refreshable` is false.
     * 
     */
    public Output<String> refreshToken() {
        return this.refreshToken;
    }
    /**
     * (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    @Export(name="refreshable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> refreshable;

    /**
     * @return (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> refreshable() {
        return Codegen.optional(this.refreshable);
    }
    /**
     * (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return (Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require `groups` to be set.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessToken(String name) {
        this(name, AccessTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessToken(String name, AccessTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessToken(String name, AccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/accessToken:AccessToken", name, args == null ? AccessTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessToken(String name, Output<String> id, @Nullable AccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/accessToken:AccessToken", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessToken",
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
    public static AccessToken get(String name, Output<String> id, @Nullable AccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessToken(name, id, state, options);
    }
}
