// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class GetFederatedChefRepositoryMemberArgs : global::Pulumi.InvokeArgs
    {
        [Input("enabled", required: true)]
        public bool Enabled { get; set; }

        [Input("url", required: true)]
        public string Url { get; set; } = null!;

        public GetFederatedChefRepositoryMemberArgs()
        {
        }
        public static new GetFederatedChefRepositoryMemberArgs Empty => new GetFederatedChefRepositoryMemberArgs();
    }
}
