// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetVirtualMavenRepository
    {
        /// <summary>
        /// Retrieves a virtual Maven repository.
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
        ///     var virtual_maven = Artifactory.GetVirtualMavenRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-maven",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVirtualMavenRepositoryResult> InvokeAsync(GetVirtualMavenRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMavenRepositoryResult>("artifactory:index/getVirtualMavenRepository:getVirtualMavenRepository", args ?? new GetVirtualMavenRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Maven repository.
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
        ///     var virtual_maven = Artifactory.GetVirtualMavenRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-maven",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualMavenRepositoryResult> Invoke(GetVirtualMavenRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMavenRepositoryResult>("artifactory:index/getVirtualMavenRepository:getVirtualMavenRepository", args ?? new GetVirtualMavenRepositoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Maven repository.
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
        ///     var virtual_maven = Artifactory.GetVirtualMavenRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-maven",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualMavenRepositoryResult> Invoke(GetVirtualMavenRepositoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMavenRepositoryResult>("artifactory:index/getVirtualMavenRepository:getVirtualMavenRepository", args ?? new GetVirtualMavenRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMavenRepositoryArgs : global::Pulumi.InvokeArgs
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
        /// (Optional) Forces authentication when fetching from remote repos.
        /// </summary>
        [Input("forceMavenAuthentication")]
        public bool? ForceMavenAuthentication { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("keyPair")]
        public string? KeyPair { get; set; }

        [Input("notes")]
        public string? Notes { get; set; }

        /// <summary>
        /// (Optional) One of: `"discard_active_reference", "discard_any_reference", "nothing"`
        /// </summary>
        [Input("pomRepositoryReferencesCleanupPolicy")]
        public string? PomRepositoryReferencesCleanupPolicy { get; set; }

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

        public GetVirtualMavenRepositoryArgs()
        {
        }
        public static new GetVirtualMavenRepositoryArgs Empty => new GetVirtualMavenRepositoryArgs();
    }

    public sealed class GetVirtualMavenRepositoryInvokeArgs : global::Pulumi.InvokeArgs
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
        /// (Optional) Forces authentication when fetching from remote repos.
        /// </summary>
        [Input("forceMavenAuthentication")]
        public Input<bool>? ForceMavenAuthentication { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// (Optional) One of: `"discard_active_reference", "discard_any_reference", "nothing"`
        /// </summary>
        [Input("pomRepositoryReferencesCleanupPolicy")]
        public Input<string>? PomRepositoryReferencesCleanupPolicy { get; set; }

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

        public GetVirtualMavenRepositoryInvokeArgs()
        {
        }
        public static new GetVirtualMavenRepositoryInvokeArgs Empty => new GetVirtualMavenRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualMavenRepositoryResult
    {
        public readonly bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts;
        public readonly string? DefaultDeploymentRepo;
        public readonly string? Description;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// (Optional) Forces authentication when fetching from remote repos.
        /// </summary>
        public readonly bool ForceMavenAuthentication;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly string? KeyPair;
        public readonly string? Notes;
        public readonly string PackageType;
        /// <summary>
        /// (Optional) One of: `"discard_active_reference", "discard_any_reference", "nothing"`
        /// </summary>
        public readonly string PomRepositoryReferencesCleanupPolicy;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly string? RepoLayoutRef;
        public readonly ImmutableArray<string> Repositories;

        [OutputConstructor]
        private GetVirtualMavenRepositoryResult(
            bool? artifactoryRequestsCanRetrieveRemoteArtifacts,

            string? defaultDeploymentRepo,

            string? description,

            string? excludesPattern,

            bool forceMavenAuthentication,

            string id,

            string? includesPattern,

            string key,

            string? keyPair,

            string? notes,

            string packageType,

            string pomRepositoryReferencesCleanupPolicy,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            string? repoLayoutRef,

            ImmutableArray<string> repositories)
        {
            ArtifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            DefaultDeploymentRepo = defaultDeploymentRepo;
            Description = description;
            ExcludesPattern = excludesPattern;
            ForceMavenAuthentication = forceMavenAuthentication;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            KeyPair = keyPair;
            Notes = notes;
            PackageType = packageType;
            PomRepositoryReferencesCleanupPolicy = pomRepositoryReferencesCleanupPolicy;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            RepoLayoutRef = repoLayoutRef;
            Repositories = repositories;
        }
    }
}
