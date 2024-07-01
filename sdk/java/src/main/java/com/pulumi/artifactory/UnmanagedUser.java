// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.UnmanagedUserArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.UnmanagedUserState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory unmanaged user resource. This can be used to create and manage Artifactory users.
 * The password is a required field by the [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser), but we made it optional in this resource to accommodate the scenario where the password is not needed and will be reset by the actual user later. When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.
 * 
 * &gt; The generated password won&#39;t be stored in the TF state and can not be recovered. The user must reset the password to be able to log in. An admin can always generate the access key for the user as well. The password change won&#39;t trigger state drift.
 * 
 * &gt; This resource is an alias for `artifactory.User` resource, it is identical and was added for clarity and compatibility purposes. We don&#39;t recommend to use this resource unless there is a specific use case for it. Recommended resource is `artifactory.ManagedUser`.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.UnmanagedUser;
 * import com.pulumi.artifactory.UnmanagedUserArgs;
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
 *         // Create a new Artifactory user called terraform
 *         var test_user = new UnmanagedUser("test-user", UnmanagedUserArgs.builder()
 *             .name("terraform")
 *             .email("test-user{@literal @}artifactory-terraform.com")
 *             .groups("logged-in-users")
 *             .password("my super secret password")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Managing groups relationship
 * 
 * See our recommendation on how to manage user-group relationship.
 * 
 * ## Import
 * 
 * Users can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/unmanagedUser:UnmanagedUser test-user myusername
 * ```
 * 
 */
@ResourceType(type="artifactory:index/unmanagedUser:UnmanagedUser")
public class UnmanagedUser extends com.pulumi.resources.CustomResource {
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
     * Username for user. May contain lowercase letters, numbers and symbols: `.-_{@literal @}` for self-hosted. For SaaS, `+` is also allowed.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Username for user. May contain lowercase letters, numbers and symbols: `.-_{@literal @}` for self-hosted. For SaaS, `+` is also allowed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return (Optional, Sensitive) Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
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
    public UnmanagedUser(String name) {
        this(name, UnmanagedUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UnmanagedUser(String name, UnmanagedUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UnmanagedUser(String name, UnmanagedUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/unmanagedUser:UnmanagedUser", name, args == null ? UnmanagedUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UnmanagedUser(String name, Output<String> id, @Nullable UnmanagedUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/unmanagedUser:UnmanagedUser", name, state, makeResourceOptions(options, id));
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
    public static UnmanagedUser get(String name, Output<String> id, @Nullable UnmanagedUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UnmanagedUser(name, id, state, options);
    }
}
