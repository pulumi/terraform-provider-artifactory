// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class XrayPolicyRuleCriteriaGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Whether or not to allow components whose license cannot be determined (`true` or `false`).
        /// </summary>
        [Input("allowUnknown")]
        public Input<bool>? AllowUnknown { get; set; }

        [Input("allowedLicenses")]
        private InputList<string>? _allowedLicenses;

        /// <summary>
        /// (Optional) A list of OSS license names that may be attached to a component.
        /// </summary>
        public InputList<string> AllowedLicenses
        {
            get => _allowedLicenses ?? (_allowedLicenses = new InputList<string>());
            set => _allowedLicenses = value;
        }

        [Input("bannedLicenses")]
        private InputList<string>? _bannedLicenses;

        /// <summary>
        /// (Optional) A list of OSS license names that may not be attached to a component.
        /// </summary>
        public InputList<string> BannedLicenses
        {
            get => _bannedLicenses ?? (_bannedLicenses = new InputList<string>());
            set => _bannedLicenses = value;
        }

        /// <summary>
        /// (Optional) Nested block describing a CVS score range to be impacted. Defined below.
        /// </summary>
        [Input("cvssRange")]
        public Input<Inputs.XrayPolicyRuleCriteriaCvssRangeGetArgs>? CvssRange { get; set; }

        /// <summary>
        /// (Optional) The minimum security vulnerability severity that will be impacted by the policy.
        /// </summary>
        [Input("minSeverity")]
        public Input<string>? MinSeverity { get; set; }

        public XrayPolicyRuleCriteriaGetArgs()
        {
        }
    }
}
