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
    /// Provides an Artifactory pull replication resource. This can be used to create and manage pull replication in Artifactory
    /// for a local or remote repo. Pull replication provides a convenient way to proactively populate a remote cache, and is very useful
    /// when waiting for new artifacts to arrive on demand (when first requested) is not desirable due to network latency.
    /// See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PullReplication).
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
    ///     // Create a replication between two artifactory local repositories
    ///     var providerTestSource = new Artifactory.LocalMavenRepository("providerTestSource", new()
    ///     {
    ///         Key = "provider_test_source",
    ///     });
    /// 
    ///     var providerTestDest = new Artifactory.RemoteMavenRepository("providerTestDest", new()
    ///     {
    ///         Key = "provider_test_dest",
    ///         Password = "bar",
    ///         Url = $"https://example.com/artifactory/{artifactory_local_maven_repository.Artifactory_local_maven_repository.Key}",
    ///         Username = "foo",
    ///     });
    /// 
    ///     var remote_rep = new Artifactory.PullReplication("remote-rep", new()
    ///     {
    ///         CronExp = "0 0 * * * ?",
    ///         EnableEventReplication = true,
    ///         RepoKey = providerTestDest.Key,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Pull replication config can be imported using its repo key, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/pullReplication:PullReplication foo-rep repository-key
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/pullReplication:PullReplication")]
    public partial class PullReplication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When true, enables distributed checksum storage. For more information, see
        /// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Output("checkBinaryExistenceInFilestore")]
        public Output<bool?> CheckBinaryExistenceInFilestore { get; private set; } = null!;

        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Output("cronExp")]
        public Output<string?> CronExp { get; private set; } = null!;

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
        /// </summary>
        [Output("enableEventReplication")]
        public Output<bool?> EnableEventReplication { get; private set; } = null!;

        /// <summary>
        /// When set, this replication will be enabled when saved.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
        /// </summary>
        [Output("pathPrefix")]
        public Output<string?> PathPrefix { get; private set; } = null!;

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Output("proxy")]
        public Output<string?> Proxy { get; private set; } = null!;

        /// <summary>
        /// Repository name.
        /// </summary>
        [Output("repoKey")]
        public Output<string> RepoKey { get; private set; } = null!;

        [Output("socketTimeoutMillis")]
        public Output<int> SocketTimeoutMillis { get; private set; } = null!;

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
        /// </summary>
        [Output("syncDeletes")]
        public Output<bool> SyncDeletes { get; private set; } = null!;

        /// <summary>
        /// When set, the task also synchronizes the properties of replicated artifacts.
        /// </summary>
        [Output("syncProperties")]
        public Output<bool> SyncProperties { get; private set; } = null!;

        /// <summary>
        /// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
        /// </summary>
        [Output("syncStatistics")]
        public Output<bool> SyncStatistics { get; private set; } = null!;

        /// <summary>
        /// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/&lt;pkg&gt;. 
        /// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a PullReplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PullReplication(string name, PullReplicationArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/pullReplication:PullReplication", name, args ?? new PullReplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PullReplication(string name, Input<string> id, PullReplicationState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/pullReplication:PullReplication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PullReplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PullReplication Get(string name, Input<string> id, PullReplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new PullReplication(name, id, state, options);
        }
    }

    public sealed class PullReplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, enables distributed checksum storage. For more information, see
        /// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Input("checkBinaryExistenceInFilestore")]
        public Input<bool>? CheckBinaryExistenceInFilestore { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
        /// </summary>
        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        /// <summary>
        /// When set, this replication will be enabled when saved.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
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
        /// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
        /// </summary>
        [Input("pathPrefix")]
        public Input<string>? PathPrefix { get; set; }

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Repository name.
        /// </summary>
        [Input("repoKey", required: true)]
        public Input<string> RepoKey { get; set; } = null!;

        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
        /// </summary>
        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        /// <summary>
        /// When set, the task also synchronizes the properties of replicated artifacts.
        /// </summary>
        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        /// <summary>
        /// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
        /// </summary>
        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        /// <summary>
        /// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/&lt;pkg&gt;. 
        /// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public PullReplicationArgs()
        {
        }
        public static new PullReplicationArgs Empty => new PullReplicationArgs();
    }

    public sealed class PullReplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, enables distributed checksum storage. For more information, see
        /// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Input("checkBinaryExistenceInFilestore")]
        public Input<bool>? CheckBinaryExistenceInFilestore { get; set; }

        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
        /// </summary>
        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        /// <summary>
        /// When set, this replication will be enabled when saved.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
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
        /// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
        /// </summary>
        [Input("pathPrefix")]
        public Input<string>? PathPrefix { get; set; }

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Repository name.
        /// </summary>
        [Input("repoKey")]
        public Input<string>? RepoKey { get; set; }

        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
        /// </summary>
        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        /// <summary>
        /// When set, the task also synchronizes the properties of replicated artifacts.
        /// </summary>
        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        /// <summary>
        /// When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
        /// </summary>
        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        /// <summary>
        /// The URL of the target local repository on a remote Artifactory server. For some package types, you need to prefix the repository key in the URL with api/&lt;pkg&gt;. 
        /// For a list of package types where this is required, see the [note](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-anchorPREFIX).
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public PullReplicationState()
        {
        }
        public static new PullReplicationState Empty => new PullReplicationState();
    }
}
