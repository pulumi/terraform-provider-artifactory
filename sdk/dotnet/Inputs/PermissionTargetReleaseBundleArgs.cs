// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetReleaseBundleArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        public Input<Inputs.PermissionTargetReleaseBundleActionsArgs>? Actions { get; set; }

        [Input("excludesPatterns")]
        private InputList<string>? _excludesPatterns;

        /// <summary>
        /// The default value will be [] if nothing is supplied
        /// </summary>
        public InputList<string> ExcludesPatterns
        {
            get => _excludesPatterns ?? (_excludesPatterns = new InputList<string>());
            set => _excludesPatterns = value;
        }

        [Input("includesPatterns")]
        private InputList<string>? _includesPatterns;

        /// <summary>
        /// The default value will be [""] if nothing is supplied
        /// </summary>
        public InputList<string> IncludesPatterns
        {
            get => _includesPatterns ?? (_includesPatterns = new InputList<string>());
            set => _includesPatterns = value;
        }

        [Input("repositories", required: true)]
        private InputList<string>? _repositories;

        /// <summary>
        /// This can only be 1 value: "artifactory-build-info", and currently, validation of sets/lists is not allowed. Artifactory will reject the request if you change this
        /// </summary>
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        public PermissionTargetReleaseBundleArgs()
        {
        }
        public static new PermissionTargetReleaseBundleArgs Empty => new PermissionTargetReleaseBundleArgs();
    }
}
