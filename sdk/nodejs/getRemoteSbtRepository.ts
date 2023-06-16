// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a remote SBT repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote-sbt = artifactory.getRemoteSbtRepository({
 *     key: "remote-sbt",
 * });
 * ```
 */
export function getRemoteSbtRepository(args: GetRemoteSbtRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteSbtRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getRemoteSbtRepository:getRemoteSbtRepository", {
        "allowAnyHostAuth": args.allowAnyHostAuth,
        "assumedOfflinePeriodSecs": args.assumedOfflinePeriodSecs,
        "blackedOut": args.blackedOut,
        "blockMismatchingMimeTypes": args.blockMismatchingMimeTypes,
        "bypassHeadRequests": args.bypassHeadRequests,
        "cdnRedirect": args.cdnRedirect,
        "clientTlsCertificate": args.clientTlsCertificate,
        "contentSynchronisation": args.contentSynchronisation,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "enableCookieManagement": args.enableCookieManagement,
        "excludesPattern": args.excludesPattern,
        "fetchJarsEagerly": args.fetchJarsEagerly,
        "fetchSourcesEagerly": args.fetchSourcesEagerly,
        "handleReleases": args.handleReleases,
        "handleSnapshots": args.handleSnapshots,
        "hardFail": args.hardFail,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "listRemoteFolderItems": args.listRemoteFolderItems,
        "localAddress": args.localAddress,
        "metadataRetrievalTimeoutSecs": args.metadataRetrievalTimeoutSecs,
        "mismatchingMimeTypesOverrideList": args.mismatchingMimeTypesOverrideList,
        "missedCachePeriodSeconds": args.missedCachePeriodSeconds,
        "notes": args.notes,
        "offline": args.offline,
        "password": args.password,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "queryParams": args.queryParams,
        "rejectInvalidJars": args.rejectInvalidJars,
        "remoteRepoChecksumPolicyType": args.remoteRepoChecksumPolicyType,
        "remoteRepoLayoutRef": args.remoteRepoLayoutRef,
        "repoLayoutRef": args.repoLayoutRef,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "shareConfiguration": args.shareConfiguration,
        "socketTimeoutMillis": args.socketTimeoutMillis,
        "storeArtifactsLocally": args.storeArtifactsLocally,
        "suppressPomConsistencyChecks": args.suppressPomConsistencyChecks,
        "synchronizeProperties": args.synchronizeProperties,
        "unusedArtifactsCleanupPeriodHours": args.unusedArtifactsCleanupPeriodHours,
        "url": args.url,
        "username": args.username,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getRemoteSbtRepository.
 */
export interface GetRemoteSbtRepositoryArgs {
    allowAnyHostAuth?: boolean;
    assumedOfflinePeriodSecs?: number;
    blackedOut?: boolean;
    blockMismatchingMimeTypes?: boolean;
    bypassHeadRequests?: boolean;
    cdnRedirect?: boolean;
    clientTlsCertificate?: string;
    contentSynchronisation?: inputs.GetRemoteSbtRepositoryContentSynchronisation;
    description?: string;
    downloadDirect?: boolean;
    enableCookieManagement?: boolean;
    excludesPattern?: string;
    /**
     * (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
     */
    fetchJarsEagerly?: boolean;
    /**
     * (Optional, Default: `false`) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
     */
    fetchSourcesEagerly?: boolean;
    /**
     * (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
     */
    handleReleases?: boolean;
    /**
     * (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
     */
    handleSnapshots?: boolean;
    hardFail?: boolean;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    listRemoteFolderItems?: boolean;
    localAddress?: string;
    metadataRetrievalTimeoutSecs?: number;
    mismatchingMimeTypesOverrideList?: string;
    missedCachePeriodSeconds?: number;
    notes?: string;
    offline?: boolean;
    password?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    proxy?: string;
    queryParams?: string;
    /**
     * (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
     */
    rejectInvalidJars?: boolean;
    /**
     * (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
     */
    remoteRepoChecksumPolicyType?: string;
    remoteRepoLayoutRef?: string;
    repoLayoutRef?: string;
    retrievalCachePeriodSeconds?: number;
    shareConfiguration?: boolean;
    socketTimeoutMillis?: number;
    storeArtifactsLocally?: boolean;
    /**
     * (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
     */
    suppressPomConsistencyChecks?: boolean;
    synchronizeProperties?: boolean;
    unusedArtifactsCleanupPeriodHours?: number;
    url?: string;
    username?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getRemoteSbtRepository.
 */
export interface GetRemoteSbtRepositoryResult {
    readonly allowAnyHostAuth?: boolean;
    readonly assumedOfflinePeriodSecs?: number;
    readonly blackedOut?: boolean;
    readonly blockMismatchingMimeTypes?: boolean;
    readonly bypassHeadRequests?: boolean;
    readonly cdnRedirect?: boolean;
    readonly clientTlsCertificate: string;
    readonly contentSynchronisation: outputs.GetRemoteSbtRepositoryContentSynchronisation;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly enableCookieManagement?: boolean;
    readonly excludesPattern?: string;
    /**
     * (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
     */
    readonly fetchJarsEagerly?: boolean;
    /**
     * (Optional, Default: `false`) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
     */
    readonly fetchSourcesEagerly?: boolean;
    /**
     * (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
     */
    readonly handleReleases?: boolean;
    /**
     * (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
     */
    readonly handleSnapshots?: boolean;
    readonly hardFail?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    readonly listRemoteFolderItems?: boolean;
    readonly localAddress?: string;
    readonly metadataRetrievalTimeoutSecs?: number;
    readonly mismatchingMimeTypesOverrideList?: string;
    readonly missedCachePeriodSeconds?: number;
    readonly notes?: string;
    readonly offline?: boolean;
    readonly packageType: string;
    readonly password?: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly proxy?: string;
    readonly queryParams?: string;
    /**
     * (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
     */
    readonly rejectInvalidJars?: boolean;
    /**
     * (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
     */
    readonly remoteRepoChecksumPolicyType?: string;
    readonly remoteRepoLayoutRef?: string;
    readonly repoLayoutRef?: string;
    readonly retrievalCachePeriodSeconds?: number;
    readonly shareConfiguration: boolean;
    readonly socketTimeoutMillis?: number;
    readonly storeArtifactsLocally?: boolean;
    /**
     * (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
     */
    readonly suppressPomConsistencyChecks?: boolean;
    readonly synchronizeProperties?: boolean;
    readonly unusedArtifactsCleanupPeriodHours?: number;
    readonly url?: string;
    readonly username?: string;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a remote SBT repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote-sbt = artifactory.getRemoteSbtRepository({
 *     key: "remote-sbt",
 * });
 * ```
 */
