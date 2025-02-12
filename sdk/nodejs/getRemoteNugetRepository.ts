// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a remote NuGet repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote_nuget = artifactory.getRemoteNugetRepository({
 *     key: "remote-nuget",
 * });
 * ```
 */
export function getRemoteNugetRepository(args: GetRemoteNugetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteNugetRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getRemoteNugetRepository:getRemoteNugetRepository", {
        "allowAnyHostAuth": args.allowAnyHostAuth,
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "assumedOfflinePeriodSecs": args.assumedOfflinePeriodSecs,
        "blackedOut": args.blackedOut,
        "blockMismatchingMimeTypes": args.blockMismatchingMimeTypes,
        "bypassHeadRequests": args.bypassHeadRequests,
        "cdnRedirect": args.cdnRedirect,
        "clientTlsCertificate": args.clientTlsCertificate,
        "contentSynchronisation": args.contentSynchronisation,
        "curated": args.curated,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "disableUrlNormalization": args.disableUrlNormalization,
        "downloadContextPath": args.downloadContextPath,
        "downloadDirect": args.downloadDirect,
        "enableCookieManagement": args.enableCookieManagement,
        "excludesPattern": args.excludesPattern,
        "feedContextPath": args.feedContextPath,
        "forceNugetAuthentication": args.forceNugetAuthentication,
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
        "remoteRepoLayoutRef": args.remoteRepoLayoutRef,
        "repoLayoutRef": args.repoLayoutRef,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "shareConfiguration": args.shareConfiguration,
        "socketTimeoutMillis": args.socketTimeoutMillis,
        "storeArtifactsLocally": args.storeArtifactsLocally,
        "symbolServerUrl": args.symbolServerUrl,
        "synchronizeProperties": args.synchronizeProperties,
        "unusedArtifactsCleanupPeriodHours": args.unusedArtifactsCleanupPeriodHours,
        "url": args.url,
        "username": args.username,
        "v3FeedUrl": args.v3FeedUrl,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getRemoteNugetRepository.
 */
export interface GetRemoteNugetRepositoryArgs {
    allowAnyHostAuth?: boolean;
    archiveBrowsingEnabled?: boolean;
    assumedOfflinePeriodSecs?: number;
    blackedOut?: boolean;
    blockMismatchingMimeTypes?: boolean;
    bypassHeadRequests?: boolean;
    cdnRedirect?: boolean;
    clientTlsCertificate?: string;
    contentSynchronisation?: inputs.GetRemoteNugetRepositoryContentSynchronisation;
    curated?: boolean;
    description?: string;
    disableProxy?: boolean;
    disableUrlNormalization?: boolean;
    /**
     * (Optional) The context path prefix through which NuGet downloads are served. For example, the NuGet Gallery download URL is `https://nuget.org/api/v2/package`, so the repository URL should be configured as `https://nuget.org` and the download context path should be configured as `api/v2/package`. Default value is `api/v2/package`.
     */
    downloadContextPath?: string;
    downloadDirect?: boolean;
    enableCookieManagement?: boolean;
    excludesPattern?: string;
    /**
     * (Optional) When proxying a remote NuGet repository, customize feed resource location using this attribute. Default value is `api/v2`.
     */
    feedContextPath?: string;
    /**
     * (Optional) Force basic authentication credentials in order to use this repository. Default value is `false`.
     */
    forceNugetAuthentication?: boolean;
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
    remoteRepoLayoutRef?: string;
    repoLayoutRef?: string;
    retrievalCachePeriodSeconds?: number;
    shareConfiguration?: boolean;
    socketTimeoutMillis?: number;
    storeArtifactsLocally?: boolean;
    /**
     * (Optional) NuGet symbol server URL. Default value is `https://symbols.nuget.org/download/symbols`.
     */
    symbolServerUrl?: string;
    synchronizeProperties?: boolean;
    unusedArtifactsCleanupPeriodHours?: number;
    url?: string;
    username?: string;
    /**
     * (Optional) The URL to the NuGet v3 feed. Default value is `https://api.nuget.org/v3/index.json`.
     */
    v3FeedUrl?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getRemoteNugetRepository.
 */
