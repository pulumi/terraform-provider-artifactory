// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PullReplicationArgs extends com.pulumi.resources.ResourceArgs {

    public static final PullReplicationArgs Empty = new PullReplicationArgs();

    /**
     * When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with
     * Checksum-Based
     * Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    @Import(name="checkBinaryExistenceInFilestore")
    private @Nullable Output<Boolean> checkBinaryExistenceInFilestore;

    /**
     * @return When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with
     * Checksum-Based
     * Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    public Optional<Output<Boolean>> checkBinaryExistenceInFilestore() {
        return Optional.ofNullable(this.checkBinaryExistenceInFilestore);
    }

    /**
     * The Cron expression that determines when the next replication will be triggered.
     * 
     */
    @Import(name="cronExp")
    private @Nullable Output<String> cronExp;

    /**
     * @return The Cron expression that determines when the next replication will be triggered.
     * 
     */
    public Optional<Output<String>> cronExp() {
        return Optional.ofNullable(this.cronExp);
    }

    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on
     * artifact, e.g. add, deleted or property change. Default value is `false`.
     * 
     */
    @Import(name="enableEventReplication")
    private @Nullable Output<Boolean> enableEventReplication;

    /**
     * @return When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on
     * artifact, e.g. add, deleted or property change. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> enableEventReplication() {
        return Optional.ofNullable(this.enableEventReplication);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Password for local repository replication. Required for local repository, but not needed for remote repository.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password for local repository replication. Required for local repository, but not needed for remote repository.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="pathPrefix")
    private @Nullable Output<String> pathPrefix;

    public Optional<Output<String>> pathPrefix() {
        return Optional.ofNullable(this.pathPrefix);
    }

    /**
     * Proxy key from Artifactory Proxies setting
     * 
     */
    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies setting
     * 
     */
    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    /**
     * Repository name.
     * 
     */
    @Import(name="repoKey", required=true)
    private Output<String> repoKey;

    /**
     * @return Repository name.
     * 
     */
    public Output<String> repoKey() {
        return this.repoKey;
    }

    @Import(name="socketTimeoutMillis")
    private @Nullable Output<Integer> socketTimeoutMillis;

    public Optional<Output<Integer>> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }

    @Import(name="syncDeletes")
    private @Nullable Output<Boolean> syncDeletes;

    public Optional<Output<Boolean>> syncDeletes() {
        return Optional.ofNullable(this.syncDeletes);
    }

    @Import(name="syncProperties")
    private @Nullable Output<Boolean> syncProperties;

    public Optional<Output<Boolean>> syncProperties() {
        return Optional.ofNullable(this.syncProperties);
    }

    @Import(name="syncStatistics")
    private @Nullable Output<Boolean> syncStatistics;

    public Optional<Output<Boolean>> syncStatistics() {
        return Optional.ofNullable(this.syncStatistics);
    }

    /**
     * URL for local repository replication. Required for local repository, but not needed for remote repository.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return URL for local repository replication. Required for local repository, but not needed for remote repository.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * Username for local repository replication. Required for local repository, but not needed for remote repository.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Username for local repository replication. Required for local repository, but not needed for remote repository.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private PullReplicationArgs() {}

    private PullReplicationArgs(PullReplicationArgs $) {
        this.checkBinaryExistenceInFilestore = $.checkBinaryExistenceInFilestore;
        this.cronExp = $.cronExp;
        this.enableEventReplication = $.enableEventReplication;
        this.enabled = $.enabled;
        this.password = $.password;
        this.pathPrefix = $.pathPrefix;
        this.proxy = $.proxy;
        this.repoKey = $.repoKey;
        this.socketTimeoutMillis = $.socketTimeoutMillis;
        this.syncDeletes = $.syncDeletes;
        this.syncProperties = $.syncProperties;
        this.syncStatistics = $.syncStatistics;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PullReplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PullReplicationArgs $;

        public Builder() {
            $ = new PullReplicationArgs();
        }

        public Builder(PullReplicationArgs defaults) {
            $ = new PullReplicationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param checkBinaryExistenceInFilestore When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with
         * Checksum-Based
         * Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
         * 
         * @return builder
         * 
         */
        public Builder checkBinaryExistenceInFilestore(@Nullable Output<Boolean> checkBinaryExistenceInFilestore) {
            $.checkBinaryExistenceInFilestore = checkBinaryExistenceInFilestore;
            return this;
        }

        /**
         * @param checkBinaryExistenceInFilestore When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with
         * Checksum-Based
         * Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
         * 
         * @return builder
         * 
         */
        public Builder checkBinaryExistenceInFilestore(Boolean checkBinaryExistenceInFilestore) {
            return checkBinaryExistenceInFilestore(Output.of(checkBinaryExistenceInFilestore));
        }

        /**
         * @param cronExp The Cron expression that determines when the next replication will be triggered.
         * 
         * @return builder
         * 
         */
        public Builder cronExp(@Nullable Output<String> cronExp) {
            $.cronExp = cronExp;
            return this;
        }

        /**
         * @param cronExp The Cron expression that determines when the next replication will be triggered.
         * 
         * @return builder
         * 
         */
        public Builder cronExp(String cronExp) {
            return cronExp(Output.of(cronExp));
        }

        /**
         * @param enableEventReplication When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on
         * artifact, e.g. add, deleted or property change. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableEventReplication(@Nullable Output<Boolean> enableEventReplication) {
            $.enableEventReplication = enableEventReplication;
            return this;
        }

        /**
         * @param enableEventReplication When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on
         * artifact, e.g. add, deleted or property change. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableEventReplication(Boolean enableEventReplication) {
            return enableEventReplication(Output.of(enableEventReplication));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param password Password for local repository replication. Required for local repository, but not needed for remote repository.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password for local repository replication. Required for local repository, but not needed for remote repository.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder pathPrefix(@Nullable Output<String> pathPrefix) {
            $.pathPrefix = pathPrefix;
            return this;
        }

        public Builder pathPrefix(String pathPrefix) {
            return pathPrefix(Output.of(pathPrefix));
        }

        /**
         * @param proxy Proxy key from Artifactory Proxies setting
         * 
         * @return builder
         * 
         */
        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        /**
         * @param proxy Proxy key from Artifactory Proxies setting
         * 
         * @return builder
         * 
         */
        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        /**
         * @param repoKey Repository name.
         * 
         * @return builder
         * 
         */
        public Builder repoKey(Output<String> repoKey) {
            $.repoKey = repoKey;
            return this;
        }

        /**
         * @param repoKey Repository name.
         * 
         * @return builder
         * 
         */
        public Builder repoKey(String repoKey) {
            return repoKey(Output.of(repoKey));
        }

        public Builder socketTimeoutMillis(@Nullable Output<Integer> socketTimeoutMillis) {
            $.socketTimeoutMillis = socketTimeoutMillis;
            return this;
        }

        public Builder socketTimeoutMillis(Integer socketTimeoutMillis) {
            return socketTimeoutMillis(Output.of(socketTimeoutMillis));
        }

        public Builder syncDeletes(@Nullable Output<Boolean> syncDeletes) {
            $.syncDeletes = syncDeletes;
            return this;
        }

        public Builder syncDeletes(Boolean syncDeletes) {
            return syncDeletes(Output.of(syncDeletes));
        }

        public Builder syncProperties(@Nullable Output<Boolean> syncProperties) {
            $.syncProperties = syncProperties;
            return this;
        }

        public Builder syncProperties(Boolean syncProperties) {
            return syncProperties(Output.of(syncProperties));
        }

        public Builder syncStatistics(@Nullable Output<Boolean> syncStatistics) {
            $.syncStatistics = syncStatistics;
            return this;
        }

        public Builder syncStatistics(Boolean syncStatistics) {
            return syncStatistics(Output.of(syncStatistics));
        }

        /**
         * @param url URL for local repository replication. Required for local repository, but not needed for remote repository.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL for local repository replication. Required for local repository, but not needed for remote repository.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username Username for local repository replication. Required for local repository, but not needed for remote repository.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username for local repository replication. Required for local repository, but not needed for remote repository.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public PullReplicationArgs build() {
            $.repoKey = Objects.requireNonNull($.repoKey, "expected parameter 'repoKey' to be non-null");
            return $;
        }
    }

}
