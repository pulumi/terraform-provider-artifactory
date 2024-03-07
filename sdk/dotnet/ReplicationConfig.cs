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
    /// &gt; This resource is deprecated in favor of `artifactory.PushReplication` resource.
    /// 
    /// Provides an Artifactory replication config resource. This can be used to create and manage Artifactory replications.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a replication between two artifactory local repositories
    ///     var providerTestSource = new Artifactory.LocalMavenRepository("providerTestSource", new()
    ///     {
    ///         Key = "provider_test_source",
    ///     });
    /// 
    ///     var providerTestDest = new Artifactory.LocalMavenRepository("providerTestDest", new()
    ///     {
    ///         Key = "provider_test_dest",
    ///     });
    /// 
    ///     var foo_rep = new Artifactory.ReplicationConfig("foo-rep", new()
    ///     {
    ///         CronExp = "0 0 * * * ?",
    ///         EnableEventReplication = true,
    ///         Replications = new[]
    ///         {
    ///             new Artifactory.Inputs.ReplicationConfigReplicationArgs
    ///             {
    ///                 Password = "$var.artifactory_password",
    ///                 Url = "$var.artifactory_url",
    ///                 Username = "$var.artifactory_username",
    ///             },
    ///         },
    ///         RepoKey = providerTestSource.Key,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Replication configs can be imported using their repo key, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/replicationConfig:ReplicationConfig foo-rep provider_test_source
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/replicationConfig:ReplicationConfig")]
    public partial class ReplicationConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cron expression to control the operation frequency.
        /// </summary>
        [Output("cronExp")]
        public Output<string> CronExp { get; private set; } = null!;

        [Output("enableEventReplication")]
        public Output<bool> EnableEventReplication { get; private set; } = null!;

        [Output("replications")]
        public Output<ImmutableArray<Outputs.ReplicationConfigReplication>> Replications { get; private set; } = null!;

        [Output("repoKey")]
        public Output<string> RepoKey { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationConfig(string name, ReplicationConfigArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/replicationConfig:ReplicationConfig", name, args ?? new ReplicationConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationConfig(string name, Input<string> id, ReplicationConfigState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/replicationConfig:ReplicationConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicationConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationConfig Get(string name, Input<string> id, ReplicationConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationConfig(name, id, state, options);
        }
    }

    public sealed class ReplicationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron expression to control the operation frequency.
        /// </summary>
        [Input("cronExp", required: true)]
        public Input<string> CronExp { get; set; } = null!;

        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        [Input("replications")]
        private InputList<Inputs.ReplicationConfigReplicationArgs>? _replications;
        public InputList<Inputs.ReplicationConfigReplicationArgs> Replications
        {
            get => _replications ?? (_replications = new InputList<Inputs.ReplicationConfigReplicationArgs>());
            set => _replications = value;
        }

        [Input("repoKey", required: true)]
        public Input<string> RepoKey { get; set; } = null!;

        public ReplicationConfigArgs()
        {
        }
        public static new ReplicationConfigArgs Empty => new ReplicationConfigArgs();
    }

    public sealed class ReplicationConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron expression to control the operation frequency.
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        [Input("replications")]
        private InputList<Inputs.ReplicationConfigReplicationGetArgs>? _replications;
        public InputList<Inputs.ReplicationConfigReplicationGetArgs> Replications
        {
            get => _replications ?? (_replications = new InputList<Inputs.ReplicationConfigReplicationGetArgs>());
            set => _replications = value;
        }

        [Input("repoKey")]
        public Input<string>? RepoKey { get; set; }

        public ReplicationConfigState()
        {
        }
        public static new ReplicationConfigState Empty => new ReplicationConfigState();
    }
}