export interface GetRemoteNugetRepositoryResult {
    readonly allowAnyHostAuth?: boolean;
    readonly archiveBrowsingEnabled?: boolean;
    readonly assumedOfflinePeriodSecs?: number;
    readonly blackedOut?: boolean;
    readonly blockMismatchingMimeTypes?: boolean;
    readonly bypassHeadRequests?: boolean;
    readonly cdnRedirect?: boolean;
    readonly clientTlsCertificate: string;
    readonly contentSynchronisation: outputs.GetRemoteNugetRepositoryContentSynchronisation;
    readonly curated?: boolean;
    readonly description?: string;
    readonly disableProxy?: boolean;
    readonly disableUrlNormalization?: boolean;
    /**
     * (Optional) The context path prefix through which NuGet downloads are served. For example, the NuGet Gallery download URL is `https://nuget.org/api/v2/package`, so the repository URL should be configured as `https://nuget.org` and the download context path should be configured as `api/v2/package`. Default value is `api/v2/package`.
     */
    readonly downloadContextPath?: string;
    readonly downloadDirect?: boolean;
    readonly enableCookieManagement?: boolean;
    readonly excludesPattern?: string;
    /**
     * (Optional) When proxying a remote NuGet repository, customize feed resource location using this attribute. Default value is `api/v2`.
     */
    readonly feedContextPath?: string;
    /**
     * (Optional) Force basic authentication credentials in order to use this repository. Default value is `false`.
     */
    readonly forceNugetAuthentication?: boolean;
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
    readonly remoteRepoLayoutRef?: string;
    readonly repoLayoutRef?: string;
    readonly retrievalCachePeriodSeconds?: number;
    readonly shareConfiguration: boolean;
    readonly socketTimeoutMillis?: number;
    readonly storeArtifactsLocally?: boolean;
    /**
     * (Optional) NuGet symbol server URL. Default value is `https://symbols.nuget.org/download/symbols`.
     */
    readonly symbolServerUrl?: string;
    readonly synchronizeProperties?: boolean;
    readonly unusedArtifactsCleanupPeriodHours?: number;
    readonly url?: string;
    readonly username?: string;
    /**
     * (Optional) The URL to the NuGet v3 feed. Default value is `https://api.nuget.org/v3/index.json`.
     */
    readonly v3FeedUrl?: string;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a remote NuGet repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote_nuget = artifactory.getRemoteNugetRepository({
 *     key: "remote-nuget",
 * });
 * ```
 */
export function getRemoteNugetRepositoryOutput(args: GetRemoteNugetRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRemoteNugetRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getRemoteNugetRepository:getRemoteNugetRepository", {
        "allowAnyHostAuth": args.allowAnyHostAuth,
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "assumedOfflinePeriodSecs": args.assumedOfflinePeriodSecs,
        "blackedOut": args.blackedOut,
        "blockMismatchingMimeTypes": args.blockMismatchingMimeTypes,
        "bypassHeadRequests": args.bypassHeadRequests,
        "cdnRedirect": args.cdnRedirect,
        "clientTlsCertificate": args.clientTlsCertificate,
        "contentSynchronisation": args.contentSynchronisation,
        "curated": args.curated,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "disableUrlNormalization": args.disableUrlNormalization,
        "downloadContextPath": args.downloadContextPath,
        "downloadDirect": args.downloadDirect,
        "enableCookieManagement": args.enableCookieManagement,
        "excludesPattern": args.excludesPattern,
        "feedContextPath": args.feedContextPath,
        "forceNugetAuthentication": args.forceNugetAuthentication,
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
        "remoteRepoLayoutRef": args.remoteRepoLayoutRef,
        "repoLayoutRef": args.repoLayoutRef,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "shareConfiguration": args.shareConfiguration,
        "socketTimeoutMillis": args.socketTimeoutMillis,
        "storeArtifactsLocally": args.storeArtifactsLocally,
        "symbolServerUrl": args.symbolServerUrl,
        "synchronizeProperties": args.synchronizeProperties,
        "unusedArtifactsCleanupPeriodHours": args.unusedArtifactsCleanupPeriodHours,
        "url": args.url,
        "username": args.username,
        "v3FeedUrl": args.v3FeedUrl,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getRemoteNugetRepository.
 */
export interface GetRemoteNugetRepositoryOutputArgs {
    allowAnyHostAuth?: pulumi.Input<boolean>;
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    blackedOut?: pulumi.Input<boolean>;
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    bypassHeadRequests?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    contentSynchronisation?: pulumi.Input<inputs.GetRemoteNugetRepositoryContentSynchronisationArgs>;
    curated?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    disableProxy?: pulumi.Input<boolean>;
    disableUrlNormalization?: pulumi.Input<boolean>;
    /**
     * (Optional) The context path prefix through which NuGet downloads are served. For example, the NuGet Gallery download URL is `https://nuget.org/api/v2/package`, so the repository URL should be configured as `https://nuget.org` and the download context path should be configured as `api/v2/package`. Default value is `api/v2/package`.
     */
    downloadContextPath?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    enableCookieManagement?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * (Optional) When proxying a remote NuGet repository, customize feed resource location using this attribute. Default value is `api/v2`.
     */
    feedContextPath?: pulumi.Input<string>;
    /**
     * (Optional) Force basic authentication credentials in order to use this repository. Default value is `false`.
     */
    forceNugetAuthentication?: pulumi.Input<boolean>;
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
    remoteRepoLayoutRef?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * (Optional) NuGet symbol server URL. Default value is `https://symbols.nuget.org/download/symbols`.
     */
    symbolServerUrl?: pulumi.Input<string>;
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    /**
     * (Optional) The URL to the NuGet v3 feed. Default value is `https://api.nuget.org/v3/index.json`.
     */
    v3FeedUrl?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
