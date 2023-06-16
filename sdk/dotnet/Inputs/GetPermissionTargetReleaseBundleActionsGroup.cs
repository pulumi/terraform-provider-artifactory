// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class GetPermissionTargetReleaseBundleActionsGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("permissions", required: true)]
        private List<string>? _permissions;
        public List<string> Permissions
        {
            get => _permissions ?? (_permissions = new List<string>());
            set => _permissions = value;
        }

        public GetPermissionTargetReleaseBundleActionsGroupArgs()
        {
        }
        public static new GetPermissionTargetReleaseBundleActionsGroupArgs Empty => new GetPermissionTargetReleaseBundleActionsGroupArgs();
    }
}
