// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class VaultConfigurationConfigAuthGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client certificate (in PEM format) for `Certificate` type.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Private key (in PEM format) for `Certificate` type.
        /// </summary>
        [Input("certificateKey")]
        public Input<string>? CertificateKey { get; set; }

        [Input("roleId")]
        private Input<string>? _roleId;

        /// <summary>
        /// Role ID for `AppRole` type
        /// </summary>
        public Input<string>? RoleId
        {
            get => _roleId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _roleId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("secretId")]
        private Input<string>? _secretId;

        /// <summary>
        /// Secret ID for `AppRole` type
        /// </summary>
        public Input<string>? SecretId
        {
            get => _secretId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public VaultConfigurationConfigAuthGetArgs()
        {
        }
        public static new VaultConfigurationConfigAuthGetArgs Empty => new VaultConfigurationConfigAuthGetArgs();
    }
}
