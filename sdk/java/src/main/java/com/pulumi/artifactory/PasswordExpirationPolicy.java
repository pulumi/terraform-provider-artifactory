// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.PasswordExpirationPolicyArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.PasswordExpirationPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory Password Expiration Policy resource.
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
 * import com.pulumi.artifactory.PasswordExpirationPolicy;
 * import com.pulumi.artifactory.PasswordExpirationPolicyArgs;
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
 *         var my_password_expiration_policy = new PasswordExpirationPolicy(&#34;my-password-expiration-policy&#34;, PasswordExpirationPolicyArgs.builder()        
 *             .name(&#34;my-password-expiration-policy&#34;)
 *             .enabled(true)
 *             .passwordMaxAge(120)
 *             .notifyByEmail(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import artifactory:index/passwordExpirationPolicy:PasswordExpirationPolicy my-password-expiration-policy my-password-expiration-policy
 * ```
 * 
 */
@ResourceType(type="artifactory:index/passwordExpirationPolicy:PasswordExpirationPolicy")
public class PasswordExpirationPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Enable Password Expiration Policy. This only applies to internal user passwords.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Enable Password Expiration Policy. This only applies to internal user passwords.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Name of the resource. Only used for importing.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. Only used for importing.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Send mail notification before password expiration. Users will receive an email notification X days before password will expire. Mail server must be enabled and configured correctly.
     * 
     */
    @Export(name="notifyByEmail", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> notifyByEmail;

    /**
     * @return Send mail notification before password expiration. Users will receive an email notification X days before password will expire. Mail server must be enabled and configured correctly.
     * 
     */
    public Output<Boolean> notifyByEmail() {
        return this.notifyByEmail;
    }
    /**
     * Password expires every N days. The time interval in which users will be obligated to change their password.
     * 
     */
    @Export(name="passwordMaxAge", refs={Integer.class}, tree="[0]")
    private Output<Integer> passwordMaxAge;

    /**
     * @return Password expires every N days. The time interval in which users will be obligated to change their password.
     * 
     */
    public Output<Integer> passwordMaxAge() {
        return this.passwordMaxAge;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PasswordExpirationPolicy(String name) {
        this(name, PasswordExpirationPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PasswordExpirationPolicy(String name, PasswordExpirationPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PasswordExpirationPolicy(String name, PasswordExpirationPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/passwordExpirationPolicy:PasswordExpirationPolicy", name, args == null ? PasswordExpirationPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PasswordExpirationPolicy(String name, Output<String> id, @Nullable PasswordExpirationPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/passwordExpirationPolicy:PasswordExpirationPolicy", name, state, makeResourceOptions(options, id));
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
    public static PasswordExpirationPolicy get(String name, Output<String> id, @Nullable PasswordExpirationPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PasswordExpirationPolicy(name, id, state, options);
    }
}
