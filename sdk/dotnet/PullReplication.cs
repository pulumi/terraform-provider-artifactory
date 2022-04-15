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
    /// ## # Artifactory Pull Replication Resource
    /// 
    /// Provides an Artifactory pull replication resource. This can be used to create and manage pull replication in Artifactory
    /// for a local or remote repo.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a replication between two artifactory local repositories
    ///         var providerTestSource = new Artifactory.LocalMavenRepository("providerTestSource", new Artifactory.LocalMavenRepositoryArgs
    ///         {
    ///             Key = "provider_test_source",
    ///         });
    ///         var providerTestDest = new Artifactory.RemoteMavenRepository("providerTestDest", new Artifactory.RemoteMavenRepositoryArgs
    ///         {
    ///             Key = "provider_test_dest",
    ///             Password = "bar",
    ///             Url = $"https://example.com/artifactory/{artifactory_local_maven_repository.Artifactory_local_maven_repository.Key}",
    ///             Username = "foo",
    ///         });
    ///         var remote_rep = new Artifactory.PullReplication("remote-rep", new Artifactory.PullReplicationArgs
    ///         {
    ///             CronExp = "0 0 * * * ?",
    ///             EnableEventReplication = true,
    ///             RepoKey = providerTestDest.Key,
    ///         });
    ///     }
    /// 
    /// }
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
    public partial class PullReplication : Pulumi.CustomResource
    {
        [Output("cronExp")]
        public Output<string> CronExp { get; private set; } = null!;

        [Output("enableEventReplication")]
        public Output<bool> EnableEventReplication { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("pathPrefix")]
        public Output<string?> PathPrefix { get; private set; } = null!;

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Output("proxy")]
        public Output<string?> Proxy { get; private set; } = null!;

        [Output("repoKey")]
        public Output<string> RepoKey { get; private set; } = null!;

        [Output("socketTimeoutMillis")]
        public Output<int> SocketTimeoutMillis { get; private set; } = null!;

        [Output("syncDeletes")]
        public Output<bool> SyncDeletes { get; private set; } = null!;

        [Output("syncProperties")]
        public Output<bool> SyncProperties { get; private set; } = null!;

        [Output("syncStatistics")]
        public Output<bool> SyncStatistics { get; private set; } = null!;

        /// <summary>
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

    public sealed class PullReplicationArgs : Pulumi.ResourceArgs
    {
        [Input("cronExp", required: true)]
        public Input<string> CronExp { get; set; } = null!;

        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("pathPrefix")]
        public Input<string>? PathPrefix { get; set; }

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("repoKey", required: true)]
        public Input<string> RepoKey { get; set; } = null!;

        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        /// <summary>
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
    }

    public sealed class PullReplicationState : Pulumi.ResourceArgs
    {
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("pathPrefix")]
        public Input<string>? PathPrefix { get; set; }

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("repoKey")]
        public Input<string>? RepoKey { get; set; }

        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        /// <summary>
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
    }
}
