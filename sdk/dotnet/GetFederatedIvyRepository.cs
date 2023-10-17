// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFederatedIvyRepository
    {
        /// <summary>
        /// Retrieves a federated Ivy repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var federated_test_ivy_repo = Artifactory.GetFederatedIvyRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-ivy-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFederatedIvyRepositoryResult> InvokeAsync(GetFederatedIvyRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFederatedIvyRepositoryResult>("artifactory:index/getFederatedIvyRepository:getFederatedIvyRepository", args ?? new GetFederatedIvyRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated Ivy repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var federated_test_ivy_repo = Artifactory.GetFederatedIvyRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-ivy-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFederatedIvyRepositoryResult> Invoke(GetFederatedIvyRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedIvyRepositoryResult>("artifactory:index/getFederatedIvyRepository:getFederatedIvyRepository", args ?? new GetFederatedIvyRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFederatedIvyRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("checksumPolicyType")]
        public string? ChecksumPolicyType { get; set; }

        [Input("cleanupOnDelete")]
        public bool? CleanupOnDelete { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        [Input("handleReleases")]
        public bool? HandleReleases { get; set; }

        [Input("handleSnapshots")]
        public bool? HandleSnapshots { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("maxUniqueSnapshots")]
        public int? MaxUniqueSnapshots { get; set; }

        [Input("members")]
        private List<Inputs.GetFederatedIvyRepositoryMemberArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public List<Inputs.GetFederatedIvyRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new List<Inputs.GetFederatedIvyRepositoryMemberArgs>());
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

        [Input("repoLayoutRef")]
        public string? RepoLayoutRef { get; set; }

        [Input("snapshotVersionBehavior")]
        public string? SnapshotVersionBehavior { get; set; }

        [Input("suppressPomConsistencyChecks")]
        public bool? SuppressPomConsistencyChecks { get; set; }

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetFederatedIvyRepositoryArgs()
        {
        }
        public static new GetFederatedIvyRepositoryArgs Empty => new GetFederatedIvyRepositoryArgs();
    }

    public sealed class GetFederatedIvyRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("checksumPolicyType")]
        public Input<string>? ChecksumPolicyType { get; set; }

        [Input("cleanupOnDelete")]
        public Input<bool>? CleanupOnDelete { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("handleReleases")]
        public Input<bool>? HandleReleases { get; set; }

        [Input("handleSnapshots")]
        public Input<bool>? HandleSnapshots { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("maxUniqueSnapshots")]
        public Input<int>? MaxUniqueSnapshots { get; set; }

        [Input("members")]
        private InputList<Inputs.GetFederatedIvyRepositoryMemberInputArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.GetFederatedIvyRepositoryMemberInputArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GetFederatedIvyRepositoryMemberInputArgs>());
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

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("snapshotVersionBehavior")]
        public Input<string>? SnapshotVersionBehavior { get; set; }

        [Input("suppressPomConsistencyChecks")]
        public Input<bool>? SuppressPomConsistencyChecks { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetFederatedIvyRepositoryInvokeArgs()
        {
        }
        public static new GetFederatedIvyRepositoryInvokeArgs Empty => new GetFederatedIvyRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetFederatedIvyRepositoryResult
    {
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly bool? BlackedOut;
        public readonly bool? CdnRedirect;
        public readonly string? ChecksumPolicyType;
        public readonly bool? CleanupOnDelete;
        public readonly string? Description;
        public readonly bool? DownloadDirect;
        public readonly string? ExcludesPattern;
        public readonly bool? HandleReleases;
        public readonly bool? HandleSnapshots;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly int? MaxUniqueSnapshots;
        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedIvyRepositoryMemberResult> Members;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        public readonly string? RepoLayoutRef;
        public readonly string? SnapshotVersionBehavior;
        public readonly bool? SuppressPomConsistencyChecks;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetFederatedIvyRepositoryResult(
            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool? cdnRedirect,

            string? checksumPolicyType,

            bool? cleanupOnDelete,

            string? description,

            bool? downloadDirect,

            string? excludesPattern,

            bool? handleReleases,

            bool? handleSnapshots,

            string id,

            string? includesPattern,

            string key,

            int? maxUniqueSnapshots,

            ImmutableArray<Outputs.GetFederatedIvyRepositoryMemberResult> members,

            string? notes,

            string packageType,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? repoLayoutRef,

            string? snapshotVersionBehavior,

            bool? suppressPomConsistencyChecks,

            bool? xrayIndex)
        {
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            CdnRedirect = cdnRedirect;
            ChecksumPolicyType = checksumPolicyType;
            CleanupOnDelete = cleanupOnDelete;
            Description = description;
            DownloadDirect = downloadDirect;
            ExcludesPattern = excludesPattern;
            HandleReleases = handleReleases;
            HandleSnapshots = handleSnapshots;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            MaxUniqueSnapshots = maxUniqueSnapshots;
            Members = members;
            Notes = notes;
            PackageType = packageType;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            RepoLayoutRef = repoLayoutRef;
            SnapshotVersionBehavior = snapshotVersionBehavior;
            SuppressPomConsistencyChecks = suppressPomConsistencyChecks;
            XrayIndex = xrayIndex;
        }
    }
}
