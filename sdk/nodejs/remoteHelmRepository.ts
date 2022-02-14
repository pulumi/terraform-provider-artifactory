// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Artifactory Remote Repository Resource
 *
 * Provides an Artifactory remote `helm` repository resource. This provides helm specific fields and is the only way to get them.
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Package+Management),
 * although helm is (currently) not listed as a supported format
 *
 * ## Example Usage
 *
 * Includes only new and relevant fields, for anything else, see: generic repo.
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const helm_remote = new artifactory.RemoteHelmRepository("helm-remote", {
 *     externalDependenciesEnabled: true,
 *     externalDependenciesPatterns: ["**github.com**"],
 *     helmChartsBaseUrl: "https://foo.com",
 *     key: "helm-remote-foo25",
 *     url: "https://repo.chartcenter.io/",
 * });
 * ```
 */
export class RemoteHelmRepository extends pulumi.CustomResource {
    /**
     * Get an existing RemoteHelmRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemoteHelmRepositoryState, opts?: pulumi.CustomResourceOptions): RemoteHelmRepository {
        return new RemoteHelmRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/remoteHelmRepository:RemoteHelmRepository';

    /**
     * Returns true if the given object is an instance of RemoteHelmRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemoteHelmRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemoteHelmRepository.__pulumiType;
    }

    /**
     * Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
     * any other host.
     */
    public readonly allowAnyHostAuth!: pulumi.Output<boolean>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline. Default to 300.
     */
    public readonly assumedOfflinePeriodSecs!: pulumi.Output<number | undefined>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     */
    public readonly blackedOut!: pulumi.Output<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    public readonly blockMismatchingMimeTypes!: pulumi.Output<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    public readonly bypassHeadRequests!: pulumi.Output<boolean>;
    public readonly clientTlsCertificate!: pulumi.Output<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    public readonly contentSynchronisation!: pulumi.Output<outputs.RemoteHelmRepositoryContentSynchronisation>;
    public readonly description!: pulumi.Output<string>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    public readonly enableCookieManagement!: pulumi.Output<boolean>;
    public readonly excludesPattern!: pulumi.Output<string>;
    /**
     * When set, external dependencies are rewritten.
     */
    public readonly externalDependenciesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * An Allow List of Ant-style path expressions that specify where external
     * dependencies may be downloaded from. By default, this is an empty list which means that no dependencies may be downloaded
     * from external sources. Note that the official documentation states the default is '**', which is correct when creating
     * repositories in the UI, but incorrect for the API.
     */
    public readonly externalDependenciesPatterns!: pulumi.Output<string[] | undefined>;
    /**
     * @deprecated This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
     */
    public /*out*/ readonly failedRetrievalCachePeriodSecs!: pulumi.Output<number>;
    public readonly hardFail!: pulumi.Output<boolean>;
    /**
     * - No documentation is available. Hopefully you know what this means
     */
    public readonly helmChartsBaseUrl!: pulumi.Output<string | undefined>;
    public readonly includesPattern!: pulumi.Output<string>;
    /**
     * The repository identifier. Must be unique system-wide
     */
    public readonly key!: pulumi.Output<string>;
    public readonly localAddress!: pulumi.Output<string | undefined>;
    /**
     * This is actually the missedRetrievalCachePeriodSecs in the API
     */
    public readonly missedCachePeriodSeconds!: pulumi.Output<number>;
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    public readonly offline!: pulumi.Output<boolean>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    public readonly priorityResolution!: pulumi.Output<boolean>;
    public readonly propagateQueryParams!: pulumi.Output<boolean | undefined>;
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    public readonly proxy!: pulumi.Output<string>;
    public readonly remoteRepoLayoutRef!: pulumi.Output<string>;
    public readonly repoLayoutRef!: pulumi.Output<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    public readonly retrievalCachePeriodSeconds!: pulumi.Output<number>;
    public readonly shareConfiguration!: pulumi.Output<boolean>;
    public readonly socketTimeoutMillis!: pulumi.Output<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     */
    public readonly storeArtifactsLocally!: pulumi.Output<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    public readonly synchronizeProperties!: pulumi.Output<boolean>;
    public readonly unusedArtifactsCleanupPeriodEnabled!: pulumi.Output<boolean>;
    public readonly unusedArtifactsCleanupPeriodHours!: pulumi.Output<number>;
    public readonly url!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string | undefined>;
    public readonly xrayIndex!: pulumi.Output<boolean>;

    /**
     * Create a RemoteHelmRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteHelmRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemoteHelmRepositoryArgs | RemoteHelmRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RemoteHelmRepositoryState | undefined;
            resourceInputs["allowAnyHostAuth"] = state ? state.allowAnyHostAuth : undefined;
            resourceInputs["assumedOfflinePeriodSecs"] = state ? state.assumedOfflinePeriodSecs : undefined;
            resourceInputs["blackedOut"] = state ? state.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = state ? state.blockMismatchingMimeTypes : undefined;
            resourceInputs["bypassHeadRequests"] = state ? state.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = state ? state.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = state ? state.contentSynchronisation : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableCookieManagement"] = state ? state.enableCookieManagement : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["externalDependenciesEnabled"] = state ? state.externalDependenciesEnabled : undefined;
            resourceInputs["externalDependenciesPatterns"] = state ? state.externalDependenciesPatterns : undefined;
            resourceInputs["failedRetrievalCachePeriodSecs"] = state ? state.failedRetrievalCachePeriodSecs : undefined;
            resourceInputs["hardFail"] = state ? state.hardFail : undefined;
            resourceInputs["helmChartsBaseUrl"] = state ? state.helmChartsBaseUrl : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["localAddress"] = state ? state.localAddress : undefined;
            resourceInputs["missedCachePeriodSeconds"] = state ? state.missedCachePeriodSeconds : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["offline"] = state ? state.offline : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["priorityResolution"] = state ? state.priorityResolution : undefined;
            resourceInputs["propagateQueryParams"] = state ? state.propagateQueryParams : undefined;
            resourceInputs["propertySets"] = state ? state.propertySets : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["remoteRepoLayoutRef"] = state ? state.remoteRepoLayoutRef : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = state ? state.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = state ? state.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = state ? state.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = state ? state.storeArtifactsLocally : undefined;
            resourceInputs["synchronizeProperties"] = state ? state.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodEnabled"] = state ? state.unusedArtifactsCleanupPeriodEnabled : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = state ? state.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as RemoteHelmRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["allowAnyHostAuth"] = args ? args.allowAnyHostAuth : undefined;
            resourceInputs["assumedOfflinePeriodSecs"] = args ? args.assumedOfflinePeriodSecs : undefined;
            resourceInputs["blackedOut"] = args ? args.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = args ? args.blockMismatchingMimeTypes : undefined;
            resourceInputs["bypassHeadRequests"] = args ? args.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = args ? args.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = args ? args.contentSynchronisation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableCookieManagement"] = args ? args.enableCookieManagement : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["externalDependenciesEnabled"] = args ? args.externalDependenciesEnabled : undefined;
            resourceInputs["externalDependenciesPatterns"] = args ? args.externalDependenciesPatterns : undefined;
            resourceInputs["hardFail"] = args ? args.hardFail : undefined;
            resourceInputs["helmChartsBaseUrl"] = args ? args.helmChartsBaseUrl : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["localAddress"] = args ? args.localAddress : undefined;
            resourceInputs["missedCachePeriodSeconds"] = args ? args.missedCachePeriodSeconds : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["offline"] = args ? args.offline : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["priorityResolution"] = args ? args.priorityResolution : undefined;
            resourceInputs["propagateQueryParams"] = args ? args.propagateQueryParams : undefined;
            resourceInputs["propertySets"] = args ? args.propertySets : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["remoteRepoLayoutRef"] = args ? args.remoteRepoLayoutRef : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = args ? args.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = args ? args.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = args ? args.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = args ? args.storeArtifactsLocally : undefined;
            resourceInputs["synchronizeProperties"] = args ? args.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodEnabled"] = args ? args.unusedArtifactsCleanupPeriodEnabled : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = args ? args.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["xrayIndex"] = args ? args.xrayIndex : undefined;
            resourceInputs["failedRetrievalCachePeriodSecs"] = undefined /*out*/;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RemoteHelmRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RemoteHelmRepository resources.
 */
export interface RemoteHelmRepositoryState {
    /**
     * Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
     * any other host.
     */
    allowAnyHostAuth?: pulumi.Input<boolean>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline. Default to 300.
     */
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    bypassHeadRequests?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    contentSynchronisation?: pulumi.Input<inputs.RemoteHelmRepositoryContentSynchronisation>;
    description?: pulumi.Input<string>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    enableCookieManagement?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * When set, external dependencies are rewritten.
     */
    externalDependenciesEnabled?: pulumi.Input<boolean>;
    /**
     * An Allow List of Ant-style path expressions that specify where external
     * dependencies may be downloaded from. By default, this is an empty list which means that no dependencies may be downloaded
     * from external sources. Note that the official documentation states the default is '**', which is correct when creating
     * repositories in the UI, but incorrect for the API.
     */
    externalDependenciesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
     */
    failedRetrievalCachePeriodSecs?: pulumi.Input<number>;
    hardFail?: pulumi.Input<boolean>;
    /**
     * - No documentation is available. Hopefully you know what this means
     */
    helmChartsBaseUrl?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * The repository identifier. Must be unique system-wide
     */
    key?: pulumi.Input<string>;
    localAddress?: pulumi.Input<string>;
    /**
     * This is actually the missedRetrievalCachePeriodSecs in the API
     */
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    offline?: pulumi.Input<boolean>;
    packageType?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    propagateQueryParams?: pulumi.Input<boolean>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    remoteRepoLayoutRef?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     */
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodEnabled?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RemoteHelmRepository resource.
 */
export interface RemoteHelmRepositoryArgs {
    /**
     * Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
     * any other host.
     */
    allowAnyHostAuth?: pulumi.Input<boolean>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline. Default to 300.
     */
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    bypassHeadRequests?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    contentSynchronisation?: pulumi.Input<inputs.RemoteHelmRepositoryContentSynchronisation>;
    description?: pulumi.Input<string>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    enableCookieManagement?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * When set, external dependencies are rewritten.
     */
    externalDependenciesEnabled?: pulumi.Input<boolean>;
    /**
     * An Allow List of Ant-style path expressions that specify where external
     * dependencies may be downloaded from. By default, this is an empty list which means that no dependencies may be downloaded
     * from external sources. Note that the official documentation states the default is '**', which is correct when creating
     * repositories in the UI, but incorrect for the API.
     */
    externalDependenciesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    hardFail?: pulumi.Input<boolean>;
    /**
     * - No documentation is available. Hopefully you know what this means
     */
    helmChartsBaseUrl?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * The repository identifier. Must be unique system-wide
     */
    key: pulumi.Input<string>;
    localAddress?: pulumi.Input<string>;
    /**
     * This is actually the missedRetrievalCachePeriodSecs in the API
     */
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    offline?: pulumi.Input<boolean>;
    password?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    propagateQueryParams?: pulumi.Input<boolean>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    remoteRepoLayoutRef?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     */
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodEnabled?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
