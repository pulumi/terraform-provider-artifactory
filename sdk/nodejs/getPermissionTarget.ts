// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # Artifactory Permission Target Data Source
 *
 * Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * //
 * const target1 = artifactory.getPermissionTarget({
 *     name: "my_permission",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPermissionTarget(args: GetPermissionTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetPermissionTargetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getPermissionTarget:getPermissionTarget", {
        "build": args.build,
        "name": args.name,
        "releaseBundle": args.releaseBundle,
        "repo": args.repo,
    }, opts);
}

/**
 * A collection of arguments for invoking getPermissionTarget.
 */
export interface GetPermissionTargetArgs {
    /**
     * Same as repo but for artifactory-build-info permissions.
     */
    build?: inputs.GetPermissionTargetBuild;
    /**
     * Name of the permission target.
     */
    name: string;
    /**
     * Same as repo but for release-bundles permissions.
     */
    releaseBundle?: inputs.GetPermissionTargetReleaseBundle;
    /**
     * Repository permission configuration.
     */
    repo?: inputs.GetPermissionTargetRepo;
}

/**
 * A collection of values returned by getPermissionTarget.
 */
export interface GetPermissionTargetResult {
    /**
     * Same as repo but for artifactory-build-info permissions.
     */
    readonly build?: outputs.GetPermissionTargetBuild;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * Same as repo but for release-bundles permissions.
     */
    readonly releaseBundle?: outputs.GetPermissionTargetReleaseBundle;
    /**
     * Repository permission configuration.
     */
    readonly repo?: outputs.GetPermissionTargetRepo;
}
/**
 * ## # Artifactory Permission Target Data Source
 *
 * Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * //
 * const target1 = artifactory.getPermissionTarget({
 *     name: "my_permission",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPermissionTargetOutput(args: GetPermissionTargetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPermissionTargetResult> {
    return pulumi.output(args).apply((a: any) => getPermissionTarget(a, opts))
}

/**
 * A collection of arguments for invoking getPermissionTarget.
 */
export interface GetPermissionTargetOutputArgs {
    /**
     * Same as repo but for artifactory-build-info permissions.
     */
    build?: pulumi.Input<inputs.GetPermissionTargetBuildArgs>;
    /**
     * Name of the permission target.
     */
    name: pulumi.Input<string>;
    /**
     * Same as repo but for release-bundles permissions.
     */
    releaseBundle?: pulumi.Input<inputs.GetPermissionTargetReleaseBundleArgs>;
    /**
     * Repository permission configuration.
     */
    repo?: pulumi.Input<inputs.GetPermissionTargetRepoArgs>;
}
