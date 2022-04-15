// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetBuildArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// </summary>
        [Input("actions")]
        public Input<Inputs.PermissionTargetBuildActionsArgs>? Actions { get; set; }

        [Input("excludesPatterns")]
        private InputList<string>? _excludesPatterns;

        /// <summary>
        /// Pattern of artifacts to exclude
        /// </summary>
        public InputList<string> ExcludesPatterns
        {
            get => _excludesPatterns ?? (_excludesPatterns = new InputList<string>());
            set => _excludesPatterns = value;
        }

        [Input("includesPatterns")]
        private InputList<string>? _includesPatterns;

        /// <summary>
        /// Pattern of artifacts to include
        /// </summary>
        public InputList<string> IncludesPatterns
        {
            get => _includesPatterns ?? (_includesPatterns = new InputList<string>());
            set => _includesPatterns = value;
        }

        [Input("repositories", required: true)]
        private InputList<string>? _repositories;

        /// <summary>
        /// List of repositories this permission target is applicable for
        /// </summary>
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        public PermissionTargetBuildArgs()
        {
        }
    }
}
