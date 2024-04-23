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
    /// Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.
    /// 
    /// &gt; This resource has been deprecated in favor of platform_permission resource.
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
    ///     // Create a new Artifactory permission target called testpermission
    ///     var test_perm = new Artifactory.PermissionTarget("test-perm", new()
    ///     {
    ///         Name = "test-perm",
    ///         Repo = new Artifactory.Inputs.PermissionTargetRepoArgs
    ///         {
    ///             IncludesPatterns = new[]
    ///             {
    ///                 "foo/**",
    ///             },
    ///             ExcludesPatterns = new[]
    ///             {
    ///                 "bar/**",
    ///             },
    ///             Repositories = new[]
    ///             {
    ///                 "example-repo-local",
    ///             },
    ///             Actions = new Artifactory.Inputs.PermissionTargetRepoActionsArgs
    ///             {
    ///                 Users = new[]
    ///                 {
    ///                     new Artifactory.Inputs.PermissionTargetRepoActionsUserArgs
    ///                     {
    ///                         Name = "anonymous",
    ///                         Permissions = new[]
    ///                         {
    ///                             "read",
    ///                             "write",
    ///                         },
    ///                     },
    ///                     new Artifactory.Inputs.PermissionTargetRepoActionsUserArgs
    ///                     {
    ///                         Name = "user1",
    ///                         Permissions = new[]
    ///                         {
    ///                             "read",
    ///                             "write",
    ///                         },
    ///                     },
    ///                 },
    ///                 Groups = new[]
    ///                 {
    ///                     new Artifactory.Inputs.PermissionTargetRepoActionsGroupArgs
    ///                     {
    ///                         Name = "readers",
    ///                         Permissions = new[]
    ///                         {
    ///                             "read",
    ///                         },
    ///                     },
    ///                     new Artifactory.Inputs.PermissionTargetRepoActionsGroupArgs
    ///                     {
    ///                         Name = "dev",
    ///                         Permissions = new[]
    ///                         {
    ///                             "read",
    ///                             "write",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Build = new Artifactory.Inputs.PermissionTargetBuildArgs
    ///         {
    ///             IncludesPatterns = new[]
    ///             {
    ///                 "**",
    ///             },
    ///             Repositories = new[]
    ///             {
    ///                 "artifactory-build-info",
    ///             },
    ///             Actions = new Artifactory.Inputs.PermissionTargetBuildActionsArgs
    ///             {
    ///                 Users = new[]
    ///                 {
    ///                     new Artifactory.Inputs.PermissionTargetBuildActionsUserArgs
    ///                     {
    ///                         Name = "anonymous",
    ///                         Permissions = new[]
    ///                         {
    ///                             "read",
    ///                         },
    ///                     },
    ///                     new Artifactory.Inputs.PermissionTargetBuildActionsUserArgs
    ///                     {
    ///                         Name = "user1",
    ///                         Permissions = new[]
    ///                         {
    ///                             "read",
    ///                             "write",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         ReleaseBundle = new Artifactory.Inputs.PermissionTargetReleaseBundleArgs
    ///         {
    ///             IncludesPatterns = new[]
    ///             {
    ///                 "**",
    ///             },
    ///             Repositories = new[]
    ///             {
    ///                 "release-bundles",
    ///             },
    ///             Actions = new Artifactory.Inputs.PermissionTargetReleaseBundleActionsArgs
    ///             {
    ///                 Users = new[]
    ///                 {
    ///                     new Artifactory.Inputs.PermissionTargetReleaseBundleActionsUserArgs
    ///                     {
    ///                         Name = "anonymous",
    ///                         Permissions = new[]
    ///                         {
    ///                             "read",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Permissions
    /// 
    /// The provider supports the following `permission` enums:
    /// 
    /// * `read`
    /// * `write`
    /// * `annotate`
    /// * `delete`
    /// * `manage`
    /// * `managedXrayMeta`
    /// * `distribute`
    /// 
    /// The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):
    /// 
    /// * `read` - matches `Read` permissions.
    /// * `write` - matches `  Deploy / Cache / Create ` permissions.
    /// * `annotate` - matches `Annotate` permissions.
    /// * `delete` - matches `Delete / Overwrite` permissions.
    /// * `manage` - matches `Manage` permissions.
    /// * `managedXrayMeta` - matches `Manage Xray Metadata` permissions.
    /// * `distribute` - matches `Distribute` permissions.
    /// 
    /// ## Import
    /// 
    /// Permission targets can be imported using their name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/permissionTarget:PermissionTarget")]
    public partial class PermissionTarget : global::Pulumi.CustomResource
    {
        /// <summary>
        /// As for repo but for artifactory-build-info permissions.
        /// </summary>
        [Output("build")]
        public Output<Outputs.PermissionTargetBuild?> Build { get; private set; } = null!;

        /// <summary>
        /// Name of permission.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// As for repo for for release-bundles permissions.
        /// </summary>
        [Output("releaseBundle")]
        public Output<Outputs.PermissionTargetReleaseBundle?> ReleaseBundle { get; private set; } = null!;

        /// <summary>
        /// Repository permission configuration.
        /// </summary>
        [Output("repo")]
        public Output<Outputs.PermissionTargetRepo?> Repo { get; private set; } = null!;


        /// <summary>
        /// Create a PermissionTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PermissionTarget(string name, PermissionTargetArgs? args = null, CustomResourceOptions? options = null)
            : base("artifactory:index/permissionTarget:PermissionTarget", name, args ?? new PermissionTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PermissionTarget(string name, Input<string> id, PermissionTargetState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/permissionTarget:PermissionTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PermissionTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PermissionTarget Get(string name, Input<string> id, PermissionTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new PermissionTarget(name, id, state, options);
        }
    }

    public sealed class PermissionTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// As for repo but for artifactory-build-info permissions.
        /// </summary>
        [Input("build")]
        public Input<Inputs.PermissionTargetBuildArgs>? Build { get; set; }

        /// <summary>
        /// Name of permission.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// As for repo for for release-bundles permissions.
        /// </summary>
        [Input("releaseBundle")]
        public Input<Inputs.PermissionTargetReleaseBundleArgs>? ReleaseBundle { get; set; }

        /// <summary>
        /// Repository permission configuration.
        /// </summary>
        [Input("repo")]
        public Input<Inputs.PermissionTargetRepoArgs>? Repo { get; set; }

        public PermissionTargetArgs()
        {
        }
        public static new PermissionTargetArgs Empty => new PermissionTargetArgs();
    }

    public sealed class PermissionTargetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// As for repo but for artifactory-build-info permissions.
        /// </summary>
        [Input("build")]
        public Input<Inputs.PermissionTargetBuildGetArgs>? Build { get; set; }

        /// <summary>
        /// Name of permission.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// As for repo for for release-bundles permissions.
        /// </summary>
        [Input("releaseBundle")]
        public Input<Inputs.PermissionTargetReleaseBundleGetArgs>? ReleaseBundle { get; set; }

        /// <summary>
        /// Repository permission configuration.
        /// </summary>
        [Input("repo")]
        public Input<Inputs.PermissionTargetRepoGetArgs>? Repo { get; set; }

        public PermissionTargetState()
        {
        }
        public static new PermissionTargetState Empty => new PermissionTargetState();
    }
}
