// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LocalRepositorySingleReplicationArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LocalRepositorySingleReplicationState;
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
 * Provides a local repository replication resource, also referred to as Artifactory push replication. This can be used to create and manage Artifactory local repository replications using [Push Replication API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-SetRepositoryReplicationConfiguration).
 * Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near end invoking a synchronization of artifacts to the far end.
 * See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).
 * This resource can create the replication of local repository to single repository on the remote server.
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
 * import com.pulumi.artifactory.LocalMavenRepository;
 * import com.pulumi.artifactory.LocalMavenRepositoryArgs;
 * import com.pulumi.artifactory.LocalRepositorySingleReplication;
 * import com.pulumi.artifactory.LocalRepositorySingleReplicationArgs;
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
 *         final var config = ctx.config();
 *         final var artifactoryUrl = config.get("artifactoryUrl");
 *         final var artifactoryUsername = config.get("artifactoryUsername");
 *         final var artifactoryPassword = config.get("artifactoryPassword");
 *         // Create a replication between two artifactory local repositories
 *         var providerTestSource = new LocalMavenRepository("providerTestSource", LocalMavenRepositoryArgs.builder()
 *             .key("provider_test_source")
 *             .build());
 * 
 *         var providerTestDest = new LocalMavenRepository("providerTestDest", LocalMavenRepositoryArgs.builder()
 *             .key("provider_test_dest")
 *             .build());
 * 
 *         var foo_rep = new LocalRepositorySingleReplication("foo-rep", LocalRepositorySingleReplicationArgs.builder()
 *             .repoKey(providerTestSource.key())
 *             .cronExp("0 0 * * * ?")
 *             .enableEventReplication(true)
 *             .url(providerTestDest.key().applyValue(key -> String.format("%s/artifactory/%s", artifactoryUrl,key)))
 *             .username("$var.artifactory_username")
 *             .password("$var.artifactory_password")
 *             .enabled(true)
 *             .socketTimeoutMillis(16000)
 *             .syncDeletes(false)
 *             .syncProperties(true)
 *             .syncStatistics(true)
 *             .includePathPrefixPattern("/some-repo/")
 *             .excludePathPrefixPattern("/some-other-repo/")
 *             .checkBinaryExistenceInFilestore(true)
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
 * Push replication configs can be imported using their repo key, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication foo-rep provider_test_source
 * ```
 * 
 */
@ResourceType(type="artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication")
public class LocalRepositorySingleReplication extends com.pulumi.resources.CustomResource {
    /**
     * Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    @Export(name="checkBinaryExistenceInFilestore", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> checkBinaryExistenceInFilestore;

    /**
     * @return Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    public Output<Optional<Boolean>> checkBinaryExistenceInFilestore() {
        return Codegen.optional(this.checkBinaryExistenceInFilestore);
    }
    /**
     * A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     * 
     */
    @Export(name="cronExp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cronExp;

    /**
     * @return A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     * 
     */
    public Output<Optional<String>> cronExp() {
        return Codegen.optional(this.cronExp);
    }
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
     * 
     */
    @Export(name="enableEventReplication", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableEventReplication;

    /**
     * @return When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> enableEventReplication() {
        return Codegen.optional(this.enableEventReplication);
    }
    /**
     * When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**{@literal /}z/*`. By default, no artifacts are excluded.
     * 
     */
    @Export(name="excludePathPrefixPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> excludePathPrefixPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**{@literal /}z/*`. By default, no artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludePathPrefixPattern() {
        return Codegen.optional(this.excludePathPrefixPattern);
    }
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**{@literal /}z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**{@literal /}*)`.
     * 
     */
    @Export(name="includePathPrefixPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> includePathPrefixPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**{@literal /}z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**{@literal /}*)`.
     * 
     */
    public Output<Optional<String>> includePathPrefixPattern() {
        return Codegen.optional(this.includePathPrefixPattern);
    }
    /**
     * Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     * 
     */
    @Export(name="proxy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     * 
     */
    public Output<Optional<String>> proxy() {
        return Codegen.optional(this.proxy);
    }
    /**
     * Replication ID, the value is unknown until the resource is created. Can&#39;t be set or updated.
     * 
     */
    @Export(name="replicationKey", refs={String.class}, tree="[0]")
    private Output<String> replicationKey;

    /**
     * @return Replication ID, the value is unknown until the resource is created. Can&#39;t be set or updated.
     * 
     */
    public Output<String> replicationKey() {
        return this.replicationKey;
    }
    /**
     * Repository name.
     * 
     */
    @Export(name="repoKey", refs={String.class}, tree="[0]")
    private Output<String> repoKey;

    /**
     * @return Repository name.
     * 
     */
    public Output<String> repoKey() {
        return this.repoKey;
    }
    /**
     * The network timeout in milliseconds to use for remote operations. Default value is `15000`.
     * 
     */
    @Export(name="socketTimeoutMillis", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> socketTimeoutMillis;

    /**
     * @return The network timeout in milliseconds to use for remote operations. Default value is `15000`.
     * 
     */
    public Output<Optional<Integer>> socketTimeoutMillis() {
        return Codegen.optional(this.socketTimeoutMillis);
    }
    /**
     * When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
     * 
     */
    @Export(name="syncDeletes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> syncDeletes;

    /**
     * @return When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> syncDeletes() {
        return Codegen.optional(this.syncDeletes);
    }
    /**
     * When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
     * 
     */
    @Export(name="syncProperties", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> syncProperties;

    /**
     * @return When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> syncProperties() {
        return Codegen.optional(this.syncProperties);
    }
    /**
     * When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
     * 
     */
    @Export(name="syncStatistics", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> syncStatistics;

    /**
     * @return When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
     * 
     */
    public Output<Optional<Boolean>> syncStatistics() {
        return Codegen.optional(this.syncStatistics);
    }
    /**
     * The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Username on the remote Artifactory instance.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return Username on the remote Artifactory instance.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LocalRepositorySingleReplication(String name) {
        this(name, LocalRepositorySingleReplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LocalRepositorySingleReplication(String name, LocalRepositorySingleReplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LocalRepositorySingleReplication(String name, LocalRepositorySingleReplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private LocalRepositorySingleReplication(String name, Output<String> id, @Nullable LocalRepositorySingleReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, state, makeResourceOptions(options, id));
    }

    private static LocalRepositorySingleReplicationArgs makeArgs(LocalRepositorySingleReplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LocalRepositorySingleReplicationArgs.Empty : args;
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
    public static LocalRepositorySingleReplication get(String name, Output<String> id, @Nullable LocalRepositorySingleReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LocalRepositorySingleReplication(name, id, state, options);
    }
}
