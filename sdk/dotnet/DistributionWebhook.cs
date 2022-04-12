// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    [ArtifactoryResourceType("artifactory:index/distributionWebhook:DistributionWebhook")]
    public partial class DistributionWebhook : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies where the webhook will be applied, on which release bundles or distributions.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.DistributionWebhookCriteria> Criteria { get; private set; } = null!;

        /// <summary>
        /// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        /// </summary>
        [Output("customHttpHeaders")]
        public Output<ImmutableDictionary<string, string>?> CustomHttpHeaders { get; private set; } = null!;

        /// <summary>
        /// Description of webhook. Max length 1000 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Status of webhook. Default to 'true'
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
        /// values: distribute_started, distribute_completed, distribute_aborted, distribute_failed, delete_started,
        /// delete_completed, delete_failed
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Output("proxy")]
        public Output<string?> Proxy { get; private set; } = null!;

        /// <summary>
        /// Secret authentication token that will be sent to the configured URL.
        /// </summary>
        [Output("secret")]
        public Output<string?> Secret { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a DistributionWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DistributionWebhook(string name, DistributionWebhookArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/distributionWebhook:DistributionWebhook", name, args ?? new DistributionWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DistributionWebhook(string name, Input<string> id, DistributionWebhookState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/distributionWebhook:DistributionWebhook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DistributionWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DistributionWebhook Get(string name, Input<string> id, DistributionWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new DistributionWebhook(name, id, state, options);
        }
    }

    public sealed class DistributionWebhookArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied, on which release bundles or distributions.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.DistributionWebhookCriteriaArgs> Criteria { get; set; } = null!;

        [Input("customHttpHeaders")]
        private InputMap<string>? _customHttpHeaders;

        /// <summary>
        /// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        /// </summary>
        public InputMap<string> CustomHttpHeaders
        {
            get => _customHttpHeaders ?? (_customHttpHeaders = new InputMap<string>());
            set => _customHttpHeaders = value;
        }

        /// <summary>
        /// Description of webhook. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to 'true'
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes", required: true)]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
        /// values: distribute_started, distribute_completed, distribute_aborted, distribute_failed, delete_started,
        /// delete_completed, delete_failed
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        /// <summary>
        /// Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Secret authentication token that will be sent to the configured URL.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public DistributionWebhookArgs()
        {
        }
    }

    public sealed class DistributionWebhookState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied, on which release bundles or distributions.
        /// </summary>
        [Input("criteria")]
        public Input<Inputs.DistributionWebhookCriteriaGetArgs>? Criteria { get; set; }

        [Input("customHttpHeaders")]
        private InputMap<string>? _customHttpHeaders;

        /// <summary>
        /// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        /// </summary>
        public InputMap<string> CustomHttpHeaders
        {
            get => _customHttpHeaders ?? (_customHttpHeaders = new InputMap<string>());
            set => _customHttpHeaders = value;
        }

        /// <summary>
        /// Description of webhook. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to 'true'
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes")]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow
        /// values: distribute_started, distribute_completed, distribute_aborted, distribute_failed, delete_started,
        /// delete_completed, delete_failed
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        /// <summary>
        /// Key of webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Proxy key from Artifactory Proxies setting
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Secret authentication token that will be sent to the configured URL.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public DistributionWebhookState()
        {
        }
    }
}
