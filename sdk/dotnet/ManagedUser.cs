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
    /// ## # Artifactory Managed User Resource
    /// 
    /// Provides an Artifactory managed user resource. This can be used to create and maintain Artifactory users. For example, service account where password is known and managed externally.
    /// 
    /// Unlike `artifactory.UnmanagedUser` and `artifactory.User`, the `password` attribute is required and cannot be empty.
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
    ///         // Create a new Artifactory user called terraform
    ///         var test_user = new Artifactory.ManagedUser("test-user", new Artifactory.ManagedUserArgs
    ///         {
    ///             Email = "test-user@artifactory-terraform.com",
    ///             Groups = 
    ///             {
    ///                 "logged-in-users",
    ///                 "readers",
    ///             },
    ///             Password = "my super secret password",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Users can be imported using their name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/managedUser:ManagedUser test-user myusername
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/managedUser:ManagedUser")]
    public partial class ManagedUser : Pulumi.CustomResource
    {
        /// <summary>
        /// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        /// </summary>
        [Output("admin")]
        public Output<bool?> Admin { get; private set; } = null!;

        /// <summary>
        /// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        /// </summary>
        [Output("disableUiAccess")]
        public Output<bool?> DisableUiAccess { get; private set; } = null!;

        /// <summary>
        /// Email for user.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// List of groups this user is a part of.
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Output("internalPasswordDisabled")]
        public Output<bool?> InternalPasswordDisabled { get; private set; } = null!;

        /// <summary>
        /// Username for user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password for the user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        /// </summary>
        [Output("profileUpdatable")]
        public Output<bool?> ProfileUpdatable { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedUser(string name, ManagedUserArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/managedUser:ManagedUser", name, args ?? new ManagedUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedUser(string name, Input<string> id, ManagedUserState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/managedUser:ManagedUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManagedUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedUser Get(string name, Input<string> id, ManagedUserState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedUser(name, id, state, options);
        }
    }

    public sealed class ManagedUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        /// </summary>
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        /// <summary>
        /// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        /// </summary>
        [Input("disableUiAccess")]
        public Input<bool>? DisableUiAccess { get; set; }

        /// <summary>
        /// Email for user.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of groups this user is a part of.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Input("internalPasswordDisabled")]
        public Input<bool>? InternalPasswordDisabled { get; set; }

        /// <summary>
        /// Username for user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for the user.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        /// </summary>
        [Input("profileUpdatable")]
        public Input<bool>? ProfileUpdatable { get; set; }

        public ManagedUserArgs()
        {
        }
    }

    public sealed class ManagedUserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        /// </summary>
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        /// <summary>
        /// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        /// </summary>
        [Input("disableUiAccess")]
        public Input<bool>? DisableUiAccess { get; set; }

        /// <summary>
        /// Email for user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of groups this user is a part of.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Input("internalPasswordDisabled")]
        public Input<bool>? InternalPasswordDisabled { get; set; }

        /// <summary>
        /// Username for user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for the user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        /// </summary>
        [Input("profileUpdatable")]
        public Input<bool>? ProfileUpdatable { get; set; }

        public ManagedUserState()
        {
        }
    }
}
