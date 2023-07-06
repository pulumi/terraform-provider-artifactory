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
    /// Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_docker_local = new Artifactory.DockerV2Repository("my-docker-local", new()
    ///     {
    ///         Key = "my-docker-local",
    ///     });
    /// 
    ///     var docker_custom_webhook = new Artifactory.DockerCustomWebhook("docker-custom-webhook", new()
    ///     {
    ///         Key = "docker-custom_webhook",
    ///         EventTypes = new[]
    ///         {
    ///             "pushed",
    ///             "deleted",
    ///             "promoted",
    ///         },
    ///         Criteria = new Artifactory.Inputs.DockerCustomWebhookCriteriaArgs
    ///         {
    ///             AnyLocal = true,
    ///             AnyRemote = false,
    ///             RepoKeys = new[]
    ///             {
    ///                 my_docker_local.Key,
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
    ///             new Artifactory.Inputs.DockerCustomWebhookHandlerArgs
    ///             {
    ///                 Url = "https://tempurl.org",
    ///                 Secrets = 
    ///                 {
    ///                     { "secretName1", "value1" },
    ///                     { "secretName2", "value2" },
    ///                 },
    ///                 HttpHeaders = 
    ///                 {
    ///                     { "headerName1", "value1" },
    ///                     { "headerName2", "value2" },
    ///                 },
    ///                 Payload = "{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }",
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             my_docker_local,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/dockerCustomWebhook:DockerCustomWebhook")]
    public partial class DockerCustomWebhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.DockerCustomWebhookCriteria> Criteria { get; private set; } = null!;

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
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// At least one is required.
        /// </summary>
        [Output("handlers")]
        public Output<ImmutableArray<Outputs.DockerCustomWebhookHandler>> Handlers { get; private set; } = null!;

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;


        /// <summary>
        /// Create a DockerCustomWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DockerCustomWebhook(string name, DockerCustomWebhookArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/dockerCustomWebhook:DockerCustomWebhook", name, args ?? new DockerCustomWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DockerCustomWebhook(string name, Input<string> id, DockerCustomWebhookState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/dockerCustomWebhook:DockerCustomWebhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DockerCustomWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DockerCustomWebhook Get(string name, Input<string> id, DockerCustomWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new DockerCustomWebhook(name, id, state, options);
        }
    }

    public sealed class DockerCustomWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.DockerCustomWebhookCriteriaArgs> Criteria { get; set; } = null!;

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
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers", required: true)]
        private InputList<Inputs.DockerCustomWebhookHandlerArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.DockerCustomWebhookHandlerArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.DockerCustomWebhookHandlerArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public DockerCustomWebhookArgs()
        {
        }
        public static new DockerCustomWebhookArgs Empty => new DockerCustomWebhookArgs();
    }

    public sealed class DockerCustomWebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Input("criteria")]
        public Input<Inputs.DockerCustomWebhookCriteriaGetArgs>? Criteria { get; set; }

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
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers")]
        private InputList<Inputs.DockerCustomWebhookHandlerGetArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.DockerCustomWebhookHandlerGetArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.DockerCustomWebhookHandlerGetArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public DockerCustomWebhookState()
        {
        }
        public static new DockerCustomWebhookState Empty => new DockerCustomWebhookState();
    }
}
