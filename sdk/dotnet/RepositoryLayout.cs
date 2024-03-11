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
    /// This resource can be used to manage Artifactory's Repository Layout settings. See [Repository Layouts](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts) in the Artifactory Wiki documentation for more details.
    /// 
    /// ~&gt;The `artifactory.RepositoryLayout` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
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
    ///     var custom_layout = new Artifactory.RepositoryLayout("custom-layout", new()
    ///     {
    ///         ArtifactPathPattern = "[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).[ext]",
    ///         DescriptorPathPattern = "[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).pom",
    ///         DistinctiveDescriptorPathPattern = true,
    ///         FileIntegrationRevisionRegexp = "Foo|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))",
    ///         FolderIntegrationRevisionRegexp = "Foo",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Repository layout can be imported using its name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/repositoryLayout:RepositoryLayout custom-layout custom-layout
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/repositoryLayout:RepositoryLayout")]
    public partial class RepositoryLayout : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
        /// </summary>
        [Output("artifactPathPattern")]
        public Output<string> ArtifactPathPattern { get; private set; } = null!;

        /// <summary>
        /// Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
        /// </summary>
        [Output("descriptorPathPattern")]
        public Output<string?> DescriptorPathPattern { get; private set; } = null!;

        /// <summary>
        /// When set, `descriptor_path_pattern` will be used. Default to `false`.
        /// </summary>
        [Output("distinctiveDescriptorPathPattern")]
        public Output<bool?> DistinctiveDescriptorPathPattern { get; private set; } = null!;

        /// <summary>
        /// A regular expression matching the integration revision string appearing in a file name as part of the artifact's path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
        /// </summary>
        [Output("fileIntegrationRevisionRegexp")]
        public Output<string> FileIntegrationRevisionRegexp { get; private set; } = null!;

        /// <summary>
        /// A regular expression matching the integration revision string appearing in a folder name as part of the artifact's path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
        /// </summary>
        [Output("folderIntegrationRevisionRegexp")]
        public Output<string> FolderIntegrationRevisionRegexp { get; private set; } = null!;

        /// <summary>
        /// Layout name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryLayout resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryLayout(string name, RepositoryLayoutArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/repositoryLayout:RepositoryLayout", name, args ?? new RepositoryLayoutArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryLayout(string name, Input<string> id, RepositoryLayoutState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/repositoryLayout:RepositoryLayout", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryLayout resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryLayout Get(string name, Input<string> id, RepositoryLayoutState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryLayout(name, id, state, options);
        }
    }

    public sealed class RepositoryLayoutArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
        /// </summary>
        [Input("artifactPathPattern", required: true)]
        public Input<string> ArtifactPathPattern { get; set; } = null!;

        /// <summary>
        /// Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
        /// </summary>
        [Input("descriptorPathPattern")]
        public Input<string>? DescriptorPathPattern { get; set; }

        /// <summary>
        /// When set, `descriptor_path_pattern` will be used. Default to `false`.
        /// </summary>
        [Input("distinctiveDescriptorPathPattern")]
        public Input<bool>? DistinctiveDescriptorPathPattern { get; set; }

        /// <summary>
        /// A regular expression matching the integration revision string appearing in a file name as part of the artifact's path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
        /// </summary>
        [Input("fileIntegrationRevisionRegexp", required: true)]
        public Input<string> FileIntegrationRevisionRegexp { get; set; } = null!;

        /// <summary>
        /// A regular expression matching the integration revision string appearing in a folder name as part of the artifact's path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
        /// </summary>
        [Input("folderIntegrationRevisionRegexp", required: true)]
        public Input<string> FolderIntegrationRevisionRegexp { get; set; } = null!;

        /// <summary>
        /// Layout name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RepositoryLayoutArgs()
        {
        }
        public static new RepositoryLayoutArgs Empty => new RepositoryLayoutArgs();
    }

    public sealed class RepositoryLayoutState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Please refer to: [Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts) in the Artifactory Wiki documentation.
        /// </summary>
        [Input("artifactPathPattern")]
        public Input<string>? ArtifactPathPattern { get; set; }

        /// <summary>
        /// Please refer to: [Descriptor Path Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in the Artifactory Wiki documentation.
        /// </summary>
        [Input("descriptorPathPattern")]
        public Input<string>? DescriptorPathPattern { get; set; }

        /// <summary>
        /// When set, `descriptor_path_pattern` will be used. Default to `false`.
        /// </summary>
        [Input("distinctiveDescriptorPathPattern")]
        public Input<bool>? DistinctiveDescriptorPathPattern { get; set; }

        /// <summary>
        /// A regular expression matching the integration revision string appearing in a file name as part of the artifact's path. For example, `SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
        /// </summary>
        [Input("fileIntegrationRevisionRegexp")]
        public Input<string>? FileIntegrationRevisionRegexp { get; set; }

        /// <summary>
        /// A regular expression matching the integration revision string appearing in a folder name as part of the artifact's path. For example, `SNAPSHOT`, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression. If not applicable use `.*`
        /// </summary>
        [Input("folderIntegrationRevisionRegexp")]
        public Input<string>? FolderIntegrationRevisionRegexp { get; set; }

        /// <summary>
        /// Layout name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RepositoryLayoutState()
        {
        }
        public static new RepositoryLayoutState Empty => new RepositoryLayoutState();
    }
}
