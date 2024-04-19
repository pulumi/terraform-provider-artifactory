// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const distribution_custom_webhook = new artifactory.DistributionCustomWebhook("distribution-custom-webhook", {
 *     key: "distribution-custom-webhook",
 *     eventTypes: [
 *         "distribute_started",
 *         "distribute_completed",
 *         "distribute_aborted",
 *         "distribute_failed",
 *         "delete_started",
 *         "delete_completed",
 *         "delete_failed",
 *     ],
 *     criteria: {
 *         anyReleaseBundle: false,
 *         registeredReleaseBundleNames: ["bundle-name"],
 *         includePatterns: ["foo/**"],
 *         excludePatterns: ["bar/**"],
 *     },
 *     handlers: [{
 *         url: "https://tempurl.org",
 *         secrets: {
 *             secretName1: "value1",
 *             secretName2: "value2",
 *         },
 *         httpHeaders: {
 *             headerName1: "value1",
 *             headerName2: "value2",
 *         },
 *         payload: "{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class DistributionCustomWebhook extends pulumi.CustomResource {
    /**
     * Get an existing DistributionCustomWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DistributionCustomWebhookState, opts?: pulumi.CustomResourceOptions): DistributionCustomWebhook {
        return new DistributionCustomWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/distributionCustomWebhook:DistributionCustomWebhook';

    /**
     * Returns true if the given object is an instance of DistributionCustomWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DistributionCustomWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DistributionCustomWebhook.__pulumiType;
    }

    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    public readonly criteria!: pulumi.Output<outputs.DistributionCustomWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Status of webhook. Default to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, `distribute_failed, `deleteStarted`, `deleteCompleted`, `deleteFailed`
     */
    public readonly eventTypes!: pulumi.Output<string[]>;
    /**
     * At least one is required.
     */
    public readonly handlers!: pulumi.Output<outputs.DistributionCustomWebhookHandler[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    public readonly key!: pulumi.Output<string>;

    /**
     * Create a DistributionCustomWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributionCustomWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DistributionCustomWebhookArgs | DistributionCustomWebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DistributionCustomWebhookState | undefined;
            resourceInputs["criteria"] = state ? state.criteria : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["eventTypes"] = state ? state.eventTypes : undefined;
            resourceInputs["handlers"] = state ? state.handlers : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
        } else {
            const args = argsOrState as DistributionCustomWebhookArgs | undefined;
            if ((!args || args.criteria === undefined) && !opts.urn) {
                throw new Error("Missing required property 'criteria'");
            }
            if ((!args || args.eventTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventTypes'");
            }
            if ((!args || args.handlers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'handlers'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["criteria"] = args ? args.criteria : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventTypes"] = args ? args.eventTypes : undefined;
            resourceInputs["handlers"] = args ? args.handlers : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DistributionCustomWebhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DistributionCustomWebhook resources.
 */
export interface DistributionCustomWebhookState {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria?: pulumi.Input<inputs.DistributionCustomWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, `distribute_failed, `deleteStarted`, `deleteCompleted`, `deleteFailed`
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.DistributionCustomWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DistributionCustomWebhook resource.
 */
export interface DistributionCustomWebhookArgs {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria: pulumi.Input<inputs.DistributionCustomWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distributeStarted`, `distributeCompleted`, `distributeAborted`, `distribute_failed, `deleteStarted`, `deleteCompleted`, `deleteFailed`
     */
    eventTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers: pulumi.Input<pulumi.Input<inputs.DistributionCustomWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key: pulumi.Input<string>;
}
