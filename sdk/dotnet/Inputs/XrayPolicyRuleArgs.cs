// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class XrayPolicyRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required) Nested block describing the actions to be applied by the policy. Described below.
        /// </summary>
        [Input("actions")]
        public Input<Inputs.XrayPolicyRuleActionsArgs>? Actions { get; set; }

        /// <summary>
        /// (Required) Nested block describing the criteria for the policy. Described below.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.XrayPolicyRuleCriteriaArgs> Criteria { get; set; } = null!;

        /// <summary>
        /// (Required) Name of the rule
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// (Required) Integer describing the rule priority
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        public XrayPolicyRuleArgs()
        {
        }
    }
}
