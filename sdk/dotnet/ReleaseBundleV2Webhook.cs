// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
    /// 
    /// ## Example Usage
    /// 
    /// .
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var release_bundle_v2_webhook = new Artifactory.ReleaseBundleV2Webhook("release-bundle-v2-webhook", new()
    ///     {
    ///         Key = "release-bundle-v2-webhook",
    ///         EventTypes = new[]
    ///         {
    ///             "release_bundle_v2_started",
    ///             "release_bundle_v2_failed",
    ///             "release_bundle_v2_completed",
    ///         },
    ///         Criteria = new Artifactory.Inputs.ReleaseBundleV2WebhookCriteriaArgs
    ///         {
    ///             AnyReleaseBundle = false,
    ///             SelectedReleaseBundles = new[]
    ///             {
    ///                 "bundle-name",
    ///             },
    ///             IncludePatterns = new[]
    ///             {
    ///                 "foo/**",
    ///             },
    ///             ExcludePatterns = new[]
    ///             {
    ///                 "bar/**",
    ///             },
    ///         },
    ///         Handlers = new[]
    ///         {
    ///             new Artifactory.Inputs.ReleaseBundleV2WebhookHandlerArgs
    ///             {
    ///                 Url = "http://tempurl.org/webhook",
    ///                 Secret = "some-secret",
    ///                 Proxy = "proxy-key",
    ///                 CustomHttpHeaders = 
    ///                 {
    ///                     { "header-1", "value-1" },
    ///                     { "header-2", "value-2" },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/releaseBundleV2Webhook:ReleaseBundleV2Webhook")]
    public partial class ReleaseBundleV2Webhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.ReleaseBundleV2WebhookCriteria> Criteria { get; private set; } = null!;

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Status of webhook. Default to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// At least one is required.
        /// </summary>
        [Output("handlers")]
        public Output<ImmutableArray<Outputs.ReleaseBundleV2WebhookHandler>> Handlers { get; private set; } = null!;

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;


        /// <summary>
        /// Create a ReleaseBundleV2Webhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReleaseBundleV2Webhook(string name, ReleaseBundleV2WebhookArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/releaseBundleV2Webhook:ReleaseBundleV2Webhook", name, args ?? new ReleaseBundleV2WebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReleaseBundleV2Webhook(string name, Input<string> id, ReleaseBundleV2WebhookState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/releaseBundleV2Webhook:ReleaseBundleV2Webhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReleaseBundleV2Webhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReleaseBundleV2Webhook Get(string name, Input<string> id, ReleaseBundleV2WebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new ReleaseBundleV2Webhook(name, id, state, options);
        }
    }

    public sealed class ReleaseBundleV2WebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.ReleaseBundleV2WebhookCriteriaArgs> Criteria { get; set; } = null!;

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes", required: true)]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers", required: true)]
        private InputList<Inputs.ReleaseBundleV2WebhookHandlerArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.ReleaseBundleV2WebhookHandlerArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.ReleaseBundleV2WebhookHandlerArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public ReleaseBundleV2WebhookArgs()
        {
        }
        public static new ReleaseBundleV2WebhookArgs Empty => new ReleaseBundleV2WebhookArgs();
    }

    public sealed class ReleaseBundleV2WebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Input("criteria")]
        public Input<Inputs.ReleaseBundleV2WebhookCriteriaGetArgs>? Criteria { get; set; }

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes")]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `release_bundle_v2_started`, `release_bundle_v2_failed`, `release_bundle_v2_completed`.
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers")]
        private InputList<Inputs.ReleaseBundleV2WebhookHandlerGetArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.ReleaseBundleV2WebhookHandlerGetArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.ReleaseBundleV2WebhookHandlerGetArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public ReleaseBundleV2WebhookState()
        {
        }
        public static new ReleaseBundleV2WebhookState Empty => new ReleaseBundleV2WebhookState();
    }
}
