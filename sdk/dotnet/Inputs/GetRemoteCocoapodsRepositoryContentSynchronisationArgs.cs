// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class GetRemoteCocoapodsRepositoryContentSynchronisationInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.
        /// </summary>
        [Input("propertiesEnabled")]
        public Input<bool>? PropertiesEnabled { get; set; }

        /// <summary>
        /// If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'
        /// </summary>
        [Input("sourceOriginAbsenceDetection")]
        public Input<bool>? SourceOriginAbsenceDetection { get; set; }

        /// <summary>
        /// If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.
        /// </summary>
        [Input("statisticsEnabled")]
        public Input<bool>? StatisticsEnabled { get; set; }

        public GetRemoteCocoapodsRepositoryContentSynchronisationInputArgs()
        {
        }
        public static new GetRemoteCocoapodsRepositoryContentSynchronisationInputArgs Empty => new GetRemoteCocoapodsRepositoryContentSynchronisationInputArgs();
    }
}
