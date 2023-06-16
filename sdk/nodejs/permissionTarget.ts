// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // Create a new Artifactory permission target called testpermission
 * const test_perm = new artifactory.PermissionTarget("test-perm", {
 *     build: {
 *         actions: {
 *             users: [{
 *                 name: "anonymous",
 *                 permissions: [
 *                     "read",
 *                     "write",
 *                 ],
 *             }],
 *         },
 *         includesPatterns: ["**"],
 *         repositories: ["artifactory-build-info"],
 *     },
 *     releaseBundle: {
 *         actions: {
 *             users: [{
 *                 name: "anonymous",
 *                 permissions: ["read"],
 *             }],
 *         },
 *         includesPatterns: ["**"],
 *         repositories: ["release-bundles"],
 *     },
 *     repo: {
 *         actions: {
 *             groups: [{
 *                 name: "readers",
 *                 permissions: ["read"],
 *             }],
 *             users: [{
 *                 name: "anonymous",
 *                 permissions: [
 *                     "read",
 *                     "write",
 *                 ],
 *             }],
 *         },
 *         excludesPatterns: ["bar/**"],
 *         includesPatterns: ["foo/**"],
 *         repositories: ["example-repo-local"],
 *     },
 * });
 * ```
 * ## Permissions
 *
 * The provider supports the following `permission` enums:
 *
 * * `read`
 * * `write`
 * * `annotate`
 * * `delete`
 * * `manage`
 * * `managedXrayMeta`
 * * `distribute`
 *
 * The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):
 *
 * * `read` - matches `Read` permissions.
 * * `write` - matches `  Deploy / Cache / Create ` permissions.
 * * `annotate` - matches `Annotate` permissions.
 * * `delete` - matches `Delete / Overwrite` permissions.
 * * `manage` - matches `Manage` permissions.
 * * `managedXrayMeta` - matches `Manage Xray Metadata` permissions.
 * * `distribute` - matches `Distribute` permissions.
 *
 * ## Import
 *
 * Permission targets can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
 * ```
 */
export class PermissionTarget extends pulumi.CustomResource {
    /**
     * Get an existing PermissionTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PermissionTargetState, opts?: pulumi.CustomResourceOptions): PermissionTarget {
        return new PermissionTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/permissionTarget:PermissionTarget';

    /**
     * Returns true if the given object is an instance of PermissionTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PermissionTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PermissionTarget.__pulumiType;
    }

    /**
     * As for repo but for artifactory-build-info permissions.
     */
    public readonly build!: pulumi.Output<outputs.PermissionTargetBuild | undefined>;
    /**
     * Name of permission.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * As for repo for for release-bundles permissions.
     */
    public readonly releaseBundle!: pulumi.Output<outputs.PermissionTargetReleaseBundle | undefined>;
    /**
     * Repository permission configuration.
     */
    public readonly repo!: pulumi.Output<outputs.PermissionTargetRepo | undefined>;

    /**
     * Create a PermissionTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PermissionTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PermissionTargetArgs | PermissionTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PermissionTargetState | undefined;
            resourceInputs["build"] = state ? state.build : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["releaseBundle"] = state ? state.releaseBundle : undefined;
            resourceInputs["repo"] = state ? state.repo : undefined;
        } else {
            const args = argsOrState as PermissionTargetArgs | undefined;
            resourceInputs["build"] = args ? args.build : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["releaseBundle"] = args ? args.releaseBundle : undefined;
            resourceInputs["repo"] = args ? args.repo : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PermissionTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PermissionTarget resources.
 */
export interface PermissionTargetState {
    /**
     * As for repo but for artifactory-build-info permissions.
     */
    build?: pulumi.Input<inputs.PermissionTargetBuild>;
    /**
     * Name of permission.
     */
    name?: pulumi.Input<string>;
    /**
     * As for repo for for release-bundles permissions.
     */
    releaseBundle?: pulumi.Input<inputs.PermissionTargetReleaseBundle>;
    /**
     * Repository permission configuration.
     */
    repo?: pulumi.Input<inputs.PermissionTargetRepo>;
}

/**
 * The set of arguments for constructing a PermissionTarget resource.
 */
export interface PermissionTargetArgs {
    /**
     * As for repo but for artifactory-build-info permissions.
     */
    build?: pulumi.Input<inputs.PermissionTargetBuild>;
    /**
     * Name of permission.
     */
    name?: pulumi.Input<string>;
    /**
     * As for repo for for release-bundles permissions.
     */
    releaseBundle?: pulumi.Input<inputs.PermissionTargetReleaseBundle>;
    /**
     * Repository permission configuration.
     */
    repo?: pulumi.Input<inputs.PermissionTargetRepo>;
}
