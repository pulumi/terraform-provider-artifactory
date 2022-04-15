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
    /// ## # Artifactory General Security Resource
    /// 
    /// This resource can be used to manage Artifactory's general security settings.
    /// 
    /// Only a single `artifactory.GeneralSecurity` resource is meant to be defined.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Configure Artifactory general security settings
    ///         var security = new Artifactory.GeneralSecurity("security", new Artifactory.GeneralSecurityArgs
    ///         {
    ///             EnableAnonymousAccess = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Current general security settings can be imported using `security` as the `ID`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/generalSecurity:GeneralSecurity security security
    /// ```
    /// 
    ///  environments, or may change without notice.
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/generalSecurity:GeneralSecurity")]
    public partial class GeneralSecurity : Pulumi.CustomResource
    {
        [Output("enableAnonymousAccess")]
        public Output<bool?> EnableAnonymousAccess { get; private set; } = null!;


        /// <summary>
        /// Create a GeneralSecurity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeneralSecurity(string name, GeneralSecurityArgs? args = null, CustomResourceOptions? options = null)
            : base("artifactory:index/generalSecurity:GeneralSecurity", name, args ?? new GeneralSecurityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeneralSecurity(string name, Input<string> id, GeneralSecurityState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/generalSecurity:GeneralSecurity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GeneralSecurity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeneralSecurity Get(string name, Input<string> id, GeneralSecurityState? state = null, CustomResourceOptions? options = null)
        {
            return new GeneralSecurity(name, id, state, options);
        }
    }

    public sealed class GeneralSecurityArgs : Pulumi.ResourceArgs
    {
        [Input("enableAnonymousAccess")]
        public Input<bool>? EnableAnonymousAccess { get; set; }

        public GeneralSecurityArgs()
        {
        }
    }

    public sealed class GeneralSecurityState : Pulumi.ResourceArgs
    {
        [Input("enableAnonymousAccess")]
        public Input<bool>? EnableAnonymousAccess { get; set; }

        public GeneralSecurityState()
        {
        }
    }
}
