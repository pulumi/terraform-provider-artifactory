// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.VirtualGradleRepositoryArgs;
import com.pulumi.artifactory.inputs.VirtualGradleRepositoryState;
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
 * Creates a virtual Gradle repository.
 * Official documentation can be found [here](https://jfrog.com/blog/how-to-set-up-a-private-remote-and-virtual-maven-gradle-registry/).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.VirtualGradleRepository;
 * import com.pulumi.artifactory.VirtualGradleRepositoryArgs;
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
 *         var foo_gradle = new VirtualGradleRepository(&#34;foo-gradle&#34;, VirtualGradleRepositoryArgs.builder()        
 *             .description(&#34;A test virtual repo&#34;)
 *             .excludesPattern(&#34;com/google/**&#34;)
 *             .includesPattern(&#34;com/jfrog/**,cloud/jfrog/**&#34;)
 *             .key(&#34;foo-gradle&#34;)
 *             .notes(&#34;Internal description&#34;)
 *             .pomRepositoryReferencesCleanupPolicy(&#34;discard_active_reference&#34;)
 *             .repositories()
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Virtual repositories can be imported using their name, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/virtualGradleRepository:VirtualGradleRepository foo-gradle foo-gradle
 * ```
 * 
 */
@ResourceType(type="artifactory:index/virtualGradleRepository:VirtualGradleRepository")
public class VirtualGradleRepository extends com.pulumi.resources.CustomResource {
    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    @Export(name="artifactoryRequestsCanRetrieveRemoteArtifacts", type=Boolean.class, parameters={})
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
    @Export(name="defaultDeploymentRepo", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultDeploymentRepo;

    /**
     * @return Default repository to deploy artifacts.
     * 
     */
    public Output<Optional<String>> defaultDeploymentRepo() {
        return Codegen.optional(this.defaultDeploymentRepo);
    }
    /**
     * A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
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
    @Export(name="excludesPattern", type=String.class, parameters={})
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
     * User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
     * is also enforced when aggregated repositories support anonymous requests.
     * 
     */
    @Export(name="forceMavenAuthentication", type=Boolean.class, parameters={})
    private Output<Boolean> forceMavenAuthentication;

    /**
     * @return User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
     * is also enforced when aggregated repositories support anonymous requests.
     * 
     */
    public Output<Boolean> forceMavenAuthentication() {
        return this.forceMavenAuthentication;
    }
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> includesPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
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
    @Export(name="key", type=String.class, parameters={})
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
     * The keypair used to sign artifacts.
     * 
     */
    @Export(name="keyPair", type=String.class, parameters={})
    private Output</* @Nullable */ String> keyPair;

    /**
     * @return The keypair used to sign artifacts.
     * 
     */
    public Output<Optional<String>> keyPair() {
        return Codegen.optional(this.keyPair);
    }
    /**
     * A free text field to add additional notes about the repository. These are only visible to the administrator.
     * 
     */
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    /**
     * @return A free text field to add additional notes about the repository. These are only visible to the administrator.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
     * 
     */
    @Export(name="packageType", type=String.class, parameters={})
    private Output<String> packageType;

    /**
     * @return The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
     * 
     */
    public Output<String> packageType() {
        return this.packageType;
    }
    /**
     * - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
     * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
     * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
     * 
     */
    @Export(name="pomRepositoryReferencesCleanupPolicy", type=String.class, parameters={})
    private Output<String> pomRepositoryReferencesCleanupPolicy;

    /**
     * @return - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
     * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
     * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
     * 
     */
    public Output<String> pomRepositoryReferencesCleanupPolicy() {
        return this.pomRepositoryReferencesCleanupPolicy;
    }
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    @Export(name="projectEnvironments", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    public Output<Optional<List<String>>> projectEnvironments() {
        return Codegen.optional(this.projectEnvironments);
    }
    /**
     * Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
     * repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
     * repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * Repository layout key for the local repository
     * 
     */
    @Export(name="repoLayoutRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * The effective list of actual repositories included in this virtual repository.
     * 
     */
    @Export(name="repositories", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> repositories;

    /**
     * @return The effective list of actual repositories included in this virtual repository.
     * 
     */
    public Output<Optional<List<String>>> repositories() {
        return Codegen.optional(this.repositories);
    }
    /**
     * This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
     * repositories. A value of 0 indicates no caching.
     * 
     */
    @Export(name="retrievalCachePeriodSeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> retrievalCachePeriodSeconds;

    /**
     * @return This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
     * repositories. A value of 0 indicates no caching.
     * 
     */
    public Output<Optional<Integer>> retrievalCachePeriodSeconds() {
        return Codegen.optional(this.retrievalCachePeriodSeconds);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualGradleRepository(String name) {
        this(name, VirtualGradleRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualGradleRepository(String name, VirtualGradleRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualGradleRepository(String name, VirtualGradleRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualGradleRepository:VirtualGradleRepository", name, args == null ? VirtualGradleRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VirtualGradleRepository(String name, Output<String> id, @Nullable VirtualGradleRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualGradleRepository:VirtualGradleRepository", name, state, makeResourceOptions(options, id));
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
    public static VirtualGradleRepository get(String name, Output<String> id, @Nullable VirtualGradleRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualGradleRepository(name, id, state, options);
    }
}
