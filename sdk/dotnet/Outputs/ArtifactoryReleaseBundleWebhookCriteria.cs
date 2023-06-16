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
    public sealed class ArtifactoryReleaseBundleWebhookCriteria
    {
        public readonly bool AnyReleaseBundle;
        public readonly ImmutableArray<string> ExcludePatterns;
        public readonly ImmutableArray<string> IncludePatterns;
        public readonly ImmutableArray<string> RegisteredReleaseBundleNames;

        [OutputConstructor]
        private ArtifactoryReleaseBundleWebhookCriteria(
            bool anyReleaseBundle,

            ImmutableArray<string> excludePatterns,

            ImmutableArray<string> includePatterns,

            ImmutableArray<string> registeredReleaseBundleNames)
        {
            AnyReleaseBundle = anyReleaseBundle;
            ExcludePatterns = excludePatterns;
            IncludePatterns = includePatterns;
            RegisteredReleaseBundleNames = registeredReleaseBundleNames;
        }
    }
}
