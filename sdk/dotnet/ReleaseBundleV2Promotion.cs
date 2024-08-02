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
    /// This resource enables you to promote Release Bundle V2 version. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/promote-a-release-bundle-v2-to-a-target-environment).
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
    ///     var my_release_bundle_v2_promotion = new Artifactory.ReleaseBundleV2Promotion("my-release-bundle-v2-promotion", new()
    ///     {
    ///         Name = "my-release-bundle-v2-artifacts",
    ///         Version = "1.0.0",
    ///         KeypairName = "my-keypair-name",
    ///         Environment = "DEV",
    ///         IncludedRepositoryKeys = new[]
    ///         {
    ///             "commons-qa-maven-local",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion")]
    public partial class ReleaseBundleV2Promotion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Timestamp when the new version was created (ISO 8601 standard).
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Timestamp when the new version was created (in milliseconds).
        /// </summary>
        [Output("createdMillis")]
        public Output<int> CreatedMillis { get; private set; } = null!;

        /// <summary>
        /// Target environment
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// Defines specific repositories to exclude from the promotion.
        /// </summary>
        [Output("excludedRepositoryKeys")]
        public Output<ImmutableArray<string>> ExcludedRepositoryKeys { get; private set; } = null!;

        /// <summary>
        /// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        /// </summary>
        [Output("includedRepositoryKeys")]
        public Output<ImmutableArray<string>> IncludedRepositoryKeys { get; private set; } = null!;

        /// <summary>
        /// Key-pair name to use for signature creation
        /// </summary>
        [Output("keypairName")]
        public Output<string> KeypairName { get; private set; } = null!;

        /// <summary>
        /// Name of Release Bundle
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Project key the Release Bundle belongs to
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// Version to promote
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ReleaseBundleV2Promotion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReleaseBundleV2Promotion(string name, ReleaseBundleV2PromotionArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion", name, args ?? new ReleaseBundleV2PromotionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReleaseBundleV2Promotion(string name, Input<string> id, ReleaseBundleV2PromotionState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReleaseBundleV2Promotion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReleaseBundleV2Promotion Get(string name, Input<string> id, ReleaseBundleV2PromotionState? state = null, CustomResourceOptions? options = null)
        {
            return new ReleaseBundleV2Promotion(name, id, state, options);
        }
    }

    public sealed class ReleaseBundleV2PromotionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target environment
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        [Input("excludedRepositoryKeys")]
        private InputList<string>? _excludedRepositoryKeys;

        /// <summary>
        /// Defines specific repositories to exclude from the promotion.
        /// </summary>
        public InputList<string> ExcludedRepositoryKeys
        {
            get => _excludedRepositoryKeys ?? (_excludedRepositoryKeys = new InputList<string>());
            set => _excludedRepositoryKeys = value;
        }

        [Input("includedRepositoryKeys")]
        private InputList<string>? _includedRepositoryKeys;

        /// <summary>
        /// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        /// </summary>
        public InputList<string> IncludedRepositoryKeys
        {
            get => _includedRepositoryKeys ?? (_includedRepositoryKeys = new InputList<string>());
            set => _includedRepositoryKeys = value;
        }

        /// <summary>
        /// Key-pair name to use for signature creation
        /// </summary>
        [Input("keypairName", required: true)]
        public Input<string> KeypairName { get; set; } = null!;

        /// <summary>
        /// Name of Release Bundle
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project key the Release Bundle belongs to
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Version to promote
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ReleaseBundleV2PromotionArgs()
        {
        }
        public static new ReleaseBundleV2PromotionArgs Empty => new ReleaseBundleV2PromotionArgs();
    }

    public sealed class ReleaseBundleV2PromotionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Timestamp when the new version was created (ISO 8601 standard).
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Timestamp when the new version was created (in milliseconds).
        /// </summary>
        [Input("createdMillis")]
        public Input<int>? CreatedMillis { get; set; }

        /// <summary>
        /// Target environment
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        [Input("excludedRepositoryKeys")]
        private InputList<string>? _excludedRepositoryKeys;

        /// <summary>
        /// Defines specific repositories to exclude from the promotion.
        /// </summary>
        public InputList<string> ExcludedRepositoryKeys
        {
            get => _excludedRepositoryKeys ?? (_excludedRepositoryKeys = new InputList<string>());
            set => _excludedRepositoryKeys = value;
        }

        [Input("includedRepositoryKeys")]
        private InputList<string>? _includedRepositoryKeys;

        /// <summary>
        /// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
        /// </summary>
        public InputList<string> IncludedRepositoryKeys
        {
            get => _includedRepositoryKeys ?? (_includedRepositoryKeys = new InputList<string>());
            set => _includedRepositoryKeys = value;
        }

        /// <summary>
        /// Key-pair name to use for signature creation
        /// </summary>
        [Input("keypairName")]
        public Input<string>? KeypairName { get; set; }

        /// <summary>
        /// Name of Release Bundle
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project key the Release Bundle belongs to
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Version to promote
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ReleaseBundleV2PromotionState()
        {
        }
        public static new ReleaseBundleV2PromotionState Empty => new ReleaseBundleV2PromotionState();
    }
}
