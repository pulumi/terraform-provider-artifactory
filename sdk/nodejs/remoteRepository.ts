// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Remote repositories can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/remoteRepository:RemoteRepository my-remote my-remote
 * ```
 */
export class RemoteRepository extends pulumi.CustomResource {
    /**
     * Get an existing RemoteRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemoteRepositoryState, opts?: pulumi.CustomResourceOptions): RemoteRepository {
        return new RemoteRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/remoteRepository:RemoteRepository';

    /**
     * Returns true if the given object is an instance of RemoteRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemoteRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemoteRepository.__pulumiType;
    }

    public readonly allowAnyHostAuth!: pulumi.Output<boolean>;
    public readonly blackedOut!: pulumi.Output<boolean>;
    public readonly blockMismatchingMimeTypes!: pulumi.Output<boolean>;
    public readonly bowerRegistryUrl!: pulumi.Output<string>;
    public readonly bypassHeadRequests!: pulumi.Output<boolean>;
    public readonly clientTlsCertificate!: pulumi.Output<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    public readonly contentSynchronisation!: pulumi.Output<outputs.RemoteRepositoryContentSynchronisation>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly downloadContextPath!: pulumi.Output<string | undefined>;
    public readonly enableCookieManagement!: pulumi.Output<boolean>;
    public readonly enableTokenAuthentication!: pulumi.Output<boolean>;
    public readonly excludesPattern!: pulumi.Output<string>;
    public readonly feedContextPath!: pulumi.Output<string | undefined>;
    public readonly fetchJarsEagerly!: pulumi.Output<boolean>;
    public readonly fetchSourcesEagerly!: pulumi.Output<boolean>;
    public readonly forceNugetAuthentication!: pulumi.Output<boolean>;
    public readonly handleReleases!: pulumi.Output<boolean>;
    public readonly handleSnapshots!: pulumi.Output<boolean>;
    public readonly hardFail!: pulumi.Output<boolean>;
    public readonly includesPattern!: pulumi.Output<string>;
    public readonly key!: pulumi.Output<string>;
    public readonly localAddress!: pulumi.Output<string | undefined>;
    public readonly maxUniqueSnapshots!: pulumi.Output<number>;
    public readonly missedCachePeriodSeconds!: pulumi.Output<number>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public readonly offline!: pulumi.Output<boolean>;
    public readonly packageType!: pulumi.Output<string | undefined>;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly propagateQueryParams!: pulumi.Output<boolean>;
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    /**
     * Proxy key from Artifactory Proxies setting
     */
    public readonly proxy!: pulumi.Output<string | undefined>;
    public readonly pypiRegistryUrl!: pulumi.Output<string>;
    public readonly remoteRepoChecksumPolicyType!: pulumi.Output<string>;
    public readonly repoLayoutRef!: pulumi.Output<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    public readonly retrievalCachePeriodSeconds!: pulumi.Output<number>;
    public readonly shareConfiguration!: pulumi.Output<boolean>;
    public readonly socketTimeoutMillis!: pulumi.Output<number>;
    public readonly storeArtifactsLocally!: pulumi.Output<boolean>;
    public readonly suppressPomConsistencyChecks!: pulumi.Output<boolean>;
    public readonly synchronizeProperties!: pulumi.Output<boolean>;
    public readonly unusedArtifactsCleanupPeriodHours!: pulumi.Output<number>;
    public readonly url!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string | undefined>;
    public readonly v3FeedUrl!: pulumi.Output<string | undefined>;
    public readonly vcsGitDownloadUrl!: pulumi.Output<string>;
    public readonly vcsGitProvider!: pulumi.Output<string>;
    public readonly vcsType!: pulumi.Output<string>;
    public readonly xrayIndex!: pulumi.Output<boolean>;

    /**
     * Create a RemoteRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemoteRepositoryArgs | RemoteRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RemoteRepositoryState | undefined;
            resourceInputs["allowAnyHostAuth"] = state ? state.allowAnyHostAuth : undefined;
            resourceInputs["blackedOut"] = state ? state.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = state ? state.blockMismatchingMimeTypes : undefined;
            resourceInputs["bowerRegistryUrl"] = state ? state.bowerRegistryUrl : undefined;
            resourceInputs["bypassHeadRequests"] = state ? state.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = state ? state.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = state ? state.contentSynchronisation : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["downloadContextPath"] = state ? state.downloadContextPath : undefined;
            resourceInputs["enableCookieManagement"] = state ? state.enableCookieManagement : undefined;
            resourceInputs["enableTokenAuthentication"] = state ? state.enableTokenAuthentication : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["feedContextPath"] = state ? state.feedContextPath : undefined;
            resourceInputs["fetchJarsEagerly"] = state ? state.fetchJarsEagerly : undefined;
            resourceInputs["fetchSourcesEagerly"] = state ? state.fetchSourcesEagerly : undefined;
            resourceInputs["forceNugetAuthentication"] = state ? state.forceNugetAuthentication : undefined;
            resourceInputs["handleReleases"] = state ? state.handleReleases : undefined;
            resourceInputs["handleSnapshots"] = state ? state.handleSnapshots : undefined;
            resourceInputs["hardFail"] = state ? state.hardFail : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["localAddress"] = state ? state.localAddress : undefined;
            resourceInputs["maxUniqueSnapshots"] = state ? state.maxUniqueSnapshots : undefined;
            resourceInputs["missedCachePeriodSeconds"] = state ? state.missedCachePeriodSeconds : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["offline"] = state ? state.offline : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["propagateQueryParams"] = state ? state.propagateQueryParams : undefined;
            resourceInputs["propertySets"] = state ? state.propertySets : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["pypiRegistryUrl"] = state ? state.pypiRegistryUrl : undefined;
            resourceInputs["remoteRepoChecksumPolicyType"] = state ? state.remoteRepoChecksumPolicyType : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = state ? state.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = state ? state.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = state ? state.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = state ? state.storeArtifactsLocally : undefined;
            resourceInputs["suppressPomConsistencyChecks"] = state ? state.suppressPomConsistencyChecks : undefined;
            resourceInputs["synchronizeProperties"] = state ? state.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = state ? state.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["v3FeedUrl"] = state ? state.v3FeedUrl : undefined;
            resourceInputs["vcsGitDownloadUrl"] = state ? state.vcsGitDownloadUrl : undefined;
            resourceInputs["vcsGitProvider"] = state ? state.vcsGitProvider : undefined;
            resourceInputs["vcsType"] = state ? state.vcsType : undefined;
            resourceInputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as RemoteRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["allowAnyHostAuth"] = args ? args.allowAnyHostAuth : undefined;
            resourceInputs["blackedOut"] = args ? args.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = args ? args.blockMismatchingMimeTypes : undefined;
            resourceInputs["bowerRegistryUrl"] = args ? args.bowerRegistryUrl : undefined;
            resourceInputs["bypassHeadRequests"] = args ? args.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = args ? args.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = args ? args.contentSynchronisation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["downloadContextPath"] = args ? args.downloadContextPath : undefined;
            resourceInputs["enableCookieManagement"] = args ? args.enableCookieManagement : undefined;
            resourceInputs["enableTokenAuthentication"] = args ? args.enableTokenAuthentication : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["feedContextPath"] = args ? args.feedContextPath : undefined;
            resourceInputs["fetchJarsEagerly"] = args ? args.fetchJarsEagerly : undefined;
            resourceInputs["fetchSourcesEagerly"] = args ? args.fetchSourcesEagerly : undefined;
            resourceInputs["forceNugetAuthentication"] = args ? args.forceNugetAuthentication : undefined;
            resourceInputs["handleReleases"] = args ? args.handleReleases : undefined;
            resourceInputs["handleSnapshots"] = args ? args.handleSnapshots : undefined;
            resourceInputs["hardFail"] = args ? args.hardFail : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["localAddress"] = args ? args.localAddress : undefined;
            resourceInputs["maxUniqueSnapshots"] = args ? args.maxUniqueSnapshots : undefined;
            resourceInputs["missedCachePeriodSeconds"] = args ? args.missedCachePeriodSeconds : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["offline"] = args ? args.offline : undefined;
            resourceInputs["packageType"] = args ? args.packageType : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["propagateQueryParams"] = args ? args.propagateQueryParams : undefined;
            resourceInputs["propertySets"] = args ? args.propertySets : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["pypiRegistryUrl"] = args ? args.pypiRegistryUrl : undefined;
            resourceInputs["remoteRepoChecksumPolicyType"] = args ? args.remoteRepoChecksumPolicyType : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = args ? args.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = args ? args.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = args ? args.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = args ? args.storeArtifactsLocally : undefined;
            resourceInputs["suppressPomConsistencyChecks"] = args ? args.suppressPomConsistencyChecks : undefined;
            resourceInputs["synchronizeProperties"] = args ? args.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = args ? args.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["v3FeedUrl"] = args ? args.v3FeedUrl : undefined;
            resourceInputs["vcsGitDownloadUrl"] = args ? args.vcsGitDownloadUrl : undefined;
            resourceInputs["vcsGitProvider"] = args ? args.vcsGitProvider : undefined;
            resourceInputs["vcsType"] = args ? args.vcsType : undefined;
            resourceInputs["xrayIndex"] = args ? args.xrayIndex : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RemoteRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RemoteRepository resources.
 */
export interface RemoteRepositoryState {
    allowAnyHostAuth?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    bowerRegistryUrl?: pulumi.Input<string>;
    bypassHeadRequests?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    contentSynchronisation?: pulumi.Input<inputs.RemoteRepositoryContentSynchronisation>;
    description?: pulumi.Input<string>;
    downloadContextPath?: pulumi.Input<string>;
    enableCookieManagement?: pulumi.Input<boolean>;
    enableTokenAuthentication?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    feedContextPath?: pulumi.Input<string>;
    fetchJarsEagerly?: pulumi.Input<boolean>;
    fetchSourcesEagerly?: pulumi.Input<boolean>;
    forceNugetAuthentication?: pulumi.Input<boolean>;
    handleReleases?: pulumi.Input<boolean>;
    handleSnapshots?: pulumi.Input<boolean>;
    hardFail?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    key?: pulumi.Input<string>;
    localAddress?: pulumi.Input<string>;
    maxUniqueSnapshots?: pulumi.Input<number>;
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    offline?: pulumi.Input<boolean>;
    packageType?: pulumi.Input<string>;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password?: pulumi.Input<string>;
    propagateQueryParams?: pulumi.Input<boolean>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Proxy key from Artifactory Proxies setting
     */
    proxy?: pulumi.Input<string>;
    pypiRegistryUrl?: pulumi.Input<string>;
    remoteRepoChecksumPolicyType?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    storeArtifactsLocally?: pulumi.Input<boolean>;
    suppressPomConsistencyChecks?: pulumi.Input<boolean>;
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    v3FeedUrl?: pulumi.Input<string>;
    vcsGitDownloadUrl?: pulumi.Input<string>;
    vcsGitProvider?: pulumi.Input<string>;
    vcsType?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RemoteRepository resource.
 */
export interface RemoteRepositoryArgs {
    allowAnyHostAuth?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    bowerRegistryUrl?: pulumi.Input<string>;
    bypassHeadRequests?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    contentSynchronisation?: pulumi.Input<inputs.RemoteRepositoryContentSynchronisation>;
    description?: pulumi.Input<string>;
    downloadContextPath?: pulumi.Input<string>;
    enableCookieManagement?: pulumi.Input<boolean>;
    enableTokenAuthentication?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    feedContextPath?: pulumi.Input<string>;
    fetchJarsEagerly?: pulumi.Input<boolean>;
    fetchSourcesEagerly?: pulumi.Input<boolean>;
    forceNugetAuthentication?: pulumi.Input<boolean>;
    handleReleases?: pulumi.Input<boolean>;
    handleSnapshots?: pulumi.Input<boolean>;
    hardFail?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    key: pulumi.Input<string>;
    localAddress?: pulumi.Input<string>;
    maxUniqueSnapshots?: pulumi.Input<number>;
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    offline?: pulumi.Input<boolean>;
    packageType?: pulumi.Input<string>;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password?: pulumi.Input<string>;
    propagateQueryParams?: pulumi.Input<boolean>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Proxy key from Artifactory Proxies setting
     */
    proxy?: pulumi.Input<string>;
    pypiRegistryUrl?: pulumi.Input<string>;
    remoteRepoChecksumPolicyType?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    storeArtifactsLocally?: pulumi.Input<boolean>;
    suppressPomConsistencyChecks?: pulumi.Input<boolean>;
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    v3FeedUrl?: pulumi.Input<string>;
    vcsGitDownloadUrl?: pulumi.Input<string>;
    vcsGitProvider?: pulumi.Input<string>;
    vcsType?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
