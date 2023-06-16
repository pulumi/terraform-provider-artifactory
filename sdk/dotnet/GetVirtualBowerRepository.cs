// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetVirtualBowerRepository
    {
        /// <summary>
        /// Retrieves a virtual Bower repository.
        /// </summary>
        public static Task<GetVirtualBowerRepositoryResult> InvokeAsync(GetVirtualBowerRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualBowerRepositoryResult>("artifactory:index/getVirtualBowerRepository:getVirtualBowerRepository", args ?? new GetVirtualBowerRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Bower repository.
        /// </summary>
        public static Output<GetVirtualBowerRepositoryResult> Invoke(GetVirtualBowerRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualBowerRepositoryResult>("artifactory:index/getVirtualBowerRepository:getVirtualBowerRepository", args ?? new GetVirtualBowerRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualBowerRepositoryArgs : global::Pulumi.InvokeArgs
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
        /// (Optional) When set, external dependencies are rewritten. Default value is false.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public bool? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private List<string>? _externalDependenciesPatterns;

        /// <summary>
        /// (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
        /// </summary>
        public List<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new List<string>());
            set => _externalDependenciesPatterns = value;
        }

        /// <summary>
        /// (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
        /// </summary>
        [Input("externalDependenciesRemoteRepo")]
        public string? ExternalDependenciesRemoteRepo { get; set; }

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

        public GetVirtualBowerRepositoryArgs()
        {
        }
        public static new GetVirtualBowerRepositoryArgs Empty => new GetVirtualBowerRepositoryArgs();
    }

    public sealed class GetVirtualBowerRepositoryInvokeArgs : global::Pulumi.InvokeArgs
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
        /// (Optional) When set, external dependencies are rewritten. Default value is false.
        /// </summary>
        [Input("externalDependenciesEnabled")]
        public Input<bool>? ExternalDependenciesEnabled { get; set; }

        [Input("externalDependenciesPatterns")]
        private InputList<string>? _externalDependenciesPatterns;

        /// <summary>
        /// (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
        /// </summary>
        public InputList<string> ExternalDependenciesPatterns
        {
            get => _externalDependenciesPatterns ?? (_externalDependenciesPatterns = new InputList<string>());
            set => _externalDependenciesPatterns = value;
        }

        /// <summary>
        /// (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
        /// </summary>
        [Input("externalDependenciesRemoteRepo")]
        public Input<string>? ExternalDependenciesRemoteRepo { get; set; }

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

        public GetVirtualBowerRepositoryInvokeArgs()
        {
        }
        public static new GetVirtualBowerRepositoryInvokeArgs Empty => new GetVirtualBowerRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualBowerRepositoryResult
    {
        public readonly bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts;
        public readonly string? DefaultDeploymentRepo;
        public readonly string? Description;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// (Optional) When set, external dependencies are rewritten. Default value is false.
        /// </summary>
        public readonly bool? ExternalDependenciesEnabled;
        /// <summary>
        /// (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
        /// </summary>
        public readonly ImmutableArray<string> ExternalDependenciesPatterns;
        /// <summary>
        /// (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
        /// </summary>
        public readonly string? ExternalDependenciesRemoteRepo;
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
        private GetVirtualBowerRepositoryResult(
            bool? artifactoryRequestsCanRetrieveRemoteArtifacts,

            string? defaultDeploymentRepo,

            string? description,

            string? excludesPattern,

            bool? externalDependenciesEnabled,

            ImmutableArray<string> externalDependenciesPatterns,

            string? externalDependenciesRemoteRepo,

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
            ExternalDependenciesRemoteRepo = externalDependenciesRemoteRepo;
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
