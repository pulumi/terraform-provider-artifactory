// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.FederatedDockerV2RepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.FederatedDockerV2RepositoryState;
import com.pulumi.artifactory.outputs.FederatedDockerV2RepositoryMember;
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
 * Creates a federated Docker repository.
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
 * import com.pulumi.artifactory.FederatedDockerV2Repository;
 * import com.pulumi.artifactory.FederatedDockerV2RepositoryArgs;
 * import com.pulumi.artifactory.inputs.FederatedDockerV2RepositoryMemberArgs;
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
 *         var terraform_federated_test_docker_repo = new FederatedDockerV2Repository("terraform-federated-test-docker-repo", FederatedDockerV2RepositoryArgs.builder()
 *             .key("terraform-federated-test-docker-repo")
 *             .members(            
 *                 FederatedDockerV2RepositoryMemberArgs.builder()
 *                     .url("http://tempurl.org/artifactory/terraform-federated-test-docker-repo")
 *                     .enabled(true)
 *                     .build(),
 *                 FederatedDockerV2RepositoryMemberArgs.builder()
 *                     .url("http://tempurl2.org/artifactory/terraform-federated-test-docker-repo-2")
 *                     .enabled(true)
 *                     .build())
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
 * Federated repositories can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/federatedDockerV2Repository:FederatedDockerV2Repository terraform-federated-test-docker-repo terraform-federated-test-docker-repo
 * ```
 * 
 */
@ResourceType(type="artifactory:index/federatedDockerV2Repository:FederatedDockerV2Repository")
public class FederatedDockerV2Repository extends com.pulumi.resources.CustomResource {
    /**
     * The Docker API version to use. This cannot be set
     * 
     */
    @Export(name="apiVersion", refs={String.class}, tree="[0]")
    private Output<String> apiVersion;

    /**
     * @return The Docker API version to use. This cannot be set
     * 
     */
    public Output<String> apiVersion() {
        return this.apiVersion;
    }
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
     * When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
     * 
     */
    @Export(name="blockPushingSchema1", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> blockPushingSchema1;

    /**
     * @return When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
     * 
     */
    public Output<Boolean> blockPushingSchema1() {
        return this.blockPushingSchema1;
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
    @Export(name="cleanupOnDelete", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> cleanupOnDelete;

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
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*.By default no
     * artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludesPattern() {
        return Codegen.optional(this.excludesPattern);
    }
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     * 
     */
    @Export(name="includesPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
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
     * The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
     * image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
     * applies to manifest v2
     * 
     */
    @Export(name="maxUniqueTags", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxUniqueTags;

    /**
     * @return The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
     * image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
     * applies to manifest v2
     * 
     */
    public Output<Optional<Integer>> maxUniqueTags() {
        return Codegen.optional(this.maxUniqueTags);
    }
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    @Export(name="members", refs={List.class,FederatedDockerV2RepositoryMember.class}, tree="[0,1]")
    private Output<List<FederatedDockerV2RepositoryMember>> members;

    /**
     * @return The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    public Output<List<FederatedDockerV2RepositoryMember>> members() {
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
     * If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
     * manifest V2
     * 
     */
    @Export(name="tagRetention", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tagRetention;

    /**
     * @return If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
     * manifest V2
     * 
     */
    public Output<Optional<Integer>> tagRetention() {
        return Codegen.optional(this.tagRetention);
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
    public FederatedDockerV2Repository(java.lang.String name) {
        this(name, FederatedDockerV2RepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FederatedDockerV2Repository(java.lang.String name, FederatedDockerV2RepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FederatedDockerV2Repository(java.lang.String name, FederatedDockerV2RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/federatedDockerV2Repository:FederatedDockerV2Repository", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FederatedDockerV2Repository(java.lang.String name, Output<java.lang.String> id, @Nullable FederatedDockerV2RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/federatedDockerV2Repository:FederatedDockerV2Repository", name, state, makeResourceOptions(options, id), false);
    }

    private static FederatedDockerV2RepositoryArgs makeArgs(FederatedDockerV2RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FederatedDockerV2RepositoryArgs.Empty : args;
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
    public static FederatedDockerV2Repository get(java.lang.String name, Output<java.lang.String> id, @Nullable FederatedDockerV2RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FederatedDockerV2Repository(name, id, state, options);
    }
}