export function getRemoteSbtRepositoryOutput(args: GetRemoteSbtRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRemoteSbtRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getRemoteSbtRepository(a, opts))
}

/**
 * A collection of arguments for invoking getRemoteSbtRepository.
 */
export interface GetRemoteSbtRepositoryOutputArgs {
    allowAnyHostAuth?: pulumi.Input<boolean>;
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    blackedOut?: pulumi.Input<boolean>;
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    bypassHeadRequests?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    contentSynchronisation?: pulumi.Input<inputs.GetRemoteSbtRepositoryContentSynchronisationArgs>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    enableCookieManagement?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
     */
    fetchJarsEagerly?: pulumi.Input<boolean>;
    /**
     * (Optional, Default: `false`) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
     */
    fetchSourcesEagerly?: pulumi.Input<boolean>;
    /**
     * (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
     */
    handleReleases?: pulumi.Input<boolean>;
    /**
     * (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
     */
    handleSnapshots?: pulumi.Input<boolean>;
    hardFail?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    listRemoteFolderItems?: pulumi.Input<boolean>;
    localAddress?: pulumi.Input<string>;
    metadataRetrievalTimeoutSecs?: pulumi.Input<number>;
    mismatchingMimeTypesOverrideList?: pulumi.Input<string>;
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    offline?: pulumi.Input<boolean>;
    password?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    queryParams?: pulumi.Input<string>;
    /**
     * (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
     */
    rejectInvalidJars?: pulumi.Input<boolean>;
    /**
     * (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
     */
    remoteRepoChecksumPolicyType?: pulumi.Input<string>;
    remoteRepoLayoutRef?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
     */
    suppressPomConsistencyChecks?: pulumi.Input<boolean>;
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
