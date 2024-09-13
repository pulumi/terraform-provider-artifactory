// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.PackageCleanupPolicyArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.PackageCleanupPolicyState;
import com.pulumi.artifactory.outputs.PackageCleanupPolicySearchCriteria;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory Package Cleanup Policy resource. This resource enable system administrators to define and customize policies based on specific criteria for removing unused binaries from across their JFrog platform. See [Rentation Policies](https://jfrog.com/help/r/jfrog-platform-administration-documentation/retention-policies) for more details.
 * 
 * ~&gt;Currently in beta and not yet globally available. A full rollout is scheduled for early October 2024.
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
 * import com.pulumi.artifactory.PackageCleanupPolicy;
 * import com.pulumi.artifactory.PackageCleanupPolicyArgs;
 * import com.pulumi.artifactory.inputs.PackageCleanupPolicySearchCriteriaArgs;
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
 *         var my_cleanup_policy = new PackageCleanupPolicy("my-cleanup-policy", PackageCleanupPolicyArgs.builder()
 *             .key("my-policy")
 *             .description("My package cleanup policy")
 *             .cronExpression("0 0 2 ? * MON-SAT *")
 *             .durationInMinutes(60)
 *             .enabled(true)
 *             .skipTrashcan(false)
 *             .projectKey("myprojkey")
 *             .searchCriteria(PackageCleanupPolicySearchCriteriaArgs.builder()
 *                 .package_types(                
 *                     "docker",
 *                     "maven")
 *                 .repos(                
 *                     "my-docker-local",
 *                     "my-maven-local")
 *                 .excluded_repos("gradle-global")
 *                 .include_all_projects(false)
 *                 .included_projects()
 *                 .included_packages("com/jfrog")
 *                 .excluded_packages("com/jfrog/latest")
 *                 .created_before_in_months(1)
 *                 .last_downloaded_before_in_months(6)
 *                 .keep_last_n_versions(0)
 *                 .build())
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
 * $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy
 * ```
 * 
 * ```sh
 * $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy:myproj
 * ```
 * 
 */
@ResourceType(type="artifactory:index/packageCleanupPolicy:PackageCleanupPolicy")
public class PackageCleanupPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
     * 
     */
    @Export(name="cronExpression", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cronExpression;

    /**
     * @return The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
     * 
     */
    public Output<Optional<String>> cronExpression() {
        return Codegen.optional(this.cronExpression);
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
     * 
     */
    @Export(name="durationInMinutes", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> durationInMinutes;

    /**
     * @return Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
     * 
     */
    public Output<Optional<Integer>> durationInMinutes() {
        return Codegen.optional(this.durationInMinutes);
    }
    /**
     * Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    @Export(name="searchCriteria", refs={PackageCleanupPolicySearchCriteria.class}, tree="[0]")
    private Output<PackageCleanupPolicySearchCriteria> searchCriteria;

    public Output<PackageCleanupPolicySearchCriteria> searchCriteria() {
        return this.searchCriteria;
    }
    /**
     * Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
     * 
     */
    @Export(name="skipTrashcan", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> skipTrashcan;

    /**
     * @return Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
     * 
     */
    public Output<Boolean> skipTrashcan() {
        return this.skipTrashcan;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PackageCleanupPolicy(java.lang.String name) {
        this(name, PackageCleanupPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PackageCleanupPolicy(java.lang.String name, PackageCleanupPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PackageCleanupPolicy(java.lang.String name, PackageCleanupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/packageCleanupPolicy:PackageCleanupPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PackageCleanupPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable PackageCleanupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/packageCleanupPolicy:PackageCleanupPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static PackageCleanupPolicyArgs makeArgs(PackageCleanupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PackageCleanupPolicyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static PackageCleanupPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable PackageCleanupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PackageCleanupPolicy(name, id, state, options);
    }
}
