// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class RemoteCondaRepositoryContentSynchronisationArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("propertiesEnabled")]
        public Input<bool>? PropertiesEnabled { get; set; }

        [Input("sourceOriginAbsenceDetection")]
        public Input<bool>? SourceOriginAbsenceDetection { get; set; }

        [Input("statisticsEnabled")]
        public Input<bool>? StatisticsEnabled { get; set; }

        public RemoteCondaRepositoryContentSynchronisationArgs()
        {
        }
        public static new RemoteCondaRepositoryContentSynchronisationArgs Empty => new RemoteCondaRepositoryContentSynchronisationArgs();
    }
}
