// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class ReplicationConfig extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicationConfigState, opts?: pulumi.CustomResourceOptions): ReplicationConfig {
        return new ReplicationConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/replicationConfig:ReplicationConfig';

    /**
     * Returns true if the given object is an instance of ReplicationConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationConfig.__pulumiType;
    }

    public readonly cronExp!: pulumi.Output<string>;
    public readonly enableEventReplication!: pulumi.Output<boolean>;
    public readonly replications!: pulumi.Output<outputs.ReplicationConfigReplication[] | undefined>;
    public readonly repoKey!: pulumi.Output<string>;

    /**
     * Create a ReplicationConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicationConfigArgs | ReplicationConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplicationConfigState | undefined;
            resourceInputs["cronExp"] = state ? state.cronExp : undefined;
            resourceInputs["enableEventReplication"] = state ? state.enableEventReplication : undefined;
            resourceInputs["replications"] = state ? state.replications : undefined;
            resourceInputs["repoKey"] = state ? state.repoKey : undefined;
        } else {
            const args = argsOrState as ReplicationConfigArgs | undefined;
            if ((!args || args.cronExp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cronExp'");
            }
            if ((!args || args.repoKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repoKey'");
            }
            resourceInputs["cronExp"] = args ? args.cronExp : undefined;
            resourceInputs["enableEventReplication"] = args ? args.enableEventReplication : undefined;
            resourceInputs["replications"] = args ? args.replications : undefined;
            resourceInputs["repoKey"] = args ? args.repoKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplicationConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicationConfig resources.
 */
export interface ReplicationConfigState {
    cronExp?: pulumi.Input<string>;
    enableEventReplication?: pulumi.Input<boolean>;
    replications?: pulumi.Input<pulumi.Input<inputs.ReplicationConfigReplication>[]>;
    repoKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicationConfig resource.
 */
export interface ReplicationConfigArgs {
    cronExp: pulumi.Input<string>;
    enableEventReplication?: pulumi.Input<boolean>;
    replications?: pulumi.Input<pulumi.Input<inputs.ReplicationConfigReplication>[]>;
    repoKey: pulumi.Input<string>;
}
