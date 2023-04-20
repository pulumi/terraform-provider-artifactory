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
    /// Creates a federated Docker repository.
    /// 
    /// ~&gt;This resource has been superseded by the `artifactory.FederatedDockerV2Repository` resource. This resource will continue to be available in the provider for backward compatibility. For documentation, please refer to the new resource.
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
    ///     var terraform_federated_test_docker_repo = new Artifactory.FederatedDockerRepository("terraform-federated-test-docker-repo", new()
    ///     {
    ///         Key = "terraform-federated-test-docker-repo",
    ///         Members = new[]
    ///         {
    ///             new Artifactory.Inputs.FederatedDockerRepositoryMemberArgs
    ///             {
    ///                 Enabled = true,
    ///                 Url = "http://tempurl.org/artifactory/terraform-federated-test-docker-repo",
    ///             },
    ///             new Artifactory.Inputs.FederatedDockerRepositoryMemberArgs
    ///             {
    ///                 Enabled = true,
    ///                 Url = "http://tempurl2.org/artifactory/terraform-federated-test-docker-repo-2",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Federated repositories can be imported using their name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/federatedDockerRepository:FederatedDockerRepository terraform-federated-test-docker-repo terraform-federated-test-docker-repo
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/federatedDockerRepository:FederatedDockerRepository")]
    public partial class FederatedDockerRepository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Docker API version to use. This cannot be set
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Output("archiveBrowsingEnabled")]
        public Output<bool?> ArchiveBrowsingEnabled { get; private set; } = null!;

        /// <summary>
        /// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
        /// </summary>
        [Output("blackedOut")]
        public Output<bool?> BlackedOut { get; private set; } = null!;

        /// <summary>
        /// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        /// </summary>
        [Output("blockPushingSchema1")]
        public Output<bool> BlockPushingSchema1 { get; private set; } = null!;

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Output("cdnRedirect")]
        public Output<bool?> CdnRedirect { get; private set; } = null!;

        /// <summary>
        /// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
        /// the federation on other Artifactory instances.
        /// </summary>
        [Output("cleanupOnDelete")]
        public Output<bool?> CleanupOnDelete { get; private set; } = null!;

        /// <summary>
        /// Public description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
        /// storage provider. Available in Enterprise+ and Edge licenses only.
        /// </summary>
        [Output("downloadDirect")]
        public Output<bool?> DownloadDirect { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
        /// artifacts are excluded.
        /// </summary>
        [Output("excludesPattern")]
        public Output<string> ExcludesPattern { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Output("includesPattern")]
        public Output<string> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
        /// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
        /// applies to manifest v2
        /// </summary>
        [Output("maxUniqueTags")]
        public Output<int?> MaxUniqueTags { get; private set; } = null!;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.FederatedDockerRepositoryMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Internal description.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("packageType")]
        public Output<string> PackageType { get; private set; } = null!;

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Output("priorityResolution")]
        public Output<bool?> PriorityResolution { get; private set; } = null!;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
        /// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
        /// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
        /// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
        /// </summary>
        [Output("projectEnvironments")]
        public Output<ImmutableArray<string>> ProjectEnvironments { get; private set; } = null!;

        /// <summary>
        /// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// List of property set name
        /// </summary>
        [Output("propertySets")]
        public Output<ImmutableArray<string>> PropertySets { get; private set; } = null!;

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Output("repoLayoutRef")]
        public Output<string?> RepoLayoutRef { get; private set; } = null!;

        /// <summary>
        /// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
        /// manifest V2
        /// </summary>
        [Output("tagRetention")]
        public Output<int?> TagRetention { get; private set; } = null!;

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Output("xrayIndex")]
        public Output<bool?> XrayIndex { get; private set; } = null!;


        /// <summary>
        /// Create a FederatedDockerRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FederatedDockerRepository(string name, FederatedDockerRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/federatedDockerRepository:FederatedDockerRepository", name, args ?? new FederatedDockerRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FederatedDockerRepository(string name, Input<string> id, FederatedDockerRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/federatedDockerRepository:FederatedDockerRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FederatedDockerRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FederatedDockerRepository Get(string name, Input<string> id, FederatedDockerRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new FederatedDockerRepository(name, id, state, options);
        }
    }

    public sealed class FederatedDockerRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        /// <summary>
        /// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
        /// </summary>
        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        /// <summary>
        /// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        /// </summary>
        [Input("blockPushingSchema1")]
        public Input<bool>? BlockPushingSchema1 { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        /// <summary>
        /// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
        /// the federation on other Artifactory instances.
        /// </summary>
        [Input("cleanupOnDelete")]
        public Input<bool>? CleanupOnDelete { get; set; }

        /// <summary>
        /// Public description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
        /// storage provider. Available in Enterprise+ and Edge licenses only.
        /// </summary>
        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
        /// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
        /// applies to manifest v2
        /// </summary>
        [Input("maxUniqueTags")]
        public Input<int>? MaxUniqueTags { get; set; }

        [Input("members", required: true)]
        private InputList<Inputs.FederatedDockerRepositoryMemberArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.FederatedDockerRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.FederatedDockerRepositoryMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

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
        /// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;

        /// <summary>
        /// List of property set name
        /// </summary>
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
        /// manifest V2
        /// </summary>
        [Input("tagRetention")]
        public Input<int>? TagRetention { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public FederatedDockerRepositoryArgs()
        {
        }
        public static new FederatedDockerRepositoryArgs Empty => new FederatedDockerRepositoryArgs();
    }

    public sealed class FederatedDockerRepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Docker API version to use. This cannot be set
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        /// <summary>
        /// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
        /// </summary>
        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        /// <summary>
        /// When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        /// </summary>
        [Input("blockPushingSchema1")]
        public Input<bool>? BlockPushingSchema1 { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        /// <summary>
        /// Delete all federated members on `terraform destroy` if set to `true`. Caution: it will delete all the repositories in
        /// the federation on other Artifactory instances.
        /// </summary>
        [Input("cleanupOnDelete")]
        public Input<bool>? CleanupOnDelete { get; set; }

        /// <summary>
        /// Public description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
        /// storage provider. Available in Enterprise+ and Edge licenses only.
        /// </summary>
        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an
        /// image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only
        /// applies to manifest v2
        /// </summary>
        [Input("maxUniqueTags")]
        public Input<int>? MaxUniqueTags { get; set; }

        [Input("members")]
        private InputList<Inputs.FederatedDockerRepositoryMemberGetArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.FederatedDockerRepositoryMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.FederatedDockerRepositoryMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

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
        /// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;

        /// <summary>
        /// List of property set name
        /// </summary>
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to
        /// manifest V2
        /// </summary>
        [Input("tagRetention")]
        public Input<int>? TagRetention { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public FederatedDockerRepositoryState()
        {
        }
        public static new FederatedDockerRepositoryState Empty => new FederatedDockerRepositoryState();
    }
}
