// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class DistributionWebhookCriteriaGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger on any release bundle
        /// </summary>
        [Input("anyReleaseBundle", required: true)]
        public Input<bool> AnyReleaseBundle { get; set; } = null!;

        [Input("excludePatterns")]
        private InputList<string>? _excludePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
        /// </summary>
        public InputList<string> ExcludePatterns
        {
            get => _excludePatterns ?? (_excludePatterns = new InputList<string>());
            set => _excludePatterns = value;
        }

        [Input("includePatterns")]
        private InputList<string>? _includePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
        /// </summary>
        public InputList<string> IncludePatterns
        {
            get => _includePatterns ?? (_includePatterns = new InputList<string>());
            set => _includePatterns = value;
        }

        [Input("registeredReleaseBundleNames", required: true)]
        private InputList<string>? _registeredReleaseBundleNames;

        /// <summary>
        /// Trigger on this list of release bundle names
        /// </summary>
        public InputList<string> RegisteredReleaseBundleNames
        {
            get => _registeredReleaseBundleNames ?? (_registeredReleaseBundleNames = new InputList<string>());
            set => _registeredReleaseBundleNames = value;
        }

        public DistributionWebhookCriteriaGetArgs()
        {
        }
    }
}
