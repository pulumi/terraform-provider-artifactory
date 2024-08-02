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
 * !&gt; **Warning:** This resource is being deprecated and replaced by `artifactory.ScopedToken` since v6.8.0.
 * 
 * Provides an Artifactory Access Token resource. This can be used to create and manage Artifactory Access Tokens.
 * 
 * &gt; **Note:** Access Tokens will be stored in the raw state as plain-text. Read more about sensitive data in
 * state.
 * 
 * ## Example Usage
 * 
 * ### S
 * ### Create a new Artifactory Access Token for an existing user
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var exisingUser = new AccessToken("exisingUser", AccessTokenArgs.builder()
 *             .username("existing-user")
 *             .endDateRelative("5m")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Note: This assumes that the user `existing-user` has already been created in Artifactory by different means, i.e. manually or in a separate pulumi up.
 * 
 * ### Create a new Artifactory User and Access token
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var newUser = new User("newUser", UserArgs.builder()
 *             .name("new_user")
 *             .email("new_user{@literal @}somewhere.com")
 *             .groups("readers")
 *             .build());
 * 
 *         var newUserAccessToken = new AccessToken("newUserAccessToken", AccessTokenArgs.builder()
 *             .username(newUser.name())
 *             .endDateRelative("5m")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Creates a new token for groups
 * This creates a transient user called `temporary-user`.
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var temporaryUser = new AccessToken("temporaryUser", AccessTokenArgs.builder()
 *             .username("temporary-user")
 *             .endDateRelative("1h")
 *             .groups("readers")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Create token with no expiry
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var noExpiry = new AccessToken("noExpiry", AccessTokenArgs.builder()
 *             .username("existing-user")
 *             .endDateRelative("0s")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Creates a refreshable token
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var refreshable = new AccessToken("refreshable", AccessTokenArgs.builder()
 *             .username("refreshable")
 *             .endDateRelative("1m")
 *             .refreshable(true)
 *             .groups("readers")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Creates an administrator token
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var admin = new AccessToken("admin", AccessTokenArgs.builder()
 *             .username("admin")
 *             .endDateRelative("1m")
 *             .adminToken(AccessTokenAdminTokenArgs.builder()
 *                 .instanceId("<instance id>")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Creates a token with an audience
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var audience = new AccessToken("audience", AccessTokenArgs.builder()
 *             .username("audience")
 *             .endDateRelative("1m")
 *             .audience("jfrt{@literal @}*")
 *             .refreshable(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Creates a token with a fixed end date
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var fixeddate = new AccessToken("fixeddate", AccessTokenArgs.builder()
 *             .username("fixeddate")
 *             .endDate("2018-01-01T01:02:03Z")
 *             .groups("readers")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Rotate token after it expires
 * This example will generate a token that will expire in 1 hour.
 * 
 * If `pulumi up` is run before 1 hour, nothing changes.
 * One an hour has passed, `pulumi up` will generate a new token.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var nowPlus1Hours = new Rotating("nowPlus1Hours", RotatingArgs.builder()
 *             .rotationHours("1")
 *             .build());
 * 
 *         var rotating = new AccessToken("rotating", AccessTokenArgs.builder()
 *             .username("rotating")
 *             .endDate(nowPlus1Hour.rotationRfc3339())
 *             .groups("readers")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Rotate token each pulumi up
 * This example will generate a token that will expire in 1 hour.
 * 
 * If `pulumi up` is run before 1 hour, a new token is generated with an expiry of 1 hour.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var nowPlus1Hours = new Rotating("nowPlus1Hours", RotatingArgs.builder()
 *             .triggers(Map.of("key", StdFunctions.timestamp().result()))
 *             .rotationHours("1")
 *             .build());
 * 
 *         var rotating = new AccessToken("rotating", AccessTokenArgs.builder()
 *             .username("rotating")
 *             .endDate(nowPlus1Hour.rotationRfc3339())
 *             .groups("readers")
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
     * (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt{@literal @}*&#34;` so the token to be accepted by all Artifactory instances.
     * 
     */
    @Export(name="audience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> audience;

    /**
     * @return (Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set `&#34;jfrt{@literal @}*&#34;` so the token to be accepted by all Artifactory instances.
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
        super("artifactory:index/accessToken:AccessToken", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private AccessToken(String name, Output<String> id, @Nullable AccessTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/accessToken:AccessToken", name, state, makeResourceOptions(options, id));
    }

    private static AccessTokenArgs makeArgs(AccessTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AccessTokenArgs.Empty : args;
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
