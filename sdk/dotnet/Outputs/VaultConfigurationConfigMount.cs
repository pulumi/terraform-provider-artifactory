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
    public sealed class VaultConfigurationConfigMount
    {
        /// <summary>
        /// Vault secret engine path
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Vault supports several secret engines, each one has different capabilities. The supported secret engine types are: `KV1` and `KV2`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private VaultConfigurationConfigMount(
            string path,

            string type)
        {
            Path = path;
            Type = type;
        }
    }
}
