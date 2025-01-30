// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.AlpineRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.AlpineRepositoryState;
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
 * Creates a local Alpine repository.
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
 * import com.pulumi.artifactory.Keypair;
 * import com.pulumi.artifactory.KeypairArgs;
 * import com.pulumi.artifactory.AlpineRepository;
 * import com.pulumi.artifactory.AlpineRepositoryArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var some_keypairRSA = new Keypair("some-keypairRSA", KeypairArgs.builder()
 *             .pairName("some-keypair")
 *             .pairType("RSA")
 *             .alias("foo-alias")
 *             .privateKey(StdFunctions.file(FileArgs.builder()
 *                 .input("samples/rsa.priv")
 *                 .build()).result())
 *             .publicKey(StdFunctions.file(FileArgs.builder()
 *                 .input("samples/rsa.pub")
 *                 .build()).result())
 *             .build());
 * 
 *         var terraform_local_test_alpine_repo_basic = new AlpineRepository("terraform-local-test-alpine-repo-basic", AlpineRepositoryArgs.builder()
 *             .key("terraform-local-test-alpine-repo-basic")
 *             .primaryKeypairRef(some_keypairRSA.pairName())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(some_keypairRSA)
 *                 .build());
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
 * $ pulumi import artifactory:index/alpineRepository:AlpineRepository terraform-local-test-alpine-repo-basic terraform-local-test-alpine-repo-basic
 * ```
 * 
 */
@ResourceType(type="artifactory:index/alpineRepository:AlpineRepository")
public class AlpineRepository extends com.pulumi.resources.CustomResource {
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
     * Used to sign index files in Alpine Linux repositories. See:
     * https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
     * 
     */
    @Export(name="primaryKeypairRef", refs={String.class}, tree="[0]")
    private Output<String> primaryKeypairRef;

    /**
     * @return Used to sign index files in Alpine Linux repositories. See:
     * https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
     * 
     */
    public Output<String> primaryKeypairRef() {
        return this.primaryKeypairRef;
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
    public AlpineRepository(java.lang.String name) {
        this(name, AlpineRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlpineRepository(java.lang.String name, AlpineRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlpineRepository(java.lang.String name, AlpineRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/alpineRepository:AlpineRepository", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AlpineRepository(java.lang.String name, Output<java.lang.String> id, @Nullable AlpineRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/alpineRepository:AlpineRepository", name, state, makeResourceOptions(options, id), false);
    }

    private static AlpineRepositoryArgs makeArgs(AlpineRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AlpineRepositoryArgs.Empty : args;
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
    public static AlpineRepository get(java.lang.String name, Output<java.lang.String> id, @Nullable AlpineRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlpineRepository(name, id, state, options);
    }
}
