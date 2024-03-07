// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetLocalDockerV1Repository
    {
        /// <summary>
        /// Retrieves a local Docker (v1) repository resource.
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
        ///     var artifactoryLocalTestDockerV1Repository = new Artifactory.DockerV1Repository("artifactoryLocalTestDockerV1Repository", new()
        ///     {
        ///         Key = "artifactory_local_test_docker_v1_repository",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLocalDockerV1RepositoryResult> InvokeAsync(GetLocalDockerV1RepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocalDockerV1RepositoryResult>("artifactory:index/getLocalDockerV1Repository:getLocalDockerV1Repository", args ?? new GetLocalDockerV1RepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a local Docker (v1) repository resource.
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
        ///     var artifactoryLocalTestDockerV1Repository = new Artifactory.DockerV1Repository("artifactoryLocalTestDockerV1Repository", new()
        ///     {
        ///         Key = "artifactory_local_test_docker_v1_repository",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLocalDockerV1RepositoryResult> Invoke(GetLocalDockerV1RepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalDockerV1RepositoryResult>("artifactory:index/getLocalDockerV1Repository:getLocalDockerV1Repository", args ?? new GetLocalDockerV1RepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocalDockerV1RepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("description")]
        public string? Description { get; set; }

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

        /// <summary>
        /// The maximum number of unique tags of a single Docker image to store in this repository. Once the
        /// number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
        /// limit. This only applies to manifest v2.
        /// </summary>
        [Input("maxUniqueTags")]
        public int? MaxUniqueTags { get; set; }

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

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetLocalDockerV1RepositoryArgs()
        {
        }
        public static new GetLocalDockerV1RepositoryArgs Empty => new GetLocalDockerV1RepositoryArgs();
    }

    public sealed class GetLocalDockerV1RepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

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

        /// <summary>
        /// The maximum number of unique tags of a single Docker image to store in this repository. Once the
        /// number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
        /// limit. This only applies to manifest v2.
        /// </summary>
        [Input("maxUniqueTags")]
        public Input<int>? MaxUniqueTags { get; set; }

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

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetLocalDockerV1RepositoryInvokeArgs()
        {
        }
        public static new GetLocalDockerV1RepositoryInvokeArgs Empty => new GetLocalDockerV1RepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocalDockerV1RepositoryResult
    {
        /// <summary>
        /// The Docker API version in use.
        /// </summary>
        public readonly string ApiVersion;
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly bool? BlackedOut;
        /// <summary>
        /// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to
        /// this repository.
        /// </summary>
        public readonly bool BlockPushingSchema1;
        public readonly bool? CdnRedirect;
        public readonly string? Description;
        public readonly bool? DownloadDirect;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        /// <summary>
        /// The maximum number of unique tags of a single Docker image to store in this repository. Once the
        /// number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
        /// limit. This only applies to manifest v2.
        /// </summary>
        public readonly int MaxUniqueTags;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        public readonly string? RepoLayoutRef;
        /// <summary>
        /// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This
        /// only applies to manifest V2.
        /// </summary>
        public readonly int TagRetention;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetLocalDockerV1RepositoryResult(
            string apiVersion,

            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool blockPushingSchema1,

            bool? cdnRedirect,

            string? description,

            bool? downloadDirect,

            string? excludesPattern,

            string id,

            string? includesPattern,

            string key,

            int maxUniqueTags,

            string? notes,

            string packageType,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? repoLayoutRef,

            int tagRetention,

            bool? xrayIndex)
        {
            ApiVersion = apiVersion;
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            BlockPushingSchema1 = blockPushingSchema1;
            CdnRedirect = cdnRedirect;
            Description = description;
            DownloadDirect = downloadDirect;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            MaxUniqueTags = maxUniqueTags;
            Notes = notes;
            PackageType = packageType;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            RepoLayoutRef = repoLayoutRef;
            TagRetention = tagRetention;
            XrayIndex = xrayIndex;
        }
    }
}
