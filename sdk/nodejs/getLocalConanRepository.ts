// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a local conan repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-conan-repo = artifactory.getLocalConanRepository({
 *     key: "local-test-conan-repo",
 * });
 * ```
 */
export function getLocalConanRepository(args: GetLocalConanRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalConanRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalConanRepository:getLocalConanRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "forceConanAuthentication": args.forceConanAuthentication,
        "includesPattern": args.includesPattern,
        "key": args.key,
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
 * A collection of arguments for invoking getLocalConanRepository.
 */
export interface GetLocalConanRepositoryArgs {
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
    forceConanAuthentication?: boolean;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getLocalConanRepository.
 */
export interface GetLocalConanRepositoryResult {
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
    readonly forceConanAuthentication?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    readonly key: string;
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
 * Retrieves a local conan repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local-test-conan-repo = artifactory.getLocalConanRepository({
 *     key: "local-test-conan-repo",
 * });
 * ```
 */
export function getLocalConanRepositoryOutput(args: GetLocalConanRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocalConanRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getLocalConanRepository:getLocalConanRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "forceConanAuthentication": args.forceConanAuthentication,
        "includesPattern": args.includesPattern,
        "key": args.key,
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
 * A collection of arguments for invoking getLocalConanRepository.
 */
export interface GetLocalConanRepositoryOutputArgs {
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
    forceConanAuthentication?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
