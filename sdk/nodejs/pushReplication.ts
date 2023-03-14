// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory push replication resource. This can be used to create and manage Artifactory push replications.
 * Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near
 * end invoking a synchronization of artifacts to the far end.
 * See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const config = new pulumi.Config();
 * const artifactoryUrl = config.require("artifactoryUrl");
 * const artifactoryUsername = config.require("artifactoryUsername");
 * const artifactoryPassword = config.require("artifactoryPassword");
 * // Create a replication between two artifactory local repositories
 * const providerTestSource = new artifactory.LocalMavenRepository("providerTestSource", {key: "provider_test_source"});
 * const providerTestDest = new artifactory.LocalMavenRepository("providerTestDest", {key: "provider_test_dest"});
 * const foo_rep = new artifactory.PushReplication("foo-rep", {
 *     repoKey: providerTestSource.key,
 *     cronExp: "0 0 * * * ?",
 *     enableEventReplication: true,
 *     replications: [{
 *         url: pulumi.interpolate`${artifactoryUrl}/${providerTestDest.key}`,
 *         username: "$var.artifactory_username",
 *         password: "$var.artifactory_password",
 *         enabled: true,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Push replication configs can be imported using their repo key, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/pushReplication:PushReplication foo-rep provider_test_source
 * ```
 */
export class PushReplication extends pulumi.CustomResource {
    /**
     * Get an existing PushReplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PushReplicationState, opts?: pulumi.CustomResourceOptions): PushReplication {
        return new PushReplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/pushReplication:PushReplication';

    /**
     * Returns true if the given object is an instance of PushReplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PushReplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PushReplication.__pulumiType;
    }

    /**
     * A valid CRON expression that you can use to control replication frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     */
    public readonly cronExp!: pulumi.Output<string>;
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
     */
    public readonly enableEventReplication!: pulumi.Output<boolean>;
    public readonly replications!: pulumi.Output<outputs.PushReplicationReplication[] | undefined>;
    /**
     * Repository name.
     */
    public readonly repoKey!: pulumi.Output<string>;

    /**
     * Create a PushReplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PushReplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PushReplicationArgs | PushReplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PushReplicationState | undefined;
            resourceInputs["cronExp"] = state ? state.cronExp : undefined;
            resourceInputs["enableEventReplication"] = state ? state.enableEventReplication : undefined;
            resourceInputs["replications"] = state ? state.replications : undefined;
            resourceInputs["repoKey"] = state ? state.repoKey : undefined;
        } else {
            const args = argsOrState as PushReplicationArgs | undefined;
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
        super(PushReplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PushReplication resources.
 */
export interface PushReplicationState {
    /**
     * A valid CRON expression that you can use to control replication frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     */
    cronExp?: pulumi.Input<string>;
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
     */
    enableEventReplication?: pulumi.Input<boolean>;
    replications?: pulumi.Input<pulumi.Input<inputs.PushReplicationReplication>[]>;
    /**
     * Repository name.
     */
    repoKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PushReplication resource.
 */
export interface PushReplicationArgs {
    /**
     * A valid CRON expression that you can use to control replication frequency. Eg: "0 0 12 * * ? *", "0 0 2 ? * MON-SAT *". Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     */
    cronExp: pulumi.Input<string>;
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
     */
    enableEventReplication?: pulumi.Input<boolean>;
    replications?: pulumi.Input<pulumi.Input<inputs.PushReplicationReplication>[]>;
    /**
     * Repository name.
     */
    repoKey: pulumi.Input<string>;
}
