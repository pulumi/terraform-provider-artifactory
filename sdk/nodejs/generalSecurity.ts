// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can be used to manage Artifactory's general security settings.
 *
 * Only a single `artifactory.GeneralSecurity` resource is meant to be defined.
 *
 * ~>The `artifactory.GeneralSecurity` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // Configure Artifactory general security settings
 * const security = new artifactory.GeneralSecurity("security", {enableAnonymousAccess: true});
 * ```
 *
 * ## Import
 *
 * Current general security settings can be imported using `security` as the `ID`, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/generalSecurity:GeneralSecurity security security
 * ```
 * ~>The `artifactory_general_security` resource uses endpoints that are undocumented and may not work with SaaS
 * environments, or may change without notice.
 */
export class GeneralSecurity extends pulumi.CustomResource {
    /**
     * Get an existing GeneralSecurity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GeneralSecurityState, opts?: pulumi.CustomResourceOptions): GeneralSecurity {
        return new GeneralSecurity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/generalSecurity:GeneralSecurity';

    /**
     * Returns true if the given object is an instance of GeneralSecurity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GeneralSecurity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GeneralSecurity.__pulumiType;
    }

    /**
     * Enable anonymous access.  Default value is `false`.
     */
    public readonly enableAnonymousAccess!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GeneralSecurity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GeneralSecurityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GeneralSecurityArgs | GeneralSecurityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GeneralSecurityState | undefined;
            resourceInputs["enableAnonymousAccess"] = state ? state.enableAnonymousAccess : undefined;
        } else {
            const args = argsOrState as GeneralSecurityArgs | undefined;
            resourceInputs["enableAnonymousAccess"] = args ? args.enableAnonymousAccess : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GeneralSecurity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GeneralSecurity resources.
 */
export interface GeneralSecurityState {
    /**
     * Enable anonymous access.  Default value is `false`.
     */
    enableAnonymousAccess?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GeneralSecurity resource.
 */
export interface GeneralSecurityArgs {
    /**
     * Enable anonymous access.  Default value is `false`.
     */
    enableAnonymousAccess?: pulumi.Input<boolean>;
}
