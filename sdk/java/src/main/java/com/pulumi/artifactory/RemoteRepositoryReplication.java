// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.RemoteRepositoryReplicationArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.RemoteRepositoryReplicationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a remote repository replication resource, also referred to as Artifactory pull replication.
 * This resource provides a convenient way to proactively populate a remote cache, and is very useful when waiting for new artifacts to arrive on demand (when first requested) is not desirable due to network latency. See [official documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PullReplication).
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
 * import com.pulumi.artifactory.RemoteMavenRepository;
 * import com.pulumi.artifactory.RemoteMavenRepositoryArgs;
 * import com.pulumi.artifactory.RemoteRepositoryReplication;
 * import com.pulumi.artifactory.RemoteRepositoryReplicationArgs;
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
 *         var providerTestSource = new LocalMavenRepository("providerTestSource", LocalMavenRepositoryArgs.builder()
 *             .key("provider_test_source")
 *             .build());
 * 
 *         var providerTestDest = new RemoteMavenRepository("providerTestDest", RemoteMavenRepositoryArgs.builder()
 *             .key("provider_test_dest")
 *             .url(String.format("%s/artifactory/%s", artifactoryUrl,artifactoryLocalMavenRepository.key()))
 *             .username("foo")
 *             .password("bar")
 *             .build());
 * 
 *         var remote_rep = new RemoteRepositoryReplication("remote-rep", RemoteRepositoryReplicationArgs.builder()
 *             .repoKey(providerTestDest.key())
 *             .cronExp("0 0 * * * ?")
 *             .enableEventReplication(true)
 *             .enabled(true)
 *             .syncDeletes(false)
 *             .syncProperties(true)
 *             .includePathPrefixPattern("/some-repo/")
 *             .excludePathPrefixPattern("/some-other-repo/")
 *             .checkBinaryExistenceInFilestore(false)
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
 * $ pulumi import artifactory:index/remoteRepositoryReplication:RemoteRepositoryReplication foo-rep provider_test_source
 * ```
 * 
 */
@ResourceType(type="artifactory:index/remoteRepositoryReplication:RemoteRepositoryReplication")
public class RemoteRepositoryReplication extends com.pulumi.resources.CustomResource {
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
     * com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     * 
     */
    @Export(name="enableEventReplication", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableEventReplication;

    /**
     * @return When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
     * com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
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
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`. By default, no artifacts are excluded.
     * 
     */
    @Export(name="excludePathPrefixPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> excludePathPrefixPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`. By default, no artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludePathPrefixPattern() {
        return Codegen.optional(this.excludePathPrefixPattern);
    }
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**&#47;*)`.
     * 
     */
    @Export(name="includePathPrefixPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> includePathPrefixPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**&#47;*)`.
     * 
     */
    public Output<Optional<String>> includePathPrefixPattern() {
        return Codegen.optional(this.includePathPrefixPattern);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RemoteRepositoryReplication(java.lang.String name) {
        this(name, RemoteRepositoryReplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RemoteRepositoryReplication(java.lang.String name, RemoteRepositoryReplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RemoteRepositoryReplication(java.lang.String name, RemoteRepositoryReplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remoteRepositoryReplication:RemoteRepositoryReplication", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RemoteRepositoryReplication(java.lang.String name, Output<java.lang.String> id, @Nullable RemoteRepositoryReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/remoteRepositoryReplication:RemoteRepositoryReplication", name, state, makeResourceOptions(options, id), false);
    }

    private static RemoteRepositoryReplicationArgs makeArgs(RemoteRepositoryReplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RemoteRepositoryReplicationArgs.Empty : args;
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
    public static RemoteRepositoryReplication get(java.lang.String name, Output<java.lang.String> id, @Nullable RemoteRepositoryReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RemoteRepositoryReplication(name, id, state, options);
    }
}
