// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetsRepoGetArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        public Input<Inputs.PermissionTargetsRepoActionsGetArgs>? Actions { get; set; }

        [Input("excludesPatterns")]
        private InputList<string>? _excludesPatterns;
        public InputList<string> ExcludesPatterns
        {
            get => _excludesPatterns ?? (_excludesPatterns = new InputList<string>());
            set => _excludesPatterns = value;
        }

        [Input("includesPatterns")]
        private InputList<string>? _includesPatterns;
        public InputList<string> IncludesPatterns
        {
            get => _includesPatterns ?? (_includesPatterns = new InputList<string>());
            set => _includesPatterns = value;
        }

        [Input("repositories", required: true)]
        private InputList<string>? _repositories;
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        public PermissionTargetsRepoGetArgs()
        {
        }
    }
}
