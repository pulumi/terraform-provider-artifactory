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

@ResourceType(type="artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication")
public class LocalRepositorySingleReplication extends com.pulumi.resources.CustomResource {
    /**
     * Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise+ license. When true, enables distributed
     * checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based
     * Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    @Export(name="checkBinaryExistenceInFilestore", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> checkBinaryExistenceInFilestore;

    /**
     * @return Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise+ license. When true, enables distributed
     * checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based
     * Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    public Output<Optional<Boolean>> checkBinaryExistenceInFilestore() {
        return Codegen.optional(this.checkBinaryExistenceInFilestore);
    }
    /**
     * The Cron expression that determines when the next replication will be triggered.
     * 
     */
    @Export(name="cronExp", type=String.class, parameters={})
    private Output</* @Nullable */ String> cronExp;

    /**
     * @return The Cron expression that determines when the next replication will be triggered.
     * 
     */
    public Output<Optional<String>> cronExp() {
        return Codegen.optional(this.cronExp);
    }
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on
     * artifact, e.g. add, deleted or property change. Default value is `false`.
     * 
     */
    @Export(name="enableEventReplication", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableEventReplication;

    /**
     * @return When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on
     * artifact, e.g. add, deleted or property change. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> enableEventReplication() {
        return Codegen.optional(this.enableEventReplication);
    }
    /**
     * When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludePathPrefixPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> excludePathPrefixPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludePathPrefixPattern() {
        return Codegen.optional(this.excludePathPrefixPattern);
    }
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includePathPrefixPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> includePathPrefixPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<Optional<String>> includePathPrefixPattern() {
        return Codegen.optional(this.includePathPrefixPattern);
    }
    /**
     * Use either the HTTP authentication password or identity token.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return Use either the HTTP authentication password or identity token.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * A proxy configuration to use when communicating with the remote instance.
     * 
     */
    @Export(name="proxy", type=String.class, parameters={})
    private Output</* @Nullable */ String> proxy;

    /**
     * @return A proxy configuration to use when communicating with the remote instance.
     * 
     */
    public Output<Optional<String>> proxy() {
        return Codegen.optional(this.proxy);
    }
    /**
     * Replication ID. The ID is known only after the replication is created, for this reason it&#39;s `Computed` and can not be
     * set by the user in HCL.
     * 
     */
    @Export(name="replicationKey", type=String.class, parameters={})
    private Output<String> replicationKey;

    /**
     * @return Replication ID. The ID is known only after the replication is created, for this reason it&#39;s `Computed` and can not be
     * set by the user in HCL.
     * 
     */
    public Output<String> replicationKey() {
        return this.replicationKey;
    }
    /**
     * Repository name.
     * 
     */
    @Export(name="repoKey", type=String.class, parameters={})
    private Output<String> repoKey;

    /**
     * @return Repository name.
     * 
     */
    public Output<String> repoKey() {
        return this.repoKey;
    }
    /**
     * The network timeout in milliseconds to use for remote operations.
     * 
     */
    @Export(name="socketTimeoutMillis", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> socketTimeoutMillis;

    /**
     * @return The network timeout in milliseconds to use for remote operations.
     * 
     */
    public Output<Optional<Integer>> socketTimeoutMillis() {
        return Codegen.optional(this.socketTimeoutMillis);
    }
    /**
     * When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note
     * that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value
     * is `false`.
     * 
     */
    @Export(name="syncDeletes", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> syncDeletes;

    /**
     * @return When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note
     * that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value
     * is `false`.
     * 
     */
    public Output<Optional<Boolean>> syncDeletes() {
        return Codegen.optional(this.syncDeletes);
    }
    /**
     * When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`
     * 
     */
    @Export(name="syncProperties", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> syncProperties;

    /**
     * @return When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`
     * 
     */
    public Output<Optional<Boolean>> syncProperties() {
        return Codegen.optional(this.syncProperties);
    }
    /**
     * When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target
     * instance when setting up replication for disaster recovery. Default value is `false`
     * 
     */
    @Export(name="syncStatistics", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> syncStatistics;

    /**
     * @return When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target
     * instance when setting up replication for disaster recovery. Default value is `false`
     * 
     */
    public Output<Optional<Boolean>> syncStatistics() {
        return Codegen.optional(this.syncStatistics);
    }
    /**
     * The URL of the target local repository on a remote Artifactory server. Use the format
     * `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
     * 
     */
    @Export(name="url", type=String.class, parameters={})
    private Output<String> url;

    /**
     * @return The URL of the target local repository on a remote Artifactory server. Use the format
     * `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The HTTP authentication username.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    /**
     * @return The HTTP authentication username.
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
        super("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, args == null ? LocalRepositorySingleReplicationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LocalRepositorySingleReplication(String name, Output<String> id, @Nullable LocalRepositorySingleReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, state, makeResourceOptions(options, id));
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
