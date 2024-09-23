// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a local Nuget repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-nuget-repo-basic = artifactory.getLocalNugetRepository({
 *     key: "local-test-nuget-repo-basic",
 * });
 * ```
 */
export function getLocalNugetRepository(args: GetLocalNugetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalNugetRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalNugetRepository:getLocalNugetRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "forceNugetAuthentication": args.forceNugetAuthentication,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueSnapshots": args.maxUniqueSnapshots,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "repoLayoutRef": args.repoLayoutRef,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalNugetRepository.
 */
export interface GetLocalNugetRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    description?: string;
    downloadDirect?: boolean;
    excludesPattern?: string;
    /**
     * Force basic authentication credentials in order to use this repository.
     * Default is `false`.
     */
    forceNugetAuthentication?: boolean;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    /**
     * The maximum number of unique snapshots of a single artifact to store Once the
     * number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no
     * limit, and unique snapshots are not cleaned up.
     */
    maxUniqueSnapshots?: number;
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getLocalNugetRepository.
 */
export interface GetLocalNugetRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly excludesPattern?: string;
    /**
     * Force basic authentication credentials in order to use this repository.
     * Default is `false`.
     */
    readonly forceNugetAuthentication?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    /**
     * The maximum number of unique snapshots of a single artifact to store Once the
     * number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no
     * limit, and unique snapshots are not cleaned up.
     */
    readonly maxUniqueSnapshots?: number;
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a local Nuget repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-nuget-repo-basic = artifactory.getLocalNugetRepository({
 *     key: "local-test-nuget-repo-basic",
 * });
 * ```
 */
export function getLocalNugetRepositoryOutput(args: GetLocalNugetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocalNugetRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getLocalNugetRepository:getLocalNugetRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "forceNugetAuthentication": args.forceNugetAuthentication,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueSnapshots": args.maxUniqueSnapshots,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "repoLayoutRef": args.repoLayoutRef,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalNugetRepository.
 */
export interface GetLocalNugetRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * Force basic authentication credentials in order to use this repository.
     * Default is `false`.
     */
    forceNugetAuthentication?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    /**
     * The maximum number of unique snapshots of a single artifact to store Once the
     * number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no
     * limit, and unique snapshots are not cleaned up.
     */
    maxUniqueSnapshots?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
