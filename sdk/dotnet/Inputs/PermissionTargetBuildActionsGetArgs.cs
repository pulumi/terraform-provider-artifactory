// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetBuildActionsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<Inputs.PermissionTargetBuildActionsGroupGetArgs>? _groups;
        public InputList<Inputs.PermissionTargetBuildActionsGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.PermissionTargetBuildActionsGroupGetArgs>());
            set => _groups = value;
        }

        [Input("users")]
        private InputList<Inputs.PermissionTargetBuildActionsUserGetArgs>? _users;
        public InputList<Inputs.PermissionTargetBuildActionsUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.PermissionTargetBuildActionsUserGetArgs>());
            set => _users = value;
        }

        public PermissionTargetBuildActionsGetArgs()
        {
        }
        public static new PermissionTargetBuildActionsGetArgs Empty => new PermissionTargetBuildActionsGetArgs();
    }
}
