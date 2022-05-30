// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PushReplicationReplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, enables distributed checksum storage. For more information, see
        /// [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Input("checkBinaryExistenceInFilestore")]
        public Input<bool>? CheckBinaryExistenceInFilestore { get; set; }

        /// <summary>
        /// When set, this replication will be enabled when saved.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
        /// </summary>
        [Input("pathPrefix")]
        public Input<string>? PathPrefix { get; set; }

        /// <summary>
        /// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// The network timeout in milliseconds to use for remote operations.
        /// </summary>
        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). 
        /// Note that enabling this option, will delete artifacts on the target that do not exist in the source repository.
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
        /// The URL of the target local repository on a remote Artifactory server. Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Required for local repository, but not needed for remote repository.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public PushReplicationReplicationArgs()
        {
        }
    }
}
