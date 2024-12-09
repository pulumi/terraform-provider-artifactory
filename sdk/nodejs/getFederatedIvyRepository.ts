// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a federated Ivy repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const federated-test-ivy-repo = artifactory.getFederatedIvyRepository({
 *     key: "federated-test-ivy-repo",
 * });
 * ```
 */
export function getFederatedIvyRepository(args: GetFederatedIvyRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetFederatedIvyRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getFederatedIvyRepository:getFederatedIvyRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "checksumPolicyType": args.checksumPolicyType,
        "cleanupOnDelete": args.cleanupOnDelete,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "handleReleases": args.handleReleases,
        "handleSnapshots": args.handleSnapshots,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueSnapshots": args.maxUniqueSnapshots,
        "members": args.members,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "repoLayoutRef": args.repoLayoutRef,
        "snapshotVersionBehavior": args.snapshotVersionBehavior,
        "suppressPomConsistencyChecks": args.suppressPomConsistencyChecks,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getFederatedIvyRepository.
 */
export interface GetFederatedIvyRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    checksumPolicyType?: string;
    cleanupOnDelete?: boolean;
    description?: string;
    /**
     * When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     */
    disableProxy?: boolean;
    downloadDirect?: boolean;
    excludesPattern?: string;
    handleReleases?: boolean;
    handleSnapshots?: boolean;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    maxUniqueSnapshots?: number;
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    members?: inputs.GetFederatedIvyRepositoryMember[];
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    /**
     * Proxy key from Artifactory Proxies settings.
     */
    proxy?: string;
    repoLayoutRef?: string;
    snapshotVersionBehavior?: string;
    suppressPomConsistencyChecks?: boolean;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getFederatedIvyRepository.
 */
export interface GetFederatedIvyRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    readonly checksumPolicyType?: string;
    readonly cleanupOnDelete?: boolean;
    readonly description?: string;
    /**
     * When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     */
    readonly disableProxy?: boolean;
    readonly downloadDirect?: boolean;
    readonly excludesPattern?: string;
    readonly handleReleases?: boolean;
    readonly handleSnapshots?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    readonly maxUniqueSnapshots?: number;
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    readonly members?: outputs.GetFederatedIvyRepositoryMember[];
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    /**
     * Proxy key from Artifactory Proxies settings.
     */
    readonly proxy?: string;
    readonly repoLayoutRef?: string;
    readonly snapshotVersionBehavior?: string;
    readonly suppressPomConsistencyChecks?: boolean;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a federated Ivy repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const federated-test-ivy-repo = artifactory.getFederatedIvyRepository({
 *     key: "federated-test-ivy-repo",
 * });
 * ```
 */
export function getFederatedIvyRepositoryOutput(args: GetFederatedIvyRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFederatedIvyRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getFederatedIvyRepository:getFederatedIvyRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "checksumPolicyType": args.checksumPolicyType,
        "cleanupOnDelete": args.cleanupOnDelete,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "handleReleases": args.handleReleases,
        "handleSnapshots": args.handleSnapshots,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueSnapshots": args.maxUniqueSnapshots,
        "members": args.members,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "repoLayoutRef": args.repoLayoutRef,
        "snapshotVersionBehavior": args.snapshotVersionBehavior,
        "suppressPomConsistencyChecks": args.suppressPomConsistencyChecks,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getFederatedIvyRepository.
 */
export interface GetFederatedIvyRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    checksumPolicyType?: pulumi.Input<string>;
    cleanupOnDelete?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    /**
     * When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
     */
    disableProxy?: pulumi.Input<boolean>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    handleReleases?: pulumi.Input<boolean>;
    handleSnapshots?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    maxUniqueSnapshots?: pulumi.Input<number>;
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    members?: pulumi.Input<pulumi.Input<inputs.GetFederatedIvyRepositoryMemberArgs>[]>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Proxy key from Artifactory Proxies settings.
     */
    proxy?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    snapshotVersionBehavior?: pulumi.Input<string>;
    suppressPomConsistencyChecks?: pulumi.Input<boolean>;
    xrayIndex?: pulumi.Input<boolean>;
}
