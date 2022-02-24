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
    /// ## # Artifactory Remote Repository Resource
    /// 
    /// Provides an Artifactory remote `helm` repository resource. This provides helm specific fields and is the only way to get them.
    /// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Package+Management),
    /// although helm is (currently) not listed as a supported format
    /// 
    /// ## Example Usage
    /// 
    /// Includes only new and relevant fields, for anything else, see: generic repo.
    /// ```csharp
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var helm_remote = new Artifactory.RemoteHelmRepository("helm-remote", new Artifactory.RemoteHelmRepositoryArgs
    ///         {
    ///             ExternalDependenciesEnabled = true,
    ///             ExternalDependenciesPatterns = 
    ///             {
    ///                 "**github.com**",
    ///             },
    ///             HelmChartsBaseUrl = "https://foo.com",
    ///             Key = "helm-remote-foo25",
    ///             Url = "https://repo.chartcenter.io/",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/remoteHelmRepository:RemoteHelmRepository")]
    public partial class RemoteHelmRepository : Pulumi.CustomResource
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

        /// <summary>
        /// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
        /// </summary>
        [Output("contentSynchronisation")]
        public Output<Outputs.RemoteHelmRepositoryContentSynchronisation> ContentSynchronisation { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Enables cookie management if the remote repository uses cookies to manage client state.
        /// </summary>
        [Output("enableCookieManagement")]
        public Output<bool> EnableCookieManagement { get; private set; } = null!;

        [Output("excludesPattern")]
        public Output<string> ExcludesPattern { get; private set; } = null!;

        /// <summary>
        /// When set, external dependencies are rewritten.
        /// </summary>
        [Output("externalDependenciesEnabled")]
        public Output<bool?> ExternalDependenciesEnabled { get; private set; } = null!;

        /// <summary>
        /// An Allow List of Ant-style path expressions that specify where external
        /// dependencies may be downloaded from. By default, this is an empty list which means that no dependencies may be downloaded
        /// from external sources. Note that the official documentation states the default is '**', which is correct when creating
        /// repositories in the UI, but incorrect for the API.
        /// </summary>
        [Output("externalDependenciesPatterns")]
        public Output<ImmutableArray<string>> ExternalDependenciesPatterns { get; private set; } = null!;

        [Output("failedRetrievalCachePeriodSecs")]
        public Output<int> FailedRetrievalCachePeriodSecs { get; private set; } = null!;

        [Output("hardFail")]
        public Output<bool> HardFail { get; private set; } = null!;

        /// <summary>
        /// - No documentation is available. Hopefully you know what this means
        /// </summary>
        [Output("helmChartsBaseUrl")]
        public Output<string?> HelmChartsBaseUrl { get; private set; } = null!;

        [Output("includesPattern")]
        public Output<string> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// The repository identifier. Must be unique system-wide
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        [Output("localAddress")]
        public Output<string?> LocalAddress { get; private set; } = null!;

        /// <summary>
        /// This is actually the missedRetrievalCachePeriodSecs in the API
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
        /// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        [Output("propagateQueryParams")]
        public Output<bool?> PropagateQueryParams { get; private set; } = null!;

        [Output("propertySets")]
        public Output<ImmutableArray<string>> PropertySets { get; private set; } = null!;

        [Output("proxy")]
        public Output<string> Proxy { get; private set; } = null!;

        [Output("remoteRepoLayoutRef")]
        public Output<string> RemoteRepoLayoutRef { get; private set; } = null!;

        [Output("repoLayoutRef")]
        public Output<string> RepoLayoutRef { get; private set; } = null!;

        /// <summary>
        /// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
        /// </summary>
        [Output("retrievalCachePeriodSeconds")]
        public Output<int> RetrievalCachePeriodSeconds { get; private set; } = null!;

        [Output("shareConfiguration")]
        public Output<bool> ShareConfiguration { get; private set; } = null!;

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

        [Output("unusedArtifactsCleanupPeriodHours")]
        public Output<int> UnusedArtifactsCleanupPeriodHours { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
        /// </summary>
        [Output("xrayIndex")]
        public Output<bool?> XrayIndex { get; private set; } = null!;


        /// <summary>
        /// Create a RemoteHelmRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RemoteHelmRepository(string name, RemoteHelmRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/remoteHelmRepository:RemoteHelmRepository", name, args ?? new RemoteHelmRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RemoteHelmRepository(string name, Input<string> id, RemoteHelmRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/remoteHelmRepository:RemoteHelmRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RemoteHelmRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RemoteHelmRepository Get(string name, Input<string> id, RemoteHelmRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new RemoteHelmRepository(name, id, state, options);
        }
    }

    public sealed class RemoteHelmRepositoryArgs : Pulumi.ResourceArgs
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

        /// <summary>
        /// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
        /// </summary>
        [Input("contentSynchronisation")]
        public Input<Inputs.RemoteHelmRepositoryContentSynchronisationArgs>? ContentSynchronisation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables cookie management if the remote repository uses cookies to manage client state.
        /// </summary>
        [Input("enableCookieManagement")]
        public Input<bool>? EnableCookieManagement { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// When set, external dependencies are rewritten.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public Input<bool>? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private InputList<string>? _externalDependenciesPatterns;

        /// <summary>
        /// An Allow List of Ant-style path expressions that specify where external
        /// dependencies may be downloaded from. By default, this is an empty list which means that no dependencies may be downloaded
        /// from external sources. Note that the official documentation states the default is '**', which is correct when creating
        /// repositories in the UI, but incorrect for the API.
        /// </summary>
        public InputList<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new InputList<string>());
            set => _externalDependenciesPatterns = value;
        }

        [Input("hardFail")]
        public Input<bool>? HardFail { get; set; }

        /// <summary>
        /// - No documentation is available. Hopefully you know what this means
        /// </summary>
        [Input("helmChartsBaseUrl")]
        public Input<string>? HelmChartsBaseUrl { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// The repository identifier. Must be unique system-wide
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        /// <summary>
        /// This is actually the missedRetrievalCachePeriodSecs in the API
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
        /// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propagateQueryParams")]
        public Input<bool>? PropagateQueryParams { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("remoteRepoLayoutRef")]
        public Input<string>? RemoteRepoLayoutRef { get; set; }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        [Input("shareConfiguration")]
        public Input<bool>? ShareConfiguration { get; set; }

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

        [Input("unusedArtifactsCleanupPeriodHours")]
        public Input<int>? UnusedArtifactsCleanupPeriodHours { get; set; }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public RemoteHelmRepositoryArgs()
        {
        }
    }

    public sealed class RemoteHelmRepositoryState : Pulumi.ResourceArgs
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

        /// <summary>
        /// Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
        /// </summary>
        [Input("contentSynchronisation")]
        public Input<Inputs.RemoteHelmRepositoryContentSynchronisationGetArgs>? ContentSynchronisation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables cookie management if the remote repository uses cookies to manage client state.
        /// </summary>
        [Input("enableCookieManagement")]
        public Input<bool>? EnableCookieManagement { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// When set, external dependencies are rewritten.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public Input<bool>? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private InputList<string>? _externalDependenciesPatterns;

        /// <summary>
        /// An Allow List of Ant-style path expressions that specify where external
        /// dependencies may be downloaded from. By default, this is an empty list which means that no dependencies may be downloaded
        /// from external sources. Note that the official documentation states the default is '**', which is correct when creating
        /// repositories in the UI, but incorrect for the API.
        /// </summary>
        public InputList<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new InputList<string>());
            set => _externalDependenciesPatterns = value;
        }

        [Input("failedRetrievalCachePeriodSecs")]
        public Input<int>? FailedRetrievalCachePeriodSecs { get; set; }

        [Input("hardFail")]
        public Input<bool>? HardFail { get; set; }

        /// <summary>
        /// - No documentation is available. Hopefully you know what this means
        /// </summary>
        [Input("helmChartsBaseUrl")]
        public Input<string>? HelmChartsBaseUrl { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// The repository identifier. Must be unique system-wide
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        /// <summary>
        /// This is actually the missedRetrievalCachePeriodSecs in the API
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
        /// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propagateQueryParams")]
        public Input<bool>? PropagateQueryParams { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("remoteRepoLayoutRef")]
        public Input<string>? RemoteRepoLayoutRef { get; set; }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        [Input("shareConfiguration")]
        public Input<bool>? ShareConfiguration { get; set; }

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

        [Input("unusedArtifactsCleanupPeriodHours")]
        public Input<int>? UnusedArtifactsCleanupPeriodHours { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public RemoteHelmRepositoryState()
        {
        }
    }
}
