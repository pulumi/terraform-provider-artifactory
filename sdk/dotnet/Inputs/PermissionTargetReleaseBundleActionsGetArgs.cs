// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetReleaseBundleActionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<Inputs.PermissionTargetReleaseBundleActionsGroupGetArgs>? _groups;
        public InputList<Inputs.PermissionTargetReleaseBundleActionsGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.PermissionTargetReleaseBundleActionsGroupGetArgs>());
            set => _groups = value;
        }

        [Input("users")]
        private InputList<Inputs.PermissionTargetReleaseBundleActionsUserGetArgs>? _users;
        public InputList<Inputs.PermissionTargetReleaseBundleActionsUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.PermissionTargetReleaseBundleActionsUserGetArgs>());
            set => _users = value;
        }

        public PermissionTargetReleaseBundleActionsGetArgs()
        {
        }
    }
}
