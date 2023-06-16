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
    public sealed class GetPermissionTargetRepoActionsResult
    {
        public readonly ImmutableArray<Outputs.GetPermissionTargetRepoActionsGroupResult> Groups;
        public readonly ImmutableArray<Outputs.GetPermissionTargetRepoActionsUserResult> Users;

        [OutputConstructor]
        private GetPermissionTargetRepoActionsResult(
            ImmutableArray<Outputs.GetPermissionTargetRepoActionsGroupResult> groups,

            ImmutableArray<Outputs.GetPermissionTargetRepoActionsUserResult> users)
        {
            Groups = groups;
            Users = users;
        }
    }
}
