// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Creates a virtual Alpine repository.
    /// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-VirtualRepositorySettingupaVirtualRepository).
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
    ///     var foo_alpine = new Artifactory.VirtualAlpineRepository("foo-alpine", new()
    ///     {
    ///         Key = "foo-alpine",
    ///         Repositories = new[] {},
    ///         Description = "A test virtual repo",
    ///         Notes = "Internal description",
    ///         IncludesPattern = "com/jfrog/**,cloud/jfrog/**",
    ///         ExcludesPattern = "com/google/**",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Virtual repositories can be imported using their name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/virtualAlpineRepository:VirtualAlpineRepository foo-alpine foo-alpine
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/virtualAlpineRepository:VirtualAlpineRepository")]
    public partial class VirtualAlpineRepository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        /// another Artifactory instance.
        /// </summary>
        [Output("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Output<bool?> ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; private set; } = null!;

        /// <summary>
        /// Default repository to deploy artifacts.
        /// </summary>
        [Output("defaultDeploymentRepo")]
        public Output<string?> DefaultDeploymentRepo { get; private set; } = null!;

        /// <summary>
        /// Public description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Output("excludesPattern")]
        public Output<string?> ExcludesPattern { get; private set; } = null!;

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Output("includesPattern")]
        public Output<string?> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        /// contain spaces or special characters.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Internal description.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("packageType")]
        public Output<string> PackageType { get; private set; } = null!;

        /// <summary>
        /// Primary keypair used to sign artifacts. Default value is empty.
        /// </summary>
        [Output("primaryKeypairRef")]
        public Output<string?> PrimaryKeypairRef { get; private set; } = null!;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
        /// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
        /// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
        /// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
        /// </summary>
        [Output("projectEnvironments")]
        public Output<ImmutableArray<string>> ProjectEnvironments { get; private set; } = null!;

        /// <summary>
        /// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// Repository layout key for the virtual repository
        /// </summary>
        [Output("repoLayoutRef")]
        public Output<string?> RepoLayoutRef { get; private set; } = null!;

        /// <summary>
        /// The effective list of actual repositories included in this virtual repository.
        /// </summary>
        [Output("repositories")]
        public Output<ImmutableArray<string>> Repositories { get; private set; } = null!;

        /// <summary>
        /// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        /// </summary>
        [Output("retrievalCachePeriodSeconds")]
        public Output<int?> RetrievalCachePeriodSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualAlpineRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualAlpineRepository(string name, VirtualAlpineRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/virtualAlpineRepository:VirtualAlpineRepository", name, args ?? new VirtualAlpineRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualAlpineRepository(string name, Input<string> id, VirtualAlpineRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/virtualAlpineRepository:VirtualAlpineRepository", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualAlpineRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualAlpineRepository Get(string name, Input<string> id, VirtualAlpineRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualAlpineRepository(name, id, state, options);
        }
    }

    public sealed class VirtualAlpineRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        /// another Artifactory instance.
        /// </summary>
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        /// <summary>
        /// Default repository to deploy artifacts.
        /// </summary>
        [Input("defaultDeploymentRepo")]
        public Input<string>? DefaultDeploymentRepo { get; set; }

        /// <summary>
        /// Public description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        /// contain spaces or special characters.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// Primary keypair used to sign artifacts. Default value is empty.
        /// </summary>
        [Input("primaryKeypairRef")]
        public Input<string>? PrimaryKeypairRef { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
        /// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
        /// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
        /// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Repository layout key for the virtual repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private InputList<string>? _repositories;

        /// <summary>
        /// The effective list of actual repositories included in this virtual repository.
        /// </summary>
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        /// <summary>
        /// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        public VirtualAlpineRepositoryArgs()
        {
        }
        public static new VirtualAlpineRepositoryArgs Empty => new VirtualAlpineRepositoryArgs();
    }

    public sealed class VirtualAlpineRepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        /// another Artifactory instance.
        /// </summary>
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        /// <summary>
        /// Default repository to deploy artifacts.
        /// </summary>
        [Input("defaultDeploymentRepo")]
        public Input<string>? DefaultDeploymentRepo { get; set; }

        /// <summary>
        /// Public description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
        /// contain spaces or special characters.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// Primary keypair used to sign artifacts. Default value is empty.
        /// </summary>
        [Input("primaryKeypairRef")]
        public Input<string>? PrimaryKeypairRef { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
        /// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
        /// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
        /// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Repository layout key for the virtual repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private InputList<string>? _repositories;

        /// <summary>
        /// The effective list of actual repositories included in this virtual repository.
        /// </summary>
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        /// <summary>
        /// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        public VirtualAlpineRepositoryState()
        {
        }
        public static new VirtualAlpineRepositoryState Empty => new VirtualAlpineRepositoryState();
    }
}
