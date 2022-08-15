// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Creates a remote Generic repository.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_remote_generic = new Artifactory.RemoteGenericRepository("my-remote-generic", new()
    ///     {
    ///         Key = "my-remote-generic",
    ///         Url = "http://testartifactory.io/artifactory/example-generic/",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Remote repositories can be imported using their name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/remoteGenericRepository:RemoteGenericRepository my-remote-generic my-remote-generic
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/remoteGenericRepository:RemoteGenericRepository")]
    public partial class RemoteGenericRepository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
        /// any other host.
        /// </summary>
        [Output("allowAnyHostAuth")]
        public Output<bool> AllowAnyHostAuth { get; private set; } = null!;

        /// <summary>
        /// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
        /// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
        /// offline. Default to 300.
        /// </summary>
        [Output("assumedOfflinePeriodSecs")]
        public Output<int?> AssumedOfflinePeriodSecs { get; private set; } = null!;

        /// <summary>
        /// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
        /// resolution.
        /// </summary>
        [Output("blackedOut")]
        public Output<bool> BlackedOut { get; private set; } = null!;

        /// <summary>
        /// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
        /// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
        /// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
        /// </summary>
        [Output("blockMismatchingMimeTypes")]
        public Output<bool> BlockMismatchingMimeTypes { get; private set; } = null!;

        /// <summary>
        /// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
        /// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
        /// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
        /// </summary>
        [Output("bypassHeadRequests")]
        public Output<bool> BypassHeadRequests { get; private set; } = null!;

        [Output("clientTlsCertificate")]
        public Output<string> ClientTlsCertificate { get; private set; } = null!;

        [Output("contentSynchronisation")]
        public Output<Outputs.RemoteGenericRepositoryContentSynchronisation> ContentSynchronisation { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Enables cookie management if the remote repository uses cookies to manage client state.
        /// </summary>
        [Output("enableCookieManagement")]
        public Output<bool> EnableCookieManagement { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
        /// artifacts are excluded.
        /// </summary>
        [Output("excludesPattern")]
        public Output<string> ExcludesPattern { get; private set; } = null!;

        [Output("failedRetrievalCachePeriodSecs")]
        public Output<int> FailedRetrievalCachePeriodSecs { get; private set; } = null!;

        /// <summary>
        /// When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
        /// communicate with this repository.
        /// </summary>
        [Output("hardFail")]
        public Output<bool> HardFail { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Output("includesPattern")]
        public Output<string> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        /// contain spaces or special characters.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
        /// the 'Retrieval Cache Period'. Default value is 'false'.
        /// </summary>
        [Output("listRemoteFolderItems")]
        public Output<bool?> ListRemoteFolderItems { get; private set; } = null!;

        /// <summary>
        /// The local address to be used when creating connections. Useful for specifying the interface to use on systems with
        /// multiple network interfaces.
        /// </summary>
        [Output("localAddress")]
        public Output<string?> LocalAddress { get; private set; } = null!;

        /// <summary>
        /// The set of mime types that should override the block_mismatching_mime_types setting. Eg:
        /// "application/json,application/xml". Default value is empty.
        /// </summary>
        [Output("mismatchingMimeTypesOverrideList")]
        public Output<string?> MismatchingMimeTypesOverrideList { get; private set; } = null!;

        /// <summary>
        /// The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
        /// </summary>
        [Output("missedCachePeriodSeconds")]
        public Output<int> MissedCachePeriodSeconds { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
        /// </summary>
        [Output("offline")]
        public Output<bool> Offline { get; private set; } = null!;

        [Output("packageType")]
        public Output<string> PackageType { get; private set; } = null!;

        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Output("priorityResolution")]
        public Output<bool> PriorityResolution { get; private set; } = null!;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        [Output("projectEnvironments")]
        public Output<ImmutableArray<string>> ProjectEnvironments { get; private set; } = null!;

        /// <summary>
        /// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
        /// </summary>
        [Output("propagateQueryParams")]
        public Output<bool?> PropagateQueryParams { get; private set; } = null!;

        /// <summary>
        /// List of property set names
        /// </summary>
        [Output("propertySets")]
        public Output<ImmutableArray<string>> PropertySets { get; private set; } = null!;

        /// <summary>
        /// Proxy key from Artifactory Proxies settings
        /// </summary>
        [Output("proxy")]
        public Output<string?> Proxy { get; private set; } = null!;

        /// <summary>
        /// Repository layout key for the remote layout mapping
        /// </summary>
        [Output("remoteRepoLayoutRef")]
        public Output<string> RemoteRepoLayoutRef { get; private set; } = null!;

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Output("repoLayoutRef")]
        public Output<string?> RepoLayoutRef { get; private set; } = null!;

        /// <summary>
        /// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
        /// </summary>
        [Output("retrievalCachePeriodSeconds")]
        public Output<int> RetrievalCachePeriodSeconds { get; private set; } = null!;

        [Output("shareConfiguration")]
        public Output<bool> ShareConfiguration { get; private set; } = null!;

        /// <summary>
        /// Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
        /// operation is considered a retrieval failure.
        /// </summary>
        [Output("socketTimeoutMillis")]
        public Output<int> SocketTimeoutMillis { get; private set; } = null!;

        /// <summary>
        /// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
        /// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
        /// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
        /// servers.
        /// </summary>
        [Output("storeArtifactsLocally")]
        public Output<bool> StoreArtifactsLocally { get; private set; } = null!;

        /// <summary>
        /// When set, remote artifacts are fetched along with their properties.
        /// </summary>
        [Output("synchronizeProperties")]
        public Output<bool> SynchronizeProperties { get; private set; } = null!;

        [Output("unusedArtifactsCleanupPeriodEnabled")]
        public Output<bool> UnusedArtifactsCleanupPeriodEnabled { get; private set; } = null!;

        /// <summary>
        /// The number of hours to wait before an artifact is deemed "unused" and eligible for cleanup from the repository. A value
        /// of 0 means automatic cleanup of cached artifacts is disabled.
        /// </summary>
        [Output("unusedArtifactsCleanupPeriodHours")]
        public Output<int> UnusedArtifactsCleanupPeriodHours { get; private set; } = null!;

        /// <summary>
        /// The remote repo URL.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Output("xrayIndex")]
        public Output<bool?> XrayIndex { get; private set; } = null!;


        /// <summary>
        /// Create a RemoteGenericRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RemoteGenericRepository(string name, RemoteGenericRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/remoteGenericRepository:RemoteGenericRepository", name, args ?? new RemoteGenericRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RemoteGenericRepository(string name, Input<string> id, RemoteGenericRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/remoteGenericRepository:RemoteGenericRepository", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RemoteGenericRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RemoteGenericRepository Get(string name, Input<string> id, RemoteGenericRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new RemoteGenericRepository(name, id, state, options);
        }
    }

    public sealed class RemoteGenericRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
        /// any other host.
        /// </summary>
        [Input("allowAnyHostAuth")]
        public Input<bool>? AllowAnyHostAuth { get; set; }

        /// <summary>
        /// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
        /// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
        /// offline. Default to 300.
        /// </summary>
        [Input("assumedOfflinePeriodSecs")]
        public Input<int>? AssumedOfflinePeriodSecs { get; set; }

        /// <summary>
        /// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
        /// resolution.
        /// </summary>
        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        /// <summary>
        /// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
        /// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
        /// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
        /// </summary>
        [Input("blockMismatchingMimeTypes")]
        public Input<bool>? BlockMismatchingMimeTypes { get; set; }

        /// <summary>
        /// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
        /// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
        /// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
        /// </summary>
        [Input("bypassHeadRequests")]
        public Input<bool>? BypassHeadRequests { get; set; }

        [Input("clientTlsCertificate")]
        public Input<string>? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Input<Inputs.RemoteGenericRepositoryContentSynchronisationArgs>? ContentSynchronisation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables cookie management if the remote repository uses cookies to manage client state.
        /// </summary>
        [Input("enableCookieManagement")]
        public Input<bool>? EnableCookieManagement { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
        /// communicate with this repository.
        /// </summary>
        [Input("hardFail")]
        public Input<bool>? HardFail { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        /// contain spaces or special characters.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
        /// the 'Retrieval Cache Period'. Default value is 'false'.
        /// </summary>
        [Input("listRemoteFolderItems")]
        public Input<bool>? ListRemoteFolderItems { get; set; }

        /// <summary>
        /// The local address to be used when creating connections. Useful for specifying the interface to use on systems with
        /// multiple network interfaces.
        /// </summary>
        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        /// <summary>
        /// The set of mime types that should override the block_mismatching_mime_types setting. Eg:
        /// "application/json,application/xml". Default value is empty.
        /// </summary>
        [Input("mismatchingMimeTypesOverrideList")]
        public Input<string>? MismatchingMimeTypesOverrideList { get; set; }

        /// <summary>
        /// The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
        /// </summary>
        [Input("missedCachePeriodSeconds")]
        public Input<int>? MissedCachePeriodSeconds { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
        /// </summary>
        [Input("offline")]
        public Input<bool>? Offline { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
        /// </summary>
        [Input("propagateQueryParams")]
        public Input<bool>? PropagateQueryParams { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;

        /// <summary>
        /// List of property set names
        /// </summary>
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Proxy key from Artifactory Proxies settings
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Repository layout key for the remote layout mapping
        /// </summary>
        [Input("remoteRepoLayoutRef")]
        public Input<string>? RemoteRepoLayoutRef { get; set; }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        [Input("shareConfiguration")]
        public Input<bool>? ShareConfiguration { get; set; }

        /// <summary>
        /// Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
        /// operation is considered a retrieval failure.
        /// </summary>
        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
        /// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
        /// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
        /// servers.
        /// </summary>
        [Input("storeArtifactsLocally")]
        public Input<bool>? StoreArtifactsLocally { get; set; }

        /// <summary>
        /// When set, remote artifacts are fetched along with their properties.
        /// </summary>
        [Input("synchronizeProperties")]
        public Input<bool>? SynchronizeProperties { get; set; }

        [Input("unusedArtifactsCleanupPeriodEnabled")]
        public Input<bool>? UnusedArtifactsCleanupPeriodEnabled { get; set; }

        /// <summary>
        /// The number of hours to wait before an artifact is deemed "unused" and eligible for cleanup from the repository. A value
        /// of 0 means automatic cleanup of cached artifacts is disabled.
        /// </summary>
        [Input("unusedArtifactsCleanupPeriodHours")]
        public Input<int>? UnusedArtifactsCleanupPeriodHours { get; set; }

        /// <summary>
        /// The remote repo URL.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public RemoteGenericRepositoryArgs()
        {
        }
        public static new RemoteGenericRepositoryArgs Empty => new RemoteGenericRepositoryArgs();
    }

    public sealed class RemoteGenericRepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
        /// any other host.
        /// </summary>
        [Input("allowAnyHostAuth")]
        public Input<bool>? AllowAnyHostAuth { get; set; }

        /// <summary>
        /// The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
        /// an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
        /// offline. Default to 300.
        /// </summary>
        [Input("assumedOfflinePeriodSecs")]
        public Input<int>? AssumedOfflinePeriodSecs { get; set; }

        /// <summary>
        /// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
        /// resolution.
        /// </summary>
        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        /// <summary>
        /// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
        /// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
        /// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
        /// </summary>
        [Input("blockMismatchingMimeTypes")]
        public Input<bool>? BlockMismatchingMimeTypes { get; set; }

        /// <summary>
        /// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
        /// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
        /// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
        /// </summary>
        [Input("bypassHeadRequests")]
        public Input<bool>? BypassHeadRequests { get; set; }

        [Input("clientTlsCertificate")]
        public Input<string>? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Input<Inputs.RemoteGenericRepositoryContentSynchronisationGetArgs>? ContentSynchronisation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables cookie management if the remote repository uses cookies to manage client state.
        /// </summary>
        [Input("enableCookieManagement")]
        public Input<bool>? EnableCookieManagement { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("failedRetrievalCachePeriodSecs")]
        public Input<int>? FailedRetrievalCachePeriodSecs { get; set; }

        /// <summary>
        /// When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
        /// communicate with this repository.
        /// </summary>
        [Input("hardFail")]
        public Input<bool>? HardFail { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        /// contain spaces or special characters.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
        /// the 'Retrieval Cache Period'. Default value is 'false'.
        /// </summary>
        [Input("listRemoteFolderItems")]
        public Input<bool>? ListRemoteFolderItems { get; set; }

        /// <summary>
        /// The local address to be used when creating connections. Useful for specifying the interface to use on systems with
        /// multiple network interfaces.
        /// </summary>
        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        /// <summary>
        /// The set of mime types that should override the block_mismatching_mime_types setting. Eg:
        /// "application/json,application/xml". Default value is empty.
        /// </summary>
        [Input("mismatchingMimeTypesOverrideList")]
        public Input<string>? MismatchingMimeTypesOverrideList { get; set; }

        /// <summary>
        /// The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
        /// </summary>
        [Input("missedCachePeriodSeconds")]
        public Input<int>? MissedCachePeriodSeconds { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
        /// </summary>
        [Input("offline")]
        public Input<bool>? Offline { get; set; }

        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
        /// </summary>
        [Input("propagateQueryParams")]
        public Input<bool>? PropagateQueryParams { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;

        /// <summary>
        /// List of property set names
        /// </summary>
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Proxy key from Artifactory Proxies settings
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Repository layout key for the remote layout mapping
        /// </summary>
        [Input("remoteRepoLayoutRef")]
        public Input<string>? RemoteRepoLayoutRef { get; set; }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        [Input("shareConfiguration")]
        public Input<bool>? ShareConfiguration { get; set; }

        /// <summary>
        /// Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
        /// operation is considered a retrieval failure.
        /// </summary>
        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
        /// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
        /// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
        /// servers.
        /// </summary>
        [Input("storeArtifactsLocally")]
        public Input<bool>? StoreArtifactsLocally { get; set; }

        /// <summary>
        /// When set, remote artifacts are fetched along with their properties.
        /// </summary>
        [Input("synchronizeProperties")]
        public Input<bool>? SynchronizeProperties { get; set; }

        [Input("unusedArtifactsCleanupPeriodEnabled")]
        public Input<bool>? UnusedArtifactsCleanupPeriodEnabled { get; set; }

        /// <summary>
        /// The number of hours to wait before an artifact is deemed "unused" and eligible for cleanup from the repository. A value
        /// of 0 means automatic cleanup of cached artifacts is disabled.
        /// </summary>
        [Input("unusedArtifactsCleanupPeriodHours")]
        public Input<int>? UnusedArtifactsCleanupPeriodHours { get; set; }

        /// <summary>
        /// The remote repo URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public RemoteGenericRepositoryState()
        {
        }
        public static new RemoteGenericRepositoryState Empty => new RemoteGenericRepositoryState();
    }
}
