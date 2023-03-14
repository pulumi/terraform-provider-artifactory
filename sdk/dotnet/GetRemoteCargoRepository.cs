// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetRemoteCargoRepository
    {
        /// <summary>
        /// Retrieves a remote Cargo repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var remote_cargo = Artifactory.GetRemoteCargoRepository.Invoke(new()
        ///     {
        ///         Key = "remote-cargo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRemoteCargoRepositoryResult> InvokeAsync(GetRemoteCargoRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRemoteCargoRepositoryResult>("artifactory:index/getRemoteCargoRepository:getRemoteCargoRepository", args ?? new GetRemoteCargoRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a remote Cargo repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var remote_cargo = Artifactory.GetRemoteCargoRepository.Invoke(new()
        ///     {
        ///         Key = "remote-cargo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRemoteCargoRepositoryResult> Invoke(GetRemoteCargoRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRemoteCargoRepositoryResult>("artifactory:index/getRemoteCargoRepository:getRemoteCargoRepository", args ?? new GetRemoteCargoRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRemoteCargoRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("allowAnyHostAuth")]
        public bool? AllowAnyHostAuth { get; set; }

        /// <summary>
        /// (Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
        /// </summary>
        [Input("anonymousAccess")]
        public bool? AnonymousAccess { get; set; }

        [Input("assumedOfflinePeriodSecs")]
        public int? AssumedOfflinePeriodSecs { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("blockMismatchingMimeTypes")]
        public bool? BlockMismatchingMimeTypes { get; set; }

        [Input("bypassHeadRequests")]
        public bool? BypassHeadRequests { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("clientTlsCertificate")]
        public string? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Inputs.GetRemoteCargoRepositoryContentSynchronisationArgs? ContentSynchronisation { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("enableCookieManagement")]
        public bool? EnableCookieManagement { get; set; }

        /// <summary>
        /// (Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
        /// </summary>
        [Input("enableSparseIndex")]
        public bool? EnableSparseIndex { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional) This is the index url, expected to be a git repository. Default value is `https://github.com/rust-lang/crates.io-index`.
        /// </summary>
        [Input("gitRegistryUrl")]
        public string? GitRegistryUrl { get; set; }

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

        public GetRemoteCargoRepositoryArgs()
        {
        }
        public static new GetRemoteCargoRepositoryArgs Empty => new GetRemoteCargoRepositoryArgs();
    }

    public sealed class GetRemoteCargoRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("allowAnyHostAuth")]
        public Input<bool>? AllowAnyHostAuth { get; set; }

        /// <summary>
        /// (Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
        /// </summary>
        [Input("anonymousAccess")]
        public Input<bool>? AnonymousAccess { get; set; }

        [Input("assumedOfflinePeriodSecs")]
        public Input<int>? AssumedOfflinePeriodSecs { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("blockMismatchingMimeTypes")]
        public Input<bool>? BlockMismatchingMimeTypes { get; set; }

        [Input("bypassHeadRequests")]
        public Input<bool>? BypassHeadRequests { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("clientTlsCertificate")]
        public Input<string>? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Input<Inputs.GetRemoteCargoRepositoryContentSynchronisationInputArgs>? ContentSynchronisation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("enableCookieManagement")]
        public Input<bool>? EnableCookieManagement { get; set; }

        /// <summary>
        /// (Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
        /// </summary>
        [Input("enableSparseIndex")]
        public Input<bool>? EnableSparseIndex { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional) This is the index url, expected to be a git repository. Default value is `https://github.com/rust-lang/crates.io-index`.
        /// </summary>
        [Input("gitRegistryUrl")]
        public Input<string>? GitRegistryUrl { get; set; }

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

        public GetRemoteCargoRepositoryInvokeArgs()
        {
        }
        public static new GetRemoteCargoRepositoryInvokeArgs Empty => new GetRemoteCargoRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRemoteCargoRepositoryResult
    {
        public readonly bool? AllowAnyHostAuth;
        /// <summary>
        /// (Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
        /// </summary>
        public readonly bool? AnonymousAccess;
        public readonly int? AssumedOfflinePeriodSecs;
        public readonly bool? BlackedOut;
        public readonly bool? BlockMismatchingMimeTypes;
        public readonly bool? BypassHeadRequests;
        public readonly bool? CdnRedirect;
        public readonly string ClientTlsCertificate;
        public readonly Outputs.GetRemoteCargoRepositoryContentSynchronisationResult ContentSynchronisation;
        public readonly string? Description;
        public readonly bool? DownloadDirect;
        public readonly bool? EnableCookieManagement;
        /// <summary>
        /// (Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
        /// </summary>
        public readonly bool? EnableSparseIndex;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// (Optional) This is the index url, expected to be a git repository. Default value is `https://github.com/rust-lang/crates.io-index`.
        /// </summary>
        public readonly string? GitRegistryUrl;
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
        private GetRemoteCargoRepositoryResult(
            bool? allowAnyHostAuth,

            bool? anonymousAccess,

            int? assumedOfflinePeriodSecs,

            bool? blackedOut,

            bool? blockMismatchingMimeTypes,

            bool? bypassHeadRequests,

            bool? cdnRedirect,

            string clientTlsCertificate,

            Outputs.GetRemoteCargoRepositoryContentSynchronisationResult contentSynchronisation,

            string? description,

            bool? downloadDirect,

            bool? enableCookieManagement,

            bool? enableSparseIndex,

            string? excludesPattern,

            string? gitRegistryUrl,

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
            AnonymousAccess = anonymousAccess;
            AssumedOfflinePeriodSecs = assumedOfflinePeriodSecs;
            BlackedOut = blackedOut;
            BlockMismatchingMimeTypes = blockMismatchingMimeTypes;
            BypassHeadRequests = bypassHeadRequests;
            CdnRedirect = cdnRedirect;
            ClientTlsCertificate = clientTlsCertificate;
            ContentSynchronisation = contentSynchronisation;
            Description = description;
            DownloadDirect = downloadDirect;
            EnableCookieManagement = enableCookieManagement;
            EnableSparseIndex = enableSparseIndex;
            ExcludesPattern = excludesPattern;
            GitRegistryUrl = gitRegistryUrl;
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
