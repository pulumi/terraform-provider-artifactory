// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LocalMavenRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LocalMavenRepositoryState;
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
 * Creates a local Maven repository.
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
 * import com.pulumi.artifactory.LocalMavenRepository;
 * import com.pulumi.artifactory.LocalMavenRepositoryArgs;
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
 *         var terraform_local_test_maven_repo_basic = new LocalMavenRepository(&#34;terraform-local-test-maven-repo-basic&#34;, LocalMavenRepositoryArgs.builder()        
 *             .key(&#34;terraform-local-test-maven-repo-basic&#34;)
 *             .checksumPolicyType(&#34;client-checksums&#34;)
 *             .snapshotVersionBehavior(&#34;unique&#34;)
 *             .maxUniqueSnapshots(10)
 *             .handleReleases(true)
 *             .handleSnapshots(true)
 *             .suppressPomConsistencyChecks(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Local repositories can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/localMavenRepository:LocalMavenRepository terraform-local-test-maven-repo-basic terraform-local-test-maven-repo-basic
 * ```
 * 
 */
@ResourceType(type="artifactory:index/localMavenRepository:LocalMavenRepository")
public class LocalMavenRepository extends com.pulumi.resources.CustomResource {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    @Export(name="archiveBrowsingEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> archiveBrowsingEnabled;

    /**
     * @return When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    public Output<Optional<Boolean>> archiveBrowsingEnabled() {
        return Codegen.optional(this.archiveBrowsingEnabled);
    }
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    @Export(name="blackedOut", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> blackedOut;

    /**
     * @return When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    public Output<Optional<Boolean>> blackedOut() {
        return Codegen.optional(this.blackedOut);
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    @Export(name="cdnRedirect", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> cdnRedirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from AWS
     * CloudFront. Available in Enterprise+ and Edge licenses only. Default value is &#39;false&#39;
     * 
     */
    public Output<Optional<Boolean>> cdnRedirect() {
        return Codegen.optional(this.cdnRedirect);
    }
    /**
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are:
     * - `client-checksums`
     * - `server-generated-checksums`.
     *   For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
     * 
     */
    @Export(name="checksumPolicyType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> checksumPolicyType;

    /**
     * @return Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are:
     * - `client-checksums`
     * - `server-generated-checksums`.
     *   For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
     * 
     */
    public Output<Optional<String>> checksumPolicyType() {
        return Codegen.optional(this.checksumPolicyType);
    }
    /**
     * Public description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Public description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    @Export(name="downloadDirect", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    public Output<Optional<Boolean>> downloadDirect() {
        return Codegen.optional(this.downloadDirect);
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludesPattern() {
        return Codegen.optional(this.excludesPattern);
    }
    /**
     * If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
     * 
     */
    @Export(name="handleReleases", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> handleReleases;

    /**
     * @return If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
     * 
     */
    public Output<Optional<Boolean>> handleReleases() {
        return Codegen.optional(this.handleReleases);
    }
    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
     * 
     */
    @Export(name="handleSnapshots", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> handleSnapshots;

    /**
     * @return If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
     * 
     */
    public Output<Optional<Boolean>> handleSnapshots() {
        return Codegen.optional(this.handleSnapshots);
    }
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<Optional<String>> includesPattern() {
        return Codegen.optional(this.includesPattern);
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
    private Output</* @Nullable */ Integer> maxUniqueSnapshots;

    /**
     * @return The maximum number of unique snapshots of a single artifact to store.
     * Once the number of snapshots exceeds this setting, older versions are removed.
     * A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     * 
     */
    public Output<Optional<Integer>> maxUniqueSnapshots() {
        return Codegen.optional(this.maxUniqueSnapshots);
    }
    /**
     * Internal description.
     * 
     */
    @Export(name="notes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notes;

    /**
     * @return Internal description.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    @Export(name="packageType", refs={String.class}, tree="[0]")
    private Output<String> packageType;

    public Output<String> packageType() {
        return this.packageType;
    }
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Export(name="priorityResolution", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Output<Optional<Boolean>> priorityResolution() {
        return Codegen.optional(this.priorityResolution);
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
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
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
     * Repository layout key for the local repository
     * 
     */
    @Export(name="repoLayoutRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * Specifies the naming convention for Maven SNAPSHOT versions.
     * The options are -
     * 
     */
    @Export(name="snapshotVersionBehavior", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snapshotVersionBehavior;

    /**
     * @return Specifies the naming convention for Maven SNAPSHOT versions.
     * The options are -
     * 
     */
    public Output<Optional<String>> snapshotVersionBehavior() {
        return Codegen.optional(this.snapshotVersionBehavior);
    }
    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
     * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
     * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
     * 
     */
    @Export(name="suppressPomConsistencyChecks", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> suppressPomConsistencyChecks;

    /**
     * @return By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
     * If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a &#34;409 Conflict&#34; error.
     * You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
     * 
     */
    public Output<Optional<Boolean>> suppressPomConsistencyChecks() {
        return Codegen.optional(this.suppressPomConsistencyChecks);
    }
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Export(name="xrayIndex", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Output<Optional<Boolean>> xrayIndex() {
        return Codegen.optional(this.xrayIndex);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LocalMavenRepository(String name) {
        this(name, LocalMavenRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LocalMavenRepository(String name, LocalMavenRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LocalMavenRepository(String name, LocalMavenRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localMavenRepository:LocalMavenRepository", name, args == null ? LocalMavenRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LocalMavenRepository(String name, Output<String> id, @Nullable LocalMavenRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localMavenRepository:LocalMavenRepository", name, state, makeResourceOptions(options, id));
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
    public static LocalMavenRepository get(String name, Output<String> id, @Nullable LocalMavenRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LocalMavenRepository(name, id, state, options);
    }
}
