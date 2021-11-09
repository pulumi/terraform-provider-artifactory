// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Xray Watch Resource
 *
 * Provides a Xray watch resource.
 *
 * ## Import
 *
 * Watches can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/xrayWatch:XrayWatch example watch-name
 * ```
 */
export class XrayWatch extends pulumi.CustomResource {
    /**
     * Get an existing XrayWatch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: XrayWatchState, opts?: pulumi.CustomResourceOptions): XrayWatch {
        return new XrayWatch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/xrayWatch:XrayWatch';

    /**
     * Returns true if the given object is an instance of XrayWatch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is XrayWatch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === XrayWatch.__pulumiType;
    }

    /**
     * Whether or not the watch will be active
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * Nested argument describing policies that will be applied. Defined below.
     */
    public readonly assignedPolicies!: pulumi.Output<outputs.XrayWatchAssignedPolicy[]>;
    /**
     * Description of the watch
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the watch (must be unique)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Nested argument describing the resources to be watched. Defined below.
     */
    public readonly resources!: pulumi.Output<outputs.XrayWatchResource[]>;
    public readonly watchRecipients!: pulumi.Output<string[] | undefined>;

    /**
     * Create a XrayWatch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: XrayWatchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: XrayWatchArgs | XrayWatchState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as XrayWatchState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["assignedPolicies"] = state ? state.assignedPolicies : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resources"] = state ? state.resources : undefined;
            inputs["watchRecipients"] = state ? state.watchRecipients : undefined;
        } else {
            const args = argsOrState as XrayWatchArgs | undefined;
            if ((!args || args.assignedPolicies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assignedPolicies'");
            }
            if ((!args || args.resources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resources'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["assignedPolicies"] = args ? args.assignedPolicies : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resources"] = args ? args.resources : undefined;
            inputs["watchRecipients"] = args ? args.watchRecipients : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(XrayWatch.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering XrayWatch resources.
 */
export interface XrayWatchState {
    /**
     * Whether or not the watch will be active
     */
    active?: pulumi.Input<boolean>;
    /**
     * Nested argument describing policies that will be applied. Defined below.
     */
    assignedPolicies?: pulumi.Input<pulumi.Input<inputs.XrayWatchAssignedPolicy>[]>;
    /**
     * Description of the watch
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the watch (must be unique)
     */
    name?: pulumi.Input<string>;
    /**
     * Nested argument describing the resources to be watched. Defined below.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.XrayWatchResource>[]>;
    watchRecipients?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a XrayWatch resource.
 */
export interface XrayWatchArgs {
    /**
     * Whether or not the watch will be active
     */
    active?: pulumi.Input<boolean>;
    /**
     * Nested argument describing policies that will be applied. Defined below.
     */
    assignedPolicies: pulumi.Input<pulumi.Input<inputs.XrayWatchAssignedPolicy>[]>;
    /**
     * Description of the watch
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the watch (must be unique)
     */
    name?: pulumi.Input<string>;
    /**
     * Nested argument describing the resources to be watched. Defined below.
     */
    resources: pulumi.Input<pulumi.Input<inputs.XrayWatchResource>[]>;
    watchRecipients?: pulumi.Input<pulumi.Input<string>[]>;
}
