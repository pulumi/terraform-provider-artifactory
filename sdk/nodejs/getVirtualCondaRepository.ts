// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a virtual Conda repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const virtual-conda = artifactory.getVirtualCondaRepository({
 *     key: "virtual-conda",
 * });
 * ```
 */
export function getVirtualCondaRepository(args: GetVirtualCondaRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualCondaRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getVirtualCondaRepository:getVirtualCondaRepository", {
        "artifactoryRequestsCanRetrieveRemoteArtifacts": args.artifactoryRequestsCanRetrieveRemoteArtifacts,
        "defaultDeploymentRepo": args.defaultDeploymentRepo,
        "description": args.description,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "notes": args.notes,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "repoLayoutRef": args.repoLayoutRef,
        "repositories": args.repositories,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualCondaRepository.
 */
export interface GetVirtualCondaRepositoryArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: boolean;
    defaultDeploymentRepo?: string;
    description?: string;
    excludesPattern?: string;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    notes?: string;
    projectEnvironments?: string[];
    projectKey?: string;
    repoLayoutRef?: string;
    repositories?: string[];
    /**
     * (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: number;
}

/**
 * A collection of values returned by getVirtualCondaRepository.
 */
export interface GetVirtualCondaRepositoryResult {
    readonly artifactoryRequestsCanRetrieveRemoteArtifacts?: boolean;
    readonly defaultDeploymentRepo?: string;
    readonly description?: string;
    readonly excludesPattern?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    readonly notes?: string;
    readonly packageType: string;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly repoLayoutRef?: string;
    readonly repositories?: string[];
    /**
     * (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     */
    readonly retrievalCachePeriodSeconds?: number;
}
/**
 * Retrieves a virtual Conda repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const virtual-conda = artifactory.getVirtualCondaRepository({
 *     key: "virtual-conda",
 * });
 * ```
 */
export function getVirtualCondaRepositoryOutput(args: GetVirtualCondaRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualCondaRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getVirtualCondaRepository:getVirtualCondaRepository", {
        "artifactoryRequestsCanRetrieveRemoteArtifacts": args.artifactoryRequestsCanRetrieveRemoteArtifacts,
        "defaultDeploymentRepo": args.defaultDeploymentRepo,
        "description": args.description,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "notes": args.notes,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "repoLayoutRef": args.repoLayoutRef,
        "repositories": args.repositories,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualCondaRepository.
 */
export interface GetVirtualCondaRepositoryOutputArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: pulumi.Input<boolean>;
    defaultDeploymentRepo?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
}
