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
    /// Creates a local Sbt repository.
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
    ///     var terraform_local_test_sbt_repo = new Artifactory.LocalSbtRepository("terraform-local-test-sbt-repo", new()
    ///     {
    ///         Key = "terraform-local-test-sbt-repo",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Local repositories can be imported using their name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/localSbtRepository:LocalSbtRepository terraform-local-test-sbt-repo terraform-local-test-sbt-repo
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/localSbtRepository:LocalSbtRepository")]
    public partial class LocalSbtRepository : global::Pulumi.CustomResource
    {
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
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Output("cdnRedirect")]
        public Output<bool?> CdnRedirect { get; private set; } = null!;

        /// <summary>
        /// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
        /// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
        /// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
        /// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
        /// </summary>
        [Output("checksumPolicyType")]
        public Output<string?> ChecksumPolicyType { get; private set; } = null!;

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
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Output("excludesPattern")]
        public Output<string?> ExcludesPattern { get; private set; } = null!;

        /// <summary>
        /// If set, Artifactory allows you to deploy release artifacts into this repository.
        /// </summary>
        [Output("handleReleases")]
        public Output<bool?> HandleReleases { get; private set; } = null!;

        /// <summary>
        /// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
        /// </summary>
        [Output("handleSnapshots")]
        public Output<bool?> HandleSnapshots { get; private set; } = null!;

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Output("includesPattern")]
        public Output<string?> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
        /// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
        /// </summary>
        [Output("maxUniqueSnapshots")]
        public Output<int?> MaxUniqueSnapshots { get; private set; } = null!;

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
        /// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
        /// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
        /// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
        /// </summary>
        [Output("snapshotVersionBehavior")]
        public Output<string?> SnapshotVersionBehavior { get; private set; } = null!;

        /// <summary>
        /// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
        /// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
        /// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
        /// checkbox.
        /// </summary>
        [Output("suppressPomConsistencyChecks")]
        public Output<bool?> SuppressPomConsistencyChecks { get; private set; } = null!;

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Output("xrayIndex")]
        public Output<bool?> XrayIndex { get; private set; } = null!;


        /// <summary>
        /// Create a LocalSbtRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocalSbtRepository(string name, LocalSbtRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/localSbtRepository:LocalSbtRepository", name, args ?? new LocalSbtRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocalSbtRepository(string name, Input<string> id, LocalSbtRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/localSbtRepository:LocalSbtRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LocalSbtRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocalSbtRepository Get(string name, Input<string> id, LocalSbtRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new LocalSbtRepository(name, id, state, options);
        }
    }

    public sealed class LocalSbtRepositoryArgs : global::Pulumi.ResourceArgs
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
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        /// <summary>
        /// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
        /// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
        /// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
        /// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
        /// </summary>
        [Input("checksumPolicyType")]
        public Input<string>? ChecksumPolicyType { get; set; }

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
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy release artifacts into this repository.
        /// </summary>
        [Input("handleReleases")]
        public Input<bool>? HandleReleases { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
        /// </summary>
        [Input("handleSnapshots")]
        public Input<bool>? HandleSnapshots { get; set; }

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
        /// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
        /// </summary>
        [Input("maxUniqueSnapshots")]
        public Input<int>? MaxUniqueSnapshots { get; set; }

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
        /// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
        /// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
        /// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
        /// </summary>
        [Input("snapshotVersionBehavior")]
        public Input<string>? SnapshotVersionBehavior { get; set; }

        /// <summary>
        /// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
        /// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
        /// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
        /// checkbox.
        /// </summary>
        [Input("suppressPomConsistencyChecks")]
        public Input<bool>? SuppressPomConsistencyChecks { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public LocalSbtRepositoryArgs()
        {
        }
        public static new LocalSbtRepositoryArgs Empty => new LocalSbtRepositoryArgs();
    }

    public sealed class LocalSbtRepositoryState : global::Pulumi.ResourceArgs
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
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        /// <summary>
        /// Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or
        /// conflicts with the locally calculated checksum (bad checksum). Options are: "client-checksums", or
        /// "server-generated-checksums". Default: "client-checksums"\n For more details, please refer to Checksum Policy -
        /// https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy
        /// </summary>
        [Input("checksumPolicyType")]
        public Input<string>? ChecksumPolicyType { get; set; }

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
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy release artifacts into this repository.
        /// </summary>
        [Input("handleReleases")]
        public Input<bool>? HandleReleases { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy snapshot artifacts into this repository.
        /// </summary>
        [Input("handleSnapshots")]
        public Input<bool>? HandleSnapshots { get; set; }

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting,
        /// older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
        /// </summary>
        [Input("maxUniqueSnapshots")]
        public Input<int>? MaxUniqueSnapshots { get; set; }

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
        /// Specifies the naming convention for Maven SNAPSHOT versions. The options are - unique: Version number is based on a
        /// time-stamp (default) non-unique: Version number uses a self-overriding naming pattern of
        /// artifactId-version-SNAPSHOT.type deployer: Respects the settings in the Maven client that is deploying the artifact.
        /// </summary>
        [Input("snapshotVersionBehavior")]
        public Input<string>? SnapshotVersionBehavior { get; set; }

        /// <summary>
        /// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the
        /// groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the
        /// deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks
        /// checkbox.
        /// </summary>
        [Input("suppressPomConsistencyChecks")]
        public Input<bool>? SuppressPomConsistencyChecks { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public LocalSbtRepositoryState()
        {
        }
        public static new LocalSbtRepositoryState Empty => new LocalSbtRepositoryState();
    }
}
