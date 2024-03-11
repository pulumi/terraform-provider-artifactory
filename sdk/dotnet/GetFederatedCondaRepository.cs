// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFederatedCondaRepository
    {
        /// <summary>
        /// Retrieves a federated Conda repository.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var federated_test_conda_repo = Artifactory.GetFederatedCondaRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-conda-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetFederatedCondaRepositoryResult> InvokeAsync(GetFederatedCondaRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFederatedCondaRepositoryResult>("artifactory:index/getFederatedCondaRepository:getFederatedCondaRepository", args ?? new GetFederatedCondaRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated Conda repository.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var federated_test_conda_repo = Artifactory.GetFederatedCondaRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-conda-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetFederatedCondaRepositoryResult> Invoke(GetFederatedCondaRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedCondaRepositoryResult>("artifactory:index/getFederatedCondaRepository:getFederatedCondaRepository", args ?? new GetFederatedCondaRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFederatedCondaRepositoryArgs : global::Pulumi.InvokeArgs
    {
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
        private List<Inputs.GetFederatedCondaRepositoryMemberArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public List<Inputs.GetFederatedCondaRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new List<Inputs.GetFederatedCondaRepositoryMemberArgs>());
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

        public GetFederatedCondaRepositoryArgs()
        {
        }
        public static new GetFederatedCondaRepositoryArgs Empty => new GetFederatedCondaRepositoryArgs();
    }

    public sealed class GetFederatedCondaRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
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
        private InputList<Inputs.GetFederatedCondaRepositoryMemberInputArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.GetFederatedCondaRepositoryMemberInputArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GetFederatedCondaRepositoryMemberInputArgs>());
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

        public GetFederatedCondaRepositoryInvokeArgs()
        {
        }
        public static new GetFederatedCondaRepositoryInvokeArgs Empty => new GetFederatedCondaRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetFederatedCondaRepositoryResult
    {
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
        public readonly string? ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedCondaRepositoryMemberResult> Members;
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
        private GetFederatedCondaRepositoryResult(
            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool? cdnRedirect,

            bool? cleanupOnDelete,

            string? description,

            bool? disableProxy,

            bool? downloadDirect,

            string? excludesPattern,

            string id,

            string? includesPattern,

            string key,

            ImmutableArray<Outputs.GetFederatedCondaRepositoryMemberResult> members,

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
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            CdnRedirect = cdnRedirect;
            CleanupOnDelete = cleanupOnDelete;
            Description = description;
            DisableProxy = disableProxy;
            DownloadDirect = downloadDirect;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
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
