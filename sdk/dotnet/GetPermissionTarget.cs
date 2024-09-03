// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetPermissionTarget
    {
        /// <summary>
        /// # Artifactory Permission Target Data Source
        /// 
        /// Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
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
        ///     //
        ///     var target1 = Artifactory.GetPermissionTarget.Invoke(new()
        ///     {
        ///         Name = "my_permission",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPermissionTargetResult> InvokeAsync(GetPermissionTargetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPermissionTargetResult>("artifactory:index/getPermissionTarget:getPermissionTarget", args ?? new GetPermissionTargetArgs(), options.WithDefaults());

        /// <summary>
        /// # Artifactory Permission Target Data Source
        /// 
        /// Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.
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
        ///     //
        ///     var target1 = Artifactory.GetPermissionTarget.Invoke(new()
        ///     {
        ///         Name = "my_permission",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPermissionTargetResult> Invoke(GetPermissionTargetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPermissionTargetResult>("artifactory:index/getPermissionTarget:getPermissionTarget", args ?? new GetPermissionTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPermissionTargetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Same as repo but for artifactory-build-info permissions.
        /// </summary>
        [Input("build")]
        public Inputs.GetPermissionTargetBuildArgs? Build { get; set; }

        /// <summary>
        /// Name of the permission target.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Same as repo but for release-bundles permissions.
        /// </summary>
        [Input("releaseBundle")]
        public Inputs.GetPermissionTargetReleaseBundleArgs? ReleaseBundle { get; set; }

        /// <summary>
        /// Repository permission configuration.
        /// </summary>
        [Input("repo")]
        public Inputs.GetPermissionTargetRepoArgs? Repo { get; set; }

        public GetPermissionTargetArgs()
        {
        }
        public static new GetPermissionTargetArgs Empty => new GetPermissionTargetArgs();
    }

    public sealed class GetPermissionTargetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Same as repo but for artifactory-build-info permissions.
        /// </summary>
        [Input("build")]
        public Input<Inputs.GetPermissionTargetBuildInputArgs>? Build { get; set; }

        /// <summary>
        /// Name of the permission target.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Same as repo but for release-bundles permissions.
        /// </summary>
        [Input("releaseBundle")]
        public Input<Inputs.GetPermissionTargetReleaseBundleInputArgs>? ReleaseBundle { get; set; }

        /// <summary>
        /// Repository permission configuration.
        /// </summary>
        [Input("repo")]
        public Input<Inputs.GetPermissionTargetRepoInputArgs>? Repo { get; set; }

        public GetPermissionTargetInvokeArgs()
        {
        }
        public static new GetPermissionTargetInvokeArgs Empty => new GetPermissionTargetInvokeArgs();
    }


    [OutputType]
    public sealed class GetPermissionTargetResult
    {
        /// <summary>
        /// Same as repo but for artifactory-build-info permissions.
        /// </summary>
        public readonly Outputs.GetPermissionTargetBuildResult? Build;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// Same as repo but for release-bundles permissions.
        /// </summary>
        public readonly Outputs.GetPermissionTargetReleaseBundleResult? ReleaseBundle;
        /// <summary>
        /// Repository permission configuration.
        /// </summary>
        public readonly Outputs.GetPermissionTargetRepoResult? Repo;

        [OutputConstructor]
        private GetPermissionTargetResult(
            Outputs.GetPermissionTargetBuildResult? build,

            string id,

            string name,

            Outputs.GetPermissionTargetReleaseBundleResult? releaseBundle,

            Outputs.GetPermissionTargetRepoResult? repo)
        {
            Build = build;
            Id = id;
            Name = name;
            ReleaseBundle = releaseBundle;
            Repo = repo;
        }
    }
}
