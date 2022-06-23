// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 */
export class ArtifactPropertyWebhook extends pulumi.CustomResource {
    /**
     * Get an existing ArtifactPropertyWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ArtifactPropertyWebhookState, opts?: pulumi.CustomResourceOptions): ArtifactPropertyWebhook {
        return new ArtifactPropertyWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/artifactPropertyWebhook:ArtifactPropertyWebhook';

    /**
     * Returns true if the given object is an instance of ArtifactPropertyWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ArtifactPropertyWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ArtifactPropertyWebhook.__pulumiType;
    }

    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    public readonly criteria!: pulumi.Output<outputs.ArtifactPropertyWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Status of webhook. Default to 'true'.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "added", "deleted".
     */
    public readonly eventTypes!: pulumi.Output<string[]>;
    /**
     * At least one is required.
     */
    public readonly handlers!: pulumi.Output<outputs.ArtifactPropertyWebhookHandler[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    public readonly key!: pulumi.Output<string>;

    /**
     * Create a ArtifactPropertyWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArtifactPropertyWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ArtifactPropertyWebhookArgs | ArtifactPropertyWebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ArtifactPropertyWebhookState | undefined;
            resourceInputs["criteria"] = state ? state.criteria : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["eventTypes"] = state ? state.eventTypes : undefined;
            resourceInputs["handlers"] = state ? state.handlers : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
        } else {
            const args = argsOrState as ArtifactPropertyWebhookArgs | undefined;
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
        super(ArtifactPropertyWebhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ArtifactPropertyWebhook resources.
 */
export interface ArtifactPropertyWebhookState {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria?: pulumi.Input<inputs.ArtifactPropertyWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to 'true'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "added", "deleted".
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.ArtifactPropertyWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ArtifactPropertyWebhook resource.
 */
export interface ArtifactPropertyWebhookArgs {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria: pulumi.Input<inputs.ArtifactPropertyWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to 'true'.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "added", "deleted".
     */
    eventTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers: pulumi.Input<pulumi.Input<inputs.ArtifactPropertyWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key: pulumi.Input<string>;
}
