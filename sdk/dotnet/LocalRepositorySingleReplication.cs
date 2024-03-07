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
    /// Provides a local repository replication resource, also referred to as Artifactory push replication. This can be used to create and manage Artifactory local repository replications using [Push Replication API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-SetRepositoryReplicationConfiguration).
    /// Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near end invoking a synchronization of artifacts to the far end.
    /// See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).
    /// This resource can create the replication of local repository to single repository on the remote server.
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
    ///     var config = new Config();
    ///     var artifactoryUrl = config.Require("artifactoryUrl");
    ///     var artifactoryUsername = config.Require("artifactoryUsername");
    ///     var artifactoryPassword = config.Require("artifactoryPassword");
    ///     // Create a replication between two artifactory local repositories
    ///     var providerTestSource = new Artifactory.LocalMavenRepository("providerTestSource", new()
    ///     {
    ///         Key = "provider_test_source",
    ///     });
    /// 
    ///     var providerTestDest = new Artifactory.LocalMavenRepository("providerTestDest", new()
    ///     {
    ///         Key = "provider_test_dest",
    ///     });
    /// 
    ///     var foo_rep = new Artifactory.LocalRepositorySingleReplication("foo-rep", new()
    ///     {
    ///         RepoKey = providerTestSource.Key,
    ///         CronExp = "0 0 * * * ?",
    ///         EnableEventReplication = true,
    ///         Url = providerTestDest.Key.Apply(key =&gt; $"{artifactoryUrl}/artifactory/{key}"),
    ///         Username = "$var.artifactory_username",
    ///         Password = "$var.artifactory_password",
    ///         Enabled = true,
    ///         SocketTimeoutMillis = 16000,
    ///         SyncDeletes = false,
    ///         SyncProperties = true,
    ///         SyncStatistics = true,
    ///         IncludePathPrefixPattern = "/some-repo/",
    ///         ExcludePathPrefixPattern = "/some-other-repo/",
    ///         CheckBinaryExistenceInFilestore = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Push replication configs can be imported using their repo key, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication foo-rep provider_test_source
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication")]
    public partial class LocalRepositorySingleReplication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Output("checkBinaryExistenceInFilestore")]
        public Output<bool?> CheckBinaryExistenceInFilestore { get; private set; } = null!;

        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Output("cronExp")]
        public Output<string?> CronExp { get; private set; } = null!;

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        /// </summary>
        [Output("enableEventReplication")]
        public Output<bool?> EnableEventReplication { get; private set; } = null!;

        /// <summary>
        /// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
        /// </summary>
        [Output("excludePathPrefixPattern")]
        public Output<string?> ExcludePathPrefixPattern { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
        /// </summary>
        [Output("includePathPrefixPattern")]
        public Output<string?> IncludePathPrefixPattern { get; private set; } = null!;

        /// <summary>
        /// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
        /// </summary>
        [Output("proxy")]
        public Output<string?> Proxy { get; private set; } = null!;

        /// <summary>
        /// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
        /// </summary>
        [Output("replicationKey")]
        public Output<string> ReplicationKey { get; private set; } = null!;

        /// <summary>
        /// Repository name.
        /// </summary>
        [Output("repoKey")]
        public Output<string> RepoKey { get; private set; } = null!;

        /// <summary>
        /// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
        /// </summary>
        [Output("socketTimeoutMillis")]
        public Output<int?> SocketTimeoutMillis { get; private set; } = null!;

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
        /// </summary>
        [Output("syncDeletes")]
        public Output<bool?> SyncDeletes { get; private set; } = null!;

        /// <summary>
        /// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
        /// </summary>
        [Output("syncProperties")]
        public Output<bool?> SyncProperties { get; private set; } = null!;

        /// <summary>
        /// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
        /// </summary>
        [Output("syncStatistics")]
        public Output<bool?> SyncStatistics { get; private set; } = null!;

        /// <summary>
        /// The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Username on the remote Artifactory instance.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a LocalRepositorySingleReplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocalRepositorySingleReplication(string name, LocalRepositorySingleReplicationArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, args ?? new LocalRepositorySingleReplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocalRepositorySingleReplication(string name, Input<string> id, LocalRepositorySingleReplicationState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LocalRepositorySingleReplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocalRepositorySingleReplication Get(string name, Input<string> id, LocalRepositorySingleReplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new LocalRepositorySingleReplication(name, id, state, options);
        }
    }

    public sealed class LocalRepositorySingleReplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Input("checkBinaryExistenceInFilestore")]
        public Input<bool>? CheckBinaryExistenceInFilestore { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        /// </summary>
        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        /// <summary>
        /// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
        /// </summary>
        [Input("excludePathPrefixPattern")]
        public Input<string>? ExcludePathPrefixPattern { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
        /// </summary>
        [Input("includePathPrefixPattern")]
        public Input<string>? IncludePathPrefixPattern { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
        /// </summary>
        [Input("replicationKey")]
        public Input<string>? ReplicationKey { get; set; }

        /// <summary>
        /// Repository name.
        /// </summary>
        [Input("repoKey", required: true)]
        public Input<string> RepoKey { get; set; } = null!;

        /// <summary>
        /// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
        /// </summary>
        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
        /// </summary>
        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        /// <summary>
        /// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
        /// </summary>
        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        /// <summary>
        /// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
        /// </summary>
        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        /// <summary>
        /// The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Username on the remote Artifactory instance.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public LocalRepositorySingleReplicationArgs()
        {
        }
        public static new LocalRepositorySingleReplicationArgs Empty => new LocalRepositorySingleReplicationArgs();
    }

    public sealed class LocalRepositorySingleReplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Input("checkBinaryExistenceInFilestore")]
        public Input<bool>? CheckBinaryExistenceInFilestore { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
        /// </summary>
        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        /// <summary>
        /// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
        /// </summary>
        [Input("excludePathPrefixPattern")]
        public Input<string>? ExcludePathPrefixPattern { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
        /// </summary>
        [Input("includePathPrefixPattern")]
        public Input<string>? IncludePathPrefixPattern { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
        /// </summary>
        [Input("replicationKey")]
        public Input<string>? ReplicationKey { get; set; }

        /// <summary>
        /// Repository name.
        /// </summary>
        [Input("repoKey")]
        public Input<string>? RepoKey { get; set; }

        /// <summary>
        /// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
        /// </summary>
        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
        /// </summary>
        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        /// <summary>
        /// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
        /// </summary>
        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        /// <summary>
        /// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
        /// </summary>
        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        /// <summary>
        /// The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Username on the remote Artifactory instance.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public LocalRepositorySingleReplicationState()
        {
        }
        public static new LocalRepositorySingleReplicationState Empty => new LocalRepositorySingleReplicationState();
    }
}
