// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory API key resource. This can be used to create and manage Artifactory API keys.
 *
 * > **Note:** API keys will be stored in the raw state as plain-text. Read more about sensitive data in state.
 *
 * !> As notified in [Artifactory 7.47.10](https://jfrog.com/help/r/jfrog-release-information/artifactory-7.47.10-cloud-self-hosted), support for API Key is slated to be removed in a future release. To ease customer migration to [reference tokens](https://jfrog.com/help/r/jfrog-platform-administration-documentation/user-profile), which replaces API key, we are disabling the ability to create new API keys at the end of Q3 2024. The ability to use API keys will be removed at the end of Q4 2024. For more information, see [JFrog API Key Deprecation Process](https://jfrog.com/help/r/jfrog-platform-administration-documentation/jfrog-api-key-deprecation-process).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // Create a new Artifactory API key for the configured user
 * const ci = new artifactory.ApiKey("ci", {});
 * ```
 *
 * ## Import
 *
 * A user's API key can be imported using any identifier, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/apiKey:ApiKey test import
 * ```
 */
export class ApiKey extends pulumi.CustomResource {
    /**
     * Get an existing ApiKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiKeyState, opts?: pulumi.CustomResourceOptions): ApiKey {
        return new ApiKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/apiKey:ApiKey';

    /**
     * Returns true if the given object is an instance of ApiKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiKey.__pulumiType;
    }

    /**
     * The API key. Deprecated.
     *
     * @deprecated Deprecated in favor of "artifactory.ScopedToken".
     */
    public /*out*/ readonly apiKey!: pulumi.Output<string>;

    /**
     * Create a ApiKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApiKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiKeyArgs | ApiKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiKeyState | undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
        } else {
            const args = argsOrState as ApiKeyArgs | undefined;
            resourceInputs["apiKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApiKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiKey resources.
 */
export interface ApiKeyState {
    /**
     * The API key. Deprecated.
     *
     * @deprecated Deprecated in favor of "artifactory.ScopedToken".
     */
    apiKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiKey resource.
 */
export interface ApiKeyArgs {
}
