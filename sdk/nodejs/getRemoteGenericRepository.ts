// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a remote Generic repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote-generic = artifactory.getRemoteGenericRepository({
 *     key: "remote-generic",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRemoteGenericRepository(args: GetRemoteGenericRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteGenericRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getRemoteGenericRepository:getRemoteGenericRepository", {
        "allowAnyHostAuth": args.allowAnyHostAuth,
        "assumedOfflinePeriodSecs": args.assumedOfflinePeriodSecs,
        "blackedOut": args.blackedOut,
        "blockMismatchingMimeTypes": args.blockMismatchingMimeTypes,
        "bypassHeadRequests": args.bypassHeadRequests,
        "cdnRedirect": args.cdnRedirect,
        "clientTlsCertificate": args.clientTlsCertificate,
        "contentSynchronisation": args.contentSynchronisation,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "disableUrlNormalization": args.disableUrlNormalization,
        "downloadDirect": args.downloadDirect,
        "enableCookieManagement": args.enableCookieManagement,
        "excludesPattern": args.excludesPattern,
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
        "propagateQueryParams": args.propagateQueryParams,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "queryParams": args.queryParams,
        "remoteRepoLayoutRef": args.remoteRepoLayoutRef,
        "repoLayoutRef": args.repoLayoutRef,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "shareConfiguration": args.shareConfiguration,
        "socketTimeoutMillis": args.socketTimeoutMillis,
        "storeArtifactsLocally": args.storeArtifactsLocally,
        "synchronizeProperties": args.synchronizeProperties,
        "unusedArtifactsCleanupPeriodHours": args.unusedArtifactsCleanupPeriodHours,
        "url": args.url,
        "username": args.username,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getRemoteGenericRepository.
 */
export interface GetRemoteGenericRepositoryArgs {
    allowAnyHostAuth?: boolean;
    assumedOfflinePeriodSecs?: number;
    blackedOut?: boolean;
    blockMismatchingMimeTypes?: boolean;
    bypassHeadRequests?: boolean;
    cdnRedirect?: boolean;
    clientTlsCertificate?: string;
    contentSynchronisation?: inputs.GetRemoteGenericRepositoryContentSynchronisation;
    description?: string;
    disableProxy?: boolean;
    disableUrlNormalization?: boolean;
    downloadDirect?: boolean;
    enableCookieManagement?: boolean;
    excludesPattern?: string;
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
    /**
     * (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    propagateQueryParams?: boolean;
    propertySets?: string[];
    proxy?: string;
    queryParams?: string;
    remoteRepoLayoutRef?: string;
    repoLayoutRef?: string;
    retrievalCachePeriodSeconds?: number;
    shareConfiguration?: boolean;
    socketTimeoutMillis?: number;
    storeArtifactsLocally?: boolean;
    synchronizeProperties?: boolean;
    unusedArtifactsCleanupPeriodHours?: number;
    url?: string;
    username?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getRemoteGenericRepository.
 */
export interface GetRemoteGenericRepositoryResult {
    readonly allowAnyHostAuth?: boolean;
    readonly assumedOfflinePeriodSecs?: number;
    readonly blackedOut?: boolean;
    readonly blockMismatchingMimeTypes?: boolean;
    readonly bypassHeadRequests?: boolean;
    readonly cdnRedirect?: boolean;
    readonly clientTlsCertificate: string;
    readonly contentSynchronisation: outputs.GetRemoteGenericRepositoryContentSynchronisation;
    readonly description?: string;
    readonly disableProxy?: boolean;
    readonly disableUrlNormalization?: boolean;
    readonly downloadDirect?: boolean;
    readonly enableCookieManagement?: boolean;
    readonly excludesPattern?: string;
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
    /**
     * (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    readonly propagateQueryParams?: boolean;
    readonly propertySets?: string[];
    readonly proxy?: string;
    readonly queryParams?: string;
    readonly remoteRepoLayoutRef?: string;
    readonly repoLayoutRef?: string;
    readonly retrievalCachePeriodSeconds?: number;
    readonly shareConfiguration: boolean;
    readonly socketTimeoutMillis?: number;
    readonly storeArtifactsLocally?: boolean;
    readonly synchronizeProperties?: boolean;
    readonly unusedArtifactsCleanupPeriodHours?: number;
    readonly url?: string;
    readonly username?: string;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a remote Generic repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote-generic = artifactory.getRemoteGenericRepository({
 *     key: "remote-generic",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRemoteGenericRepositoryOutput(args: GetRemoteGenericRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRemoteGenericRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getRemoteGenericRepository(a, opts))
}

/**
 * A collection of arguments for invoking getRemoteGenericRepository.
 */
export interface GetRemoteGenericRepositoryOutputArgs {
    allowAnyHostAuth?: pulumi.Input<boolean>;
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    blackedOut?: pulumi.Input<boolean>;
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    bypassHeadRequests?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    contentSynchronisation?: pulumi.Input<inputs.GetRemoteGenericRepositoryContentSynchronisationArgs>;
    description?: pulumi.Input<string>;
    disableProxy?: pulumi.Input<boolean>;
    disableUrlNormalization?: pulumi.Input<boolean>;
    downloadDirect?: pulumi.Input<boolean>;
    enableCookieManagement?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
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
    /**
     * (Optional, Default: `false`) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    propagateQueryParams?: pulumi.Input<boolean>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    queryParams?: pulumi.Input<string>;
    remoteRepoLayoutRef?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    storeArtifactsLocally?: pulumi.Input<boolean>;
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
