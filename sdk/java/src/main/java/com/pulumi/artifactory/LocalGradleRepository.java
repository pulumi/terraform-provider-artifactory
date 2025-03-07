// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LocalGradleRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LocalGradleRepositoryState;
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
 * Creates a local Gradle repository.
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
 * import com.pulumi.artifactory.LocalGradleRepository;
 * import com.pulumi.artifactory.LocalGradleRepositoryArgs;
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
 *         var terraform_local_test_gradle_repo_basic = new LocalGradleRepository("terraform-local-test-gradle-repo-basic", LocalGradleRepositoryArgs.builder()
 *             .key("terraform-local-test-gradle-repo-basic")
 *             .checksumPolicyType("client-checksums")
 *             .snapshotVersionBehavior("unique")
 *             .maxUniqueSnapshots(10)
 *             .handleReleases(true)
 *             .handleSnapshots(true)
 *             .suppressPomConsistencyChecks(true)
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
 * Local repositories can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/localGradleRepository:LocalGradleRepository terraform-local-test-gradle-repo-basic terraform-local-test-gradle-repo-basic
 * ```
 * 
 */
@ResourceType(type="artifactory:index/localGradleRepository:LocalGradleRepository")
public class LocalGradleRepository extends com.pulumi.resources.CustomResource {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    @Export(name="archiveBrowsingEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> archiveBrowsingEnabled;

    /**
     * @return When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    public Output<Boolean> archiveBrowsingEnabled() {
        return this.archiveBrowsingEnabled;
    }
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    @Export(name="blackedOut", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> blackedOut;

    /**
     * @return When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    public Output<Boolean> blackedOut() {
        return this.blackedOut;
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    @Export(name="cdnRedirect", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> cdnRedirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    public Output<Boolean> cdnRedirect() {
        return this.cdnRedirect;
    }
    /**
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details,
     * please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
     * 
     */
    @Export(name="checksumPolicyType", refs={String.class}, tree="[0]")
    private Output<String> checksumPolicyType;

    /**
     * @return Checksum policy determines how Artifactory behaves when a client checksum for a deployed
     * resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
     * `client-checksums` and `generated-checksums`. For more details,
     * please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
     * 
     */
    public Output<String> checksumPolicyType() {
        return this.checksumPolicyType;
    }
    /**
     * Public description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Public description.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    @Export(name="downloadDirect", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    public Output<Boolean> downloadDirect() {
        return this.downloadDirect;
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", refs={String.class}, tree="[0]")
    private Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
     * artifacts are excluded.
     * 
     */
    public Output<String> excludesPattern() {
        return this.excludesPattern;
    }
    /**
     * If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
     * 
     */
    @Export(name="handleReleases", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> handleReleases;

    /**
     * @return If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
     * 
     */
    public Output<Boolean> handleReleases() {
        return this.handleReleases;
    }
    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
     * 
     */
    @Export(name="handleSnapshots", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> handleSnapshots;

    /**
     * @return If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
     * 
     */
    public Output<Boolean> handleSnapshots() {
        return this.handleSnapshots;
    }
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
     * 
     */
    @Export(name="includesPattern", refs={String.class}, tree="[0]")
    private Output<String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
     * 
     */
    public Output<String> includesPattern() {
        return this.includesPattern;
    }
    /**
     * the identity key of the repo.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     * 
     */
    @Export(name="maxUniqueSnapshots", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxUniqueSnapshots;

    /**
     * @return The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     * 
     */
    public Output<Integer> maxUniqueSnapshots() {
        return this.maxUniqueSnapshots;
    }
    /**
     * Internal description.
     * 
     */
    @Export(name="notes", refs={String.class}, tree="[0]")
    private Output<String> notes;

    /**
     * @return Internal description.
     * 
     */
    public Output<String> notes() {
        return this.notes;
    }
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Export(name="priorityResolution", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Output<Boolean> priorityResolution() {
        return this.priorityResolution;
    }
    @Export(name="projectEnvironments", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> projectEnvironments;

    public Output<List<String>> projectEnvironments() {
        return this.projectEnvironments;
    }
    /**
     * Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output<String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<String> projectKey() {
        return this.projectKey;
    }
    /**
     * List of property set name
     * 
     */
    @Export(name="propertySets", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> propertySets;

    /**
     * @return List of property set name
     * 
     */
    public Output<Optional<List<String>>> propertySets() {
        return Codegen.optional(this.propertySets);
    }
    /**
     * Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
     * corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
     * 
     */
    @Export(name="repoLayoutRef", refs={String.class}, tree="[0]")
    private Output<String> repoLayoutRef;

    /**
     * @return Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
     * corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
     * 
     */
    public Output<String> repoLayoutRef() {
        return this.repoLayoutRef;
    }
    /**
     * Specifies the naming convention for Maven SNAPSHOT versions.
     * The options are -
     * 
     */
    @Export(name="snapshotVersionBehavior", refs={String.class}, tree="[0]")
    private Output<String> snapshotVersionBehavior;

    /**
     * @return Specifies the naming convention for Maven SNAPSHOT versions.
     * The options are -
     * 
     */
    public Output<String> snapshotVersionBehavior() {
        return this.snapshotVersionBehavior;
    }
    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
     * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
     * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     * 
     */
    @Export(name="suppressPomConsistencyChecks", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> suppressPomConsistencyChecks;

    /**
     * @return By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
     * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
     * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
     * 
     */
    public Output<Boolean> suppressPomConsistencyChecks() {
        return this.suppressPomConsistencyChecks;
    }
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Export(name="xrayIndex", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Output<Boolean> xrayIndex() {
        return this.xrayIndex;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LocalGradleRepository(java.lang.String name) {
        this(name, LocalGradleRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LocalGradleRepository(java.lang.String name, LocalGradleRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LocalGradleRepository(java.lang.String name, LocalGradleRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localGradleRepository:LocalGradleRepository", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private LocalGradleRepository(java.lang.String name, Output<java.lang.String> id, @Nullable LocalGradleRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localGradleRepository:LocalGradleRepository", name, state, makeResourceOptions(options, id), false);
    }

    private static LocalGradleRepositoryArgs makeArgs(LocalGradleRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LocalGradleRepositoryArgs.Empty : args;
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
    public static LocalGradleRepository get(java.lang.String name, Output<java.lang.String> id, @Nullable LocalGradleRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LocalGradleRepository(name, id, state, options);
    }
}
