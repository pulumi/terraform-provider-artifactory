// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetBuildActionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<Inputs.PermissionTargetBuildActionsGroupGetArgs>? _groups;

        /// <summary>
        /// Groups this permission applies for.
        /// </summary>
        public InputList<Inputs.PermissionTargetBuildActionsGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.PermissionTargetBuildActionsGroupGetArgs>());
            set => _groups = value;
        }

        [Input("users")]
        private InputList<Inputs.PermissionTargetBuildActionsUserGetArgs>? _users;

        /// <summary>
        /// Users this permission target applies for.
        /// </summary>
        public InputList<Inputs.PermissionTargetBuildActionsUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.PermissionTargetBuildActionsUserGetArgs>());
            set => _users = value;
        }

        public PermissionTargetBuildActionsGetArgs()
        {
        }
    }
}
