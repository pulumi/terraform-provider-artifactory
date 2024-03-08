// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetVirtualGoRepository
    {
        /// <summary>
        /// Retrieves a virtual Go repository.
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
        ///     var virtual_go = Artifactory.GetVirtualGoRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-go",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetVirtualGoRepositoryResult> InvokeAsync(GetVirtualGoRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualGoRepositoryResult>("artifactory:index/getVirtualGoRepository:getVirtualGoRepository", args ?? new GetVirtualGoRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Go repository.
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
        ///     var virtual_go = Artifactory.GetVirtualGoRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-go",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetVirtualGoRepositoryResult> Invoke(GetVirtualGoRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualGoRepositoryResult>("artifactory:index/getVirtualGoRepository:getVirtualGoRepository", args ?? new GetVirtualGoRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualGoRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        [Input("defaultDeploymentRepo")]
        public string? DefaultDeploymentRepo { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list. 
        /// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public bool? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private List<string>? _externalDependenciesPatterns;

        /// <summary>
        /// (Optional) 'go-import' Allow List on the UI.
        /// </summary>
        public List<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new List<string>());
            set => _externalDependenciesPatterns = value;
        }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("notes")]
        public string? Notes { get; set; }

        [Input("projectEnvironments")]
        private List<string>? _projectEnvironments;
        public List<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new List<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectKey")]
        public string? ProjectKey { get; set; }

        [Input("repoLayoutRef")]
        public string? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private List<string>? _repositories;
        public List<string> Repositories
        {
            get => _repositories ?? (_repositories = new List<string>());
            set => _repositories = value;
        }

        public GetVirtualGoRepositoryArgs()
        {
        }
        public static new GetVirtualGoRepositoryArgs Empty => new GetVirtualGoRepositoryArgs();
    }

    public sealed class GetVirtualGoRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        [Input("defaultDeploymentRepo")]
        public Input<string>? DefaultDeploymentRepo { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// (Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list. 
        /// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public Input<bool>? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private InputList<string>? _externalDependenciesPatterns;

        /// <summary>
        /// (Optional) 'go-import' Allow List on the UI.
        /// </summary>
        public InputList<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new InputList<string>());
            set => _externalDependenciesPatterns = value;
        }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private InputList<string>? _repositories;
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        public GetVirtualGoRepositoryInvokeArgs()
        {
        }
        public static new GetVirtualGoRepositoryInvokeArgs Empty => new GetVirtualGoRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualGoRepositoryResult
    {
        public readonly bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts;
        public readonly string? DefaultDeploymentRepo;
        public readonly string? Description;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// (Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list. 
        /// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
        /// </summary>
        public readonly bool? ExternalDependenciesEnabled;
        /// <summary>
        /// (Optional) 'go-import' Allow List on the UI.
        /// </summary>
        public readonly ImmutableArray<string> ExternalDependenciesPatterns;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly string? RepoLayoutRef;
        public readonly ImmutableArray<string> Repositories;

        [OutputConstructor]
        private GetVirtualGoRepositoryResult(
            bool? artifactoryRequestsCanRetrieveRemoteArtifacts,

            string? defaultDeploymentRepo,

            string? description,

            string? excludesPattern,

            bool? externalDependenciesEnabled,

            ImmutableArray<string> externalDependenciesPatterns,

            string id,

            string? includesPattern,

            string key,

            string? notes,

            string packageType,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            string? repoLayoutRef,

            ImmutableArray<string> repositories)
        {
            ArtifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            DefaultDeploymentRepo = defaultDeploymentRepo;
            Description = description;
            ExcludesPattern = excludesPattern;
            ExternalDependenciesEnabled = externalDependenciesEnabled;
            ExternalDependenciesPatterns = externalDependenciesPatterns;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            Notes = notes;
            PackageType = packageType;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            RepoLayoutRef = repoLayoutRef;
            Repositories = repositories;
        }
    }
}
