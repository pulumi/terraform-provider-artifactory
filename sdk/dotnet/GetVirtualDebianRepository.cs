// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetVirtualDebianRepository
    {
        /// <summary>
        /// Retrieves a virtual Debian repository.
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
        ///     var virtual_debian = Artifactory.GetVirtualDebianRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-debian",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVirtualDebianRepositoryResult> InvokeAsync(GetVirtualDebianRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualDebianRepositoryResult>("artifactory:index/getVirtualDebianRepository:getVirtualDebianRepository", args ?? new GetVirtualDebianRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Debian repository.
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
        ///     var virtual_debian = Artifactory.GetVirtualDebianRepository.Invoke(new()
        ///     {
        ///         Key = "virtual-debian",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualDebianRepositoryResult> Invoke(GetVirtualDebianRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualDebianRepositoryResult>("artifactory:index/getVirtualDebianRepository:getVirtualDebianRepository", args ?? new GetVirtualDebianRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualDebianRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        /// <summary>
        /// (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
        /// </summary>
        [Input("debianDefaultArchitectures")]
        public string? DebianDefaultArchitectures { get; set; }

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

        [Input("optionalIndexCompressionFormats")]
        private List<string>? _optionalIndexCompressionFormats;

        /// <summary>
        /// (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
        /// </summary>
        public List<string> OptionalIndexCompressionFormats
        {
            get => _optionalIndexCompressionFormats ?? (_optionalIndexCompressionFormats = new List<string>());
            set => _optionalIndexCompressionFormats = value;
        }

        /// <summary>
        /// (Optional) Primary keypair used to sign artifacts. Default is empty.
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

        /// <summary>
        /// (Optional) Secondary keypair used to sign artifacts. Default is empty.
        /// </summary>
        [Input("secondaryKeypairRef")]
        public string? SecondaryKeypairRef { get; set; }

        public GetVirtualDebianRepositoryArgs()
        {
        }
        public static new GetVirtualDebianRepositoryArgs Empty => new GetVirtualDebianRepositoryArgs();
    }

    public sealed class GetVirtualDebianRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        /// <summary>
        /// (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
        /// </summary>
        [Input("debianDefaultArchitectures")]
        public Input<string>? DebianDefaultArchitectures { get; set; }

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

        [Input("optionalIndexCompressionFormats")]
        private InputList<string>? _optionalIndexCompressionFormats;

        /// <summary>
        /// (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
        /// </summary>
        public InputList<string> OptionalIndexCompressionFormats
        {
            get => _optionalIndexCompressionFormats ?? (_optionalIndexCompressionFormats = new InputList<string>());
            set => _optionalIndexCompressionFormats = value;
        }

        /// <summary>
        /// (Optional) Primary keypair used to sign artifacts. Default is empty.
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

        /// <summary>
        /// (Optional) Secondary keypair used to sign artifacts. Default is empty.
        /// </summary>
        [Input("secondaryKeypairRef")]
        public Input<string>? SecondaryKeypairRef { get; set; }

        public GetVirtualDebianRepositoryInvokeArgs()
        {
        }
        public static new GetVirtualDebianRepositoryInvokeArgs Empty => new GetVirtualDebianRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualDebianRepositoryResult
    {
        public readonly bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts;
        /// <summary>
        /// (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
        /// </summary>
        public readonly string? DebianDefaultArchitectures;
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
        /// <summary>
        /// (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
        /// </summary>
        public readonly ImmutableArray<string> OptionalIndexCompressionFormats;
        public readonly string PackageType;
        /// <summary>
        /// (Optional) Primary keypair used to sign artifacts. Default is empty.
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
        /// <summary>
        /// (Optional) Secondary keypair used to sign artifacts. Default is empty.
        /// </summary>
        public readonly string? SecondaryKeypairRef;

        [OutputConstructor]
        private GetVirtualDebianRepositoryResult(
            bool? artifactoryRequestsCanRetrieveRemoteArtifacts,

            string? debianDefaultArchitectures,

            string? defaultDeploymentRepo,

            string? description,

            string? excludesPattern,

            string id,

            string? includesPattern,

            string key,

            string? notes,

            ImmutableArray<string> optionalIndexCompressionFormats,

            string packageType,

            string? primaryKeypairRef,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            string? repoLayoutRef,

            ImmutableArray<string> repositories,

            int? retrievalCachePeriodSeconds,

            string? secondaryKeypairRef)
        {
            ArtifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            DebianDefaultArchitectures = debianDefaultArchitectures;
            DefaultDeploymentRepo = defaultDeploymentRepo;
            Description = description;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            Notes = notes;
            OptionalIndexCompressionFormats = optionalIndexCompressionFormats;
            PackageType = packageType;
            PrimaryKeypairRef = primaryKeypairRef;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            RepoLayoutRef = repoLayoutRef;
            Repositories = repositories;
            RetrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            SecondaryKeypairRef = secondaryKeypairRef;
        }
    }
}
