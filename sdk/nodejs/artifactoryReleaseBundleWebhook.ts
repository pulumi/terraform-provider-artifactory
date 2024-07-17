// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 *
 * ~>This resource is being deprecated. Use `artifactory.DestinationWebhook` instead.
 *
 * ## Example Usage
 *
 * .
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const artifactory_release_bundle_webhook = new artifactory.ArtifactoryReleaseBundleWebhook("artifactory-release-bundle-webhook", {
 *     key: "artifactory-release-bundle-webhook",
 *     eventTypes: [
 *         "received",
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
 *         url: "http://tempurl.org/webhook",
 *         secret: "some-secret",
 *         proxy: "proxy-key",
 *         customHttpHeaders: {
 *             "header-1": "value-1",
 *             "header-2": "value-2",
 *         },
 *     }],
 * });
 * ```
 */
export class ArtifactoryReleaseBundleWebhook extends pulumi.CustomResource {
    /**
     * Get an existing ArtifactoryReleaseBundleWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ArtifactoryReleaseBundleWebhookState, opts?: pulumi.CustomResourceOptions): ArtifactoryReleaseBundleWebhook {
        return new ArtifactoryReleaseBundleWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/artifactoryReleaseBundleWebhook:ArtifactoryReleaseBundleWebhook';

    /**
     * Returns true if the given object is an instance of ArtifactoryReleaseBundleWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ArtifactoryReleaseBundleWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ArtifactoryReleaseBundleWebhook.__pulumiType;
    }

    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    public readonly criteria!: pulumi.Output<outputs.ArtifactoryReleaseBundleWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Status of webhook. Default to `true`
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
     */
    public readonly eventTypes!: pulumi.Output<string[]>;
    /**
     * At least one is required.
     */
    public readonly handlers!: pulumi.Output<outputs.ArtifactoryReleaseBundleWebhookHandler[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    public readonly key!: pulumi.Output<string>;

    /**
     * Create a ArtifactoryReleaseBundleWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArtifactoryReleaseBundleWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ArtifactoryReleaseBundleWebhookArgs | ArtifactoryReleaseBundleWebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ArtifactoryReleaseBundleWebhookState | undefined;
            resourceInputs["criteria"] = state ? state.criteria : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["eventTypes"] = state ? state.eventTypes : undefined;
            resourceInputs["handlers"] = state ? state.handlers : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
        } else {
            const args = argsOrState as ArtifactoryReleaseBundleWebhookArgs | undefined;
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
        super(ArtifactoryReleaseBundleWebhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ArtifactoryReleaseBundleWebhook resources.
 */
export interface ArtifactoryReleaseBundleWebhookState {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria?: pulumi.Input<inputs.ArtifactoryReleaseBundleWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.ArtifactoryReleaseBundleWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ArtifactoryReleaseBundleWebhook resource.
 */
export interface ArtifactoryReleaseBundleWebhookArgs {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria: pulumi.Input<inputs.ArtifactoryReleaseBundleWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `received`, `deleteStarted`, `deleteCompleted`, `deleteFailed`
     */
    eventTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers: pulumi.Input<pulumi.Input<inputs.ArtifactoryReleaseBundleWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key: pulumi.Input<string>;
}
