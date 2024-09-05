// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class FederatedDebianRepositoryMember
    {
        /// <summary>
        /// Admin access token for this member Artifactory instance. Used in conjunction with `cleanup_on_delete` attribute when Access Federation for access tokens is not enabled.
        /// </summary>
        public readonly string? AccessToken;
        /// <summary>
        /// Represents the active state of the federated member. It is supported to change the enabled
        /// status of my own member. The config will be updated on the other federated members automatically.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Full URL to ending with the repository name.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private FederatedDebianRepositoryMember(
            string? accessToken,

            bool enabled,

            string url)
        {
            AccessToken = accessToken;
            Enabled = enabled;
            Url = url;
        }
    }
}
