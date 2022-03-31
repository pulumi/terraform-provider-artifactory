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
    public sealed class RemoteAlpineRepositoryContentSynchronisation
    {
        public readonly bool? Enabled;
        public readonly bool? PropertiesEnabled;
        public readonly bool? SourceOriginAbsenceDetection;
        public readonly bool? StatisticsEnabled;

        [OutputConstructor]
        private RemoteAlpineRepositoryContentSynchronisation(
            bool? enabled,

            bool? propertiesEnabled,

            bool? sourceOriginAbsenceDetection,

            bool? statisticsEnabled)
        {
            Enabled = enabled;
            PropertiesEnabled = propertiesEnabled;
            SourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            StatisticsEnabled = statisticsEnabled;
        }
    }
}
