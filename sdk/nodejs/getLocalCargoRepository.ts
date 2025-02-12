// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a local cargo repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local_test_cargo_repo_basic = artifactory.getLocalCargoRepository({
 *     key: "local-test-cargo-repo-basic",
 * });
 * ```
 */
export function getLocalCargoRepository(args: GetLocalCargoRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalCargoRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalCargoRepository:getLocalCargoRepository", {
        "anonymousAccess": args.anonymousAccess,
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "enableSparseIndex": args.enableSparseIndex,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "indexCompressionFormats": args.indexCompressionFormats,
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
 * A collection of arguments for invoking getLocalCargoRepository.
 */
export interface GetLocalCargoRepositoryArgs {
    /**
     * Cargo client does not send credentials when performing download and search for crates.
     * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
     * access option. Default value is `false`.
     */
    anonymousAccess?: boolean;
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    description?: string;
    downloadDirect?: boolean;
    /**
     * Enable internal index support based on Cargo sparse index specifications, instead
     * of the default git index. Default value is `false`.
     */
    enableSparseIndex?: boolean;
    excludesPattern?: string;
    includesPattern?: string;
    indexCompressionFormats?: string[];
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
 * A collection of values returned by getLocalCargoRepository.
 */
export interface GetLocalCargoRepositoryResult {
    /**
     * Cargo client does not send credentials when performing download and search for crates.
     * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
     * access option. Default value is `false`.
     */
    readonly anonymousAccess?: boolean;
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    /**
     * Enable internal index support based on Cargo sparse index specifications, instead
     * of the default git index. Default value is `false`.
     */
    readonly enableSparseIndex?: boolean;
    readonly excludesPattern?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly indexCompressionFormats?: string[];
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
 * Retrieves a local cargo repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const local_test_cargo_repo_basic = artifactory.getLocalCargoRepository({
 *     key: "local-test-cargo-repo-basic",
 * });
 * ```
 */
export function getLocalCargoRepositoryOutput(args: GetLocalCargoRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLocalCargoRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getLocalCargoRepository:getLocalCargoRepository", {
        "anonymousAccess": args.anonymousAccess,
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "enableSparseIndex": args.enableSparseIndex,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "indexCompressionFormats": args.indexCompressionFormats,
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
 * A collection of arguments for invoking getLocalCargoRepository.
 */
export interface GetLocalCargoRepositoryOutputArgs {
    /**
     * Cargo client does not send credentials when performing download and search for crates.
     * Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous
     * access option. Default value is `false`.
     */
    anonymousAccess?: pulumi.Input<boolean>;
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    /**
     * Enable internal index support based on Cargo sparse index specifications, instead
     * of the default git index. Default value is `false`.
     */
    enableSparseIndex?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    indexCompressionFormats?: pulumi.Input<pulumi.Input<string>[]>;
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
