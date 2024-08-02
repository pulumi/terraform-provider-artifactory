// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ReleaseBundleV2SourceBuildArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether to include build dependencies in the Release Bundle. The default value is `false`.
        /// </summary>
        [Input("includeDependencies")]
        public Input<bool>? IncludeDependencies { get; set; }

        /// <summary>
        /// Name of the build.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Number (run) of the build.
        /// </summary>
        [Input("number", required: true)]
        public Input<string> Number { get; set; } = null!;

        /// <summary>
        /// The repository key of the build. If omitted, the system uses the default built-in repository, `artifactory-build-info`.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// Timestamp when the build was created. If omitted, the system uses the latest build run, as identified by the `name` and `number` combination. The timestamp is provided according to the ISO 8601 standard.
        /// </summary>
        [Input("started")]
        public Input<string>? Started { get; set; }

        public ReleaseBundleV2SourceBuildArgs()
        {
        }
        public static new ReleaseBundleV2SourceBuildArgs Empty => new ReleaseBundleV2SourceBuildArgs();
    }
}
