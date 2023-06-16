// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class LocalRepositoryMultiReplicationReplicationArgs : global::Pulumi.ResourceArgs
    {
        [Input("checkBinaryExistenceInFilestore")]
        public Input<bool>? CheckBinaryExistenceInFilestore { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("excludePathPrefixPattern")]
        public Input<string>? ExcludePathPrefixPattern { get; set; }

        [Input("includePathPrefixPattern")]
        public Input<string>? IncludePathPrefixPattern { get; set; }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("replicationKey")]
        public Input<string>? ReplicationKey { get; set; }

        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public LocalRepositoryMultiReplicationReplicationArgs()
        {
        }
        public static new LocalRepositoryMultiReplicationReplicationArgs Empty => new LocalRepositoryMultiReplicationReplicationArgs();
    }
}
