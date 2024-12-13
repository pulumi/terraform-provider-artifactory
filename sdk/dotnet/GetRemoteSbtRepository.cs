// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetRemoteSbtRepository
    {
        /// <summary>
        /// Retrieves a remote SBT repository.
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
        ///     var remote_sbt = Artifactory.GetRemoteSbtRepository.Invoke(new()
        ///     {
        ///         Key = "remote-sbt",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRemoteSbtRepositoryResult> InvokeAsync(GetRemoteSbtRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRemoteSbtRepositoryResult>("artifactory:index/getRemoteSbtRepository:getRemoteSbtRepository", args ?? new GetRemoteSbtRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a remote SBT repository.
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
        ///     var remote_sbt = Artifactory.GetRemoteSbtRepository.Invoke(new()
        ///     {
        ///         Key = "remote-sbt",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRemoteSbtRepositoryResult> Invoke(GetRemoteSbtRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRemoteSbtRepositoryResult>("artifactory:index/getRemoteSbtRepository:getRemoteSbtRepository", args ?? new GetRemoteSbtRepositoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a remote SBT repository.
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
        ///     var remote_sbt = Artifactory.GetRemoteSbtRepository.Invoke(new()
        ///     {
        ///         Key = "remote-sbt",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRemoteSbtRepositoryResult> Invoke(GetRemoteSbtRepositoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRemoteSbtRepositoryResult>("artifactory:index/getRemoteSbtRepository:getRemoteSbtRepository", args ?? new GetRemoteSbtRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRemoteSbtRepositoryArgs : global::Pulumi.InvokeArgs
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

        [Input("bypassHeadRequests")]
        public bool? BypassHeadRequests { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("clientTlsCertificate")]
        public string? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Inputs.GetRemoteSbtRepositoryContentSynchronisationArgs? ContentSynchronisation { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("disableProxy")]
        public bool? DisableProxy { get; set; }

        [Input("disableUrlNormalization")]
        public bool? DisableUrlNormalization { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("enableCookieManagement")]
        public bool? EnableCookieManagement { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
        /// </summary>
        [Input("fetchJarsEagerly")]
        public bool? FetchJarsEagerly { get; set; }

        /// <summary>
        /// (Optional, Default: `false`) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
        /// </summary>
        [Input("fetchSourcesEagerly")]
        public bool? FetchSourcesEagerly { get; set; }

        /// <summary>
        /// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
        /// </summary>
        [Input("handleReleases")]
        public bool? HandleReleases { get; set; }

        /// <summary>
        /// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
        /// </summary>
        [Input("handleSnapshots")]
        public bool? HandleSnapshots { get; set; }

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

        [Input("maxUniqueSnapshots")]
        public int? MaxUniqueSnapshots { get; set; }

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

        /// <summary>
        /// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
        /// </summary>
        [Input("rejectInvalidJars")]
        public bool? RejectInvalidJars { get; set; }

        /// <summary>
        /// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
        /// </summary>
        [Input("remoteRepoChecksumPolicyType")]
        public string? RemoteRepoChecksumPolicyType { get; set; }

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

        /// <summary>
        /// (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
        /// </summary>
        [Input("suppressPomConsistencyChecks")]
        public bool? SuppressPomConsistencyChecks { get; set; }

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

        public GetRemoteSbtRepositoryArgs()
        {
        }
        public static new GetRemoteSbtRepositoryArgs Empty => new GetRemoteSbtRepositoryArgs();
    }

    public sealed class GetRemoteSbtRepositoryInvokeArgs : global::Pulumi.InvokeArgs
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

        [Input("bypassHeadRequests")]
        public Input<bool>? BypassHeadRequests { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("clientTlsCertificate")]
        public Input<string>? ClientTlsCertificate { get; set; }

        [Input("contentSynchronisation")]
        public Input<Inputs.GetRemoteSbtRepositoryContentSynchronisationInputArgs>? ContentSynchronisation { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disableProxy")]
        public Input<bool>? DisableProxy { get; set; }

        [Input("disableUrlNormalization")]
        public Input<bool>? DisableUrlNormalization { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("enableCookieManagement")]
        public Input<bool>? EnableCookieManagement { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
        /// </summary>
        [Input("fetchJarsEagerly")]
        public Input<bool>? FetchJarsEagerly { get; set; }

        /// <summary>
        /// (Optional, Default: `false`) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
        /// </summary>
        [Input("fetchSourcesEagerly")]
        public Input<bool>? FetchSourcesEagerly { get; set; }

        /// <summary>
        /// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
        /// </summary>
        [Input("handleReleases")]
        public Input<bool>? HandleReleases { get; set; }

        /// <summary>
        /// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
        /// </summary>
        [Input("handleSnapshots")]
        public Input<bool>? HandleSnapshots { get; set; }

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

        [Input("maxUniqueSnapshots")]
        public Input<int>? MaxUniqueSnapshots { get; set; }

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

        /// <summary>
        /// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
        /// </summary>
        [Input("rejectInvalidJars")]
        public Input<bool>? RejectInvalidJars { get; set; }

        /// <summary>
        /// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
        /// </summary>
        [Input("remoteRepoChecksumPolicyType")]
        public Input<string>? RemoteRepoChecksumPolicyType { get; set; }

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

        /// <summary>
        /// (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
        /// </summary>
        [Input("suppressPomConsistencyChecks")]
        public Input<bool>? SuppressPomConsistencyChecks { get; set; }

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

        public GetRemoteSbtRepositoryInvokeArgs()
        {
        }
        public static new GetRemoteSbtRepositoryInvokeArgs Empty => new GetRemoteSbtRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRemoteSbtRepositoryResult
    {
        public readonly bool? AllowAnyHostAuth;
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly int? AssumedOfflinePeriodSecs;
        public readonly bool? BlackedOut;
        public readonly bool? BlockMismatchingMimeTypes;
        public readonly bool? BypassHeadRequests;
        public readonly bool? CdnRedirect;
        public readonly string ClientTlsCertificate;
        public readonly Outputs.GetRemoteSbtRepositoryContentSynchronisationResult ContentSynchronisation;
        public readonly string? Description;
        public readonly bool? DisableProxy;
        public readonly bool? DisableUrlNormalization;
        public readonly bool? DownloadDirect;
        public readonly bool? EnableCookieManagement;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
        /// </summary>
        public readonly bool? FetchJarsEagerly;
        /// <summary>
        /// (Optional, Default: `false`) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
        /// </summary>
        public readonly bool? FetchSourcesEagerly;
        /// <summary>
        /// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
        /// </summary>
        public readonly bool? HandleReleases;
        /// <summary>
        /// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
        /// </summary>
        public readonly bool? HandleSnapshots;
        public readonly bool? HardFail;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly bool? ListRemoteFolderItems;
        public readonly string? LocalAddress;
        public readonly int? MaxUniqueSnapshots;
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
        /// <summary>
        /// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
        /// </summary>
        public readonly bool? RejectInvalidJars;
        /// <summary>
        /// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
        /// </summary>
        public readonly string? RemoteRepoChecksumPolicyType;
        public readonly string? RemoteRepoLayoutRef;
        public readonly string? RepoLayoutRef;
        public readonly int? RetrievalCachePeriodSeconds;
        public readonly bool ShareConfiguration;
        public readonly int? SocketTimeoutMillis;
        public readonly bool? StoreArtifactsLocally;
        /// <summary>
        /// (Optional, Default: `true`) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`.
        /// </summary>
        public readonly bool? SuppressPomConsistencyChecks;
        public readonly bool? SynchronizeProperties;
        public readonly int? UnusedArtifactsCleanupPeriodHours;
        public readonly string? Url;
        public readonly string? Username;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetRemoteSbtRepositoryResult(
            bool? allowAnyHostAuth,

            bool? archiveBrowsingEnabled,

            int? assumedOfflinePeriodSecs,

            bool? blackedOut,

            bool? blockMismatchingMimeTypes,

            bool? bypassHeadRequests,

            bool? cdnRedirect,

            string clientTlsCertificate,

            Outputs.GetRemoteSbtRepositoryContentSynchronisationResult contentSynchronisation,

            string? description,

            bool? disableProxy,

            bool? disableUrlNormalization,

            bool? downloadDirect,

            bool? enableCookieManagement,

            string? excludesPattern,

            bool? fetchJarsEagerly,

            bool? fetchSourcesEagerly,

            bool? handleReleases,

            bool? handleSnapshots,

            bool? hardFail,

            string id,

            string? includesPattern,

            string key,

            bool? listRemoteFolderItems,

            string? localAddress,

            int? maxUniqueSnapshots,

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

            bool? rejectInvalidJars,

            string? remoteRepoChecksumPolicyType,

            string? remoteRepoLayoutRef,

            string? repoLayoutRef,

            int? retrievalCachePeriodSeconds,

            bool shareConfiguration,

            int? socketTimeoutMillis,

            bool? storeArtifactsLocally,

            bool? suppressPomConsistencyChecks,

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
            BypassHeadRequests = bypassHeadRequests;
            CdnRedirect = cdnRedirect;
            ClientTlsCertificate = clientTlsCertificate;
            ContentSynchronisation = contentSynchronisation;
            Description = description;
            DisableProxy = disableProxy;
            DisableUrlNormalization = disableUrlNormalization;
            DownloadDirect = downloadDirect;
            EnableCookieManagement = enableCookieManagement;
            ExcludesPattern = excludesPattern;
            FetchJarsEagerly = fetchJarsEagerly;
            FetchSourcesEagerly = fetchSourcesEagerly;
            HandleReleases = handleReleases;
            HandleSnapshots = handleSnapshots;
            HardFail = hardFail;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            ListRemoteFolderItems = listRemoteFolderItems;
            LocalAddress = localAddress;
            MaxUniqueSnapshots = maxUniqueSnapshots;
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
            RejectInvalidJars = rejectInvalidJars;
            RemoteRepoChecksumPolicyType = remoteRepoChecksumPolicyType;
            RemoteRepoLayoutRef = remoteRepoLayoutRef;
            RepoLayoutRef = repoLayoutRef;
            RetrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            ShareConfiguration = shareConfiguration;
            SocketTimeoutMillis = socketTimeoutMillis;
            StoreArtifactsLocally = storeArtifactsLocally;
            SuppressPomConsistencyChecks = suppressPomConsistencyChecks;
            SynchronizeProperties = synchronizeProperties;
            UnusedArtifactsCleanupPeriodHours = unusedArtifactsCleanupPeriodHours;
            Url = url;
            Username = username;
            XrayIndex = xrayIndex;
        }
    }
}
