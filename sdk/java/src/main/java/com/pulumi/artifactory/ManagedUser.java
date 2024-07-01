// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ManagedUserArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ManagedUserState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.artifactory.ManagedUser;
 * import com.pulumi.artifactory.ManagedUserArgs;
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
 *         var test_user = new ManagedUser("test-user", ManagedUserArgs.builder()
 *             .name("terraform")
 *             .password("my super secret password")
 *             .email("test-user{@literal @}artifactory-terraform.com")
 *             .groups(            
 *                 "readers",
 *                 "logged-in-users")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import artifactory:index/managedUser:ManagedUser test-user myusername
 * ```
 * 
 */
@ResourceType(type="artifactory:index/managedUser:ManagedUser")
public class ManagedUser extends com.pulumi.resources.CustomResource {
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
     * (Optional, Sensitive) Password for the user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return (Optional, Sensitive) Password for the user.
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
    public ManagedUser(String name) {
        this(name, ManagedUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ManagedUser(String name, ManagedUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ManagedUser(String name, ManagedUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/managedUser:ManagedUser", name, args == null ? ManagedUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ManagedUser(String name, Output<String> id, @Nullable ManagedUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/managedUser:ManagedUser", name, state, makeResourceOptions(options, id));
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
    public static ManagedUser get(String name, Output<String> id, @Nullable ManagedUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ManagedUser(name, id, state, options);
    }
}
