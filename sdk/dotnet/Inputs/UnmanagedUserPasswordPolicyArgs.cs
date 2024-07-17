// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class UnmanagedUserPasswordPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum number of digits that the password must contain
        /// </summary>
        [Input("digit")]
        public Input<int>? Digit { get; set; }

        /// <summary>
        /// Minimum length of the password
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// Minimum number of lowercase letters that the password must contain
        /// </summary>
        [Input("lowercase")]
        public Input<int>? Lowercase { get; set; }

        /// <summary>
        /// Minimum number of special char that the password must contain. Special chars list: `!"#$%&amp;'()*+,-./:;&lt;=&gt;?@[\]^_``{|}~`
        /// </summary>
        [Input("specialChar")]
        public Input<int>? SpecialChar { get; set; }

        /// <summary>
        /// Minimum number of uppercase letters that the password must contain
        /// </summary>
        [Input("uppercase")]
        public Input<int>? Uppercase { get; set; }

        public UnmanagedUserPasswordPolicyArgs()
        {
        }
        public static new UnmanagedUserPasswordPolicyArgs Empty => new UnmanagedUserPasswordPolicyArgs();
    }
}
