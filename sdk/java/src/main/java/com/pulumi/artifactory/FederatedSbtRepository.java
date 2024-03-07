// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.FederatedSbtRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.FederatedSbtRepositoryState;
import com.pulumi.artifactory.outputs.FederatedSbtRepositoryMember;
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
 * Creates a federated SBT repository.
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
 * import com.pulumi.artifactory.FederatedSbtRepository;
 * import com.pulumi.artifactory.FederatedSbtRepositoryArgs;
 * import com.pulumi.artifactory.inputs.FederatedSbtRepositoryMemberArgs;
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
 *         var terraform_federated_test_sbt_repo = new FederatedSbtRepository(&#34;terraform-federated-test-sbt-repo&#34;, FederatedSbtRepositoryArgs.builder()        
 *             .key(&#34;terraform-federated-test-sbt-repo&#34;)
 *             .members(            
 *                 FederatedSbtRepositoryMemberArgs.builder()
 *                     .enabled(true)
 *                     .url(&#34;http://tempurl.org/artifactory/terraform-federated-test-sbt-repo&#34;)
 *                     .build(),
 *                 FederatedSbtRepositoryMemberArgs.builder()
 *                     .enabled(true)
 *                     .url(&#34;http://tempurl2.org/artifactory/terraform-federated-test-sbt-repo-2&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Federated repositories can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/federatedSbtRepository:FederatedSbtRepository terraform-federated-test-sbt-repo terraform-federated-test-sbt-repo
 * ```
 * 
 */
@ResourceType(type="artifactory:index/federatedSbtRepository:FederatedSbtRepository")
public class FederatedSbtRepository extends com.pulumi.resources.CustomResource {
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
     * Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
     * conflicts with the locally calculated checksum (bad checksum). Options are: &#34;client-checksums&#34;, or
     * &#34;server-generated-checksums&#34;. Default: &#34;client-checksums&#34;\n For more details, please refer to Checksum Policy -
     * https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
     * 
     */
    @Export(name="checksumPolicyType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> checksumPolicyType;

    /**
     * @return Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
     * conflicts with the locally calculated checksum (bad checksum). Options are: &#34;client-checksums&#34;, or
     * &#34;server-generated-checksums&#34;. Default: &#34;client-checksums&#34;\n For more details, please refer to Checksum Policy -
     * https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
     * 
     */
    public Output<Optional<String>> checksumPolicyType() {
        return Codegen.optional(this.checksumPolicyType);
    }
    /**
     * Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
     * the federation on other Artifactory instances.
     * 
     */
    @Export(name="cleanupOnDelete", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> cleanupOnDelete;

    /**
     * @return Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
     * the federation on other Artifactory instances.
     * 
     */
    public Output<Optional<Boolean>> cleanupOnDelete() {
        return Codegen.optional(this.cleanupOnDelete);
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
     * When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     * 
     */
    @Export(name="disableProxy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableProxy;

    /**
     * @return When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     * 
     */
    public Output<Optional<Boolean>> disableProxy() {
        return Codegen.optional(this.disableProxy);
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
     * If set, Artifactory allows you to deploy release artifacts into this repository.
     * 
     */
    @Export(name="handleReleases", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> handleReleases;

    /**
     * @return If set, Artifactory allows you to deploy release artifacts into this repository.
     * 
     */
    public Output<Optional<Boolean>> handleReleases() {
        return Codegen.optional(this.handleReleases);
    }
    /**
     * If set, Artifactory allows you to deploy snapshot artifacts into this repository.
     * 
     */
    @Export(name="handleSnapshots", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> handleSnapshots;

    /**
     * @return If set, Artifactory allows you to deploy snapshot artifacts into this repository.
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
     * The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
     * older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     * 
     */
    @Export(name="maxUniqueSnapshots", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxUniqueSnapshots;

    /**
     * @return The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
     * older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
     * 
     */
    public Output<Optional<Integer>> maxUniqueSnapshots() {
        return Codegen.optional(this.maxUniqueSnapshots);
    }
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    @Export(name="members", refs={List.class,FederatedSbtRepositoryMember.class}, tree="[0,1]")
    private Output<List<FederatedSbtRepositoryMember>> members;

    /**
     * @return The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    public Output<List<FederatedSbtRepositoryMember>> members() {
        return this.members;
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
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    @Export(name="projectEnvironments", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34;, &#34;PROD&#34;, or one of custom environment. Before
     * Artifactory 7.53.1, up to 2 values (&#34;DEV&#34; and &#34;PROD&#34;) are allowed. From 7.53.1 onward, only one value is allowed. The
     * attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
     * be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    public Output<List<String>> projectEnvironments() {
        return this.projectEnvironments;
    }
    /**
     * Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
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
     * Proxy key from Artifactory Proxies settings. Default is empty field. Can&#39;t be set if `disable_proxy = true`.
     * 
     */
    @Export(name="proxy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies settings. Default is empty field. Can&#39;t be set if `disable_proxy = true`.
     * 
     */
    public Output<Optional<String>> proxy() {
        return Codegen.optional(this.proxy);
    }
    /**
     * Repository layout key for the federated repository
     * 
     */
    @Export(name="repoLayoutRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the federated repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
     * time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
     * artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
     * 
     */
    @Export(name="snapshotVersionBehavior", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snapshotVersionBehavior;

    /**
     * @return Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
     * time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
     * artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
     * 
     */
    public Output<Optional<String>> snapshotVersionBehavior() {
        return Codegen.optional(this.snapshotVersionBehavior);
    }
    /**
     * By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
     * groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
     * deployment with a &#34;409 Conflict&#34; error. You can disable this behavior by setting the Suppress POM Consistency Checks
     * checkbox.
     * 
     */
    @Export(name="suppressPomConsistencyChecks", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> suppressPomConsistencyChecks;

    /**
     * @return By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
     * groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
     * deployment with a &#34;409 Conflict&#34; error. You can disable this behavior by setting the Suppress POM Consistency Checks
     * checkbox.
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
    public FederatedSbtRepository(String name) {
        this(name, FederatedSbtRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FederatedSbtRepository(String name, FederatedSbtRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FederatedSbtRepository(String name, FederatedSbtRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/federatedSbtRepository:FederatedSbtRepository", name, args == null ? FederatedSbtRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FederatedSbtRepository(String name, Output<String> id, @Nullable FederatedSbtRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/federatedSbtRepository:FederatedSbtRepository", name, state, makeResourceOptions(options, id));
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
    public static FederatedSbtRepository get(String name, Output<String> id, @Nullable FederatedSbtRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FederatedSbtRepository(name, id, state, options);
    }
}
