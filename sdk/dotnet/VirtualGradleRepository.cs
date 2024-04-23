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
    /// Creates a virtual Gradle repository.
    /// Official documentation can be found [here](https://jfrog.com/blog/how-to-set-up-a-private-remote-and-virtual-maven-gradle-registry/).
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
    ///     var foo_gradle = new Artifactory.VirtualGradleRepository("foo-gradle", new()
    ///     {
    ///         Key = "foo-gradle",
    ///         Repositories = new[] {},
    ///         Description = "A test virtual repo",
    ///         Notes = "Internal description",
    ///         IncludesPattern = "com/jfrog/**,cloud/jfrog/**",
    ///         ExcludesPattern = "com/google/**",
    ///         PomRepositoryReferencesCleanupPolicy = "discard_active_reference",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual repositories can be imported using their name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/virtualGradleRepository:VirtualGradleRepository foo-gradle foo-gradle
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/virtualGradleRepository:VirtualGradleRepository")]
    public partial class VirtualGradleRepository : global::Pulumi.CustomResource
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
        /// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
        /// is also enforced when aggregated repositories support anonymous requests.
        /// </summary>
        [Output("forceMavenAuthentication")]
        public Output<bool> ForceMavenAuthentication { get; private set; } = null!;

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
        /// The keypair used to sign artifacts.
        /// </summary>
        [Output("keyPair")]
        public Output<string?> KeyPair { get; private set; } = null!;

        /// <summary>
        /// Internal description.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("packageType")]
        public Output<string> PackageType { get; private set; } = null!;

        /// <summary>
        /// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
        /// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
        /// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
        /// </summary>
        [Output("pomRepositoryReferencesCleanupPolicy")]
        public Output<string> PomRepositoryReferencesCleanupPolicy { get; private set; } = null!;

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
        /// Create a VirtualGradleRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualGradleRepository(string name, VirtualGradleRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/virtualGradleRepository:VirtualGradleRepository", name, args ?? new VirtualGradleRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualGradleRepository(string name, Input<string> id, VirtualGradleRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/virtualGradleRepository:VirtualGradleRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualGradleRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualGradleRepository Get(string name, Input<string> id, VirtualGradleRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualGradleRepository(name, id, state, options);
        }
    }

    public sealed class VirtualGradleRepositoryArgs : global::Pulumi.ResourceArgs
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
        /// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
        /// is also enforced when aggregated repositories support anonymous requests.
        /// </summary>
        [Input("forceMavenAuthentication")]
        public Input<bool>? ForceMavenAuthentication { get; set; }

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
        /// The keypair used to sign artifacts.
        /// </summary>
        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
        /// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
        /// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
        /// </summary>
        [Input("pomRepositoryReferencesCleanupPolicy")]
        public Input<string>? PomRepositoryReferencesCleanupPolicy { get; set; }

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

        public VirtualGradleRepositoryArgs()
        {
        }
        public static new VirtualGradleRepositoryArgs Empty => new VirtualGradleRepositoryArgs();
    }

    public sealed class VirtualGradleRepositoryState : global::Pulumi.ResourceArgs
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
        /// User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
        /// is also enforced when aggregated repositories support anonymous requests.
        /// </summary>
        [Input("forceMavenAuthentication")]
        public Input<bool>? ForceMavenAuthentication { get; set; }

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
        /// The keypair used to sign artifacts.
        /// </summary>
        [Input("keyPair")]
        public Input<string>? KeyPair { get; set; }

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
        /// - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
        /// - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
        /// </summary>
        [Input("pomRepositoryReferencesCleanupPolicy")]
        public Input<string>? PomRepositoryReferencesCleanupPolicy { get; set; }

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

        public VirtualGradleRepositoryState()
        {
        }
        public static new VirtualGradleRepositoryState Empty => new VirtualGradleRepositoryState();
    }
}
