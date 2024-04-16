// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.UserArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.UserState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory user resource. This can be used to create and manage Artifactory users.
 * The password is a required field by the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser), but we made it optional in this resource to accommodate the scenario where the password is not needed and will be reset by the actual user later. When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.
 * 
 * &gt; The generated password won&#39;t be stored in the TF state and can not be recovered. The user must reset the password to be able to log in. An admin can always generate the access key for the user as well. The password change won&#39;t trigger state drift. We don&#39;t recommend to use this resource unless there is a specific use case for it. Recommended resource is `artifactory.ManagedUser`.
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
 *         var test_user = new User(&#34;test-user&#34;, UserArgs.builder()        
 *             .name(&#34;terraform&#34;)
 *             .password(&#34;my super secret password&#34;)
 *             .email(&#34;test-user@artifactory-terraform.com&#34;)
 *             .admin(false)
 *             .profileUpdatable(true)
 *             .disableUiAccess(false)
 *             .internalPasswordDisabled(false)
 *             .groups(&#34;logged-in-users&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Managing groups relationship
 * 
 * See our recommendation on how to manage user-group relationship.
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import artifactory:index/user:User test-user myusername
 * ```
 * 
 */
@ResourceType(type="artifactory:index/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
     * 
     */
    @Export(name="admin", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> admin;

    /**
     * @return (Optional, Default: false) When enabled, this user is an administrator with all the ensuing privileges.
     * 
     */
    public Output<Boolean> admin() {
        return this.admin;
    }
    /**
     * (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
     * 
     */
    @Export(name="disableUiAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableUiAccess;

    /**
     * @return (Optional, Default: true) When enabled, this user can only access the system through the REST API. This option cannot be set if the user has Admin privileges.
     * 
     */
    public Output<Boolean> disableUiAccess() {
        return this.disableUiAccess;
    }
    /**
     * Email for user.
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output<String> email;

    /**
     * @return Email for user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * List of groups this user is a part of. **Notes:** If this attribute is not specified then user&#39;s group membership is set to empty. User will not be part of default &#34;readers&#34; group automatically.
     * 
     */
    @Export(name="groups", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> groups;

    /**
     * @return List of groups this user is a part of. **Notes:** If this attribute is not specified then user&#39;s group membership is set to empty. User will not be part of default &#34;readers&#34; group automatically.
     * 
     */
    public Output<List<String>> groups() {
        return this.groups;
    }
    /**
     * (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    @Export(name="internalPasswordDisabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> internalPasswordDisabled;

    /**
     * @return (Optional, Default: false) When enabled, disables the fallback mechanism for using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    public Output<Boolean> internalPasswordDisabled() {
        return this.internalPasswordDisabled;
    }
    /**
     * Username for user. May contain lowercase letters, numbers and symbols: &#39;.-_@&#39;
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Username for user. May contain lowercase letters, numbers and symbols: &#39;.-_@&#39;
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
     * 
     */
    @Export(name="profileUpdatable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> profileUpdatable;

    /**
     * @return (Optional, Default: true) When enabled, this user can update their profile details (except for the password. Only an administrator can update the password). There may be cases in which you want to leave this unset to prevent users from updating their profile. For example, a departmental user with a single password shared between all department members.
     * 
     */
    public Output<Boolean> profileUpdatable() {
        return this.profileUpdatable;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
