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
    public sealed class UserPasswordPolicy
    {
        /// <summary>
        /// Minimum number of digits that the password must contain
        /// </summary>
        public readonly int? Digit;
        /// <summary>
        /// Minimum length of the password
        /// </summary>
        public readonly int? Length;
        /// <summary>
        /// Minimum number of lowercase letters that the password must contain
        /// </summary>
        public readonly int? Lowercase;
        /// <summary>
        /// Minimum number of special char that the password must contain. Special chars list: `!"#$%&amp;'()*+,-./:;&lt;=&gt;?@[\]^_``{|}~`
        /// </summary>
        public readonly int? SpecialChar;
        /// <summary>
        /// Minimum number of uppercase letters that the password must contain
        /// </summary>
        public readonly int? Uppercase;

        [OutputConstructor]
        private UserPasswordPolicy(
            int? digit,

            int? length,

            int? lowercase,

            int? specialChar,

            int? uppercase)
        {
            Digit = digit;
            Length = length;
            Lowercase = lowercase;
            SpecialChar = specialChar;
            Uppercase = uppercase;
        }
    }
}
