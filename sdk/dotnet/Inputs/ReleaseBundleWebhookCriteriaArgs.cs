// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ReleaseBundleWebhookCriteriaArgs : Pulumi.ResourceArgs
    {
        [Input("anyReleaseBundle", required: true)]
        public Input<bool> AnyReleaseBundle { get; set; } = null!;

        [Input("excludePatterns")]
        private InputList<string>? _excludePatterns;
        public InputList<string> ExcludePatterns
        {
            get => _excludePatterns ?? (_excludePatterns = new InputList<string>());
            set => _excludePatterns = value;
        }

        [Input("includePatterns")]
        private InputList<string>? _includePatterns;
        public InputList<string> IncludePatterns
        {
            get => _includePatterns ?? (_includePatterns = new InputList<string>());
            set => _includePatterns = value;
        }

        [Input("registeredReleaseBundleNames", required: true)]
        private InputList<string>? _registeredReleaseBundleNames;
        public InputList<string> RegisteredReleaseBundleNames
        {
            get => _registeredReleaseBundleNames ?? (_registeredReleaseBundleNames = new InputList<string>());
            set => _registeredReleaseBundleNames = value;
        }

        public ReleaseBundleWebhookCriteriaArgs()
        {
        }
    }
}
