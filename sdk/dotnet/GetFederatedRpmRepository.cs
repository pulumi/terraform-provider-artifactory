// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFederatedRpmRepository
    {
        /// <summary>
        /// Retrieves a federated Rpm repository.
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
        ///     var federated_test_rpm_repo = Artifactory.GetFederatedRpmRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-rpm-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFederatedRpmRepositoryResult> InvokeAsync(GetFederatedRpmRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFederatedRpmRepositoryResult>("artifactory:index/getFederatedRpmRepository:getFederatedRpmRepository", args ?? new GetFederatedRpmRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated Rpm repository.
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
        ///     var federated_test_rpm_repo = Artifactory.GetFederatedRpmRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-rpm-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFederatedRpmRepositoryResult> Invoke(GetFederatedRpmRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedRpmRepositoryResult>("artifactory:index/getFederatedRpmRepository:getFederatedRpmRepository", args ?? new GetFederatedRpmRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFederatedRpmRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("calculateYumMetadata")]
        public bool? CalculateYumMetadata { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("enableFileListsIndexing")]
        public bool? EnableFileListsIndexing { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("members")]
        private List<Inputs.GetFederatedRpmRepositoryMemberArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public List<Inputs.GetFederatedRpmRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new List<Inputs.GetFederatedRpmRepositoryMemberArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public string? Notes { get; set; }

        [Input("primaryKeypairRef")]
        public string? PrimaryKeypairRef { get; set; }

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

        [Input("secondaryKeypairRef")]
        public string? SecondaryKeypairRef { get; set; }

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        [Input("yumGroupFileNames")]
        public string? YumGroupFileNames { get; set; }

        [Input("yumRootDepth")]
        public int? YumRootDepth { get; set; }

        public GetFederatedRpmRepositoryArgs()
        {
        }
        public static new GetFederatedRpmRepositoryArgs Empty => new GetFederatedRpmRepositoryArgs();
    }

    public sealed class GetFederatedRpmRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("calculateYumMetadata")]
        public Input<bool>? CalculateYumMetadata { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("enableFileListsIndexing")]
        public Input<bool>? EnableFileListsIndexing { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("members")]
        private InputList<Inputs.GetFederatedRpmRepositoryMemberInputArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.GetFederatedRpmRepositoryMemberInputArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GetFederatedRpmRepositoryMemberInputArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("primaryKeypairRef")]
        public Input<string>? PrimaryKeypairRef { get; set; }

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

        [Input("secondaryKeypairRef")]
        public Input<string>? SecondaryKeypairRef { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        [Input("yumGroupFileNames")]
        public Input<string>? YumGroupFileNames { get; set; }

        [Input("yumRootDepth")]
        public Input<int>? YumRootDepth { get; set; }

        public GetFederatedRpmRepositoryInvokeArgs()
        {
        }
        public static new GetFederatedRpmRepositoryInvokeArgs Empty => new GetFederatedRpmRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetFederatedRpmRepositoryResult
    {
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly bool? BlackedOut;
        public readonly bool? CalculateYumMetadata;
        public readonly bool? CdnRedirect;
        public readonly string? Description;
        public readonly bool? DownloadDirect;
        public readonly bool? EnableFileListsIndexing;
        public readonly string ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IncludesPattern;
        public readonly string Key;
        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedRpmRepositoryMemberResult> Members;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly string? PrimaryKeypairRef;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        public readonly string? RepoLayoutRef;
        public readonly string? SecondaryKeypairRef;
        public readonly bool? XrayIndex;
        public readonly string? YumGroupFileNames;
        public readonly int? YumRootDepth;

        [OutputConstructor]
        private GetFederatedRpmRepositoryResult(
            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool? calculateYumMetadata,

            bool? cdnRedirect,

            string? description,

            bool? downloadDirect,

            bool? enableFileListsIndexing,

            string excludesPattern,

            string id,

            string includesPattern,

            string key,

            ImmutableArray<Outputs.GetFederatedRpmRepositoryMemberResult> members,

            string? notes,

            string packageType,

            string? primaryKeypairRef,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? repoLayoutRef,

            string? secondaryKeypairRef,

            bool? xrayIndex,

            string? yumGroupFileNames,

            int? yumRootDepth)
        {
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            CalculateYumMetadata = calculateYumMetadata;
            CdnRedirect = cdnRedirect;
            Description = description;
            DownloadDirect = downloadDirect;
            EnableFileListsIndexing = enableFileListsIndexing;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            Members = members;
            Notes = notes;
            PackageType = packageType;
            PrimaryKeypairRef = primaryKeypairRef;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            RepoLayoutRef = repoLayoutRef;
            SecondaryKeypairRef = secondaryKeypairRef;
            XrayIndex = xrayIndex;
            YumGroupFileNames = yumGroupFileNames;
            YumRootDepth = yumRootDepth;
        }
    }
}
