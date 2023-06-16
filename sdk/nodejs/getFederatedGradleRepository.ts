// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a federated Gradle repository.
 */
export function getFederatedGradleRepository(args: GetFederatedGradleRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetFederatedGradleRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getFederatedGradleRepository:getFederatedGradleRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "checksumPolicyType": args.checksumPolicyType,
        "cleanupOnDelete": args.cleanupOnDelete,
        "description": args.description,
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
        "repoLayoutRef": args.repoLayoutRef,
        "snapshotVersionBehavior": args.snapshotVersionBehavior,
        "suppressPomConsistencyChecks": args.suppressPomConsistencyChecks,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getFederatedGradleRepository.
 */
export interface GetFederatedGradleRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    checksumPolicyType?: string;
    cleanupOnDelete?: boolean;
    description?: string;
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
    members?: inputs.GetFederatedGradleRepositoryMember[];
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    snapshotVersionBehavior?: string;
    suppressPomConsistencyChecks?: boolean;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getFederatedGradleRepository.
 */
export interface GetFederatedGradleRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    readonly checksumPolicyType?: string;
    readonly cleanupOnDelete?: boolean;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly excludesPattern: string;
    readonly handleReleases?: boolean;
    readonly handleSnapshots?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern: string;
    readonly key: string;
    readonly maxUniqueSnapshots?: number;
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    readonly members?: outputs.GetFederatedGradleRepositoryMember[];
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    readonly snapshotVersionBehavior?: string;
    readonly suppressPomConsistencyChecks?: boolean;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a federated Gradle repository.
 */
export function getFederatedGradleRepositoryOutput(args: GetFederatedGradleRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFederatedGradleRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getFederatedGradleRepository(a, opts))
}

/**
 * A collection of arguments for invoking getFederatedGradleRepository.
 */
export interface GetFederatedGradleRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    checksumPolicyType?: pulumi.Input<string>;
    cleanupOnDelete?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
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
    members?: pulumi.Input<pulumi.Input<inputs.GetFederatedGradleRepositoryMemberArgs>[]>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    snapshotVersionBehavior?: pulumi.Input<string>;
    suppressPomConsistencyChecks?: pulumi.Input<boolean>;
    xrayIndex?: pulumi.Input<boolean>;
}
