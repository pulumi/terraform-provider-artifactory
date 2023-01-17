// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.inputs.ReplicationConfigReplicationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReplicationConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReplicationConfigArgs Empty = new ReplicationConfigArgs();

    /**
     * Cron expression to control the operation frequency.
     * 
     */
    @Import(name="cronExp", required=true)
    private Output<String> cronExp;

    /**
     * @return Cron expression to control the operation frequency.
     * 
     */
    public Output<String> cronExp() {
        return this.cronExp;
    }

    @Import(name="enableEventReplication")
    private @Nullable Output<Boolean> enableEventReplication;

    public Optional<Output<Boolean>> enableEventReplication() {
        return Optional.ofNullable(this.enableEventReplication);
    }

    @Import(name="replications")
    private @Nullable Output<List<ReplicationConfigReplicationArgs>> replications;

    public Optional<Output<List<ReplicationConfigReplicationArgs>>> replications() {
        return Optional.ofNullable(this.replications);
    }

    @Import(name="repoKey", required=true)
    private Output<String> repoKey;

    public Output<String> repoKey() {
        return this.repoKey;
    }

    private ReplicationConfigArgs() {}

    private ReplicationConfigArgs(ReplicationConfigArgs $) {
        this.cronExp = $.cronExp;
        this.enableEventReplication = $.enableEventReplication;
        this.replications = $.replications;
        this.repoKey = $.repoKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReplicationConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReplicationConfigArgs $;

        public Builder() {
            $ = new ReplicationConfigArgs();
        }

        public Builder(ReplicationConfigArgs defaults) {
            $ = new ReplicationConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cronExp Cron expression to control the operation frequency.
         * 
         * @return builder
         * 
         */
        public Builder cronExp(Output<String> cronExp) {
            $.cronExp = cronExp;
            return this;
        }

        /**
         * @param cronExp Cron expression to control the operation frequency.
         * 
         * @return builder
         * 
         */
        public Builder cronExp(String cronExp) {
            return cronExp(Output.of(cronExp));
        }

        public Builder enableEventReplication(@Nullable Output<Boolean> enableEventReplication) {
            $.enableEventReplication = enableEventReplication;
            return this;
        }

        public Builder enableEventReplication(Boolean enableEventReplication) {
            return enableEventReplication(Output.of(enableEventReplication));
        }

        public Builder replications(@Nullable Output<List<ReplicationConfigReplicationArgs>> replications) {
            $.replications = replications;
            return this;
        }

        public Builder replications(List<ReplicationConfigReplicationArgs> replications) {
            return replications(Output.of(replications));
        }

        public Builder replications(ReplicationConfigReplicationArgs... replications) {
            return replications(List.of(replications));
        }

        public Builder repoKey(Output<String> repoKey) {
            $.repoKey = repoKey;
            return this;
        }

        public Builder repoKey(String repoKey) {
            return repoKey(Output.of(repoKey));
        }

        public ReplicationConfigArgs build() {
            $.cronExp = Objects.requireNonNull($.cronExp, "expected parameter 'cronExp' to be non-null");
            $.repoKey = Objects.requireNonNull($.repoKey, "expected parameter 'repoKey' to be non-null");
            return $;
        }
    }

}
