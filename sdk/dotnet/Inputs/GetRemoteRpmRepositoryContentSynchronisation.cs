// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class GetRemoteRpmRepositoryContentSynchronisationArgs : global::Pulumi.InvokeArgs
    {
        [Input("enabled")]
        public bool? Enabled { get; set; }

        [Input("propertiesEnabled")]
        public bool? PropertiesEnabled { get; set; }

        [Input("sourceOriginAbsenceDetection")]
        public bool? SourceOriginAbsenceDetection { get; set; }

        [Input("statisticsEnabled")]
        public bool? StatisticsEnabled { get; set; }

        public GetRemoteRpmRepositoryContentSynchronisationArgs()
        {
        }
        public static new GetRemoteRpmRepositoryContentSynchronisationArgs Empty => new GetRemoteRpmRepositoryContentSynchronisationArgs();
    }
}
