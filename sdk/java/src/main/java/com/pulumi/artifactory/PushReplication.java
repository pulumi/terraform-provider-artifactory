// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.PushReplicationArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.PushReplicationState;
import com.pulumi.artifactory.outputs.PushReplicationReplication;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="artifactory:index/pushReplication:PushReplication")
public class PushReplication extends com.pulumi.resources.CustomResource {
    /**
     * Cron expression to control the operation frequency.
     * 
     */
    @Export(name="cronExp", type=String.class, parameters={})
    private Output<String> cronExp;

    /**
     * @return Cron expression to control the operation frequency.
     * 
     */
    public Output<String> cronExp() {
        return this.cronExp;
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
    @Export(name="replications", type=List.class, parameters={PushReplicationReplication.class})
    private Output</* @Nullable */ List<PushReplicationReplication>> replications;

    public Output<Optional<List<PushReplicationReplication>>> replications() {
        return Codegen.optional(this.replications);
    }
    @Export(name="repoKey", type=String.class, parameters={})
    private Output<String> repoKey;

    public Output<String> repoKey() {
        return this.repoKey;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PushReplication(String name) {
        this(name, PushReplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PushReplication(String name, PushReplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PushReplication(String name, PushReplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/pushReplication:PushReplication", name, args == null ? PushReplicationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PushReplication(String name, Output<String> id, @Nullable PushReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/pushReplication:PushReplication", name, state, makeResourceOptions(options, id));
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
    public static PushReplication get(String name, Output<String> id, @Nullable PushReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PushReplication(name, id, state, options);
    }
}
