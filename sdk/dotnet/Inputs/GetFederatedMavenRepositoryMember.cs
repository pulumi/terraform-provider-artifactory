// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class GetFederatedMavenRepositoryMemberArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessToken")]
        private string? _accessToken;

        /// <summary>
        /// Admin access token for this member Artifactory instance. Used in conjunction with `cleanup_on_delete` attribute when Access Federation for access tokens is not enabled.
        /// </summary>
        public string? AccessToken
        {
            get => _accessToken;
            set => _accessToken = value;
        }

        /// <summary>
        /// Represents the active state of the federated member. It is supported to change the enabled
        /// status of my own member. The config will be updated on the other federated members automatically.
        /// </summary>
        [Input("enabled", required: true)]
        public bool Enabled { get; set; }

        /// <summary>
        /// Full URL to ending with the repository name.
        /// </summary>
        [Input("url", required: true)]
        public string Url { get; set; } = null!;

        public GetFederatedMavenRepositoryMemberArgs()
        {
        }
        public static new GetFederatedMavenRepositoryMemberArgs Empty => new GetFederatedMavenRepositoryMemberArgs();
    }
}
