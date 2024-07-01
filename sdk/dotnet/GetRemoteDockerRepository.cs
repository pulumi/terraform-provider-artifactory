// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetRemoteDockerRepository
    {
        /// <summary>
        /// Retrieves a remote Docker repository.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var remote_docker = Artifactory.GetRemoteDockerRepository.Invoke(new()
        ///     {
        ///         Key = "remote-docker",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRemoteDockerRepositoryResult> InvokeAsync(GetRemoteDockerRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRemoteDockerRepositoryResult>("artifactory:index/getRemoteDockerRepository:getRemoteDockerRepository", args ?? new GetRemoteDockerRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a remote Docker repository.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var remote_docker = Artifactory.GetRemoteDockerRepository.Invoke(new()
        ///     {
        ///         Key = "remote-docker",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRemoteDockerRepositoryResult> Invoke(GetRemoteDockerRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRemoteDockerRepositoryResult>("artifactory:index/getRemoteDockerRepository:getRemoteDockerRepository", args ?? new GetRemoteDockerRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRemoteDockerRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("allowAnyHostAuth")]
        public bool? AllowAnyHostAuth { get; set; }

        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("assumedOfflinePeriodSecs")]
        public int? AssumedOfflinePeriodSecs { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("blockMismatchingMimeTypes")]
        public bool? BlockMismatchingMimeTypes { get; set; }

        /// <summary>
        /// (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
        /// </summary>
        [Input("blockPushingSchema1")]
        public bool? BlockPushingSchema1 { get; set; }

        [Input("bypassHeadRequests")]
        public bool? BypassHeadRequests { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("clientTlsCertificate")]
        public string? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Inputs.GetRemoteDockerRepositoryContentSynchronisationArgs? ContentSynchronisation { get; set; }

        [Input("curated")]
        public bool? Curated { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("disableProxy")]
        public bool? DisableProxy { get; set; }

        /// <summary>
        /// (Optional) Whether to disable URL normalization.
        /// </summary>
        [Input("disableUrlNormalization")]
        public bool? DisableUrlNormalization { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("enableCookieManagement")]
        public bool? EnableCookieManagement { get; set; }

        /// <summary>
        /// (Optional) Enable token (Bearer) based authentication.
        /// </summary>
        [Input("enableTokenAuthentication")]
        public bool? EnableTokenAuthentication { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional) Also known as 'Foreign Layers Caching' on the UI.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public bool? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private List<string>? _externalDependenciesPatterns;

        /// <summary>
        /// (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
        /// </summary>
        public List<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new List<string>());
            set => _externalDependenciesPatterns = value;
        }

        [Input("hardFail")]
        public bool? HardFail { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("listRemoteFolderItems")]
        public bool? ListRemoteFolderItems { get; set; }

        [Input("localAddress")]
        public string? LocalAddress { get; set; }

        [Input("metadataRetrievalTimeoutSecs")]
        public int? MetadataRetrievalTimeoutSecs { get; set; }

        [Input("mismatchingMimeTypesOverrideList")]
        public string? MismatchingMimeTypesOverrideList { get; set; }

        [Input("missedCachePeriodSeconds")]
        public int? MissedCachePeriodSeconds { get; set; }

        [Input("notes")]
        public string? Notes { get; set; }

        [Input("offline")]
        public bool? Offline { get; set; }

        [Input("password")]
        private string? _password;
        public string? Password
        {
            get => _password;
            set => _password = value;
        }

        [Input("priorityResolution")]
        public bool? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private List<string>? _projectEnvironments;
        public List<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new List<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("projectKey")]
        public string? ProjectKey { get; set; }

        [Input("propertySets")]
        private List<string>? _propertySets;
        public List<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new List<string>());
            set => _propertySets = value;
        }

        [Input("proxy")]
        public string? Proxy { get; set; }

        [Input("queryParams")]
        public string? QueryParams { get; set; }

        [Input("remoteRepoLayoutRef")]
        public string? RemoteRepoLayoutRef { get; set; }

        [Input("repoLayoutRef")]
        public string? RepoLayoutRef { get; set; }

        [Input("retrievalCachePeriodSeconds")]
        public int? RetrievalCachePeriodSeconds { get; set; }

        [Input("shareConfiguration")]
        public bool? ShareConfiguration { get; set; }

        [Input("socketTimeoutMillis")]
        public int? SocketTimeoutMillis { get; set; }

        [Input("storeArtifactsLocally")]
        public bool? StoreArtifactsLocally { get; set; }

        [Input("synchronizeProperties")]
        public bool? SynchronizeProperties { get; set; }

        [Input("unusedArtifactsCleanupPeriodHours")]
        public int? UnusedArtifactsCleanupPeriodHours { get; set; }

        [Input("url")]
        public string? Url { get; set; }

        [Input("username")]
        public string? Username { get; set; }

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetRemoteDockerRepositoryArgs()
        {
        }
        public static new GetRemoteDockerRepositoryArgs Empty => new GetRemoteDockerRepositoryArgs();
    }

    public sealed class GetRemoteDockerRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("allowAnyHostAuth")]
        public Input<bool>? AllowAnyHostAuth { get; set; }

        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("assumedOfflinePeriodSecs")]
        public Input<int>? AssumedOfflinePeriodSecs { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("blockMismatchingMimeTypes")]
        public Input<bool>? BlockMismatchingMimeTypes { get; set; }

        /// <summary>
        /// (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
        /// </summary>
        [Input("blockPushingSchema1")]
        public Input<bool>? BlockPushingSchema1 { get; set; }

        [Input("bypassHeadRequests")]
        public Input<bool>? BypassHeadRequests { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("clientTlsCertificate")]
        public Input<string>? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Input<Inputs.GetRemoteDockerRepositoryContentSynchronisationInputArgs>? ContentSynchronisation { get; set; }

        [Input("curated")]
        public Input<bool>? Curated { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disableProxy")]
        public Input<bool>? DisableProxy { get; set; }

        /// <summary>
        /// (Optional) Whether to disable URL normalization.
        /// </summary>
        [Input("disableUrlNormalization")]
        public Input<bool>? DisableUrlNormalization { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("enableCookieManagement")]
        public Input<bool>? EnableCookieManagement { get; set; }

        /// <summary>
        /// (Optional) Enable token (Bearer) based authentication.
        /// </summary>
        [Input("enableTokenAuthentication")]
        public Input<bool>? EnableTokenAuthentication { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional) Also known as 'Foreign Layers Caching' on the UI.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public Input<bool>? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private InputList<string>? _externalDependenciesPatterns;

        /// <summary>
        /// (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
        /// </summary>
        public InputList<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new InputList<string>());
            set => _externalDependenciesPatterns = value;
        }

        [Input("hardFail")]
        public Input<bool>? HardFail { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("listRemoteFolderItems")]
        public Input<bool>? ListRemoteFolderItems { get; set; }

        [Input("localAddress")]
        public Input<string>? LocalAddress { get; set; }

        [Input("metadataRetrievalTimeoutSecs")]
        public Input<int>? MetadataRetrievalTimeoutSecs { get; set; }

        [Input("mismatchingMimeTypesOverrideList")]
        public Input<string>? MismatchingMimeTypesOverrideList { get; set; }

        [Input("missedCachePeriodSeconds")]
        public Input<int>? MissedCachePeriodSeconds { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("offline")]
        public Input<bool>? Offline { get; set; }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("queryParams")]
        public Input<string>? QueryParams { get; set; }

        [Input("remoteRepoLayoutRef")]
        public Input<string>? RemoteRepoLayoutRef { get; set; }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        [Input("shareConfiguration")]
        public Input<bool>? ShareConfiguration { get; set; }

        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        [Input("storeArtifactsLocally")]
        public Input<bool>? StoreArtifactsLocally { get; set; }

        [Input("synchronizeProperties")]
        public Input<bool>? SynchronizeProperties { get; set; }

        [Input("unusedArtifactsCleanupPeriodHours")]
        public Input<int>? UnusedArtifactsCleanupPeriodHours { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetRemoteDockerRepositoryInvokeArgs()
        {
        }
        public static new GetRemoteDockerRepositoryInvokeArgs Empty => new GetRemoteDockerRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRemoteDockerRepositoryResult
    {
        public readonly bool? AllowAnyHostAuth;
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly int? AssumedOfflinePeriodSecs;
        public readonly bool? BlackedOut;
        public readonly bool? BlockMismatchingMimeTypes;
        /// <summary>
        /// (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
        /// </summary>
        public readonly bool BlockPushingSchema1;
        public readonly bool? BypassHeadRequests;
        public readonly bool? CdnRedirect;
        public readonly string ClientTlsCertificate;
        public readonly Outputs.GetRemoteDockerRepositoryContentSynchronisationResult ContentSynchronisation;
        public readonly bool? Curated;
        public readonly string? Description;
        public readonly bool? DisableProxy;
        /// <summary>
        /// (Optional) Whether to disable URL normalization.
        /// </summary>
        public readonly bool? DisableUrlNormalization;
        public readonly bool? DownloadDirect;
        public readonly bool? EnableCookieManagement;
        /// <summary>
        /// (Optional) Enable token (Bearer) based authentication.
        /// </summary>
        public readonly bool EnableTokenAuthentication;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// (Optional) Also known as 'Foreign Layers Caching' on the UI.
        /// </summary>
        public readonly bool? ExternalDependenciesEnabled;
        /// <summary>
        /// (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
        /// </summary>
        public readonly ImmutableArray<string> ExternalDependenciesPatterns;
        public readonly bool? HardFail;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly bool? ListRemoteFolderItems;
        public readonly string? LocalAddress;
        public readonly int? MetadataRetrievalTimeoutSecs;
        public readonly string? MismatchingMimeTypesOverrideList;
        public readonly int? MissedCachePeriodSeconds;
        public readonly string? Notes;
        public readonly bool? Offline;
        public readonly string PackageType;
        public readonly string? Password;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectId;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        public readonly string? Proxy;
        public readonly string? QueryParams;
        public readonly string? RemoteRepoLayoutRef;
        public readonly string? RepoLayoutRef;
        public readonly int? RetrievalCachePeriodSeconds;
        public readonly bool ShareConfiguration;
        public readonly int? SocketTimeoutMillis;
        public readonly bool? StoreArtifactsLocally;
        public readonly bool? SynchronizeProperties;
        public readonly int? UnusedArtifactsCleanupPeriodHours;
        public readonly string? Url;
        public readonly string? Username;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetRemoteDockerRepositoryResult(
            bool? allowAnyHostAuth,

            bool? archiveBrowsingEnabled,

            int? assumedOfflinePeriodSecs,

            bool? blackedOut,

            bool? blockMismatchingMimeTypes,

            bool blockPushingSchema1,

            bool? bypassHeadRequests,

            bool? cdnRedirect,

            string clientTlsCertificate,

            Outputs.GetRemoteDockerRepositoryContentSynchronisationResult contentSynchronisation,

            bool? curated,

            string? description,

            bool? disableProxy,

            bool? disableUrlNormalization,

            bool? downloadDirect,

            bool? enableCookieManagement,

            bool enableTokenAuthentication,

            string? excludesPattern,

            bool? externalDependenciesEnabled,

            ImmutableArray<string> externalDependenciesPatterns,

            bool? hardFail,

            string id,

            string? includesPattern,

            string key,

            bool? listRemoteFolderItems,

            string? localAddress,

            int? metadataRetrievalTimeoutSecs,

            string? mismatchingMimeTypesOverrideList,

            int? missedCachePeriodSeconds,

            string? notes,

            bool? offline,

            string packageType,

            string? password,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectId,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? proxy,

            string? queryParams,

            string? remoteRepoLayoutRef,

            string? repoLayoutRef,

            int? retrievalCachePeriodSeconds,

            bool shareConfiguration,

            int? socketTimeoutMillis,

            bool? storeArtifactsLocally,

            bool? synchronizeProperties,

            int? unusedArtifactsCleanupPeriodHours,

            string? url,

            string? username,

            bool? xrayIndex)
        {
            AllowAnyHostAuth = allowAnyHostAuth;
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            AssumedOfflinePeriodSecs = assumedOfflinePeriodSecs;
            BlackedOut = blackedOut;
            BlockMismatchingMimeTypes = blockMismatchingMimeTypes;
            BlockPushingSchema1 = blockPushingSchema1;
            BypassHeadRequests = bypassHeadRequests;
            CdnRedirect = cdnRedirect;
            ClientTlsCertificate = clientTlsCertificate;
            ContentSynchronisation = contentSynchronisation;
            Curated = curated;
            Description = description;
            DisableProxy = disableProxy;
            DisableUrlNormalization = disableUrlNormalization;
            DownloadDirect = downloadDirect;
            EnableCookieManagement = enableCookieManagement;
            EnableTokenAuthentication = enableTokenAuthentication;
            ExcludesPattern = excludesPattern;
            ExternalDependenciesEnabled = externalDependenciesEnabled;
            ExternalDependenciesPatterns = externalDependenciesPatterns;
            HardFail = hardFail;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            ListRemoteFolderItems = listRemoteFolderItems;
            LocalAddress = localAddress;
            MetadataRetrievalTimeoutSecs = metadataRetrievalTimeoutSecs;
            MismatchingMimeTypesOverrideList = mismatchingMimeTypesOverrideList;
            MissedCachePeriodSeconds = missedCachePeriodSeconds;
            Notes = notes;
            Offline = offline;
            PackageType = packageType;
            Password = password;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectId = projectId;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            Proxy = proxy;
            QueryParams = queryParams;
            RemoteRepoLayoutRef = remoteRepoLayoutRef;
            RepoLayoutRef = repoLayoutRef;
            RetrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            ShareConfiguration = shareConfiguration;
            SocketTimeoutMillis = socketTimeoutMillis;
            StoreArtifactsLocally = storeArtifactsLocally;
            SynchronizeProperties = synchronizeProperties;
            UnusedArtifactsCleanupPeriodHours = unusedArtifactsCleanupPeriodHours;
            Url = url;
            Username = username;
            XrayIndex = xrayIndex;
        }
    }
}
