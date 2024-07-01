// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFederatedCargoRepository
    {
        /// <summary>
        /// Retrieves a federated Cargo repository.
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
        ///     var federated_test_cargo_repo = Artifactory.GetFederatedCargoRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-cargo-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFederatedCargoRepositoryResult> InvokeAsync(GetFederatedCargoRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFederatedCargoRepositoryResult>("artifactory:index/getFederatedCargoRepository:getFederatedCargoRepository", args ?? new GetFederatedCargoRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated Cargo repository.
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
        ///     var federated_test_cargo_repo = Artifactory.GetFederatedCargoRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-cargo-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFederatedCargoRepositoryResult> Invoke(GetFederatedCargoRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedCargoRepositoryResult>("artifactory:index/getFederatedCargoRepository:getFederatedCargoRepository", args ?? new GetFederatedCargoRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFederatedCargoRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("anonymousAccess")]
        public bool? AnonymousAccess { get; set; }

        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("cleanupOnDelete")]
        public bool? CleanupOnDelete { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
        /// </summary>
        [Input("disableProxy")]
        public bool? DisableProxy { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("enableSparseIndex")]
        public bool? EnableSparseIndex { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        [Input("indexCompressionFormats")]
        private List<string>? _indexCompressionFormats;
        public List<string> IndexCompressionFormats
        {
            get => _indexCompressionFormats ?? (_indexCompressionFormats = new List<string>());
            set => _indexCompressionFormats = value;
        }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("members")]
        private List<Inputs.GetFederatedCargoRepositoryMemberArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public List<Inputs.GetFederatedCargoRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new List<Inputs.GetFederatedCargoRepositoryMemberArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public string? Notes { get; set; }

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

        /// <summary>
        /// Proxy key from Artifactory Proxies settings.
        /// </summary>
        [Input("proxy")]
        public string? Proxy { get; set; }

        [Input("repoLayoutRef")]
        public string? RepoLayoutRef { get; set; }

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetFederatedCargoRepositoryArgs()
        {
        }
        public static new GetFederatedCargoRepositoryArgs Empty => new GetFederatedCargoRepositoryArgs();
    }

    public sealed class GetFederatedCargoRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("anonymousAccess")]
        public Input<bool>? AnonymousAccess { get; set; }

        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("cleanupOnDelete")]
        public Input<bool>? CleanupOnDelete { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
        /// </summary>
        [Input("disableProxy")]
        public Input<bool>? DisableProxy { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("enableSparseIndex")]
        public Input<bool>? EnableSparseIndex { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        [Input("indexCompressionFormats")]
        private InputList<string>? _indexCompressionFormats;
        public InputList<string> IndexCompressionFormats
        {
            get => _indexCompressionFormats ?? (_indexCompressionFormats = new InputList<string>());
            set => _indexCompressionFormats = value;
        }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("members")]
        private InputList<Inputs.GetFederatedCargoRepositoryMemberInputArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.GetFederatedCargoRepositoryMemberInputArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GetFederatedCargoRepositoryMemberInputArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

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

        /// <summary>
        /// Proxy key from Artifactory Proxies settings.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetFederatedCargoRepositoryInvokeArgs()
        {
        }
        public static new GetFederatedCargoRepositoryInvokeArgs Empty => new GetFederatedCargoRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetFederatedCargoRepositoryResult
    {
        public readonly bool? AnonymousAccess;
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly bool? BlackedOut;
        public readonly bool? CdnRedirect;
        public readonly bool? CleanupOnDelete;
        public readonly string? Description;
        /// <summary>
        /// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
        /// </summary>
        public readonly bool? DisableProxy;
        public readonly bool? DownloadDirect;
        public readonly bool? EnableSparseIndex;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly ImmutableArray<string> IndexCompressionFormats;
        public readonly string Key;
        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedCargoRepositoryMemberResult> Members;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        /// <summary>
        /// Proxy key from Artifactory Proxies settings.
        /// </summary>
        public readonly string? Proxy;
        public readonly string? RepoLayoutRef;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetFederatedCargoRepositoryResult(
            bool? anonymousAccess,

            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool? cdnRedirect,

            bool? cleanupOnDelete,

            string? description,

            bool? disableProxy,

            bool? downloadDirect,

            bool? enableSparseIndex,

            string? excludesPattern,

            string id,

            string? includesPattern,

            ImmutableArray<string> indexCompressionFormats,

            string key,

            ImmutableArray<Outputs.GetFederatedCargoRepositoryMemberResult> members,

            string? notes,

            string packageType,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? proxy,

            string? repoLayoutRef,

            bool? xrayIndex)
        {
            AnonymousAccess = anonymousAccess;
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            CdnRedirect = cdnRedirect;
            CleanupOnDelete = cleanupOnDelete;
            Description = description;
            DisableProxy = disableProxy;
            DownloadDirect = downloadDirect;
            EnableSparseIndex = enableSparseIndex;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            IndexCompressionFormats = indexCompressionFormats;
            Key = key;
            Members = members;
            Notes = notes;
            PackageType = packageType;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            Proxy = proxy;
            RepoLayoutRef = repoLayoutRef;
            XrayIndex = xrayIndex;
        }
    }
}
