// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetVirtualAlpineRepository
    {
        /// <summary>
        /// Retrieves a virtual Alpine repository.
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
        ///     var virtual_alpine = Artifactory.GetVirtualAlpineRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-alpine",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVirtualAlpineRepositoryResult> InvokeAsync(GetVirtualAlpineRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualAlpineRepositoryResult>("artifactory:index/getVirtualAlpineRepository:getVirtualAlpineRepository", args ?? new GetVirtualAlpineRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Alpine repository.
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
        ///     var virtual_alpine = Artifactory.GetVirtualAlpineRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-alpine",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualAlpineRepositoryResult> Invoke(GetVirtualAlpineRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualAlpineRepositoryResult>("artifactory:index/getVirtualAlpineRepository:getVirtualAlpineRepository", args ?? new GetVirtualAlpineRepositoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Alpine repository.
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
        ///     var virtual_alpine = Artifactory.GetVirtualAlpineRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-alpine",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualAlpineRepositoryResult> Invoke(GetVirtualAlpineRepositoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualAlpineRepositoryResult>("artifactory:index/getVirtualAlpineRepository:getVirtualAlpineRepository", args ?? new GetVirtualAlpineRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualAlpineRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        [Input("defaultDeploymentRepo")]
        public string? DefaultDeploymentRepo { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("notes")]
        public string? Notes { get; set; }

        /// <summary>
        /// (Optional) Primary keypair used to sign artifacts. Default value is empty.
        /// </summary>
        [Input("primaryKeypairRef")]
        public string? PrimaryKeypairRef { get; set; }

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

        /// <summary>
        /// (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public int? RetrievalCachePeriodSeconds { get; set; }

        public GetVirtualAlpineRepositoryArgs()
        {
        }
        public static new GetVirtualAlpineRepositoryArgs Empty => new GetVirtualAlpineRepositoryArgs();
    }

    public sealed class GetVirtualAlpineRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        [Input("defaultDeploymentRepo")]
        public Input<string>? DefaultDeploymentRepo { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// (Optional) Primary keypair used to sign artifacts. Default value is empty.
        /// </summary>
        [Input("primaryKeypairRef")]
        public Input<string>? PrimaryKeypairRef { get; set; }

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

        /// <summary>
        /// (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        public GetVirtualAlpineRepositoryInvokeArgs()
        {
        }
        public static new GetVirtualAlpineRepositoryInvokeArgs Empty => new GetVirtualAlpineRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualAlpineRepositoryResult
    {
        public readonly bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts;
        public readonly string? DefaultDeploymentRepo;
        public readonly string? Description;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly string? Notes;
        public readonly string PackageType;
        /// <summary>
        /// (Optional) Primary keypair used to sign artifacts. Default value is empty.
        /// </summary>
        public readonly string? PrimaryKeypairRef;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly string? RepoLayoutRef;
        public readonly ImmutableArray<string> Repositories;
        /// <summary>
        /// (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        /// </summary>
        public readonly int? RetrievalCachePeriodSeconds;

        [OutputConstructor]
        private GetVirtualAlpineRepositoryResult(
            bool? artifactoryRequestsCanRetrieveRemoteArtifacts,

            string? defaultDeploymentRepo,

            string? description,

            string? excludesPattern,

            string id,

            string? includesPattern,

            string key,

            string? notes,

            string packageType,

            string? primaryKeypairRef,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            string? repoLayoutRef,

            ImmutableArray<string> repositories,

            int? retrievalCachePeriodSeconds)
        {
            ArtifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            DefaultDeploymentRepo = defaultDeploymentRepo;
            Description = description;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            Notes = notes;
            PackageType = packageType;
            PrimaryKeypairRef = primaryKeypairRef;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            RepoLayoutRef = repoLayoutRef;
            Repositories = repositories;
            RetrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
        }
    }
}
