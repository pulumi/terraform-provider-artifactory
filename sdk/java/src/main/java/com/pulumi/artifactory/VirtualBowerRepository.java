// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.VirtualBowerRepositoryArgs;
import com.pulumi.artifactory.inputs.VirtualBowerRepositoryState;
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
 * Creates a virtual Bower repository.
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Bower+Repositories#BowerRepositories-VirtualRepositories).
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
 * import com.pulumi.artifactory.VirtualBowerRepository;
 * import com.pulumi.artifactory.VirtualBowerRepositoryArgs;
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
 *         var foo_bower = new VirtualBowerRepository("foo-bower", VirtualBowerRepositoryArgs.builder()
 *             .key("foo-bower")
 *             .repositories()
 *             .description("A test virtual repo")
 *             .notes("Internal description")
 *             .includesPattern("com/jfrog/**,cloud/jfrog/**")
 *             .excludesPattern("com/google/**")
 *             .externalDependenciesEnabled(false)
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
 * Virtual repositories can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/virtualBowerRepository:VirtualBowerRepository foo-bower foo-bower
 * ```
 * 
 */
@ResourceType(type="artifactory:index/virtualBowerRepository:VirtualBowerRepository")
public class VirtualBowerRepository extends com.pulumi.resources.CustomResource {
    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    @Export(name="artifactoryRequestsCanRetrieveRemoteArtifacts", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    /**
     * @return Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    public Output<Optional<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Codegen.optional(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }
    /**
     * Default repository to deploy artifacts.
     * 
     */
    @Export(name="defaultDeploymentRepo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultDeploymentRepo;

    /**
     * @return Default repository to deploy artifacts.
     * 
     */
    public Output<Optional<String>> defaultDeploymentRepo() {
        return Codegen.optional(this.defaultDeploymentRepo);
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
     * When set, external dependencies are rewritten. Default value is false.
     * 
     */
    @Export(name="externalDependenciesEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> externalDependenciesEnabled;

    /**
     * @return When set, external dependencies are rewritten. Default value is false.
     * 
     */
    public Output<Optional<Boolean>> externalDependenciesEnabled() {
        return Codegen.optional(this.externalDependenciesEnabled);
    }
    /**
     * An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    @Export(name="externalDependenciesPatterns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> externalDependenciesPatterns;

    /**
     * @return An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    public Output<Optional<List<String>>> externalDependenciesPatterns() {
        return Codegen.optional(this.externalDependenciesPatterns);
    }
    /**
     * The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    @Export(name="externalDependenciesRemoteRepo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> externalDependenciesRemoteRepo;

    /**
     * @return The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    public Output<Optional<String>> externalDependenciesRemoteRepo() {
        return Codegen.optional(this.externalDependenciesRemoteRepo);
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
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
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
     * Repository layout key for the virtual repository
     * 
     */
    @Export(name="repoLayoutRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the virtual repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * The effective list of actual repositories included in this virtual repository.
     * 
     */
    @Export(name="repositories", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> repositories;

    /**
     * @return The effective list of actual repositories included in this virtual repository.
     * 
     */
    public Output<Optional<List<String>>> repositories() {
        return Codegen.optional(this.repositories);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualBowerRepository(String name) {
        this(name, VirtualBowerRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualBowerRepository(String name, VirtualBowerRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualBowerRepository(String name, VirtualBowerRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualBowerRepository:VirtualBowerRepository", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private VirtualBowerRepository(String name, Output<String> id, @Nullable VirtualBowerRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualBowerRepository:VirtualBowerRepository", name, state, makeResourceOptions(options, id));
    }

    private static VirtualBowerRepositoryArgs makeArgs(VirtualBowerRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VirtualBowerRepositoryArgs.Empty : args;
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
    public static VirtualBowerRepository get(String name, Output<String> id, @Nullable VirtualBowerRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualBowerRepository(name, id, state, options);
    }
}
