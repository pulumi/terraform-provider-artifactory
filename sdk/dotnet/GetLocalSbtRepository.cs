// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetLocalSbtRepository
    {
        /// <summary>
        /// Retrieves a local Sbt repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var local_test_sbt_repo = Artifactory.GetLocalSbtRepository.Invoke(new()
        ///     {
        ///         Key = "local-test-sbt-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLocalSbtRepositoryResult> InvokeAsync(GetLocalSbtRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocalSbtRepositoryResult>("artifactory:index/getLocalSbtRepository:getLocalSbtRepository", args ?? new GetLocalSbtRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a local Sbt repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var local_test_sbt_repo = Artifactory.GetLocalSbtRepository.Invoke(new()
        ///     {
        ///         Key = "local-test-sbt-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLocalSbtRepositoryResult> Invoke(GetLocalSbtRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalSbtRepositoryResult>("artifactory:index/getLocalSbtRepository:getLocalSbtRepository", args ?? new GetLocalSbtRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocalSbtRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        /// <summary>
        /// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
        /// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
        /// `client-checksums` and `generated-checksums`. For more details, please refer
        /// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
        /// .
        /// </summary>
        [Input("checksumPolicyType")]
        public string? ChecksumPolicyType { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
        /// .
        /// </summary>
        [Input("handleReleases")]
        public bool? HandleReleases { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
        /// is `true`.
        /// </summary>
        [Input("handleSnapshots")]
        public bool? HandleSnapshots { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// The maximum number of unique snapshots of a single artifact to store. Once the number of
        /// snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
        /// unique snapshots are not cleaned up.
        /// </summary>
        [Input("maxUniqueSnapshots")]
        public int? MaxUniqueSnapshots { get; set; }

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

        /// <summary>
        /// Specifies the naming convention for Maven SNAPSHOT versions. The options are
        /// -
        /// </summary>
        [Input("snapshotVersionBehavior")]
        public string? SnapshotVersionBehavior { get; set; }

        /// <summary>
        /// By default, Artifactory keeps your repositories healthy by refusing POMs with
        /// incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
        /// path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
        /// Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
        /// </summary>
        [Input("suppressPomConsistencyChecks")]
        public bool? SuppressPomConsistencyChecks { get; set; }

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetLocalSbtRepositoryArgs()
        {
        }
        public static new GetLocalSbtRepositoryArgs Empty => new GetLocalSbtRepositoryArgs();
    }

    public sealed class GetLocalSbtRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        /// <summary>
        /// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
        /// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
        /// `client-checksums` and `generated-checksums`. For more details, please refer
        /// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
        /// .
        /// </summary>
        [Input("checksumPolicyType")]
        public Input<string>? ChecksumPolicyType { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
        /// .
        /// </summary>
        [Input("handleReleases")]
        public Input<bool>? HandleReleases { get; set; }

        /// <summary>
        /// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
        /// is `true`.
        /// </summary>
        [Input("handleSnapshots")]
        public Input<bool>? HandleSnapshots { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The maximum number of unique snapshots of a single artifact to store. Once the number of
        /// snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
        /// unique snapshots are not cleaned up.
        /// </summary>
        [Input("maxUniqueSnapshots")]
        public Input<int>? MaxUniqueSnapshots { get; set; }

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

        /// <summary>
        /// Specifies the naming convention for Maven SNAPSHOT versions. The options are
        /// -
        /// </summary>
        [Input("snapshotVersionBehavior")]
        public Input<string>? SnapshotVersionBehavior { get; set; }

        /// <summary>
        /// By default, Artifactory keeps your repositories healthy by refusing POMs with
        /// incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
        /// path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
        /// Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
        /// </summary>
        [Input("suppressPomConsistencyChecks")]
        public Input<bool>? SuppressPomConsistencyChecks { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetLocalSbtRepositoryInvokeArgs()
        {
        }
        public static new GetLocalSbtRepositoryInvokeArgs Empty => new GetLocalSbtRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocalSbtRepositoryResult
    {
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly bool? BlackedOut;
        public readonly bool? CdnRedirect;
        /// <summary>
        /// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
        /// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
        /// `client-checksums` and `generated-checksums`. For more details, please refer
        /// to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy)
        /// .
        /// </summary>
        public readonly string? ChecksumPolicyType;
        public readonly string? Description;
        public readonly bool? DownloadDirect;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`
        /// .
        /// </summary>
        public readonly bool? HandleReleases;
        /// <summary>
        /// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default
        /// is `true`.
        /// </summary>
        public readonly bool? HandleSnapshots;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        /// <summary>
        /// The maximum number of unique snapshots of a single artifact to store. Once the number of
        /// snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and
        /// unique snapshots are not cleaned up.
        /// </summary>
        public readonly int? MaxUniqueSnapshots;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        public readonly string? RepoLayoutRef;
        /// <summary>
        /// Specifies the naming convention for Maven SNAPSHOT versions. The options are
        /// -
        /// </summary>
        public readonly string? SnapshotVersionBehavior;
        /// <summary>
        /// By default, Artifactory keeps your repositories healthy by refusing POMs with
        /// incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed
        /// path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the
        /// Suppress POM Consistency Checks checkbox. True by default for Gradle repository.
        /// </summary>
        public readonly bool? SuppressPomConsistencyChecks;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetLocalSbtRepositoryResult(
            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool? cdnRedirect,

            string? checksumPolicyType,

            string? description,

            bool? downloadDirect,

            string? excludesPattern,

            bool? handleReleases,

            bool? handleSnapshots,

            string id,

            string? includesPattern,

            string key,

            int? maxUniqueSnapshots,

            string? notes,

            string packageType,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? repoLayoutRef,

            string? snapshotVersionBehavior,

            bool? suppressPomConsistencyChecks,

            bool? xrayIndex)
        {
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            CdnRedirect = cdnRedirect;
            ChecksumPolicyType = checksumPolicyType;
            Description = description;
            DownloadDirect = downloadDirect;
            ExcludesPattern = excludesPattern;
            HandleReleases = handleReleases;
            HandleSnapshots = handleSnapshots;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            MaxUniqueSnapshots = maxUniqueSnapshots;
            Notes = notes;
            PackageType = packageType;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            RepoLayoutRef = repoLayoutRef;
            SnapshotVersionBehavior = snapshotVersionBehavior;
            SuppressPomConsistencyChecks = suppressPomConsistencyChecks;
            XrayIndex = xrayIndex;
        }
    }
}
