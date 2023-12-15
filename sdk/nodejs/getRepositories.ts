// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Returns a list of minimal repository details for all repositories of the specified type.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const all-alpine-local = artifactory.getRepositories({
 *     packageType: "alpine",
 *     repositoryType: "local",
 * });
 * ```
 */
export function getRepositories(args?: GetRepositoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoriesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getRepositories:getRepositories", {
        "packageType": args.packageType,
        "projectKey": args.projectKey,
        "repositoryType": args.repositoryType,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositories.
 */
export interface GetRepositoriesArgs {
    packageType?: string;
    /**
     * Filter for repositories assigned to a specific project.
     */
    projectKey?: string;
    /**
     * Filter for repositories of a specific type. Allowed values are: local, remote, virtual, federated, distribution
     */
    repositoryType?: string;
}

/**
 * A collection of values returned by getRepositories.
 */
export interface GetRepositoriesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly packageType?: string;
    /**
     * Filter for repositories assigned to a specific project.
     */
    readonly projectKey?: string;
    /**
     * A list of repositories.
     */
    readonly repos: outputs.GetRepositoriesRepo[];
    /**
     * Filter for repositories of a specific type. Allowed values are: local, remote, virtual, federated, distribution
     */
    readonly repositoryType?: string;
}
/**
 * Returns a list of minimal repository details for all repositories of the specified type.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const all-alpine-local = artifactory.getRepositories({
 *     packageType: "alpine",
 *     repositoryType: "local",
 * });
 * ```
 */
export function getRepositoriesOutput(args?: GetRepositoriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoriesResult> {
    return pulumi.output(args).apply((a: any) => getRepositories(a, opts))
}

/**
 * A collection of arguments for invoking getRepositories.
 */
export interface GetRepositoriesOutputArgs {
    packageType?: pulumi.Input<string>;
    /**
     * Filter for repositories assigned to a specific project.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * Filter for repositories of a specific type. Allowed values are: local, remote, virtual, federated, distribution
     */
    repositoryType?: pulumi.Input<string>;
}
